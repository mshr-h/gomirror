# gomirror

This is a local directory mirroring tool written in Golang.

- Ubuntu/Windows support (probably works on macOS but not tested.)
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

WARNING: Files/directories in the destination directory will be deleted if they doesn't exists in the source directory.

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
Deleting dst/dir2
Deleting dst/file4
Copying src/dir1/file1 -> dst/dir1/file1
Copying src/file2 -> dst/file2

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

Note that `dst/dir2/`, `dst/dir2/file3`, `dst/file4` was deleted due to mentioned on Usage section.

## Contribution

1. Fork it
1. Create your feature branch (git checkout -b my-new-feature)
1. Commit your changes (git commit -am 'Add some feature')
1. Push to the branch (git push origin my-new-feature)
1. Create new Pull Request

## License

MIT: [LICENSE](LICENSE)

