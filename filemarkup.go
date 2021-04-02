// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package mdgo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type fileMarkup struct {
	path     string      // path contains full path to markup file.
	info     os.FileInfo // info contains FileInfo of markup file.
	basePath string      // basePath contains full path to file without markup extension.
	fhtml    *fileHTML   // The HTML output of this markup.
}

func newFileMarkup(filePath string, fi os.FileInfo) (fmarkup *fileMarkup, err error) {
	logp := "newFileMarkup"
	if len(filePath) == 0 {
		return nil, fmt.Errorf("%s: empty path", logp)
	}
	if fi == nil {
		fi, err = os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", logp, err)
		}
	}

	ext := strings.ToLower(filepath.Ext(filePath))

	fmarkup = &fileMarkup{
		path:     filePath,
		info:     fi,
		basePath: strings.TrimSuffix(filePath, ext),
		fhtml:    &fileHTML{},
	}

	fmarkup.fhtml.path = fmarkup.basePath + ".html"

	return fmarkup, nil
}
