// typerr let's you practice typing! Pass it a file via the -f flag and it'll ask you
// to type it line by line, giving you info on the mistakes you made.
//
// To give it a try, pass it "example.txt" located in the root directory of this package.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/isaacthefallenapple/typerr/internal/typing"
	"github.com/isaacthefallenapple/unbuffered"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

var (
	pathPtr   = flag.String("f", "", "Use this file to train.")
	stringPtr = flag.String("s", "", "Use this string to train.")
)

func main() {
	flag.Parse()

	str := *stringPtr
	path := *pathPtr

	var input io.Reader
	switch {
	case len(str) > 0 && len(path) > 0:
		fmt.Println("-s and -f are mutually exclusive.")
		return
	case len(str) > 0:
		input = strings.NewReader(str)
	case len(path) > 0:
		path = filepath.Clean(*pathPtr)
		if !strings.HasSuffix(path, ".txt") {
			fmt.Println("Not a text (.txt) file: ", path)
			os.Exit(1)
		}
		f, err := os.OpenFile(path, os.O_RDONLY, 0400)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	default:
		return
	}

	interruptions := make(chan os.Signal, 1)
	signal.Notify(interruptions, os.Interrupt)

	reset, err := unbuffered.SetUpConsole()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer reset()

	go func() {
		for range interruptions {
			reset()
			os.Exit(1)
		}
	}()

	r := fromReader(input)
	if err = r.err; err != nil {
		fmt.Println(">", err)
	}

	fmt.Println(r)
}

func fromReader(reader io.Reader) (result result) {
	scanner := bufio.NewScanner(reader)
	start := time.Now()
	defer func() {
		result.time = time.Since(start)
	}()
	for scanner.Scan() {
		result.Add(typing.TypeLine(strings.TrimSpace(scanner.Text())))
		if result.err != nil {
			return
		}
	}
	return
}
