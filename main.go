package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	if err := out.Sync(); err != nil {
		return err
	}

	si, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.Chmod(dst, si.Mode()); err != nil {
		return err
	}

	if err := os.Chtimes(dst, time.Now(), si.ModTime()); err != nil {
		return err
	}

	return nil
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

	if err := os.MkdirAll(dst, sstat.Mode()); err != nil {
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

		// if srcPath doesn't exist, delete it.
		if _, err := os.Stat(srcPath); os.IsNotExist(err) {
			log.Printf("Deleting %s\n", dstPath)
			if err := os.RemoveAll(dstPath); err != nil {
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
		if srcentry.IsDir() {
			if err := mirrorDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			si, err := os.Stat(srcPath)
			if err != nil {
				return err
			}

			if di, err := os.Stat(dstPath); os.IsNotExist(err) {
				log.Printf("Copying  %s -> %s\n", srcPath, dstPath)
				if err := copyFile(srcPath, dstPath); err != nil {
					// check if srcPath is whether in-use or not
					if strings.Contains(err.Error(), "The process cannot access the file because it is being used by another process.") {
						log.Printf("Ignore   %s because it's being used by another process.\n", srcPath)
						continue
					}
					return err
				}
			} else if isModified(si, di) {
				if err := os.RemoveAll(dstPath); err != nil {
					return nil
				}

				log.Printf("Updating %s -> %s\n", srcPath, dstPath)
				if err := copyFile(srcPath, dstPath); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func isModified(src, dst os.FileInfo) bool {
	return src.Size() != dst.Size() || !src.ModTime().Equal(dst.ModTime())
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

	if err := mirrorDir(from, to); err != nil {
		log.Fatal(err)
	}
}
