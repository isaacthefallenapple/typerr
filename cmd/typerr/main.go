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
	"path/filepath"
	"strings"
	"time"
)

var pathPtr = flag.String("f", "", "Use this file to train.")

func main() {
	flag.Parse()

	path := *pathPtr

	path = filepath.Clean(*pathPtr)
	if !strings.HasSuffix(path, ".txt") {
		log.Println("Not a text (.txt) file: ", path)
		return
	}
	input, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		log.Println(err)
		return
	}
	defer input.Close()

	reset, err := unbuffered.SetUpConsole()
	if err != nil {
		log.Println(err)
		return
	}
	defer reset()

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
