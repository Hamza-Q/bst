# bst
Experimental generics-aware implementation of a simple binary search tree in Go.

This project is intended as an example of generics, and is strongly discouraged for production use. It relies on the [dev.typeparams](https://github.com/golang/go/tree/dev.typeparams) branch of the go language repository, which is experimental and subject to change.

## Building locally
See https://github.com/golang/tools/blob/master/gopls/doc/advanced.md#working-with-generic-code for instructions on how to build a generics-aware go toolchain.

The generics-aware go tool (recommended to be built by gotip in a separate location from existing go tools) must be invoked with the build flag `-gcflags=-G=3` in order for the above compiler to recognize generics. For example, to run unit tests from the top level directory, run:
```
go test -gcflags=-G=3
```
