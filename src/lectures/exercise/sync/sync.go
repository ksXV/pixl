//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Letters struct {
	totalLetters int
	sync.Mutex
}

var wg sync.WaitGroup

func countLetters(text string, waitgroup *sync.WaitGroup, letters *Letters) {
	letters.Lock()
	defer waitgroup.Done()
	defer letters.Unlock()
	sum := 0
	for i := 0; i < len(text); i++ {
		if unicode.IsLetter(rune(text[i])) {
			sum += 1
		}
	}
	letters.totalLetters += sum
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	n := strings.TrimSpace(input)

	brokenString := strings.Split(n, " ")
	totalLetters := Letters{}

	for i := 0; i < len(brokenString); i++ {
		wg.Add(1)
		go countLetters(brokenString[i], &wg, &totalLetters)
	}
	wg.Wait()

	fmt.Println("Total letters are:", totalLetters.totalLetters)
}
