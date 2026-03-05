// Package main is the entry point for the bunchotext CLI application.
// It delegates all command handling to the cmd package following Cobra conventions.
package main

import "github.com/quadrocorp/bunchotext/internal/cmd"

// main initializes and executes the Cobra command tree.
// All flag parsing, validation, and business logic is handled in cmd/*.go.
func main() {
	cmd.Execute()
}
