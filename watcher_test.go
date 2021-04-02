// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package mdgo

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/shuLhan/share/lib/test"
)

var (
	testWatcher *watcher
	testFileMd  string
	testMdFile  *os.File
)

func TestWatcher(t *testing.T) {
	testDir := "testdata/watcher"

	err := os.RemoveAll(testDir)
	if err != nil {
		t.Fatal(err)
	}

	err = os.MkdirAll(testDir, 0700)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})

	htmlg, err := newHTMLGenerator(nil, "testdata/html.tmpl", true)
	if err != nil {
		t.Fatal(err)
	}

	testWatcher, err = newWatcher(htmlg, testDir)
	if err != nil {
		t.Fatal(err)
	}

	err = testWatcher.start()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("createMdFile", testCreate)
	t.Run("updateMdFile", testUpdate)
	t.Run("deleteMdFile", testDelete)
}

func testCreate(t *testing.T) {
	var (
		err error
	)
	testFileMd = filepath.Join(testWatcher.dir, "index.md")
	testMdFile, err = os.Create(testFileMd)
	if err != nil {
		t.Fatal(err)
	}

	got := waitChanges()

	test.Assert(t, "New md file created", testFileMd, got.path, true)

	expBody := ``
	gotBody := removeFooter(string(got.fhtml.Body))
	test.Assert(t, "HTML body", expBody, gotBody, true)
}

func testUpdate(t *testing.T) {
	_, err := testMdFile.WriteString("# Hello")
	if err != nil {
		t.Fatal(err)
	}
	err = testMdFile.Sync()
	if err != nil {
		t.Fatal(err)
	}

	got := waitChanges()
	test.Assert(t, "md file updated", testFileMd, got.path, true)

	expBody := `<h1>Hello</h1>
`
	gotBody := removeFooter(string(got.fhtml.Body))
	test.Assert(t, "HTML body", expBody, gotBody, true)
}

func testDelete(t *testing.T) {
	err := testMdFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(testFileMd)
	if err != nil {
		t.Fatal(err)
	}

	got := waitChanges()
	test.Assert(t, "md file updated", testFileMd, got.path, true)

	_, gotIsExist := testWatcher.fileMarkups[testFileMd]
	test.Assert(t, "md file deleted", false, gotIsExist, true)
}

//
// removeFooter remove the footer from generated HTML. The footer is 4 lines
// at the bottom.
//
func removeFooter(in string) string {
	lines := strings.Split(in, "\n")
	n := len(lines)
	if n > 5 {
		lines = lines[:n-5]
	}
	return strings.Join(lines, "\n")
}

func waitChanges() (fmarkup *fileMarkup) {
	var (
		ok bool
	)

	for {
		time.Sleep(5000)
		fmarkup, ok = testWatcher.changes.Pop().(*fileMarkup)
		if ok {
			break
		}
	}
	return fmarkup
}
