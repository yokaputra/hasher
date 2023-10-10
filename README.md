# Go Hasher

The Hasher package provides a simple and secure way to perform password hashing and verification using bcrypt. It also offers a utility for generating random strings. This package is a wrapper for [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).

## Installation

To install this package, you can use the go get command:

```shell
go get -u github.com/yokaputra/hasher
```

Then import the validator package into your own code.

```go
import "github.com/yokaputra/hasher"
```

## Usage

Here's an example of how to use hashing a password:

```go
// Create a hasher instance
h := hasher.NewHasher()

// Hash a password with a specified cost (optional)
hashedPassword, err := h.Hash([]byte("myPassword"), hasher.DefaultCost)
if err != nil {
    // Handle error
}

// Store hashedPassword in your database
```

Verifying a password:

```go
// Verify a plain text password against a hashed password
isValid := h.Verify("myPassword", hashedPassword)
if isValid {
    // Password is valid
} else {
    // Password is invalid
}
```

Generating a random string:

```go
// Generate a random string of a specified length
randomString, err := h.Generate(10)
if err != nil {
    // Handle error
}
```

## Contribution

---

To contrib to this project, you can open a PR or an issue.

## License

This project is licensed under the Apache License 2.0. Please refer to the LICENSE file for more information.
