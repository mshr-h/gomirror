package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

func mirrorDir(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	sstat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sstat.IsDir() {
		return fmt.Errorf("%s is not directory", src)
	}

	err = os.MkdirAll(dst, sstat.Mode())
	if err != nil {
		return err
	}

	dstentries, err := ioutil.ReadDir(dst)
	if err != nil {
		return err
	}

	// Cleanup the destination directory.
	for _, dstentry := range dstentries {
		srcPath := filepath.Join(src, dstentry.Name())
		dstPath := filepath.Join(dst, dstentry.Name())

		// if dstPath doesn't exist, delete it.
		_, err := os.Stat(srcPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		if err != nil {
			fmt.Printf("Deleting %s\n", dstPath)
			err := os.RemoveAll(dstPath)
			if err != nil {
				return nil
			}
		}
	}

	// Copy all of the contents in the source directory to the destination directory
	// if the content doesn't exists or isn't newest.
	srcentries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, srcentry := range srcentries {
		srcPath := filepath.Join(src, srcentry.Name())
		dstPath := filepath.Join(dst, srcentry.Name())
		// fmt.Println(srcPath, dstPath)
		if srcentry.IsDir() {
			err := mirrorDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Copying %s -> %s\n", srcPath, dstPath)
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return nil
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()

	var (
		from = flag.Arg(0)
		to   = flag.Arg(1)
	)

	if from == "" || to == "" {
		fmt.Println("gomirror path/to/source path/to/destination")
		return
	}

	err := mirrorDir(from, to)
	if err != nil {
		log.Fatal(err)
	}
}
