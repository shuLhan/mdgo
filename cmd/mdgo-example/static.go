// Code generated by github.com/shuLhan/share/lib/memfs DO NOT EDIT.

package main

import (
	"github.com/shuLhan/share/lib/memfs"
)

func generate__example() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example",
		Path:            "/",
		ContentType:     "",
		ContentEncoding: "",
	}
	node.SetMode(2147484141)
	node.SetName("/")
	node.SetSize(0)
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/LICENSE", generate__example_LICENSE))
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/custom.css", generate__example_custom_css))
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/favicon.ico", generate__example_favicon_ico))
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/html.tmpl", generate__example_html_tmpl))
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/index.css", generate__example_index_css))
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/index.html", generate__example_index_html))
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/sub", generate__example_sub))
	node.AddChild(_mdgoFS_getNode(mdgoFS, "_internal", generate__internal))
	return node
}

func generate__example_LICENSE() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/LICENSE",
		Path:            "/LICENSE",
		ContentType:     "text/plain; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x94\x95\x5F\x6F\xDB\x36\x17\xC6\xAF\x5F\x7E\x8A\x07\xBD\x6A\x01\xC5\xFD\xF3\x02\x03\x96\xDE\x8C\x96\xE8\x98\x98\x2C\x6A\x24\x15\xD7\x97\x8A\x45\xD7\xC4\x2C\xD1\xA0\xE8\x14\xFD\xF6\x03\x29\xA7\x71\x9D\x6C\xC3\x08\x04\x39\xB6\xCE\xEF\x9C\xE7\x3C\x3C\x86\x72\x77\xFC\xEE\xED\xD7\x7D\xC0\xA7\x0F\x1F\x7F\xB9\xF9\xF4\xE1\xE3\xAF\x19\x56\x33\xA8\xFD\xE9\xB0\x6F\x07\xBC\xED\xC7\xDF\xFE\xB4\x87\xF6\xC1\x86\x99\x1D\x76\xEE\xDD\x8C\xD0\xC3\x01\x09\x19\xE1\xCD\x68\xFC\xA3\xE9\x66\x84\x48\xD3\xD9\x31\x78\xFB\x70\x0A\xD6\x0D\x68\x87\x0E\xA7\xD1\xC0\x0E\x18\xDD\xC9\x6F\x4D\xFA\xE6\xC1\x0E\xAD\xFF\x8E\x9D\xF3\xFD\x98\xE1\x9B\x0D\x7B\x38\x9F\xFE\xBB\x53\x20\xBD\xEB\xEC\xCE\x6E\xDB\x58\x20\x43\xEB\x0D\x8E\xC6\xF7\x36\x04\xD3\xE1\xE8\xDD\xA3\xED\x4C\x87\xB0\x6F\x03\xC2\xDE\x60\xE7\x0E\x07\xF7\xCD\x0E\x5F\xB1\x75\x43\x67\x23\x34\x26\xA8\x37\xE1\x96\x90\x8F\x33\xFC\x2C\x69\x84\xDB\x3D\x69\xD9\xBA\xCE\xA0\x3F\x8D\x01\xDE\x84\xD6\x0E\xA9\x60\xFB\xE0\x1E\xE3\xA3\x27\x43\x06\x17\xEC\xD6\x64\x08\x7B\x3B\x12\x00\x07\x3B\x86\x58\xE3\xB2\xDD\xD0\x5D\x69\xE9\xEC\xB8\x3D\xB4\xB6\x37\x7E\x46\xC8\xA7\x97\x1A\xEC\x70\x69\xC2\x93\x86\xA3\x77\xDD\x69\x6B\xFE\x49\x46\x54\x10\x95\xFC\x57\x19\x38\x4F\xD7\xB9\xED\xA9\x37\x43\x48\xEE\xC6\x62\xED\xD0\xBD\x77\x1E\x2E\xEC\x8D\x47\xDF\x06\xE3\x6D\x7B\x18\x9F\x8D\x4E\xB7\x93\xC8\x8B\x01\x66\x84\xFC\x7F\x86\xCA\xD8\x44\xC5\xA7\x43\xDB\x9B\x49\xCE\x93\xE0\xBD\x3B\x74\xC6\x63\x70\xCF\x09\xC9\x7B\x1B\xC6\x28\x7A\xAA\xE5\xFC\x88\xBE\xFD\x1E\x85\x3C\x98\xB8\x29\x1D\x82\x83\x19\x3A\xE7\x47\x13\x97\xE2\xE8\x5D\xEF\x82\xC1\x64\x4D\x18\xD1\x19\x6F\x1F\x4D\x87\x9D\x77\xFD\xE4\xC4\xE8\x76\xE1\x5B\xBC\xF1\xA7\x05\x02\x30\x1E\xCD\x36\x2E\x11\x8E\xDE\xC6\xD5\xF2\x71\x7D\x86\x69\x91\xC6\x71\x9A\x40\x2F\xB9\x82\x12\x0B\xBD\xA6\x92\x81\x2B\xD4\x52\xDC\xF3\x82\x15\x98\x6F\xA0\x97\x0C\xB9\xA8\x37\x92\xDF\x2D\x35\x96\xA2\x2C\x98\x54\xA0\x55\x81\x5C\x54\x5A\xF2\x79\xA3\x85\x54\x78\x43\x15\xB8\x7A\x43\xE2\x03\x5A\x6D\xC0\xBE\xD4\x92\x29\x05\x21\xC1\x57\x75\xC9\x59\x81\x35\x95\x92\x56\x9A\x33\x95\x81\x57\x79\xD9\x14\xBC\xBA\xCB\x30\x6F\x34\x2A\xA1\x51\xF2\x15\xD7\xAC\x80\x16\x59\x6C\x4A\x5E\x62\x10\x0B\xAC\x98\xCC\x97\xB4\xD2\x74\xCE\x4B\xAE\x37\x49\xC8\x82\xEB\x2A\xF6\x5A\x08\x09\x8A\x9A\x4A\xCD\xF3\xA6\xA4\x12\x75\x23\x6B\xA1\x18\xA8\x64\xA4\xE0\x2A\x2F\x29\x5F\xB1\x62\x06\x5E\xA1\x12\x60\xF7\xAC\xD2\x50\x4B\x5A\x96\x7F\x33\xA5\x90\x3F\x0F\x39\x67\x28\x39\x9D\x97\x8C\xA4\x56\xD5\x06\x05\x97\x2C\xD7\x71\x9E\xE7\x28\xE7\x05\xAB\x34\x2D\x33\xA8\x9A\xE5\x3C\x06\xEC\x0B\x5B\xD5\x25\x95\x9B\xEC\x5C\x53\xB1\x3F\x1A\x56\x69\x4E\x4B\x52\xD0\x15\xBD\x63\x0A\x6F\xFF\xC5\x93\x5A\x8A\xBC\x91\x6C\x15\x45\x8B\x05\x54\x33\x57\x9A\xEB\x46\x33\xDC\x09\x51\x44\xB1\x44\x31\x79\xCF\x73\xA6\x3E\xA3\x14\x2A\xD9\xD5\x28\x96\xA1\xA0\x9A\xA6\xC6\xB5\x14\x0B\xAE\xD5\xE7\x18\xCF\x1B\xC5\x93\x6B\xBC\xD2\x4C\xCA\xA6\xD6\x5C\x54\xEF\xB0\x14\x6B\x76\xCF\x24\xC9\x69\xA3\x58\x91\xEC\x15\x55\x1A\x55\x2F\x99\x90\x9B\x58\x34\x7A\x90\xDC\xCF\xB0\x5E\x32\xBD\x64\x32\x3A\x9A\x9C\xA2\xD1\x02\xA5\x25\xCF\xF5\x45\x1A\x11\x12\x5A\x48\x7D\x31\x23\x2A\x76\x57\xF2\x3B\x56\xE5\x2C\xAA\x11\xB1\xCA\x9A\x2B\xF6\x0E\x54\x72\x15\x13\xF8\xD4\x76\x4D\x37\x10\x4D\x1A\x39\x5E\x52\xA3\x18\x49\xE1\xC5\xCA\x66\xE9\x2A\xC1\x17\xA0\xC5\x3D\x8F\xB2\xCF\xC9\xB5\x50\x8A\x9F\x17\x25\x59\x96\x2F\x31\xD9\x3D\x23\xF1\xC7\x91\xCE\xCD\xCD\xCD\x73\xF0\xE3\xC3\x6B\xD1\x39\x81\x90\xFF\x69\x0D\xC4\x3F\x70\x0E\xCC\xE7\xF8\x71\x28\xA5\x34\x05\x65\x3A\x53\xC2\xEF\xE7\x93\xB8\x84\x5D\x73\x11\x4A\x60\x44\xF0\xC4\x21\xA2\x11\x9A\x98\x2B\x68\x6A\xF4\x83\x7A\x01\xBD\xDA\x89\x9E\xCF\x95\xC2\x67\xE8\xB5\xB1\x9E\x3B\xBD\x2E\x0F\x13\x35\x41\xD3\xB9\x92\x77\xDD\x89\xAC\xCD\xC3\x68\x83\xB9\xC5\x3E\x84\xE3\xED\xFB\xF7\x97\xAF\x54\x92\xBB\x21\xB4\xDB\x70\x8B\xAB\x77\x2D\xF9\x2B\x00\x00\xFF\xFF\x6A\xEC\xBB\x57\x9D\x07\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("LICENSE")
	node.SetSize(957)
	return node
}

