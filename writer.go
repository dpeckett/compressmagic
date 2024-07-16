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
	"io"
	"path/filepath"

	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4"
	"github.com/ulikunitz/xz"
)

// NewWriter returns a writer that compresses the output stream based on the file extension.
func NewWriter(w io.Writer, filename string) (io.WriteCloser, error) {
	ext := filepath.Ext(filename)

	switch ext {
	case ".gz", ".gzip":
		return gzip.NewWriter(w), nil
	case ".lz4":
		return lz4.NewWriter(w), nil
	case ".xz":
		xzWriter, err := xz.NewWriter(w)
		if err != nil {
			return nil, err
		}

		return &writerWithCloser{
			Writer: xzWriter,
			close: func() error {
				return xzWriter.Close()
			},
		}, nil
	case ".zst", ".zstd":
		zstdWriter, err := zstd.NewWriter(w)
		if err != nil {
			return nil, err
		}

		return &writerWithCloser{
			Writer: zstdWriter,
			close: func() error {
				zstdWriter.Close()
				return nil
			},
		}, nil
	default:
		return &writerWithCloser{
			Writer: w,
			close: func() error {
				return nil
			},
		}, nil
	}
}

type writerWithCloser struct {
	io.Writer
	close func() error
}

func (w *writerWithCloser) Close() error {
	return w.close()
}
