package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mischief/xosd"
)

var (
	timeout = flag.Int("t", 5, "timeout in seconds")
)

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func main() {
	var lines []string

	flag.Parse()

	args := flag.Args()

	if len(args) != 0 {
		lines = []string{strings.Join(args, " ")}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
	}

	osd, err := xosd.NewXOSD(len(lines))
	if err != nil {
		die("Failed to create XOSD: %v", err)
	}

	if err := osd.SetTimeout(*timeout); err != nil {
		die("Failed to set XOSD timeout: %v", err)
	}

	if err := osd.SetFont("-misc-fixed-*-*-*-*-64-*-*-*-*-*-*-*"); err != nil {
		die("Failed to set XOSD font: %v", err)
	}

	if err := osd.SetVerticalOffset(48); err != nil {
		die("Failed to set XOSD vertical offset: %v", err)
	}

	if err := osd.SetAlign(xosd.Center); err != nil {
		die("Failed to set XOSD alignment: %v", err)
	}

	for i, l := range lines {
		if err := osd.DisplayString(i, l); err != nil {
			die("Failed to display XOSD string: %v", err)
		}
	}

	if err := osd.Wait(); err != nil {
		die("Failed to display XOSD string: %v", err)
	}
}