func generate__example_custom_css() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/custom.css",
		Path:            "/custom.css",
		ContentType:     "text/css; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xCA\x30\x54\xA8\xE6\xE2\x4C\xCE\xCF\xC9\x2F\xB2\x52\x28\xCE\x4C\xCD\xCB\x4B\xB4\xE6\xAA\xE5\x02\x04\x00\x00\xFF\xFF\x1C\x35\x0D\x2C\x17\x00\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("custom.css")
	node.SetSize(47)
	return node
}

func generate__example_favicon_ico() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/favicon.ico",
		Path:            "/favicon.ico",
		ContentType:     "image/x-icon",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xEC\x98\x0D\x4C\x55\x65\x18\xC7\x1F\xE0\x1A\x04\x85\x18\x9A\x5D\x4B\xEF\x69\xA6\x32\xE7\xDC\x2C\x95\x73\x67\x26\xAE\x9A\xB9\xE4\x5C\xF5\x9E\x13\x73\xE8\x6E\xE9\x0C\xD4\xCC\x12\xB8\x88\x9A\x62\x69\xCD\xCD\x4F\x30\x44\xCA\x89\xB9\x34\x5D\x4D\x14\x9C\x6B\x6A\xDD\x0B\x9B\x29\x61\x1A\xF3\x13\xF9\xF4\x8A\x22\x2E\x94\xE0\xF2\xE1\x85\xF7\xDF\xDE\x73\x3F\x52\xEE\x65\x92\x0B\x5B\xDB\xFD\x6F\xCF\x38\xCF\x39\xCF\xFB\xFC\xDE\x8F\xE7\x39\x67\x5C\xA2\x00\x0A\x24\x41\xE0\x7F\x05\xFA\x2E\x82\x68\x38\x11\x45\x44\x70\x3F\x84\x12\xFB\x10\xFD\x16\x41\x14\x45\x44\x02\x11\xC5\x90\x33\xEE\x61\x92\x25\x11\xDC\x58\x8D\xB9\x57\xCD\xCD\xF1\xF3\xFD\x7C\x3F\xDF\xCF\xF7\xF3\xFF\x9F\xFC\x32\xEB\x7C\x9C\xCC\x9B\xAD\x5A\xD5\xC9\xC4\xC7\xCE\xBF\x51\x73\x06\x8E\x0E\xA8\x76\xAD\xEC\xF8\x63\xE3\xB7\x95\x27\xE1\x46\xC9\x22\x34\xDE\xAD\xF7\xF0\x6F\xD5\x5E\xC2\xF5\xE2\x85\xEA\xB3\xDE\xE4\x3B\x2A\x53\xB0\x3F\xE7\x5D\x6C\xCB\x5C\x89\x6D\xBB\x8E\x22\x67\xAF\x15\x5F\xED\x2B\x44\xE6\xD7\x79\xC8\xCA\x5C\x8E\xBC\x5D\xA6\x5E\xE5\xB7\x95\xA7\xC1\xBC\xD1\x82\xE4\x0D\x45\x3E\x6D\xE5\xE6\xA3\xFF\x2E\xFF\xDA\x0A\xB0\x86\x43\x40\xEB\x15\xA0\xDD\x06\xD8\xCF\x22\x77\xFF\x11\x95\x65\x9C\xBB\x0E\xAF\xE8\xA7\x62\xDC\xAB\xD3\x30\x6B\xE1\x16\xF5\x5E\x5E\xC1\x61\xE0\xCF\x22\xA0\xE9\x34\x58\xC3\x41\x30\xDB\xEA\x47\xE7\x5F\x5F\x0B\xDC\xAB\x43\x57\x31\x00\x25\xC5\x27\xA0\x13\x86\x7A\x6C\x44\xD4\x48\x54\x54\x5E\xF2\x8A\x45\x67\x33\x58\x5D\xF6\x23\xF1\xD1\x56\xE9\x9D\xEF\x3E\xE5\x6E\x37\x63\xDA\x94\xB1\x98\xF1\xF6\x38\xE4\xFF\x90\xD3\x7D\x60\x67\x2B\xD8\xF5\x75\x3D\xE1\x37\xF1\xFB\x6D\xE5\xC9\x60\x75\x5F\x7A\xA5\x69\x69\x69\x81\xE9\xBD\x79\x88\x1A\x39\x1A\x2B\x3E\x59\x0D\xE6\xB8\x03\x56\x93\x0A\x76\x73\xAB\x57\xEC\xF9\xF3\x17\x10\x33\xF9\x0D\x8C\x79\x79\x3C\xF2\x0B\x8E\x00\x8D\x56\x57\xFD\x24\xBB\xF9\x4D\x3E\xF8\x65\xFC\xD9\xED\x73\x4B\xC0\x1A\xF2\xBD\x72\x6E\xCF\xCE\x41\xBF\xC8\x81\x1E\xB3\x16\x16\x81\xD5\x7E\xEE\xAC\x8F\x2E\x9A\x3E\x43\xF6\xC4\xE9\x5E\x1C\x86\x0E\x7B\x85\xCA\xAF\x3F\xB7\x44\xE5\x1B\x25\xF1\x4A\x57\xBE\x62\x88\x3E\xC8\x9F\x15\x1F\x99\x07\x76\xA7\xC0\x2B\x67\x56\x56\xF6\x03\x7C\x8B\xC5\x0A\x56\xFB\x85\x4F\xBE\x61\xBA\xD1\x13\x37\x44\x78\x09\x8E\xE6\x72\x95\x7F\x3A\x7F\xAE\x7B\xFD\x87\x1E\x5C\xBB\x5E\x89\x9B\xA9\x6F\xDE\xB2\x46\xC2\xB5\xD3\x1F\x80\xD5\x6D\xF7\xCA\x69\xB7\xDB\x11\x3F\xDB\x84\x61\xC3\x47\xC2\x9C\x9A\x06\xE6\x68\x74\xED\x7F\xA6\x57\xEC\xEF\xA5\xA5\x98\x30\x31\x06\xA3\x46\x8F\x41\x5E\x1E\xEF\x89\x42\x95\x7F\xD9\x92\x88\x65\x8B\xDF\xE4\xFC\x36\xD9\xA0\x4F\x50\xD9\xD3\xF4\xAF\xC7\x4D\xD7\xB3\xD2\x63\xEF\xE3\x8F\xD2\x8F\xD0\x51\x65\x46\x67\x35\xAF\xBF\x6A\x1F\xC5\xF4\xB7\xF8\x19\x79\x6A\x95\xF7\x67\x77\xEA\x6C\x73\x9E\x93\x2B\xB6\xF5\x6A\x12\xBE\xD9\x2A\xAB\xFB\xA0\x18\xC4\x38\x45\x12\x8F\x7D\x6A\x9E\x8A\xA6\xCB\x49\xD8\xB9\xD1\x88\x9D\x9B\x66\xA2\xA5\x2C\x09\x9D\x36\xDE\x7F\xF5\xBE\x73\x36\xFF\xEA\x5C\xBB\xBB\xAF\x6C\x6B\x7C\xCF\xB7\xB3\x05\xEC\x56\x0E\x58\xB5\x19\x17\x7F\x4A\x40\xF6\xFA\x19\x38\xB4\x6B\x16\xEC\x57\x92\x9C\x7C\x49\x3C\x27\x1B\xC4\xEA\x45\x73\x63\x70\xAF\x32\x19\xED\x15\x29\x9E\x73\xCA\xF8\x4C\x82\xFD\x6A\x9A\xB3\x16\x5A\x2B\x80\xF6\x5A\xC0\x5E\x0A\x56\x9F\xDB\xCD\xBB\x2D\x15\xEC\xF6\x1E\xA0\xE9\x17\xA0\xB9\xC4\xB9\x3F\xB6\x74\x95\xB5\x7B\xAB\x8C\x5B\x67\x3F\x54\xD7\xCE\xF7\xF8\x6A\xD1\x02\x77\x1D\xD8\x15\x29\xDA\xCA\xAF\x57\x2D\x9D\x82\xC3\xB9\xF1\xD8\xB0\x2A\xD6\xD3\xA3\x89\xA6\x49\xB0\x7C\x6F\x42\x3B\xEF\xC9\x7F\xF8\x9D\xE7\xBD\xC6\xC7\x26\x98\x26\xA9\xB9\xE6\x28\x13\xB0\x27\x43\xC1\x81\x1D\x71\x98\x3F\xE7\x35\x67\x1F\xC4\x8A\x17\xF9\xF9\xCB\x92\xD8\xE1\x66\x7A\xCC\x20\xDA\xDC\xD7\xA6\xB8\x09\xC8\x58\x6B\x40\xC1\xEE\x78\x9C\x3F\x91\x80\x9B\x67\x16\xA3\xF1\xE2\xC7\xEA\x9E\x71\xE3\xD7\xFC\x1E\x7F\xC6\x63\x78\x2C\x1F\xE3\x1E\xAF\x48\xFA\x72\xAF\xFC\x92\xC8\xF8\xF9\xAB\x35\x68\x18\x3F\x51\x36\xE8\x7F\x94\x25\xB1\x52\x91\xC4\x62\x45\x12\x17\x10\x51\xA0\x31\x56\x3F\xCB\x28\x89\x55\x3E\xC6\xF6\xC8\x8C\x92\x58\x6E\x8C\x15\xDF\xE1\xFF\x02\xCB\x92\x18\x2F\x4B\x62\x91\x2B\x9F\x45\x8E\x15\xA7\x76\x7D\x07\xF8\x92\xA2\x28\x41\x33\xA5\xE8\xC9\xB2\x41\xCC\x90\x0D\xE2\x71\x3E\x47\x59\x12\x1B\x5C\xE6\x66\x35\xBA\x7C\x9B\x33\x26\x7A\x33\x5F\x53\x7A\x3A\x05\xF6\x84\xF1\x5F\xCA\x42\x14\xD4\x5B\x96\x4E\x14\xE0\x8B\xC9\xEF\x83\x48\x6D\xD1\x28\x22\x8A\x70\xFD\x56\x11\x40\x21\xDD\xCE\xB3\x5F\xE4\x40\xD2\x09\x43\x69\xF4\xA8\x11\x14\x27\x8D\xE1\xDF\x0B\x4A\x99\x1F\x4D\xCB\x12\xC7\xD2\x8A\x24\x89\x0E\x64\xBC\x45\xDF\x66\xC8\x94\xBC\xA1\x88\x4E\xED\x93\x68\xD3\x8E\xBD\xB4\x77\xCF\x7A\x3A\x95\xBF\x94\x0A\x0F\x2F\xA7\xB3\xC7\xD2\xA8\xF8\xE7\xCD\x54\x62\xCD\xA2\xA6\x0B\xC9\x84\x8A\x54\x42\x15\xB7\x65\x84\x9A\x14\x42\x75\x1A\xB1\x1A\x33\xDD\xAC\x29\x25\x7B\x8B\x83\x5A\xEF\x81\x1C\x1D\x20\x07\x03\x9F\x68\x6F\x1E\x81\x5F\x7E\xF9\xF5\x10\x09\x1A\xED\x7D\xD2\x08\xBD\xE4\x47\xF4\xD5\x0E\xE8\xDF\xDF\xED\x3F\x17\x3A\x78\x70\xA8\xD6\xED\x0F\xD0\x6A\x9F\x0E\x0D\x7D\x4A\x3B\xF0\x59\xA7\x1F\xA9\xD3\x85\x07\x04\x84\xEB\x74\x91\x4E\x7F\xD0\x10\x0A\xD4\x6A\x07\x0D\xA6\x40\xAD\x46\xD0\x3C\x11\x36\x48\xA7\xFA\x3A\x0A\x0C\xEB\xA3\xD1\x84\x85\x84\x3F\xFF\x42\xB8\x56\x1B\xAE\xD3\x85\x87\x84\x69\x04\x4D\x70\x48\x58\x18\x47\x3E\xD3\xEF\xC9\x60\x8D\xC0\x15\x1C\x1C\xA4\xD1\x04\x05\x07\x0B\x3E\x94\x1E\x40\xD4\x53\x73\x8B\x5F\xDF\xED\xCB\x5F\xB4\x44\x7F\x05\x00\x00\xFF\xFF\xF7\x87\x31\xC7\x36\x16\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("favicon.ico")
	node.SetSize(1316)
	return node
}

