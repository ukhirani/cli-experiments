// Source - https://stackoverflow.com/a
// Posted by Zeke Lu
// Retrieved 2025-11-13, License - CC BY-SA 4.0

package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func main() {
	// See "man pdftotext" for more options.
	args := []string{
		"-layout",    // Maintain (as best as possible) the original physical layout of the text.
		"-nopgbrk",   // Don't insert page breaks (form feed characters) between pages.
		"ddia.pdf",   // The input file.
		"output.txt", // Send the output to stdout.
	}
	cmd := exec.CommandContext(context.Background(), "pdftotext", args...)

	var buf bytes.Buffer
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(buf.String())
}
