package main

import (
	"bufio"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func sum1(in []byte) []byte {
	sum := sha1.Sum(in)
	return sum[:]
}

func sum256(in []byte) []byte {
	sum := sha256.Sum256(in)
	return sum[:]
}

func sum384(in []byte) []byte {
	sum := sha512.Sum384(in)
	return sum[:]
}

func sum512(in []byte) []byte {
	sum := sha512.Sum512(in)
	return sum[:]
}

var (
	progName  string
	sumHidden bool
	sumFunc   func([]byte) []byte
	fileNames = make([]string, 0)
	outWriter = os.Stdout
)

func walker(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	} else if _, fname := filepath.Split(path); fname[0] == '.' && !sumHidden {
		return nil
	}
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	fmt.Fprintf(outWriter, "%x  %s\n", sumFunc(fileData), path)
	return nil
}

func main() {
	algo := flag.Int("a", 256, "SHA hash algorithm (one of 256, 384, or 512)")
	check := flag.Bool("c", false, "read SHA sums from the files and check them")
	hidden := flag.Bool("d", false, "sum hidden (dot) files")
	flag.Parse()

	progName = filepath.Base(os.Args[0])
	sumHidden = *hidden

	switch *algo {
	case 1:
		sumFunc = sum1
	case 256:
		sumFunc = sum256
	case 384:
		sumFunc = sum384
	case 512:
		sumFunc = sum512
	default:
		fmt.Println("Unrecognised algorithm.")
		os.Exit(1)
	}

	if *check {
		checkFile(flag.Args())
	} else {
		for _, root := range flag.Args() {
			err := filepath.Walk(root, walker)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}
}

func checkFile(files []string) {
	var failed int

	for _, sumFilename := range files {
		sumFile, err := os.Open(sumFilename)
		if err != nil {
			fmt.Printf("%s: %s: failed to read file\n", progName, sumFilename)
			failed++
		}

		scanner := bufio.NewScanner(sumFile)
		var badLines, failedLines, failedFiles int
		var file, digest string
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			switch len(line) {
			case 0:
				break
			case 2:
				digest = line[0]
				file = line[1]
			case 3:
				digest = line[0]
				file = line[2]
			default:
				badLines++
				continue
			}
			fileData, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("%s: %s:\n", progName, file)
				fmt.Printf("%s: FAILED open or read\n", file)
				failedLines++
				continue
			}
			actualDigest := fmt.Sprintf("%x", sumFunc(fileData))
			if digest != actualDigest {
				fmt.Printf("%s: FAILED\n", file)
				failedFiles++
				continue
			}
			fmt.Printf("%s: OK\n", file)
		}
		sumFile.Close()
		if badLines > 0 {
			fmt.Printf("%s: WARNING: %d ", progName, badLines)
			if badLines == 1 {
				fmt.Printf("line is ")
			} else {
				fmt.Printf("lines are ")
			}
			fmt.Println("improperly formatted")
		}

		if failedLines > 0 {
			fmt.Printf("%s: WARNING: %d listed ", progName, failedLines)
			if badLines == 1 {
				fmt.Printf("file ")
			} else {
				fmt.Printf("files ")
			}
			fmt.Println("could not be read")
		}

		if failedFiles > 0 {
			fmt.Printf("%s: WARNING: %d computed ", progName, failedLines)
			if badLines == 1 {
				fmt.Printf("checksum ")
			} else {
				fmt.Printf("checksums ")
			}
			fmt.Println("did NOT match")
		}
	}
}