func generate__example_html_tmpl() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/html.tmpl",
		Path:            "/html.tmpl",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x8C\x53\x4D\x8F\xDB\x2C\x10\x3E\x67\x7F\x05\xCB\x7B\x7D\x31\xAA\xAA\xAA\x55\x8B\x23\x75\xB7\x7B\xEE\x4A\xC9\xA5\xA7\x8A\xC0\x24\xA0\x62\xF0\xC2\x38\x59\xCB\x72\x7F\x7B\x85\x71\x52\x37\x5D\x69\x7B\x32\xF0\xCC\xF3\x31\x83\x11\xB7\x5F\xBE\xDE\x6F\xBF\x3D\x3E\x10\x83\x8D\x5B\xDF\x88\xF2\x59\x09\x03\x52\xAF\x6F\x56\x2B\xD1\x00\x4A\x62\x10\x5B\x06\x4F\x9D\x3D\xD6\xF4\x3E\x78\x04\x8F\x6C\xDB\xB7\x40\x89\x2A\xBB\x9A\x22\x3C\x23\xCF\xEC\x4F\x44\x19\x19\x13\x60\xDD\xE1\x9E\x7D\xA0\x84\xFF\xD6\xF1\xB2\x81\x9A\x1E\x2D\x9C\xDA\x10\x71\xC1\x3E\x59\x8D\xA6\xD6\x70\xB4\x0A\xD8\xB4\xF9\x9F\x58\x6F\xD1\x4A\xC7\x92\x92\x0E\xEA\x37\x2F\x08\xA1\x81\x06\x98\x0A\x2E\xC4\x85\xD6\x7F\x6F\xDF\xBF\x7B\xF8\x7C\x37\xD5\x67\x02\x5A\x74\xB0\x1E\x86\x6A\x9B\x17\xE3\x28\x78\x39\xC9\x98\xB3\xFE\x07\x89\xE0\x6A\x9A\xB0\x77\x90\x0C\x00\x52\x62\x22\xEC\x6B\xCA\xAD\xD7\xF0\x5C\xA9\x94\x66\xEB\x61\x60\x24\x4A\x7F\x00\x52\x6D\xA6\xEA\x71\x7C\x4D\x63\x18\xAA\x71\x5C\xD0\xC1\xEB\x4C\x12\x7C\x9E\xAF\xD8\x05\xDD\x4F\x49\xB4\x3D\x12\xE5\x64\x4A\x35\xC5\xD0\xEE\x64\xA4\xF9\xF8\x8F\xF3\xDC\xA1\xB4\x1E\x66\xE8\x9A\xC3\xB2\xA6\xF5\x87\x19\x5D\x09\x79\x6E\x84\xAE\x95\xB5\x87\x20\xB8\x9C\x89\x5C\xDB\xE3\xDF\x1A\x0D\xF8\xEE\x42\xDE\x87\xD8\x9C\x01\x8B\xD0\x50\x22\x15\xDA\xE0\x6B\xCA\xBF\x5B\x8F\x10\xBD\x74\x3C\x81\x8C\xCA\x9C\x39\x2B\x61\x7D\xDB\x21\xC1\xBE\x85\xF2\x43\xD0\xF9\xA2\x9E\x28\x69\x9D\x54\x60\x82\xD3\x10\x6B\xBA\x29\xBC\x32\x97\x29\x50\xB6\xFB\xF7\x70\x97\xCE\x52\xB7\xA3\xEB\x4D\xB7\x7B\xA1\xB5\xCB\x6A\x5E\x5C\x0D\xB9\x95\x07\x78\x7D\xC4\xC3\x50\xDD\x05\xDD\x4F\x17\xBD\xD4\xBE\x65\x8C\x54\x97\x62\xC2\xD8\xC2\xE8\x0C\x67\x87\x09\xB9\x72\xDE\x87\x80\x67\x83\xC7\x70\x82\x08\x9A\xEC\x7A\x22\xE4\xE4\x58\x1A\xCB\x0F\x2E\x7D\xE4\x3C\xC5\xCA\x20\xFF\x99\x4C\xE7\x8C\xF4\x7C\xBA\x46\x9A\xEB\x4A\xBC\x69\x5F\xA2\xC9\x65\x02\xC1\xCB\x7F\x25\x78\x79\xCF\xBF\x02\x00\x00\xFF\xFF\x24\xFC\xFC\x47\xE7\x03\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("html.tmpl")
	node.SetSize(482)
	return node
}

