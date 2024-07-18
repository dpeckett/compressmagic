// SPDX-License-Identifier: MPL-2.0
/*
 * Copyright (C) 2024 Damian Peckett <damian@pecke.tt>.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package compressmagic

import (
	"bufio"
	"bytes"
	"compress/bzip2"
	"io"

	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v4"
	"github.com/ulikunitz/xz"
)

// NewReader returns a reader that decompresses the input stream if it is compressed.
func NewReader(r io.Reader) (io.ReadCloser, error) {
	bufioReader := bufio.NewReader(r)

	buf, err := bufioReader.Peek(8)
	if err != nil {
		return nil, err
	}

	switch {
	case bytes.HasPrefix(buf, []byte{0x42, 0x5A, 0x68}): // BZIP2
		return io.NopCloser(bzip2.NewReader(bufioReader)), nil
	case bytes.HasPrefix(buf, []byte{0x1F, 0x8B}): // GZIP
		return gzip.NewReader(bufioReader)
	case bytes.HasPrefix(buf, []byte{0x04, 0x22, 0x4D, 0x18}): // LZ4
		return io.NopCloser(lz4.NewReader(bufioReader)), nil
	case bytes.HasPrefix(buf, []byte{0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00}): // XZ
		xzReader, err := xz.NewReader(bufioReader)
		if err != nil {
			return nil, err
		}

		return io.NopCloser(xzReader), nil
	case bytes.HasPrefix(buf, []byte{0x28, 0xB5, 0x2F, 0xFD}): // ZSTD
		zstdReader, err := zstd.NewReader(bufioReader)
		if err != nil {
			return nil, err
		}

		return &readerWithCloser{
			Reader: zstdReader,
			close: func() error {
				zstdReader.Close()
				return nil
			},
		}, nil
	default:
		return io.NopCloser(bufioReader), nil
	}
}

type readerWithCloser struct {
	io.Reader
	close func() error
}

func (w *readerWithCloser) Close() error {
	return w.close()
}
