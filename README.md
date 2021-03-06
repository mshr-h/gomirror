# gomirror
[![Build Status](https://travis-ci.org/mshr-h/gomirror.svg?branch=master)](https://travis-ci.org/mshr-h/gomirror)
[![license](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)

This is a local directory mirroring tool written in Golang.

- Ubuntu/Windows support (probably works on macOS but not tested)
- Support differential copy
- No config file
- No external library dependency

## Installation

You can simply install by typing

```
go get github.com/mshr-h/gomirror
```

Make sure to use Go 1.7 or above.

## Usage

```
gomirror path/to/source_directory path/to/destination_directory
```

WARNING: Files/directories in the destination directory will be deleted if they don't exist in the source directory.

## Example

```
$ tree
.
├── dst
│   ├── dir2
│   │   └── file3
│   └── file4
└── src
    ├── dir1
    │   └── file1
    └── file2

4 directories, 4 files

$ gomirror src dst
2016/09/21 22:15:20 Deleting dst/dir2
2016/09/21 22:15:20 Deleting dst/file4
2016/09/21 22:15:20 Copying  src/dir1/file1 -> dst/dir1/file1
2016/09/21 22:15:20 Copying  src/file2 -> dst/file2

$ tree
.
├── dst
│   ├── dir1
│   │   └── file1
│   └── file2
└── src
    ├── dir1
    │   └── file1
    └── file2

4 directories, 4 files
```

Note that `dst/dir2/`, `dst/dir2/file3` and `dst/file4` were deleted due to mentioned on [Usage](#usage) section.

## Contribution

1. Fork it
1. Create your feature branch (git checkout -b my-new-feature)
1. Commit your changes (git commit -am 'Add some feature')
1. Push to the branch (git push origin my-new-feature)
1. Create new Pull Request