func generate__example_index_css() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/index.css",
		Path:            "/index.css",
		ContentType:     "text/css; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x94\x55\x5D\x8F\xAB\x36\x14\x7C\x8E\x7F\xC5\x91\x56\x7D\x89\x80\x05\xB2\x24\x2D\x48\x2B\x55\x7D\xEE\x8F\x30\xF8\x10\xAC\x35\x36\x32\xA6\xC9\xF6\x2A\xFF\xFD\xCA\x60\xBE\x02\xB9\xD2\x15\x0F\x89\xEC\xB1\xCF\x78\xE6\x78\x9C\x2B\xF6\x0D\x3F\xC8\xA1\xA6\xFA\xCA\x65\x0A\x61\x46\x0E\xA5\x92\xC6\x2F\x69\xCD\xC5\x77\x0A\x7F\x6B\x4E\x85\x07\x2D\x95\xAD\xDF\xA2\xE6\x65\x46\x0E\x39\x2D\xBE\xAE\x5A\x75\x92\xF9\x85\x12\x4A\xA7\xF0\x56\x96\x76\x42\x70\x89\x7E\x85\xFC\x5A\x99\x14\xA2\xE0\x94\x91\x83\xC1\xBB\xF1\xA9\xE0\x57\x99\x42\x81\xD2\xA0\xCE\xC8\x61\x5C\x15\xC7\x71\x46\x1E\xA4\xD1\x68\x39\xAC\xEA\xFE\x8B\x52\x28\x0F\x6A\x25\x55\xDB\xD0\x02\x47\x5E\x2D\xFF\x1F\x53\x08\x83\x3F\x2F\x89\xC6\x7A\x5B\xF3\x23\x23\x07\xF5\x1F\xEA\x52\xA8\x9B\x7F\x4F\x81\x76\x46\xAD\x28\xA7\xF0\x86\xA5\xFD\x32\x72\x68\x28\x63\x5C\x5E\xED\x7E\xE7\xD8\xED\x97\x2B\xCD\x50\xFB\x9A\x32\xDE\xB5\x76\xE6\x14\xB9\xA9\x07\xA1\x96\xE6\x48\xFE\x74\x49\x90\xE6\xE3\x11\x19\x16\x4A\x53\xC3\x95\x4C\x41\x2A\x89\x3D\x3C\xAD\x2C\x15\xBB\x68\x83\xE9\x24\x43\x6D\xB9\x5B\x20\x69\x3C\x22\xF8\xE0\xC3\xDD\xBF\x71\x66\xAA\x14\x92\x70\x20\x74\x53\x9A\xF9\x37\x4D\x9B\x14\x72\x8D\xF4\xCB\xB7\x03\xBD\x6C\x9E\x55\xCE\x23\x9D\xF0\x88\x12\x4B\x17\xA3\x60\xA2\x4C\xAA\xC8\x23\x55\xEC\x91\xEA\xE4\x91\xEA\x63\x07\x05\xE1\x8C\x5F\x28\x92\xED\x9C\xB4\x77\xE0\xE6\xB4\xCE\x95\xE8\x69\x54\xD1\x64\xDE\x60\x4E\x14\xEC\x7B\xE3\xC0\x41\xAF\x45\xDD\x19\x64\x4B\x39\x2F\x97\x4B\x0F\x88\xFB\x41\x81\x54\xA7\xA0\xED\xCA\xEC\x69\xF3\xC9\xA8\x95\xA5\x21\xE6\x65\xB2\xB6\x74\xBF\x41\xE2\xE4\xF9\x20\x52\xE9\x9A\x8A\x45\xDF\x4C\x5D\x33\x35\xD2\x9E\xFC\x55\x0C\x74\x3A\xF9\x46\x94\xD3\x46\x94\xF8\x17\x7C\x7E\xB3\xF2\x6B\x2B\x97\xDD\x3A\x40\x96\x1C\xA6\x9E\x88\xE1\x13\xDA\x86\x4A\xDB\x16\xEE\x6F\x8F\x15\x8A\x9A\x59\xF6\x29\x14\x20\x4E\x9A\x3B\x84\x73\x3A\x6C\xA4\x1B\x5D\x4C\xE2\xCB\x5F\x45\x6F\x24\x61\xAF\x5A\x92\xB1\x55\xE2\xF4\xDF\x72\x5A\x78\x0E\xB2\x7B\xE1\x1F\x84\xBC\x1F\x8F\x04\x8E\xF0\x4F\xD7\x1A\x55\x43\x21\x68\xDB\x62\x0B\xA5\xD2\xD0\xD0\x2B\xB6\x04\x8E\xEF\x84\x04\x46\x35\x39\xED\x2F\xDF\x7E\xA7\x8C\x16\x7C\x0C\xB6\xCC\x16\x54\x9C\x31\x94\x7D\xA5\x71\x13\xFB\xEB\x57\x48\x6D\x6B\x79\xF3\x68\x8D\xB2\xB3\x05\xA6\xA6\x8B\x82\x53\x74\xEA\x8D\x78\xD5\xB5\xBB\xF2\xBD\xA8\xB4\xB0\x44\x60\x69\x2C\x6E\x17\xB6\x8A\xA5\x3E\x53\x5F\x67\xD2\x96\xFC\xDA\xF4\xC7\x13\x80\xAE\xBC\x72\x31\x09\x61\x10\x6D\x13\x63\xCE\x50\xC7\xE5\x56\x71\x83\xCF\x37\x75\x0A\x93\x21\x68\xED\xBA\xD0\xED\xDA\x2A\xC1\xD9\x06\x31\x45\x71\xD2\xDC\xB7\xFC\x4A\xA5\xEB\x99\xA2\x6F\x75\x5A\x52\x79\x90\xC0\xF6\x84\x45\xB8\x58\x8D\xC2\xF0\x8F\x41\x88\x7E\xE2\x13\x82\x42\x49\x43\xB9\x44\x3D\x3B\xBB\x1C\x9D\xE2\xDB\xBD\x62\x83\x15\xEB\x82\xEE\xD6\xBA\x31\x3D\xB8\xEB\x06\x67\x81\x56\xB9\xBC\x28\x10\xD4\x68\xE6\x28\x69\xCD\xB7\xC0\x14\xB8\xA1\x82\x17\xD9\x7E\xA8\x07\xA5\x52\x66\x4B\x6D\xF3\xC0\x9E\xCF\xE7\x97\x0F\xE7\xB8\x6F\x1C\x24\xAE\x63\x1F\xE4\x67\x00\x00\x00\xFF\xFF\x2F\x17\x7F\x31\x0F\x08\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.css")
	node.SetSize(739)
	return node
}

