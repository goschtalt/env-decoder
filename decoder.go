// SPDX-FileCopyrightText: 2022 Weston Schmidt <weston_schmidt@alumni.purdue.edu>
// SPDX-License-Identifier: Apache-2.0

// env package is a goschtalt decoder package for processing environment variables.
//
// See the example for how to use this extension package.
package env

import (
	"os"
	"strings"

	"github.com/goschtalt/goschtalt"
	"github.com/goschtalt/goschtalt/pkg/meta"
)

// envToObjs convert from environment variables into the meta.Object tree.
func envToObjs(prefix, delimiter, recordName string) (any, error) {
	tree := meta.Object{
		Map: make(map[string]meta.Object),
	}
	list := os.Environ()
	for _, item := range list {
		kvp := strings.Split(item, "=")
		if len(kvp) > 1 && strings.HasPrefix(kvp[0], prefix) {
			key := kvp[0]
			val := os.Getenv(key)
			key = strings.TrimPrefix(key, prefix)
			var err error
			tree, err = tree.Add(delimiter, key, meta.StringToBestType(val))
			if err != nil {
				return nil, err
			}
		}
	}

	return tree.ConvertMapsToArrays().ToRaw(), nil
}

// EnvVarConfig provides a way to collect configuration values from environment
// variables passed into the program.  The recordName is used to sort prior to
// the merge step, allowing the order of operations to be specified.  The prefix
// is the environment variable name prefix to look for when collecting them.
// The delimiter is the string used to split the tree structure on.
//
// For some environment variable environments like bash the allowable characters
// in the names is limited to: `[a-zA-Z_][a-zA-Z0-9_]*`
//
// If you need multiple prefix values, this option is safe to use multiple times.
func EnvVarConfig(recordName, prefix, delimiter string) goschtalt.Option {
	return goschtalt.AddValueGetter(recordName, goschtalt.Root,
		goschtalt.ValueGetterFunc(
			func(rn string, _ goschtalt.Unmarshaler) (any, error) {
				return envToObjs(prefix, delimiter, rn)
			},
		),
	)
}
