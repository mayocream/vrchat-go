# VRChat API Library for Go

A Go client to interact with the unofficial VRChat API. Supports all REST calls specified in the [API specification](https://github.com/vrchatapi/specification).

## Disclaimer

> Use of the API using applications other than the approved methods (website, VRChat application) are not officially supported. You may use the API for your own application, but keep these guidelines in mind:
> * We do not provide documentation or support for the API.
> * Do not make queries to the API more than once per 60 seconds.
> * Abuse of the API may result in account termination.
> * Access to API endpoints may break at any given time, with no warning.

## Getting Started

import the library

```go
import "github.com/mayocream/vrchat-go"
```

Create a new client

```go
client := vrchat.NewClient("https://vrchat.com/api/1")
```