func generate__example_index_html() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/index.html",
		Path:            "/index.html",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xA4\x56\x5F\x8F\xDB\x36\x0C\x7F\x4E\x3E\x05\xA7\x01\x7B\xD9\x39\xCE\xA5\xEB\xBF\xAB\xED\xFD\xB9\x76\xDD\x80\x0E\x2B\x70\x05\x86\x3D\x1D\x68\x89\xB1\x84\x93\x25\x57\xA2\x93\xE6\x65\x9F\x7D\x90\xED\xA4\x49\x9A\x6D\xC5\xFA\x64\x91\x22\x29\xF1\xA7\x1F\x49\x17\x5F\xBD\xFC\xFD\xF6\xDD\x9F\x6F\x5F\x81\xE6\xD6\x56\xF3\x62\xFC\xCC\x0A\x4D\xA8\xAA\xF9\x6C\x56\xB4\xC4\x08\x9A\xB9\xCB\xE8\x7D\x6F\x36\xA5\xB8\xF5\x8E\xC9\x71\xF6\x6E\xD7\x91\x00\x39\x4A\xA5\x60\xFA\xC0\x79\xF2\x7E\x01\x52\x63\x88\xC4\x65\xCF\xEB\xEC\x99\x80\xFC\x63\x1C\x87\x2D\x95\x62\x63\x68\xDB\xF9\xC0\x47\xDE\x5B\xA3\x58\x97\x8A\x36\x46\x52\x36\x08\x57\x60\x9C\x61\x83\x36\x8B\x12\x2D\x95\xD7\x17\x02\xB1\xA6\x96\x32\xE9\xAD\x0F\x47\xB1\xBE\x7E\xF4\xF4\xF1\xAB\x1F\x7F\x1A\xEC\x93\x03\x1B\xB6\x54\xFD\x41\x56\xFA\x96\x80\x3D\x48\x63\x1A\x5F\xE4\xA3\x3E\x59\x58\xE3\x1E\x20\x90\x2D\x45\xE4\x9D\xA5\xA8\x89\x58\x80\x0E\xB4\x2E\x45\x6E\x9C\xA2\x0F\x0B\x19\xE3\x78\x81\x22\x9F\xA0\x29\x6A\xAF\x76\x83\xBB\x32\x1B\x90\x16\x63\x2C\x05\xFB\xAE\xC6\x20\x92\xFA\x44\x9F\x2E\x87\xC6\xD1\xB4\x75\xEE\x93\xA5\x98\xC6\x35\xD3\xEE\xAC\xC0\xFD\xE9\xA2\x9A\x6E\x8B\x93\x63\xAE\xCC\xE6\xD3\x18\x2D\xB9\xFE\xE0\xBC\xF6\xA1\xDD\x6F\x18\xA6\x56\x00\x4A\x36\xDE\x95\x22\xBF\x37\x8E\x29\x38\xB4\x79\x24\x0C\x52\xEF\x7D\x66\x85\x71\x5D\xCF\xC0\xBB\x8E\xC6\xB7\x14\x13\xC6\xEF\x05\x74\x16\x25\x69\x6F\x15\x85\x52\xDC\x8D\x7E\xE3\x63\x0C\x17\x4A\xC7\x7D\xFE\xE5\x0E\x99\xC5\xBE\x16\xD5\x5D\x5F\x5F\x48\xED\xB0\x9A\x16\x67\x20\x77\xD8\xD0\x7F\x43\x3C\x1F\x36\x8D\x2A\x45\x02\x77\x50\x17\xFA\xFA\x02\x0F\xF4\x75\x35\x3F\x0E\xA4\x88\xD1\xD8\x98\xEC\x63\x87\x6E\x08\x81\x3D\xEB\x81\x64\xA3\xC9\x24\x56\x77\xBA\xB7\x1A\x5D\x91\x27\xC3\xAA\xA8\xC3\xB1\x0F\xB5\x68\xEC\xC1\x65\x94\xAA\x43\xFE\x49\x64\x7F\xD3\xC6\x1F\x1E\x8C\xC5\xDA\xF0\xC2\xB8\xB5\x17\xD5\x99\x22\xA1\x73\x39\x7C\xA0\x8D\x42\x26\x51\xAD\x1E\xC3\x1D\x75\x4C\x6D\x4D\x01\x56\xCB\xEB\xE7\x93\xFD\x7C\x42\x6F\xFF\xD9\xC3\x31\x15\x8A\x38\x52\x75\x81\xB0\xAD\x2D\x89\x53\x24\x22\x0D\xBC\x49\x44\x3F\xDB\xE9\x30\x60\x13\xB0\x4B\x04\x2A\xBA\x6A\xBC\xD5\x9E\x72\xED\xF0\x3E\x05\x9E\x2A\xA6\xBC\x53\x2B\x89\x37\x79\xDE\x3D\x34\x8B\xC6\x2F\x14\x6D\xF2\xC6\xF0\x22\x86\x85\xE6\xFC\xAF\x38\x02\x9A\x0F\x4F\x23\xAA\xC2\xB4\x0D\xC4\x20\x3F\xBA\x99\xB6\x59\x44\x6D\xC8\xAA\xB8\x30\x3E\xAF\x51\x35\x94\x8F\x71\xB2\x40\x6B\x0A\xE4\x24\x65\xCB\xE5\x53\xF5\x5C\x7E\x6F\x7D\xE3\xCB\xC6\x7F\x83\x6D\xF7\x22\xAD\x6F\x53\xA7\x28\xB7\xDA\x30\x0D\xBA\xA1\xDA\xCB\xB5\x45\xCE\xE2\xFB\x1E\x03\x09\x40\xCB\xA5\x78\xED\x5F\x7A\x29\xAA\x63\xEC\xF3\xEE\x14\xC8\x7F\xC0\x41\x7A\x45\xFB\x8A\x1D\xD6\x60\x22\x20\x58\x53\x07\x0C\x3B\x40\xA7\x00\xA1\x0B\xBE\x09\xD8\x26\x0A\x6E\x83\x61\x82\xC8\xC8\x46\xC2\x96\x6A\x88\x14\x36\x14\x60\x6B\x58\x43\x7A\x51\xA5\x48\xCD\xD7\xC6\x52\x84\x3E\x1A\xD7\xCC\x0F\x04\xDA\x23\x82\x51\x1A\xA3\xBC\x64\x1F\x16\x3E\x34\xB9\xF2\x32\xE6\x5B\x8D\x9C\x99\x98\xED\x37\x73\x51\xED\x97\x43\xBD\xB5\x18\x1E\xFA\x0E\x52\xE9\x22\x2F\x4E\xB2\xFB\x94\x34\x47\x6C\xB8\x1E\xAA\x68\x35\x90\xE6\x7E\xC8\xF3\x1E\xE3\xFD\x94\xDE\xD4\xAB\x00\xE3\x3E\xE1\x22\xD7\xAB\xFF\x49\xA9\x9F\x7D\x00\x74\xD0\x77\x09\xA6\x44\x74\x50\x5E\xF6\x2D\xB9\x84\x95\x77\xE0\xD7\xA0\xFD\x36\x6D\xF6\x91\x80\x35\x1D\x40\x8E\x34\xC8\x9F\x22\xF5\x59\x94\x7B\xED\xCF\x0E\x4A\xDD\x26\x61\xF6\xC5\x28\x49\x6B\x8E\x10\xBA\x7D\xF3\xEB\x17\xA0\x73\x81\x68\x63\x4C\x90\xA9\x10\xBD\xDB\x50\xE0\x2B\x68\xC8\x51\x40\xA6\xAB\xC4\xBC\xDC\x87\x91\x5E\x80\xA0\x4C\xA0\xC4\x98\x1D\xB0\x46\x86\xA9\x77\xC6\x03\x2F\x12\xE1\xAE\x52\xC8\x5F\xDE\xFD\xF6\x66\x14\xFF\x2D\xFF\xB3\x0E\xB3\xF6\x9E\xC7\x86\x7B\xAA\xC9\x86\xC1\x52\xCD\xDF\x60\x64\xE8\xBB\xF4\xAA\x0A\x56\xCB\xD5\x32\xBB\x5E\x65\xCB\x27\xB0\x7C\x74\xF3\xF8\xC9\xCD\x77\xCF\xE0\xDB\xE5\xD3\xE5\xF2\x2C\xFA\xC9\x84\x98\x1F\x09\xE7\xE3\xE1\x70\xFA\x6C\x36\x7B\xEB\xB7\x14\x48\x41\xBD\x83\x02\x87\xB1\x70\x4A\x89\x8B\x1C\x48\x76\xE3\x0C\x19\xE4\xF1\x64\x3C\x9A\x47\xB3\x22\x1F\x87\x7F\x91\x8F\xFF\x4B\x7F\x07\x00\x00\xFF\xFF\x2F\xF3\xA1\x50\x47\x09\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.html")
	node.SetSize(997)
	return node
}

