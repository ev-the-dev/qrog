package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"rsc.io/qr"
)

func main() {
	level := flag.String("level", "M", "Error correction level: L, M, Q, H")
	output := flag.String("output", "qr.png", "Output file path (default: ./qr.png)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: qrog [flags] <text>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	text := flag.Arg(0)

	var correctionLevel qr.Level
	switch *level {
	case "L":
		correctionLevel = qr.L
	case "M":
		correctionLevel = qr.M
	case "Q":
		correctionLevel = qr.Q
	case "H":
		correctionLevel = qr.H
	default:
		fmt.Fprintf(os.Stderr, "invalid level %q: must be L, M, Q, or H\n", *level)
		os.Exit(1)
	}

	code, err := qr.Encode(text, correctionLevel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode: %v\n", err)
		os.Exit(1)
	}

	outPath := *output
	if !filepath.IsAbs(outPath) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get working directory: %v\n", err)
			os.Exit(1)
		}
		outPath = filepath.Join(cwd, outPath)
	}

	if err := os.WriteFile(outPath, code.PNG(), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("QR code written to %s\n", outPath)
}
