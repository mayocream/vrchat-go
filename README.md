# VRChat API Library for Go

A Go client to interact with the unofficial VRChat API. Supports all REST calls specified in the [API specification](https://github.com/vrchatapi/specification).

## Disclaimer

> Use of the API using applications other than the approved methods (website, VRChat application) are not officially supported. You may use the API for your own application, but keep these guidelines in mind:
> * We do not provide documentation or support for the API.
> * Do not make queries to the API more than once per 60 seconds.
> * Abuse of the API may result in account termination.
> * Access to API endpoints may break at any given time, with no warning.

## Getting Started

```go
package main

import (
	"github.com/mayocream/vrchat-go"
)

func main() {
	client := vrchat.NewClient("https://vrchat.com/api/1")

	err := client.Authenticate("username", "password", "totp")
	if err != nil {
		panic(err)
	}

	user, err := client.GetCurrentUser()
	if err != nil {
		panic(err)
	}

	println("logged in as ", user.DisplayName)
}
```

Read full example [here](examples/main.go).
