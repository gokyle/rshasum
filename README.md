rshasum is a recursive SHA sum program.

### Usage

A SHA algorithm may be specified with the `-a` flag; valid algorithms
are 1, 256, 384, and 512; these select SHA-1, SHA-256, SHA-384, and
SHA-512, respectively. The default is SHA-256.

Typically, hidden (dot) files are skipped. For example, given the
following directory tree
```
	testdir/
	├── .hidden
	└── file
```
the `.hidden` file will not be hashed. The `-d` flag (for dotfiles)
will cause `rshasum` to hash these files.

To compare a digest file, the `-c` flag should be used. In this
case, the file arguments are taken to be files containing digests
and filenames. While it can match the standard `shasum` format for
text files, it will not recognise portable or binary file formats.
It will also recognise the OpenBSD `sha` digest utilities, provided
they were called in reverse format (ex. `sha1 -r`).

### LICENSE

> Copyright (c) 2014 Kyle Isom <kyle@tyrfingr.is>
> 
> Permission to use, copy, modify, and distribute this software for any
> purpose with or without fee is hereby granted, provided that the above 
> copyright notice and this permission notice appear in all copies.
> 
> THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
> WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
> MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
> ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
> WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
> ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
> OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE. 
