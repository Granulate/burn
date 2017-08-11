# StackO
StackO is a CLI tool to convert performance profiles (perf, etc) to a JSON hierarchical data structure that can be consumed by [d3-flame-graph](https://github.com/spiermar/d3-flame-graph). StackO can also generate a fully contained HTML flame graph from the same data.

## Getting Started

Just download _StackO_ from the _golang_ source or binary and use it!

### From go

Make sure you have [golang](https://golang.org/) installed and `GOPATH` correctly set.

```bash
$ go get github.com/spiermar/stacko
$ stacko $GOPATH/src/github.com/spiermar/stacko/examples/out.perf
```

### From binary

Binaries are being provided for both `linux` and `darwin`, in `amd64` arch.

#### darwin/amd64

```bash
$ curl -L "https://dl.bintray.com/mspier/stacko/darwin/amd64/stacko" -o stacko
$ ./stacko <perf_output_file>
```

#### linux/amd64

```bash
$ curl -L "https://dl.bintray.com/mspier/stacko/linux/amd64/stacko" -o stacko
$ ./stacko <perf_output_file>
```

## Options

```
$ stacko convert --help

Convert a performance profile to a JSON.

Examples:
  stacko convert examples/out.perf
  stacko convert --folded examples/out.perf-folded

Usage:
  stacko convert [flags] <input>

Flags:
  -f, --folded   input is a folded stack
  -h, --help     help for convert
  -p, --pretty   json output is pretty printed

Global Flags:
      --config string       config file (default is $HOME/.stacko.yaml)
      --cpuprofile string   write CPU profile to file
      --memprofile string   write heap profile to fil
```

```
$ stacko html --help

Convert a performance profile to HTML flame graph.

Examples:
  stacko html examples/out.perf
  stacko html --folded --output=flame.html examples/out.perf-folded

Usage:
  stacko html [flags] <input>

Flags:
  -f, --folded          input is a folded stack
  -h, --help            help for html
      --output string   output file

Global Flags:
      --config string       config file (default is $HOME/.stacko.yaml)
      --cpuprofile string   write CPU profile to file
      --memprofile string   write heap profile to file
```

## Examples

Input and output examples can be found in the [examples](/examples) directory.

## Building from Source

Make sure you have [golang](https://golang.org/) installed and `GOPATH` correctly set.

```bash
$ go get github.com/spiermar/stacko
$ cd $GOPATH/src/github.com/spiermar/stacko
$ go build
$ ./stacko examples/out.perf
```

## Issues

For bugs, questions and discussions please use the [GitHub Issues](https://github.com/spiermar/stacko/issues).

## Contributing

We love contributions! But in order to avoid total chaos, we have a few guidelines.

If you found a bug, have questions or feature requests, don't hesitate to open an [issue](https://github.com/spiermar/stacko/issues).

If you're working on an issue, please comment on it so we can assign you to it.

If you have code to submit, follow the general pull request format. Fork the repo, make your changes, and submit a [pull request](https://github.com/spiermar/stacko/pulls).

## License

Copyright 2017 Martin Spier. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.