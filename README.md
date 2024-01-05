<!--
SPDX-FileCopyrightText: 2022 Weston Schmidt <weston_schmidt@alumni.purdue.edu>
SPDX-License-Identifier: Apache-2.0
-->

# env-decoder
An optional environment variable based decoder for goschtalt.

[![Build Status](https://github.com/goschtalt/env-decoder/actions/workflows/ci.yml/badge.svg)](https://github.com/goschtalt/env-decoder/actions/workflows/ci.yml)
[![codecov.io](http://codecov.io/github/goschtalt/env-decoder/coverage.svg?branch=main)](http://codecov.io/github/goschtalt/env-decoder?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/goschtalt/env-decoder)](https://goreportcard.com/report/github.com/goschtalt/env-decoder)
[![GitHub Release](https://img.shields.io/github/release/goschtalt/env-decoder.svg)](https://github.com/goschtalt/env-decoder/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/goschtalt/env-decoder)](https://pkg.go.dev/github.com/goschtalt/env-decoder)

This decoder provides both an example of using the `goschtalt.AddValueGetter()` option
as well as a simple way to read in structured environment variables.

While this approach might seem like a great path, a far better way to interact
with environment variables is by creating a configuration file & using the built
in substitution approach via `ExpandEnv()`.  The reason for this is environment
variables often have limitations on the key case that may make their usage hard
on some systems, or across systems.
