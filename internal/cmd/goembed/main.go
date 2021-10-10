// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate_main.go

package main

import (
	"log"

	"git.sr.ht/~shulhan/mdgo"
)

func main() {
	opts := mdgo.EmbedOptions{
		ConvertOptions: mdgo.ConvertOptions{
			Root:         "_example",
			HtmlTemplate: "_example/html.tmpl",
		},
		PackageName: "main",
		VarName:     "mdgoFS",
		GoFileName:  "cmd/mdgo-example/static.go",
	}
	err := mdgo.GoEmbed(&opts)
	if err != nil {
		log.Fatal(err)
	}
}
