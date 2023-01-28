# Go Build Script

A bash script used to build a go binary with some configuration options

## Usage

### Build your binary
```
./bin/build
```

### Build and package your binaries for release
```
./bin/build --all --package
```

### Display help options
```
./bin/build --help
```

## Configuration

This script will attempt to source a `build.config` file located in either a `.ci` or `ci` directory.

| Options                | Description                                                                                                         |
|------------------------|---------------------------------------------------------------------------------------------------------------------|
| `BUILD_OS`             | A comma separated list of valid GOOS operating systems to build for when using the --all flag (i.e. "darwin,linux") |
| `BUILD_ARCH`           | A comma separated list of valid GOARCH architectures to build for when using the --all flag (i.e. "amd64,arm64")    |
| `SKIP_BUILD`           | A comma separated list of build combinations that should not be built (i.e. "windows-arm,darwin-arm64")             |
| `GOLANG_BINARY_NAME`   | The binary name that will be used                                                                                   |
| `GOLANG_LDFLAGS`       | The build LDFLAGS to add. This script supports adding the `VERSION` and `GIT_COMMIT` variables                      |
| `GOLANG_PACKAGE`       | The go package name (i.e. go module name or github path)                                                            |
| `GOLANG_VERSION_PKG`   | The location of the version package in the code base. (i.e. internal/version)                                       |
| `RELEASE_EXTRA_FILES`  | A comma separated list of files to include in the release tarball (i.e. "LICENSE,README.md")                        |

> **Note:** The default `LDFLAGS` that will be used are `-s -w -X main.version=${VERSION} -X main.gitCommit=${GIT_COMMIT}`
