// Package cmd defines the Cobra command structure for the bunchotext CLI.
// It contains the root command and shared flag definitions.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/quadrocorp/bunchotext/internal/core"
	"github.com/spf13/cobra"
)

// Global flag variables shared across commands via persistent flags.
// These are populated by Cobra's flag parsing before command execution.
var (
	outputFile string // Destination file path for bundled output
	filetype   string // File type preset for basic mode (e.g., "go", "py")
	directory  string // Root directory to scan for files
)

// rootCmd represents the base command when called without subcommands.
// This implements the "basic mode" which requires a -t/--type flag.
var rootCmd = &cobra.Command{
	Use:     "bunchotext -t type [-d dir] [-o outputFile]",
	Short:   "Extract codebase into a single .txt file",
	Long:    `bunchotext is a CLI tool used to extract all your codebase into a single .txt file, optimized for LLM context preparation.`,
	Args:    cobra.NoArgs, // No positional arguments expected
	Version: "1.1.0",      // Semantic version following SemVer 2.0

	// PreRunE executes after flag parsing but before RunE.
	// Used here for validation of required flags and preset values.
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Ensure the required -t flag was provided
		if filetype == "" {
			return fmt.Errorf("required flag -t/--type not set. Available types: %s", getAvailableTypes())
		}
		// Validate that the provided type exists in our configuration
		if _, exists := core.FilePatterns[filetype]; !exists {
			return fmt.Errorf("invalid type '%s'. Available types: %s", filetype, getAvailableTypes())
		}
		return nil
	},

	// RunE contains the main execution logic for basic mode.
	// Returns error to allow Cobra to handle exit codes and error messaging.
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Scanning directory: %s\n", directory)
		fmt.Printf("Filtering for type: %s\n", filetype)
		fmt.Printf("Writing to: %s\n", outputFile)

		if err := core.ProcessDirectory(directory, filetype, outputFile); err != nil {
			return fmt.Errorf("processing failed: %w", err)
		}
		fmt.Println("Done!")
		return nil
	},
}

// Execute runs the root command, handling any errors by printing and exiting.
// This is the entry point called from main.go.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Cobra already prints the error; we just ensure non-zero exit
		os.Exit(1)
	}
}

// init registers flags and subcommands when the package is imported.
// Cobra convention: use init() for command setup, not main().
func init() {
	// Persistent flags are inherited by all subcommands (auto, all)
	rootCmd.PersistentFlags().StringVarP(&directory, "dir", "d", ".", "Directory to search for files")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "output.txt", "Output file path for bundled result")

	// Local flag: only applies to root command (basic mode)
	rootCmd.Flags().StringVarP(&filetype, "type", "t", "", fmt.Sprintf("File type preset (%s)", getAvailableTypes()))
	_ = rootCmd.MarkFlagRequired("type") // Enforce -t as required for basic mode
}

// getAvailableTypes returns a comma-separated list of valid preset names.
// Used in help text and error messages for user guidance.
func getAvailableTypes() string {
	keys := make([]string, 0, len(core.FilePatterns))
	for k := range core.FilePatterns {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}
