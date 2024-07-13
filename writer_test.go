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
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/dpeckett/compressmagic"
	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {
	t.Run("GZIP", func(t *testing.T) {
		h := sha256.New()

		w, err := compressmagic.NewWriter(h, "hello.gz")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, w.Close())
		})

		_, err = w.Write([]byte("Hello, World!"))
		require.NoError(t, err)

		expected := "e6e37cbb76b0b47e6c6b5a0c2106718b4954c6181bf917a12ea735738060a588"

		require.Equal(t, expected, hex.EncodeToString(h.Sum(nil)))
	})

	t.Run("LZ4", func(t *testing.T) {
		h := sha256.New()

		w, err := compressmagic.NewWriter(h, "hello.lz4")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, w.Close())
		})

		_, err = w.Write([]byte("Hello, World!"))
		require.NoError(t, err)

		expected := "eb6204a7f84a0b1e53716a54793786a9da17331e3786150da533a6211b153937"

		require.Equal(t, expected, hex.EncodeToString(h.Sum(nil)))
	})

	t.Run("XZ", func(t *testing.T) {
		h := sha256.New()

		w, err := compressmagic.NewWriter(h, "hello.xz")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, w.Close())
		})

		_, err = w.Write([]byte("Hello, World!"))
		require.NoError(t, err)

		expected := "e933718ef5810ecbd2e5e436e62f9707d38f509e5b67f24242784996d57cc957"

		require.Equal(t, expected, hex.EncodeToString(h.Sum(nil)))
	})

	t.Run("ZSTD", func(t *testing.T) {
		h := sha256.New()

		w, err := compressmagic.NewWriter(h, "hello.zst")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, w.Close())
		})

		_, err = w.Write([]byte("Hello, World!"))
		require.NoError(t, err)

		expected := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

		require.Equal(t, expected, hex.EncodeToString(h.Sum(nil)))
	})

	t.Run("None", func(t *testing.T) {
		h := sha256.New()

		w, err := compressmagic.NewWriter(h, "hello.txt")
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, w.Close())
		})

		_, err = w.Write([]byte("Hello, World!"))
		require.NoError(t, err)

		expected := "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"

		require.Equal(t, expected, hex.EncodeToString(h.Sum(nil)))
	})
}
