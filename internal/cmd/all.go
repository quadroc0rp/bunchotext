package cmd

import (
	"fmt"

	"github.com/quadroc0rp/bunchotext/internal/core"
	"github.com/spf13/cobra"
)

// Flag variables specific to the 'all' subcommand.
var (
	respectGitignore bool // Enable .gitignore pattern matching
	useIgnoreDirs    bool // Enable skipping of standard ignored directories
)

// allCmd implements the "all" subcommand which bundles every text file
// regardless of extension, with optional filtering controls.
var allCmd = &cobra.Command{
	Use:   "all [-d dir] [-o outputFile] [--no-ignore-dirs] [--gitignore]",
	Short: "Bundle all files regardless of type",
	Long: `Bundles all text files in the directory with clear separators.

By default, excludes directories in the global IgnoreDirs list (.git, node_modules, vendor, etc.).
Use --no-ignore-dirs to include them.
Use --gitignore to also respect .gitignore rules from the root directory.

Ideal for complete codebase backups or comprehensive AI context preparation.`,
	Args: cobra.NoArgs,

	// RunE executes the all-files bundling with user-selected filters.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Set mode-specific default output filename
		if !cmd.Flags().Changed("output") {
			outputFile = "codebase_all.txt"
		}

		// Provide clear feedback about active filtering options
		fmt.Printf("Scanning directory: %s\n", directory)
		if useIgnoreDirs {
			fmt.Printf("Excluding standard dirs: .git, node_modules, vendor, etc.\n")
		} else {
			fmt.Printf("Including ALL directories (no IgnoreDirs filter)\n")
		}
		if respectGitignore {
			fmt.Printf("Also respecting .gitignore rules\n")
		}
		fmt.Printf("Including all file types\n")
		fmt.Printf("Writing to: %s\n", outputFile)

		if err := core.ProcessDirectoryAll(directory, outputFile, useIgnoreDirs, respectGitignore); err != nil {
			return fmt.Errorf("processing failed: %w", err)
		}
		fmt.Println("Done!")
		return nil
	},
}

// init registers the all subcommand and its local flags.
func init() {
	rootCmd.AddCommand(allCmd)

	// Local flags for 'all' mode only
	allCmd.Flags().BoolVar(&respectGitignore, "gitignore", false, "Respect .gitignore rules when scanning")
	allCmd.Flags().BoolVar(&useIgnoreDirs, "use-ignore-dirs", true, "Exclude standard directories like .git, node_modules, vendor (default: true)")
	allCmd.Flags().BoolVar(&useIgnoreDirs, "no-ignore-dirs", false, "Include all directories, even .git, node_modules, etc. (overrides --use-ignore-dirs)")

	// Handle mutual exclusivity: --no-ignore-dirs explicitly disables IgnoreDirs
	allCmd.PreRun = func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("no-ignore-dirs") {
			useIgnoreDirs = false
		}
	}
}
