// SPDX-FileCopyrightText: 2022 Weston Schmidt <weston_schmidt@alumni.purdue.edu>
// SPDX-License-Identifier: Apache-2.0

package env

import (
	"errors"
	"os"
	"testing"

	"github.com/goschtalt/goschtalt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEndToEnd(t *testing.T) {
	type kvp struct {
		k string
		v string
	}

	type input struct {
		filename  string
		prefix    string
		delimiter string
	}

	type expect struct {
		key      string
		val      string
		filename string
	}

	vars := []kvp{
		{k: "GOSCHTALT_foo_bar_0", v: "zero"},
		{k: "GOSCHTALT_foo_bar_1", v: "one"},
		{k: "GOSCHTALT_a", v: "one"},
		{k: "GOSCHTALT_b", v: "two"},
		{k: "DEFAULT_b", v: "one"},
		{k: "DEFAULT_c", v: "three"},
	}
	unknown := errors.New("unknown")
	tests := []struct {
		description string
		vars        []kvp
		input       [2]input
		expected    []expect
		expectedErr error
	}{
		{
			description: "A small test.",
			vars:        vars,
			input: [2]input{
				{filename: "1", prefix: "GOSCHTALT_", delimiter: "_"},
				{filename: "2", prefix: "DEFAULT_", delimiter: "_"},
			},
			expected: []expect{
				{key: "foo.bar.0", val: "zero", filename: "1"},
				{key: "foo.bar.1", val: "one", filename: "1"},
				{key: "a", val: "one", filename: "1"},
				{key: "b", val: "one", filename: "2"},
				{key: "c", val: "three", filename: "2"},
			},
		}, {
			description: "A small test, filenames swapped.",
			vars:        vars,
			input: [2]input{
				{filename: "2", prefix: "GOSCHTALT_", delimiter: "_"},
				{filename: "1", prefix: "DEFAULT_", delimiter: "_"},
			},
			expected: []expect{
				{key: "foo.bar.0", val: "zero", filename: "2"},
				{key: "foo.bar.1", val: "one", filename: "2"},
				{key: "a", val: "one", filename: "2"},
				{key: "b", val: "two", filename: "2"},
				{key: "c", val: "three", filename: "1"},
			},
		}, {
			description: "A small test, reversed order.",
			vars:        vars,
			input: [2]input{
				{filename: "1", prefix: "DEFAULT_", delimiter: "_"},
				{filename: "2", prefix: "GOSCHTALT_", delimiter: "_"},
			},
			expected: []expect{
				{key: "foo.bar.0", val: "zero", filename: "2"},
				{key: "foo.bar.1", val: "one", filename: "2"},
				{key: "a", val: "one", filename: "2"},
				{key: "b", val: "two", filename: "2"},
				{key: "c", val: "three", filename: "1"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)
			for _, val := range tc.vars {
				os.Setenv(val.k, val.v)
				defer os.Unsetenv(val.k)
			}
			c, err := goschtalt.New(
				EnvVarConfig(tc.input[0].filename, tc.input[0].prefix, tc.input[0].delimiter),
				EnvVarConfig(tc.input[1].filename, tc.input[1].prefix, tc.input[1].delimiter),
			)

			require.NoError(err)
			err = c.Compile()

			if tc.expectedErr == nil {
				assert.NoError(err)
				/*
					for _, val := range tc.expected {
							got, origin, err := c.FetchWithOrigin(val.key)
							assert.Equal(val.val, got)
							assert.NoError(err)
							assert.Equal(len(origin), 1)
							fn := fmt.Sprintf("%s.environ", val.filename)
							assert.Equal(fn, origin[0].File)
							assert.Equal(0, origin[0].Line)
							assert.Equal(0, origin[0].Col)
					}
				*/
				return
			}

			if errors.Is(unknown, tc.expectedErr) {
				assert.Error(err)
				return
			}

			assert.ErrorIs(err, tc.expectedErr)
		})
	}
}
