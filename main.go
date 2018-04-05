package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		n = flag.Int("n", 10, "lines")
		s = flag.Int("s", 0, "skip lines")
	)
	flag.Parse()
	filepath := flag.Arg(0)
	if !Exists(filepath) {
		// exit code 1 if file not exists
		fmt.Printf("ghd: %s: No such file or directory\n", filepath)
		os.Exit(1)
	}

	file, err := os.Open(filepath)
	if err != nil {
		// exit code 1 unable to open target
		fmt.Printf("ghd: %s: Permission denied\n", filepath)
		os.Exit(1)
	}
	defer file.Close()

	showed = 0
	sc := bufio.NewScanner(file)
	for lines := 1; sc.Scan(); lines++ {
		if err := sc.Err(); err != nil {
			os.Exit(1)
		}
		// skip showing until skip number
		if lines < *s {
			continue
		}
		fmt.Println(sc.Text())
		showed++
		// show 10 lines by default
		if showed == *n {
			os.Exit(0)
		}
	}
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