func generate__example_sub() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/sub",
		Path:            "/sub",
		ContentType:     "",
		ContentEncoding: "",
	}
	node.SetMode(2147484141)
	node.SetName("sub")
	node.SetSize(0)
	node.AddChild(_mdgoFS_getNode(mdgoFS, "/sub/index.html", generate__example_sub_index_html))
	return node
}

func generate__example_sub_index_html() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/sub/index.html",
		Path:            "/sub/index.html",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x8C\x54\x4F\x73\xFB\x34\x10\x3D\x3B\x9F\x42\x88\x23\x38\x4A\xD2\xBF\x13\x64\x0F\xA5\xF4\xC6\x40\x67\xDA\x0B\x27\x66\x6D\x6F\x22\x0D\xB2\xA4\x4A\xEB\xB4\xB9\xF0\xD9\x19\xD9\x4E\xEA\xA4\x19\xF8\x9D\xEC\xFD\xF3\x76\xA5\xB7\x6F\x25\xBF\xFB\xF5\x8F\xC7\xD7\x3F\x9F\x9F\x98\xA2\xD6\x94\x33\x39\x7C\x32\xA9\x10\x9A\x72\x96\x65\xB2\x45\x02\xA6\x88\x7C\x8E\x6F\x9D\xDE\x15\xFC\xD1\x59\x42\x4B\xF9\xEB\xDE\x23\x67\xF5\x60\x15\x9C\xF0\x83\x44\x42\xFF\xC4\x6A\x05\x21\x22\x15\x1D\x6D\xF2\x7B\xCE\xC4\x67\x1D\x0B\x2D\x16\x7C\xA7\xF1\xDD\xBB\x40\x13\xF4\xBB\x6E\x48\x15\x0D\xEE\x74\x8D\x79\x6F\xFC\xC8\xB4\xD5\xA4\xC1\xE4\xB1\x06\x83\xC5\xF2\x42\x21\x52\xD8\x62\x5E\x3B\xE3\xC2\xA4\xD6\xF7\x57\x77\x37\x4F\x0F\xBF\xF4\xF9\x09\x40\x9A\x0C\x96\x2F\x5D\xC5\x1A\x1D\xB0\x26\x17\xF6\x52\x0C\xCE\x14\x36\xDA\xFE\xCD\x02\x9A\x82\x47\xDA\x1B\x8C\x0A\x91\x38\x53\x01\x37\x05\x17\xDA\x36\xF8\x31\xAF\x63\x3C\x74\xFF\xCF\xEC\xBA\x8B\xE4\xDA\xCF\x74\x29\x46\x1A\x65\xE5\x9A\x7D\x8F\x6F\xF4\x8E\xD5\x06\x62\x2C\x38\x39\x5F\x41\xE0\xC9\x7D\xE2\x4F\x17\x01\x6D\x71\x0C\x9D\x63\xF2\x54\x53\xDB\xED\x18\xCD\x24\x1C\xDA\xF3\xB2\xD6\x7A\xEB\xA4\x80\x11\x28\x1A\xBD\xFB\x5A\xA3\x45\xDB\x1D\xC1\x1B\x17\xDA\x43\x40\x13\xB6\x9C\x41\x4D\xDA\xD9\x82\x8B\xBF\xB4\x25\x0C\x16\x8C\x88\x08\xA1\x56\x07\x4C\x26\xB5\xF5\x1D\x31\xDA\x7B\x1C\xE6\xCE\xC7\x79\xBC\x71\xE6\x0D\xD4\xA8\x9C\x69\x30\x14\xFC\x65\xC0\x0D\xD4\xF5\x07\x4A\xED\xBE\xFD\x70\xC7\x9B\xC5\xAE\xE2\x69\x82\x17\xAE\x76\xFC\x1B\x7F\xCE\x48\xF6\xB0\xC5\xFF\xA7\x78\xD6\x07\x75\x53\xF0\x44\x6E\xEF\x96\x6A\x79\xAE\x19\xB5\x2C\x67\xD3\x2A\x0D\x12\x68\x13\x53\x72\xF4\x60\x7B\x3C\x74\xA4\x7A\x35\x0E\x29\xA3\x59\xBE\x60\xED\x6C\xC3\x1E\x7A\x53\x8A\x94\x5E\xCA\x2A\x4C\x91\xD8\x82\x36\x47\xE0\x60\x95\x47\x0A\x92\x49\x6E\xDD\xCE\xA3\xEA\x8C\x02\xFB\xF3\x36\x79\xE6\xB5\x6B\x79\x79\xC1\x99\x88\xBA\xDC\x26\xE0\xAE\x01\x42\x5E\xAE\xAE\xD9\xEF\x6E\x87\x6D\x85\x81\xAD\x16\xAB\xC5\x98\x3E\x1B\x79\x3C\x7C\x0E\xC4\x8C\xEB\xC5\x27\x2E\x1F\x10\xDA\xCA\x20\x3F\xA5\x25\x62\xAF\xA0\x24\xF9\xB3\x88\x87\x00\xDB\x00\x3E\x49\x49\xFA\xF2\x55\xE9\xC8\x74\x64\x60\x19\x7E\x40\xEB\x0D\x32\xB7\x39\xEC\x31\xD3\x96\xC5\x29\xFF\xAC\x8B\xDA\x6E\xD9\xB0\x63\xEC\x73\xFB\xE6\x52\xF8\x2F\xA7\xBE\x7C\x87\x8D\x73\x34\x0C\xF7\xD4\x93\xF7\x22\x2E\x67\xBF\x41\x24\xD6\xF9\xC4\x4F\xD3\x73\x92\x2F\x57\xF9\xE2\x96\x2D\xAE\xD6\x37\xB7\xEB\xEB\x7B\xF6\xC3\xE2\x6E\xB1\x38\xAB\x7E\xA2\xC6\xD9\xC4\x38\x97\xE2\xB1\x7B\x96\x65\xCF\xEE\x1D\x03\x36\xAC\xDA\x33\x09\xBD\x04\x87\x31\xA7\x87\x36\xAE\x85\x88\x61\xAE\x48\xFC\x33\xCE\x55\xF4\x7B\xCD\x53\xDE\xA0\xD7\xDE\x1E\x3A\xC3\x44\xFB\x99\x14\xC3\x43\x23\xC5\xF0\x8E\xFF\x1B\x00\x00\xFF\xFF\x04\x81\x5B\xCD\xDF\x05\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.html")
	node.SetSize(673)
	return node
}

