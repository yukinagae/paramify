<div align="center">
  <img src="https://raw.githubusercontent.com/yukinagae/paramify/refs/heads/main/docs/resources/GO_BUILD.png" title="Gopher Build" alt="Gopher Build" style="width: 400px;">
</div>

# paramify

[![Go Reference](https://pkg.go.dev/badge/github.com/yukinagae/paramify.svg)](https://pkg.go.dev/github.com/yukinagae/paramify)
[![Go Report Card](https://goreportcard.com/badge/github.com/yukinagae/paramify)](https://goreportcard.com/report/github.com/yukinagae/paramify)
[![GitHub License](https://img.shields.io/github/license/yukinagae/paramify)](https://github.com/yukinagae/paramify/blob/main/LICENSE)
[![Github version](https://img.shields.io/github/v/release/yukinagae/paramify)](https://github.com/yukinagae/paramify/releases)
[![Static Badge](https://img.shields.io/badge/yes-a?label=maintained)](https://github.com/yukinagae/paramify/pulse)
[![GitHub Issues](https://img.shields.io/github/issues/yukinagae/paramify?color=blue)](https://github.com/yukinagae/paramify/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/yukinagae/paramify?color=blue)](https://github.com/yukinagae/paramify/pulls)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/m/yukinagae/paramify)](https://github.com/yukinagae/paramify/activity)

`paramify` is a build-function generator tool based on the [Functional Options Pattern](https://github.com/uber-go/guide/blob/master/style.md#functional-options), allowing you to build structs with intuitive, flexible, and type-safe APIs.

## Installation

```bash
go install github.com/yukinagae/paramify/cmd/paramify@latest
```

## Usage

Create a struct with the fields you need. Use struct tags to mark optional fields with `omitempty`.

```go
//go:generate paramify -type=User
type User struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Age     uint     `json:"age,omitempty"`
	Address *Address `json:"address,omitempty"`
	Friends []string `json:"friends,omitempty"`
}
```

Run the following command to generate the necessary functions for building instances of your struct using the Functional Options Pattern.

```bash
$ go generate
```

Use the generated functions to create instances of your struct. Required fields are passed as arguments to the constructor function, while optional fields are set using functional options.

```go
func main() {
	john := NewUser(
		"1",                                // Required: ID
		"John",                             // Required: Name
	)

	sam := NewUser(
		"2",                                // Required: ID
		"Sam",                              // Required: Name
		WithUserAge(20),                    // Optional: Age
		WithUserAddress(Address{"street"}), // Optional: Address
		WithUserFriends([]string{"Jane"}),  // Optional: Friends
	)
}
```

### Example

Here's a complete example demonstrating the usage:

```go
package main

import "fmt"

//go:generate paramify -type=User
type User struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Age     uint     `json:"age,omitempty"`
	Address *Address `json:"address,omitempty"`
	Friends []string `json:"friends,omitempty"`
}

type Address struct {
	Street string `json:"street"`
}

func main() {
	john := NewUser(
		"1",                                // Required: ID
		"John",                             // Required: Name
	)

	sam := NewUser(
		"2",                                // Required: ID
		"Sam",                              // Required: Name
		WithUserAge(20),                    // Optional: Age
		WithUserAddress(Address{"street"}), // Optional: Address
		WithUserFriends([]string{"Jane"}),  // Optional: Friends
	)

	fmt.Printf("%+v\n", john)
	fmt.Printf("%+v\n", sam)
}
```

## Contributing

We welcome contributions to this project! To get started, please refer to our [Contribution Guide](https://github.com/yukinagae/paramify/blob/main/CONTRIBUTING.md).
