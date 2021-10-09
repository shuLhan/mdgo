// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Program mdgo-example provide an example on how to build a binary that
// include the static, generated .go file.
//
package main

import (
	"log"

	"github.com/shuLhan/share/lib/memfs"

	"git.sr.ht/~shulhan/mdgo"
)

var mdgoFS *memfs.MemFS

func main() {
	opts := &mdgo.ServeOptions{
		ConvertOptions: mdgo.ConvertOptions{
			Root:         "_example",
			HtmlTemplate: "_example/html.tmpl",
		},
		Mfs:     mdgoFS,
		Address: ":8080",
	}

	err := mdgo.Serve(opts)
	if err != nil {
		log.Fatal(err)
	}
}
