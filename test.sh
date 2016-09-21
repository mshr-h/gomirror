#!/bin/bash
rm -rf testdir/

# prepare testdir
mkdir -p testdir/source/directory1
mkdir -p testdir/source/directory2
echo     "This is file1." > testdir/source/directory1/file1
echo     "This is file2." > testdir/source/directory1/file2
echo     "This is file3." > testdir/source/directory2/file3
echo     "This is file4." > testdir/source/directory2/file4
echo     "This is file5." > testdir/source/file5
echo     "This is file6." > testdir/source/file6

mkdir -p testdir/destination/dummydirectory
echo     "This is dummyfile1." > testdir/destination/dummydirectory/dummyfile1
echo     "This is dummyfile2." > testdir/destination/dummyfile2

go run main.go testdir/source testdir/destination

if [ -n "$(diff -r -q testdir/source testdir/destination)" ]; then
  echo "NG"
  exit 1
fi

echo     "Append" >> testdir/source/directory1/file1
go run main.go testdir/source testdir/destination

if [ -n "$(diff -r -q testdir/source testdir/destination)" ]; then
  echo "NG"
  exit 1
else
  echo "OK"
  rm -rf testdir
  exit 0
fi

