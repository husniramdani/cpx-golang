//go:build ignore

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const asciiArt = `
// ░░░░░░░░░░░░░░░SPINDYZEL░░░░░░░░░░░░░░░
// ░░░░░░░░░░▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄░░░░░░░░░
// ░░░░░░░░▄▀░░░░░░░░░░░░▄░░░░░░░▀▄░░░░░░░
// ░░░░░░░░█░░░░░░░░░░░░▄█▄▄░░▄░░░█░▄▄▄░░░
// ░▄▄▄▄▄░░█░░░░░░▀░░░░▀█░░▀▄░░░░░█▀▀░██░░
// ░██▄▀██▄█░░░▄░░░░░░░██░░░░▀▀▀▀▀░░░░██░░
// ░░▀██▄▀██░░░░░░░░▀░██▀░░░░░░░░░░░░░▀██░
// ░░░░▀████░▀░░░░▄░░░██░░░▄█░░░░▄░▄█░░██░
// ░░░░░░░▀█░░░░▄░░░░░██░░░░▄░░░▄░░▄░░░██░
// ░░░░░░░▄█▄░░░░░░░░░░░▀▄░░▀▀▀▀▀▀▀▀░░▄▀░░
// ░░░░░░█▀▀█████████▀▀▀▀████████████▀░░░░
// ░░░░░░████▀░░███▀░░░░░░▀███░░▀██▀░░░░░░
`

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, path := range args {
			err := processFile(path)
			if err != nil {
				panic(err)
			}
		}
	} else {
		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
				err := processFile(path)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}
}

func processFile(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if strings.Contains(string(content), asciiArt) {
		return nil
	}

	var newContent string
	if strings.HasPrefix(string(content), "//go:build ignore") {
		splitContent := strings.SplitN(string(content), "\n", 2)
		newContent = splitContent[0] + "\n" + asciiArt + "\n" + splitContent[1]
	} else {
		newContent = asciiArt + "\n" + string(content)
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(newContent), info.Mode())
	if err != nil {
		return err
	}

	// Output the name of the file that has been signed
	println("Sign => " + path)

	return nil
}
