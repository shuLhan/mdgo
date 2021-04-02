// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package mdgo

import (
	"fmt"
	"regexp"

	"github.com/shuLhan/share/lib/memfs"
)

const (
	defAddress = ":8080"
	defRoot    = "."
)

//
// ServeOptions define the options for Serve function.
//
type ServeOptions struct {
	// Mfs contains pointer to variable generated from Generate.
	// This option is used to use embedded files for serving on HTTP.
	Mfs *memfs.MemFS

	// Address to listen and serve for HTTP request.
	Address string

	// HTMLTemplate the HTML template to be used when converting markup
	// file into HTML.
	// If empty it will default to use embedded HTML template.
	// See template_index_html.go for template format.
	HtmlTemplate string

	// Root directory where its content will be embedded into Go source
	// code.
	// Default to DefaultRoot if its empty.
	Root string

	// Exclude contains regular expresion to exclude certain paths from
	// being scanned.
	Exclude string
	excRE   []*regexp.Regexp
}

func (opts *ServeOptions) init() (err error) {
	var (
		logp = "ServeOptions.init"
	)

	if len(opts.Address) == 0 {
		opts.Address = defAddress
	}
	if len(opts.Root) == 0 {
		opts.Root = defRoot
	}
	if len(opts.Exclude) > 0 {
		re, err := regexp.Compile(opts.Exclude)
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
		opts.excRE = append(opts.excRE, re)
		defExcludes = append(defExcludes, opts.Exclude)
	}
	return nil
}
