/*
 * Copyright (c) 2021-2022 Vadim Vygonets <vadik@vygo.net>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
ihexcat concatenates Intel HEX files.

Usage:
	ihexcat [file.hex ...] >out.hex

Data and start addresses from subsequent files overwrite those from
previous files.  Merging data from 16-bit and 32-bit Intel HEX files
is not possible.
*/
package main

import (
	"log"
	"os"

	"github.com/unixdj/ihex"
)

func main() {
	var ix ihex.IHex
	for _, v := range os.Args[1:] {
		f, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		if err = ix.ReadFrom(f); err != nil {
			log.Fatal(v+":", err)
		}
		f.Close()
	}
	if ix.Format == ihex.FormatAuto {
		ix.Format = ihex.Format32Bit
	}
	if err := ix.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
