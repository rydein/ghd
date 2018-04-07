package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout,
			`Usage of %s:
   %s [OPTIONS] [file ...]
Options
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	var (
		n = flag.Int("n", 10, "lines")
		s = flag.Int("s", 0, "skip lines")
		w = flag.String("w", "", "start from given word")
	)
	flag.Parse()
	for i := 0; i < len(flag.Args()); i++ {
		filepath := flag.Arg(i)
		if len(flag.Args()) > 1 {
			fmt.Printf("==> %s <==\n", filepath)
		}
		if !Exists(filepath) {
			// exit code 1 if file not exists
			fmt.Printf("ghd: %s: No such file or directory\n", filepath)
			os.Exit(1)
		}

		Open(filepath, n, s, w)
	}
	os.Exit(0)
}

func Exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func Open(filepath string, n *int, s *int, w *string) {
	file, err := os.Open(filepath)
	if err != nil {
		// exit code 1 unable to open target
		fmt.Printf("ghd: %s: Permission denied\n", filepath)
		os.Exit(1)
	}
	defer file.Close()

	showed := 0
	sc := bufio.NewScanner(file)
	search := *w
	for lines := 1; sc.Scan(); lines++ {
		if err := sc.Err(); err != nil {
			os.Exit(1)
		}
		// skip showing until skip number
		if lines < *s {
			continue
		}
		if search != "" {
			// -1 if not matches
			if strings.Index(sc.Text(), search) == -1 {
				continue
			}
			search = ""
		}
		fmt.Println(sc.Text())
		showed++
		// show 10 lines by default
		if showed == *n {
			break
		}
	}
}