func generate__internal() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_internal",
		Path:            "_internal",
		ContentType:     "",
		ContentEncoding: "",
	}
	node.SetMode(2147483648)
	node.SetName("_internal")
	node.SetSize(0)
	node.AddChild(_mdgoFS_getNode(mdgoFS, "_internal/.template", generate__internal__template))
	return node
}

func generate__internal__template() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/html.tmpl",
		Path:            "_internal/.template",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x8C\x53\x4D\x8F\xDB\x2C\x10\x3E\x67\x7F\x05\xCB\x7B\x7D\x31\xAA\xAA\xAA\x55\x8B\x23\x75\xB7\x7B\xEE\x4A\xC9\xA5\xA7\x8A\xC0\x24\xA0\x62\xF0\xC2\x38\x59\xCB\x72\x7F\x7B\x85\x71\x52\x37\x5D\x69\x7B\x32\xF0\xCC\xF3\x31\x83\x11\xB7\x5F\xBE\xDE\x6F\xBF\x3D\x3E\x10\x83\x8D\x5B\xDF\x88\xF2\x59\x09\x03\x52\xAF\x6F\x56\x2B\xD1\x00\x4A\x62\x10\x5B\x06\x4F\x9D\x3D\xD6\xF4\x3E\x78\x04\x8F\x6C\xDB\xB7\x40\x89\x2A\xBB\x9A\x22\x3C\x23\xCF\xEC\x4F\x44\x19\x19\x13\x60\xDD\xE1\x9E\x7D\xA0\x84\xFF\xD6\xF1\xB2\x81\x9A\x1E\x2D\x9C\xDA\x10\x71\xC1\x3E\x59\x8D\xA6\xD6\x70\xB4\x0A\xD8\xB4\xF9\x9F\x58\x6F\xD1\x4A\xC7\x92\x92\x0E\xEA\x37\x2F\x08\xA1\x81\x06\x98\x0A\x2E\xC4\x85\xD6\x7F\x6F\xDF\xBF\x7B\xF8\x7C\x37\xD5\x67\x02\x5A\x74\xB0\x1E\x86\x6A\x9B\x17\xE3\x28\x78\x39\xC9\x98\xB3\xFE\x07\x89\xE0\x6A\x9A\xB0\x77\x90\x0C\x00\x52\x62\x22\xEC\x6B\xCA\xAD\xD7\xF0\x5C\xA9\x94\x66\xEB\x61\x60\x24\x4A\x7F\x00\x52\x6D\xA6\xEA\x71\x7C\x4D\x63\x18\xAA\x71\x5C\xD0\xC1\xEB\x4C\x12\x7C\x9E\xAF\xD8\x05\xDD\x4F\x49\xB4\x3D\x12\xE5\x64\x4A\x35\xC5\xD0\xEE\x64\xA4\xF9\xF8\x8F\xF3\xDC\xA1\xB4\x1E\x66\xE8\x9A\xC3\xB2\xA6\xF5\x87\x19\x5D\x09\x79\x6E\x84\xAE\x95\xB5\x87\x20\xB8\x9C\x89\x5C\xDB\xE3\xDF\x1A\x0D\xF8\xEE\x42\xDE\x87\xD8\x9C\x01\x8B\xD0\x50\x22\x15\xDA\xE0\x6B\xCA\xBF\x5B\x8F\x10\xBD\x74\x3C\x81\x8C\xCA\x9C\x39\x2B\x61\x7D\xDB\x21\xC1\xBE\x85\xF2\x43\xD0\xF9\xA2\x9E\x28\x69\x9D\x54\x60\x82\xD3\x10\x6B\xBA\x29\xBC\x32\x97\x29\x50\xB6\xFB\xF7\x70\x97\xCE\x52\xB7\xA3\xEB\x4D\xB7\x7B\xA1\xB5\xCB\x6A\x5E\x5C\x0D\xB9\x95\x07\x78\x7D\xC4\xC3\x50\xDD\x05\xDD\x4F\x17\xBD\xD4\xBE\x65\x8C\x54\x97\x62\xC2\xD8\xC2\xE8\x0C\x67\x87\x09\xB9\x72\xDE\x87\x80\x67\x83\xC7\x70\x82\x08\x9A\xEC\x7A\x22\xE4\xE4\x58\x1A\xCB\x0F\x2E\x7D\xE4\x3C\xC5\xCA\x20\xFF\x99\x4C\xE7\x8C\xF4\x7C\xBA\x46\x9A\xEB\x4A\xBC\x69\x5F\xA2\xC9\x65\x02\xC1\xCB\x7F\x25\x78\x79\xCF\xBF\x02\x00\x00\xFF\xFF\x24\xFC\xFC\x47\xE7\x03\x00\x00"),
	}
	node.SetMode(420)
	node.SetName(".template")
	node.SetSize(482)
	return node
}

