package cmd

import (
	"fmt"

	"github.com/quadrocorp/bunchotext/internal/core"
	"github.com/spf13/cobra"
)

// autoCmd implements the "auto" subcommand which automatically detects
// the dominant file type in the target directory and bundles accordingly.
var autoCmd = &cobra.Command{
	Use:   "auto [-d dir] [-o outputFile]",
	Short: "Auto-detect file type and bundle",
	Long:  `Scans the directory, detects the most common file extension based on configured presets, and bundles files of that type. Ideal for quick context preparation without manual type selection.`,
	Args:  cobra.NoArgs,

	// PreRunE performs auto-detection before main execution.
	// Sets the global filetype variable for reuse in RunE.
	PreRunE: func(cmd *cobra.Command, args []string) error {
		detectedType, err := core.DetectDominantType(directory)
		if err != nil {
			return fmt.Errorf("failed to detect file type: %w", err)
		}
		if detectedType == "" {
			return fmt.Errorf("no recognizable files found in %s", directory)
		}
		filetype = detectedType // Set global for ProcessDirectory call
		fmt.Printf("Auto-detected type: %s\n", filetype)
		return nil
	},

	// RunE executes the bundling with auto-detected type.
	// Handles default output filename if user didn't specify.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Set sensible default output filename for auto mode
		if !cmd.Flags().Changed("output") {
			outputFile = "codebase.txt"
		}

		fmt.Printf("Scanning directory: %s\n", directory)
		fmt.Printf("Using auto-detected type: %s\n", filetype)
		fmt.Printf("Writing to: %s\n", outputFile)

		if err := core.ProcessDirectory(directory, filetype, outputFile); err != nil {
			return fmt.Errorf("processing failed: %w", err)
		}
		fmt.Println("Done!")
		return nil
	},
}

// init registers the auto subcommand with the root command.
// Persistent flags (-d, -o) are inherited automatically.
func init() {
	rootCmd.AddCommand(autoCmd)
	// No local flags needed: auto mode uses persistent flags only
}
