package util
import (
    "strconv"
    "bufio"
    "os"
    "fmt"
)

func OpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
	    os.Exit(1)
	}
	return f
}

func ReadLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	} else {
		return scanner.Text()
	}
}

func Atoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func Atof(s string) float64 {
	if f, err := strconv.ParseFloat(s, 32); err != nil {
		panic(err)
	} else {
		return f
	}
}