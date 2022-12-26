// SPDX-FileCopyrightText: 2022 Weston Schmidt <weston_schmidt@alumni.purdue.edu>
// SPDX-License-Identifier: Apache-2.0

package env_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/goschtalt/env-decoder"
	"github.com/goschtalt/goschtalt"
)

func Example() {
	_ = os.Setenv("EXAMExample_Version", "1")
	_ = os.Setenv("EXAMExample_Colors_0", "red")
	_ = os.Setenv("EXAMExample_Colors_1", "green")
	_ = os.Setenv("EXAMExample_Colors_2", "blue")
	g, err := goschtalt.New(
		goschtalt.AutoCompile(),
		env.EnvVarConfig("record", "EXAM", "_"),
	)
	if err != nil {
		panic(err)
	}

	var cfg struct {
		Example struct {
			Version int
			Colors  []string
		}
	}

	err = g.Unmarshal(goschtalt.Root, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("example")
	fmt.Printf("    version = %d\n", cfg.Example.Version)
	fmt.Printf("    colors  = [ %s ]\n", strings.Join(cfg.Example.Colors, ", "))

	// Output:
	// example
	//     version = 1
	//     colors  = [ red, green, blue ]
}
