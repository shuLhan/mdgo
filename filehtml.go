// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package mdgo

import (
	"fmt"
	"html/template"
	"strings"

	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

//
// fileHTML represent an HTML metadata for header and its body.
//
type fileHTML struct {
	Title       string
	EmbeddedCSS *template.CSS
	Styles      []string
	Body        template.HTML
	Metadata    map[string]string

	path    string
	rawBody strings.Builder
}

func (fhtml *fileHTML) unpackMarkdownMetadata(ctx parser.Context) {
	metadata := meta.Get(ctx)
	for k, v := range metadata {
		vstr, ok := v.(string)
		if !ok {
			vstr = fmt.Sprintf("%s", v)
		}

		switch k {
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, vstr)
		default:
			fhtml.Metadata[k] = vstr
		}
	}

	if len(fhtml.Styles) == 0 {
		fhtml.EmbeddedCSS = embeddedCSS()
	}

	fhtml.Body = template.HTML(fhtml.rawBody.String())
}
