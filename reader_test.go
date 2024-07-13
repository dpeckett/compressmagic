// SPDX-License-Identifier: MPL-2.0
/*
 * Copyright (C) 2024 Damian Peckett <damian@pecke.tt>.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package compressmagic_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/dpeckett/compressmagic"
	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	t.Run("BZIP2", func(t *testing.T) {
		f, err := os.Open("testdata/hello.bz2")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, f.Close())
		})

		r, err := compressmagic.NewReader(f)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, r.Close())
		})

		expected := "Hello, World!"

		buf := make([]byte, len(expected))
		_, err = r.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			require.NoError(t, err)
		}

		require.Equal(t, expected, string(buf[:]))
	})

	t.Run("GZIP", func(t *testing.T) {
		f, err := os.Open("testdata/hello.gz")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, f.Close())
		})

		r, err := compressmagic.NewReader(f)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, r.Close())
		})

		expected := "Hello, World!"

		buf := make([]byte, len(expected))
		_, err = r.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			require.NoError(t, err)
		}

		require.Equal(t, expected, string(buf[:]))
	})

	t.Run("LZ4", func(t *testing.T) {
		f, err := os.Open("testdata/hello.lz4")
		require.NoError(t, err)

		r, err := compressmagic.NewReader(f)
		require.NoError(t, err)

		expected := "Hello, World!"

		buf := make([]byte, len(expected))
		_, err = r.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			require.NoError(t, err)
		}

		require.Equal(t, expected, string(buf[:]))
	})

	t.Run("XZ", func(t *testing.T) {
		f, err := os.Open("testdata/hello.xz")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, f.Close())
		})

		r, err := compressmagic.NewReader(f)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, r.Close())
		})

		expected := "Hello, World!"

		buf := make([]byte, len(expected))
		_, err = r.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			require.NoError(t, err)
		}

		require.Equal(t, expected, string(buf[:]))
	})

	t.Run("ZSTD", func(t *testing.T) {
		f, err := os.Open("testdata/hello.zst")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, f.Close())
		})

		r, err := compressmagic.NewReader(f)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, r.Close())
		})

		expected := "Hello, World!"

		buf := make([]byte, len(expected))
		_, err = r.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			require.NoError(t, err)
		}

		require.Equal(t, expected, string(buf[:]))
	})

	t.Run("None", func(t *testing.T) {
		f, err := os.Open("testdata/hello.txt")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, f.Close())
		})

		r, err := compressmagic.NewReader(f)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, r.Close())
		})

		expected := "Hello, World!"

		buf := make([]byte, len(expected))
		_, err = r.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			require.NoError(t, err)
		}

		require.Equal(t, expected, string(buf[:]))
	})
}
