#  mdgo

[![PkgGoDev](https://pkg.go.dev/badge/git.sr.ht/~shulhan/mdgo)](https://pkg.go.dev/git.sr.ht/~shulhan/mdgo)

`mdgo` is a library and a program to write static web server with embedded
files using markdown markup language.

For working with asciidoc markup language with Go see
[ciigo](https://git.sr.ht/~shulhan/ciigo).


##  mdgo as command

mdgo provide a command line interface (CLI) can convert, generate, and/or
serve a directory that contains markup files, as HTML files.

###  Usage

```
$ mdgo [-template <file>] convert <dir>
```

Scan the "dir" recursively to find markup files (.md) and convert them into
HTML files.
The template "file" is optional, default to embedded HTML template.

```
$ mdgo [-template <file>] [-out <file>] generate <dir>
```

Convert all markup files inside directory "dir" recursively and then
embed them into ".go" source file.
The output file is optional, default to "mdgo_static.go" in current
directory.

```
$ mdgo [-template <file>] [-address <ip:port>] serve <dir>
```

Serve all files inside directory "dir" using HTTP server, watch
changes on markup files and convert them to HTML files automatically.
If the address is not set, its default to ":8080".


##  mdgo as library

This section describe step by step instructions on how to build and create
pages to be viewed for local development using `mdgo`.

First, clone the `mdgo` repository.
Let says that we have cloned the `mdgo` repository into
`$HOME/go/src/git.sr.ht/~shulhan/mdgo`.

Create new Go repository for building a website.
For example, in directory `$HOME/go/src/remote.tld/user/mysite`.
Replace "remote.tld/user/mysite" with your private or public repository.

```
$ mkdir -p $HOME/go/src/remote.tld/user/mysite
$ cd $HOME/go/src/remote.tld/user/mysite
```

Initialize the Go module,

```
$ go mod init remote.tld/user/mysite
```

Create directories for storing our content and a package binary.

```
$ mkdir -p cmd/mysite
$ mkdir -p _contents
```

Copy the example of stylesheet and HTML template from `mdgo` repository,

```
$ cp $HOME/go/src/git.sr.ht/~shulhan/mdgo/_example/index.css ./_contents/
$ cp $HOME/go/src/git.sr.ht/~shulhan/mdgo/_example/html.tmpl ./_contents/
```

Create the main Go code inside `cmd/mysite`,

```
package main

import (
	"git.sr.ht/~shulhan/mdgo"
	"github.com/shuLhan/share/lib/memfs"
)

var mysiteFS *memfs.MemFS

func main() {
	mdgo.Serve(mysiteFS, "./_contents", ":8080", "_contents/html.tmpl")
}
```

Create a new markup file `index.md` inside the "_contents" directory.
Each directory, or sub directory, should have `index.md` to be able to
accessed by browser,

```
=  Test

Hello, world!
```

Now run the `./cmd/mysite` with `DEBUG` environment variable set to non-zero,

```
$ DEBUG=1 go run ./cmd/mysite
```

Any non zero value on `DEBUG` environment signal the running program to watch
changes in ".md" files inside "_contents" directory and serve the generated
HTML directly.

Open the web browser at `localhost:8080` to view the generated HTML.
You should see "Hello, world!" as the main page.

Thats it!

Create or update any ".md" files inside "_contents" directory, the
program will automatically generated the HTML file.
Refresh the web browser to load the new generated file.


###  Deployment

First, we need to make sure that all markup files inside "_contents" are
converted to HTML and embed it into the static Go code.

Create another Go source code, lets save it in `internal/generate.go` with the
following content,

```
package main

import (
	"git.sr.ht/~shulhan/mdgo"
)

func main() {
	opts := &mdgo.GenerateOptions{
		Root:           "./_contents",
		HTMLTemplate:   "_contents/html.tmpl",
		GenPackageName: "main",
		GenVarName:     "mysiteFS",
		GenGoFileName:  "cmd/mysite/static.go",
	}
	mdgo.Generate(opts)
}
```

And then run,

```
$ go run ./internal
```

The above command will generate Go source code `cmd/mysite/static.go` that
embed all files inside the "_contents" directory.

Second, build the web server that serve static contents in `static.go`,

```
$ go build cmd/mysite
```

Third, test the web server by running the program and opening `localhost:8080`
on web browser,

```
$ ./mysite
```

Finally, deploy the program to your server.

*NOTE:* By default, server will listen on address `0.0.0.0` at port `8080`.
If you need to use another port, you can change it at `cmd/mysite/main.go`.

That's it!

##  Limitations and Known Bugs

`mdgo` will not handle automatic certificate (e.g. using LetsEncrypt), its
up to the user how the certificate are gathered, generated, or served.

Using symlink on ".md" file inside `content` directory is not supported yet.


##  Resources

The source code for this software can be viewed at
https://git.sr.ht/~shulhan/mdgo
under custom [BSD license](/LICENSE).


##  Credits

This software use the following third party libraries,

* github.com/yuin/goldmark. MIT License. Copyright (c) 2019 Yusuke Inuzuka
* github.com/yuin/goldmark-meta. MIT License. Copyright (c) 2019 Yusuke Inuzuka
