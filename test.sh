#!/bin/bash
rm -rf testdir/

# prepare testdir
mkdir -p testdir/source/directory1
mkdir -p testdir/source/directory2
touch    testdir/source/directory1/file1
touch    testdir/source/directory1/file2
touch    testdir/source/directory2/file3
touch    testdir/source/directory2/file4
touch    testdir/source/file5
touch    testdir/source/file6

mkdir -p testdir/destination/dummydirectory
touch    testdir/destination/dummydirectory/dummyfile1
touch    testdir/destination/dummyfile2

go run main.go testdir/source testdir/destination

if [ -z "$(diff -r -q testdir/source testdir/destination)" ]; then
  echo "OK"
  rm -rf testdir
  exit 0
else
  echo "NG"
  exit 1
fi

