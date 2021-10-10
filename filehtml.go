// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package mdgo

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"strings"

	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

//
// fileHtml represent an HTML metadata for header and its body.
//
type fileHtml struct {
	Title       string
	EmbeddedCSS *template.CSS
	Styles      []string
	Body        template.HTML
	Metadata    map[string]string

	path    string
	finfo   fs.FileInfo
	rawBody strings.Builder
}

func newFileHtml(path string) (fhtml *fileHtml) {
	fhtml = &fileHtml{
		path: path,
	}
	fhtml.finfo, _ = os.Stat(path)
	return fhtml
}

func (fhtml *fileHtml) unpackMetadata(ctx parser.Context) {
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
