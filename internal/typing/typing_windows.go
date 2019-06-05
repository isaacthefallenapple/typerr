package typing

import (
	"fmt"
	"github.com/isaacthefallenapple/ansi"
	"github.com/isaacthefallenapple/unbuffered"
)

func TypeLine(text string) (input string, n int, err error) {
	inputBytes := make([]byte, len(text))
	var c, char rune
	fmt.Print(text)
	fmt.Print("\r")
	defer fmt.Println()
	i := 0
	for i < len(text) {
		c = rune(text[i])
		switch char, _ := unbuffered.ReadRune(); char {
		case 13, 26:
			input = string(inputBytes[:i])
			err = fmt.Errorf("input terminated")
			return
		case 8, 10:
			if i > 0 {
				i--
				ansi.Term()
				fmt.Printf("\b%c\b", text[i])
			}
			continue
		case c:
			fmt.Print(ansi.FG_Green.Paint(string(char)))
		default:
			n++
			fmt.Print(ansi.Chain(ansi.FG_Red, ansi.Bold, ansi.Underline).Paint(string(char)))
		}
		inputBytes[i] = byte(char)
		i++
	}
	input = string(inputBytes)
	return
}