//
// _getNode is internal function to minimize duplicate node created on
// Node.AddChild() and on generatedPathNode.Set().
//
func _mdgoFS_getNode(mfs *memfs.MemFS, path string, fn func() *memfs.Node) (node *memfs.Node) {
	node = mfs.PathNodes.Get(path)
	if node != nil {
		return node
	}
	return fn()
}

func init() {
	mdgoFS = &memfs.MemFS{
		PathNodes: memfs.NewPathNode(),
		Opts: &memfs.Options{
			Root:        "_example",
			MaxFileSize: 5242880,
			Includes:    []string{},
			Excludes: []string{
				`.*\.md$`,
				`^\..*`,
			},
		},
	}
	mdgoFS.PathNodes.Set("/",
		_mdgoFS_getNode(mdgoFS, "/", generate__example))
	mdgoFS.PathNodes.Set("/LICENSE",
		_mdgoFS_getNode(mdgoFS, "/LICENSE", generate__example_LICENSE))
	mdgoFS.PathNodes.Set("/custom.css",
		_mdgoFS_getNode(mdgoFS, "/custom.css", generate__example_custom_css))
	mdgoFS.PathNodes.Set("/favicon.ico",
		_mdgoFS_getNode(mdgoFS, "/favicon.ico", generate__example_favicon_ico))
	mdgoFS.PathNodes.Set("/html.tmpl",
		_mdgoFS_getNode(mdgoFS, "/html.tmpl", generate__example_html_tmpl))
	mdgoFS.PathNodes.Set("/index.css",
		_mdgoFS_getNode(mdgoFS, "/index.css", generate__example_index_css))
	mdgoFS.PathNodes.Set("/index.html",
		_mdgoFS_getNode(mdgoFS, "/index.html", generate__example_index_html))
	mdgoFS.PathNodes.Set("/sub",
		_mdgoFS_getNode(mdgoFS, "/sub", generate__example_sub))
	mdgoFS.PathNodes.Set("/sub/index.html",
		_mdgoFS_getNode(mdgoFS, "/sub/index.html", generate__example_sub_index_html))
	mdgoFS.PathNodes.Set("_internal",
		_mdgoFS_getNode(mdgoFS, "_internal", generate__internal))
	mdgoFS.PathNodes.Set("_internal/.template",
		_mdgoFS_getNode(mdgoFS, "_internal/.template", generate__internal__template))

	mdgoFS.Root = mdgoFS.PathNodes.Get("/")
}
