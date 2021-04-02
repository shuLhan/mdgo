// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Package mdgo is a program to write static web server with embedded files
// using markdown markup languages.
//
// For more information see the README file at the page repository
// https://sr.ht/~shulhan/mdgo.
//
package mdgo

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/shuLhan/share/lib/memfs"
)

const (
	extMarkdown          = ".md"
	internalTemplatePath = "_internal/.template"
)

const (
	metadataStylesheet = "stylesheet"
)

//nolint: gochecknoglobals
var (
	defExcludes = []string{
		`.*\.md$`,
		`^\..*`,
	}
)

//
// Convert all markup files inside directory "dir" recursively into HTML
// files using "htmlTemplate" file as template.
// If htmlTemplate is empty it will default to use embedded HTML template.
// See template_index_html.go for template format.
//
func Convert(htmlTemplate, dir, exclude string) (err error) {
	var (
		logp     = "Convert"
		excludes []*regexp.Regexp
	)

	if len(dir) == 0 {
		dir = "."
	}
	if len(exclude) > 0 {
		re, err := regexp.Compile(exclude)
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
		excludes = append(excludes, re)
	}

	htmlg, err := newHTMLGenerator(nil, htmlTemplate, true)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups, err := listFileMarkups(dir, excludes)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg.convertFileMarkups(fileMarkups)

	return nil
}

//
// Generate a static Go file to be used for building binary.
//
// It will convert all markup files inside directory "dir" into HTML files,
// recursively; and read all the HTML files and files in "content/assets" and
// convert them into Go file in "out".
//
// If htmlTemplate is empty it will default to use embedded HTML template.
// See template_index_html.go for template format.
//
func Generate(opts *GenerateOptions) (err error) {
	var (
		logp = "Generate"
	)

	if opts == nil {
		opts = &GenerateOptions{}
	}
	err = opts.init()
	if err != nil {
		return err
	}

	htmlg, err := newHTMLGenerator(nil, opts.HTMLTemplate, true)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups, err := listFileMarkups(opts.Root, opts.excRE)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg.convertFileMarkups(fileMarkups)

	memfsOpts := &memfs.Options{
		Root:     opts.Root,
		Excludes: defExcludes,
	}
	mfs, err := memfs.New(memfsOpts)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	if len(opts.HTMLTemplate) > 0 {
		_, err = mfs.AddFile(internalTemplatePath, opts.HTMLTemplate)
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
	}

	err = mfs.GoGenerate(opts.GenPackageName, opts.GenVarName,
		opts.GenGoFileName, memfs.EncodingGzip)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	return nil
}

//
// Serve the content at directory "dir" using HTTP server at specific
// "address".
//
func Serve(opts *ServeOptions) (err error) {
	var (
		logp = "Serve"
	)

	if opts == nil {
		opts = &ServeOptions{}
	}

	err = opts.init()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}
	srv, err := newServer(opts)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}
	err = srv.start()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}
	return nil
}

func isExtensionMarkup(ext string) bool {
	return ext == extMarkdown
}

//
// listFileMarkups find any markup files inside the content directory,
// recursively.
//
func listFileMarkups(dir string, excRE []*regexp.Regexp) (fileMarkups map[string]*fileMarkup, err error) {
	logp := "listFileMarkups"

	d, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fis, err := d.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups = make(map[string]*fileMarkup)

	for _, fi := range fis {
		name := fi.Name()

		if fi.IsDir() && name[0] != '.' {
			newdir := filepath.Join(dir, fi.Name())
			fmarkups, err := listFileMarkups(newdir, excRE)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", logp, err)
			}
			for k, v := range fmarkups {
				fileMarkups[k] = v
			}
			continue
		}

		ext := strings.ToLower(filepath.Ext(name))
		if !isExtensionMarkup(ext) {
			continue
		}
		if fi.Size() == 0 {
			continue
		}

		filePath := filepath.Join(dir, name)

		if isExcluded(filePath, excRE) {
			continue
		}

		fmarkup := &fileMarkup{
			path:     filePath,
			info:     fi,
			basePath: strings.TrimSuffix(filePath, ext),
			fhtml:    &fileHTML{},
		}

		fmarkup.fhtml.path = fmarkup.basePath + ".html"

		fileMarkups[filePath] = fmarkup
	}

	return fileMarkups, nil
}

func isExcluded(path string, excs []*regexp.Regexp) bool {
	if len(excs) == 0 {
		return false
	}
	for _, re := range excs {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}
