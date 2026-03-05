// Package core provides the file traversal and bundling functionality.
package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	gitignore "github.com/sabhiram/go-gitignore"
)

// ProcessDirectory scans the given rootDir for files matching the specified patternKey,
// concatenates their contents with clear separators, and writes the result to outFile.
// Returns an error if file creation or traversal fails.
func ProcessDirectory(rootDir, patternKey, outFile string) error {
	// Validate that the requested pattern exists in our configuration
	extensions, ok := FilePatterns[patternKey]
	if !ok {
		return fmt.Errorf("unknown pattern key: %s", patternKey)
	}

	// Create the output file, returning early if we cannot write to it
	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close() // Ensure file handle is released when done

	// Walk the directory tree, processing each file that matches our criteria
	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err // Propagate filesystem errors
		}

		// Skip directories that are in our ignore list
		if d.IsDir() {
			if IgnoreDirs[d.Name()] {
				return filepath.SkipDir // Don't descend into ignored directories
			}
		}

		// Skip files that don't match our target extensions
		if !hasExtension(path, extensions) {
			return nil
		}

		// Read file content, logging warnings but continuing on read errors
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not read %s: %v\n", path, err)
			return nil // Continue processing other files
		}

		// Write formatted header with file path and separator lines
		headerText := fmt.Sprintf("# %s", path)
		separator := strings.Repeat("=", len(headerText))
		newLine := ""

		fmt.Fprintln(f, separator)
		fmt.Fprintln(f, headerText)
		fmt.Fprintln(f, separator)
		fmt.Fprintln(f, newLine)

		// Write file content, ensuring it ends with a newline for clean formatting
		f.Write(content)
		if !bytesEndWithNewLine(content) {
			f.Write([]byte("\n"))
		}
		fmt.Fprintln(f) // Add blank line between files for readability

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}
	return nil
}

// DetectDominantType analyzes the directory tree to find which file type preset
// has the most matching files. This enables the auto mode to intelligently select
// the appropriate language filter without user input.
// Returns the preset name with the highest count, or empty string if none found.
func DetectDominantType(rootDir string) (string, error) {
	typeCount := make(map[string]int) // Track occurrence count per preset

	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if IgnoreDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Match file against each preset; count first match only
		for patternKey, extensions := range FilePatterns {
			if hasExtension(path, extensions) {
				typeCount[patternKey]++
				break // Prevent double-counting files with multiple matching extensions
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	// Find the preset with maximum occurrences
	var dominantType string
	maxCount := 0
	for patternKey, count := range typeCount {
		if count > maxCount {
			maxCount = count
			dominantType = patternKey
		}
	}
	return dominantType, nil
}

// ProcessDirectoryAll bundles ALL text files regardless of extension, with optional
// filtering via IgnoreDirs and .gitignore support. This is used by the 'all' subcommand.
// Parameters:
//   - rootDir: directory to scan
//   - outFile: destination file path
//   - useIgnoreDirs: if true, skip directories listed in IgnoreDirs config
//   - respectGitignore: if true, also apply .gitignore rules from rootDir
func ProcessDirectoryAll(rootDir, outFile string, useIgnoreDirs bool, respectGitignore bool) error {
	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	// Optionally compile .gitignore rules if requested and file exists
	var ignoreCompiler *gitignore.GitIgnore
	if respectGitignore {
		gitignorePath := filepath.Join(rootDir, ".gitignore")
		if _, err := os.Stat(gitignorePath); err == nil {
			ignoreCompiler, err = gitignore.CompileIgnoreFile(gitignorePath)
			if err != nil {
				return fmt.Errorf("failed to compile .gitignore: %w", err)
			}
			fmt.Fprintf(os.Stderr, "Loaded .gitignore from %s\n", gitignorePath)
		}
	}

	// Traverse directory tree with configurable filtering
	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Handle directory-level filtering
		if d.IsDir() {
			// Apply IgnoreDirs filter if enabled
			if useIgnoreDirs && IgnoreDirs[d.Name()] {
				return filepath.SkipDir
			}
			// Apply .gitignore filter to directories if enabled
			if respectGitignore && ignoreCompiler != nil {
				relPath, _ := filepath.Rel(rootDir, path)
				if ignoreCompiler.MatchesPath(relPath) {
					return filepath.SkipDir
				}
			}
			return nil
		}

		// Compute relative path for gitignore matching on files
		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			relPath = path // Fallback to absolute path if relative fails
		}

		// Skip files matching .gitignore patterns if enabled
		if respectGitignore && ignoreCompiler != nil && ignoreCompiler.MatchesPath(relPath) {
			return nil
		}

		// Read and process file content
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not read %s: %v\n", path, err)
			return nil
		}

		// Skip binary files to avoid corrupting output with non-text data
		if isBinary(content) {
			return nil
		}

		// Write formatted output with clear file boundaries
		headerText := fmt.Sprintf("# %s", path)
		separator := strings.Repeat("=", len(headerText))
		newLine := ""

		fmt.Fprintln(f, separator)
		fmt.Fprintln(f, headerText)
		fmt.Fprintln(f, separator)
		fmt.Fprintln(f, newLine)
		f.Write(content)
		if !bytesEndWithNewLine(content) {
			f.Write([]byte("\n"))
		}
		fmt.Fprintln(f)

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}
	return nil
}

// isBinary detects if a file is likely binary by checking for null bytes
// in the first 8KB of data. This heuristic efficiently skips images,
// compiled binaries, and other non-text files.
func isBinary(data []byte) bool {
	limit := 8192
	if len(data) < limit {
		limit = len(data)
	}
	for i := 0; i < limit; i++ {
		if data[i] == 0 {
			return true
		}
	}
	return false
}

// hasExtension checks if a file path ends with any of the allowed extensions.
// Used for filtering files by type preset.
func hasExtension(path string, allowed []string) bool {
	for _, ext := range allowed {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

// bytesEndWithNewLine returns true if the byte slice ends with a newline character.
// Used to ensure clean formatting when concatenating files.
func bytesEndWithNewLine(b []byte) bool {
	return len(b) > 0 && b[len(b)-1] == '\n'
}
