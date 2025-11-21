// Create a program that gets input from cla and prints the words
// that apear more than once in the input and its count.

package main 
import (
	"fmt"
	"bufio"
	"os"
)

//func main() {
	counts := make(map[string] int)
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
