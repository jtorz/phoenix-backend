package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var countall bool

func main() {
	co := flag.Bool("countall", true, "true cuenta todas las lineas de codigo incluyendo lineas vacias (default:true)")
	flag.Parse()
	countall = *co
	c, err := countLinesDir(".", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Lineas: %d\n", c)
}

func shouldIgnoreDir(d string) bool {
	dirs := []string{
		"web",
		"dist",
		"extra",
		"cryptopasta",
		"rql-goqu",
		".git",
	}
	for _, ignored := range dirs {
		if d == ignored {
			return true
		}
	}
	return false
}

func countLinesDir(previusDir, d string) (c int, err error) {
	if shouldIgnoreDir(d) {
		return
	}
	var dirPath string
	if d == "" {
		dirPath = previusDir
	} else {
		dirPath = previusDir + "/" + d
	}
	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return
	}
	var dc int
	for _, file := range files {
		if file.IsDir() {
			dc, err = countLinesDir(dirPath, file.Name())
		} else {
			dc, err = countLinesFile(dirPath, file.Name())
		}
		if err != nil {
			return
		}
		c += dc
	}
	return
}

func countLinesFile(dir, f string) (int, error) {
	if f == "go.mod" ||
		f == "go.sum" {
		return 0, nil
	}
	dc := 0
	/* if f[len(f)-2:] != "go" {
		fmt.Println(f)
	} */
	file, err := os.Open(dir + "/" + f) // For read access.
	if err != nil {
		return 0, fmt.Errorf("%v; %s", err, f)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	notIms := true
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "import (" {
			notIms = false
		}
		if len(s) > 2 &&
			s[:2] != "//" &&
			!strings.HasPrefix(s, "package") &&
			notIms &&
			countall {
			dc++
		}
		if s == ")" {
			notIms = true
		}
	}
	fmt.Println(dc, f)
	//l.Info().Int("lines", dc).Msg(f)
	err = scanner.Err()
	if err != nil {
		return dc, fmt.Errorf("%v: %s ", err, f)
	}
	return dc, nil
}
