# K3OS Pulumi Provider

This is a very simple pulumi native provider for K3OS.

Resources:
* Node: It manages file located in `/var/lib/rancher/k3os/config.yaml`, i.e. if the resource craated then config will be created also (and vise-versa).
*note: your server will be rebooted!"* A typescript example of using the single resource defined in `examples/simple`. *note: required Vagrant*

``Read`` call is not implemented so commands like `pulumi refresh` or `pulumi import` are not working right now.

## Installing

This package is available only for JS/TS or Golang.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @spigell/pulumi-k3os

or `yarn`:

    $ yarn add @spigell/pulumi-k3os

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/spigell/pulumi-k3os/sdk

## Building

### Dependencies

- Go 1.15
- NodeJS 10.X.X or later

### Local Build and Test

Most of the code for the provider implementation is in `provider/pkg` directory.  

A code generator is available which generates SDKs in TypeScript, Python, Go and .NET which are also checked in to the `sdk` folder.  The SDKs are generated from a schema in `schema.json`.  This file should be kept aligned with the resources, functions and types supported by the provider implementation.

Note that the generated provider plugin (`pulumi-resource-k3os`) must be on your `PATH` to be used by Pulumi deployments.

```bash
# build the plugin and provider
$ VERSION=0.0.1 make build

# test
$ cd examples/local-go
$ go mod tidy
$ pulumi stack init test
$ PATH=$PATH:../../bin pulumi up
```
Your rendered yaml config will be in `yaml` directory.