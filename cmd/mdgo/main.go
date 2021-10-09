// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// mdgo is a command line interface (CLI) to convert, generate, and/or serve a
// directory that contains markup files, as HTML files.
//
// Usage
//
// The following section describe how to use mdgo CLI.
//
//	mdgo [-template <file>] [-exclude <regex>] convert <dir>
//
// Scan the "dir" recursively to find markup files and convert them into HTML
// files.
// The template "file" is optional, default to embedded HTML template.
//
//	mdgo [-template <file>] [-exclude <regex>] [-out <file>] generate <dir>
//
// Convert all the markup files inside directory "dir" recursively and then
// embed them into ".go" source file.
// The output file is optional, default to "mdgo_static.go" in current
// directory.
//
//	mdgo [-template <file>] [-exclude <regex>] [-address <ip:port>] serve <dir>
//
// Serve all files inside directory "dir" using HTTP server, watch changes on
// markup files and convert them to HTML files.
// If the address is not set, its default to ":8080".
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"git.sr.ht/~shulhan/mdgo"
	"github.com/shuLhan/share/lib/debug"
)

func main() {
	isHelp := flag.Bool("help", false, "print help")

	htmlTemplate := flag.String("template", "", "path to HTML template")
	outputFile := flag.String("out", "mdgo_static.go",
		"path to output of .go generated file")
	address := flag.String("address", ":8080",
		"the binding address for HTTP server")
	exclude := flag.String("exclude", "",
		"a regex to exclude certain paths from being scanned during covert, generate, watch, or serve")

	flag.Parse()

	if *isHelp {
		usage()
		os.Exit(0)
	}

	command := flag.Arg(0)
	if len(command) == 0 {
		usage()
		os.Exit(1)
	}

	dir := flag.Arg(1)
	if len(dir) == 0 {
		dir = mdgo.DefaultRoot
	}

	var err error
	command = strings.ToLower(command)

	switch command {
	case "convert":
		opts := mdgo.ConvertOptions{
			Root:         dir,
			HtmlTemplate: *htmlTemplate,
			Exclude:      *exclude,
		}
		err = mdgo.Convert(&opts)

	case "generate":
		genOpts := mdgo.GenerateOptions{
			ConvertOptions: mdgo.ConvertOptions{
				Root:         dir,
				HtmlTemplate: *htmlTemplate,
				Exclude:      *exclude,
			},
			GenGoFileName: *outputFile,
		}
		err = mdgo.Generate(&genOpts)

	case "serve":
		debug.Value = 1
		opts := mdgo.ServeOptions{
			ConvertOptions: mdgo.ConvertOptions{
				Root:         dir,
				HtmlTemplate: *htmlTemplate,
				Exclude:      *exclude,
			},
			Address: *address,
		}
		err = mdgo.Serve(&opts)

	default:
		usage()
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println(`
=  mdgo

A CLI to convert, generate, and/or serve a directory that contains markup
files, as HTML files.

==  Usage

mdgo [-template <file>] [-exclude <regex>] convert <dir>

	Scan the "dir" recursively to find markup files.
	and convert them into HTML files.
	The template "file" is optional, default to embedded HTML template.

mdgo [-template <file>] [-exclude <regex>] [-out <file>] generate <dir>

	Convert all markup files inside directory "dir" recursively and then
	embed them into ".go" source file.
	The output file is optional, default to "mdgo_static.go" in current
	directory.

mdgo [-template <file>] [-exclude <regex>] [-address <ip:port>] serve <dir>

	Serve all files inside directory "dir" using HTTP server, watch
	changes on markup files and convert them to HTML files automatically.
	If the address is not set, its default to ":8080".`)
}
