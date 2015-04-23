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

// Use only when you know how many lines of text you are going to need
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

func Btoi(s string) int64 {
	if i, err := strconv.ParseInt(s, 2, 0); err != nil {
		panic(err)
	} else {
		return i
	}
}

func Max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}

type IntQueue []int

func (q *IntQueue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *IntQueue) Dequeue() (i int) {
	i = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *IntQueue) Size() int {
	return len(*q)
}

func (q *IntQueue) Empty() bool {
	return q.Size() == 0
}

type IntStack []int

func (s *IntStack) Push(i int) {
	*s = append(*s, i)
}

func (s *IntStack) Pop() (i int) {
	x := s.Size() - 1
	i = (*s)[x]
	*s = (*s)[:x]
	return
}
func (s *IntStack) Size() int {
	return len(*s)
}

func (s *IntStack) Empty() bool {
	return s.Size() == 0
}