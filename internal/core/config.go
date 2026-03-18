// Package core contains the fundamental configuration and processing logic for bunchotext.
package core

// FilePatterns defines the mapping between preset names and their associated file extensions.
// These presets are used by the basic and auto modes to filter which files to include.
// Key: preset name (e.g., "go", "py"), Value: slice of file extensions to match.
var FilePatterns = map[string][]string{
	"go":   {".go", ".mod", ".sum"},         // Go source files and module definitions
	"py":   {".py", ".pyw", ".ipynb"},       // Python scripts and Jupyter notebooks
	"js":   {".js", ".jsx", ".mjs", ".cjs"}, // JavaScript/Node.js files
	"ts":   {".ts", ".tsx", ".d.ts"},        // TypeScript source and declaration files
	"json": {".json", ".jsonc"},             // JSON maps
}

// IgnoreDirs contains directory names that should be skipped during traversal by default.
// These are commonly used for dependencies, version control, IDE metadata, and build artifacts.
// The all mode can optionally bypass this list via the --no-ignore-dirs flag.
var IgnoreDirs = map[string]bool{
	".git":         true, // Git repository metadata
	"node_modules": true, // Node.js dependencies (often very large)
	"vendor":       true, // Go vendored dependencies
	"__pycache__":  true, // Python bytecode cache
	"dist":         true, // Build output directory
	".idea":        true, // JetBrains IDE configuration
	".vscode":      true, // VS Code workspace settings
	".vscode-test": true, // VS Code test instance dir
	".obsidian":    true, // Obsidian editor settings
}
