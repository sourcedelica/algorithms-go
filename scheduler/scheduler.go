package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Job struct {
	Weight int
	Length int
}

// Sort interface: by weight / length, descending
type ByRatio []Job

func (a ByRatio) Len() int           { return len(a) }
func (a ByRatio) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRatio) Less(i, j int) bool {
	return float32(a[j].Weight) / float32(a[j].Length) < float32(a[i].Weight) / float32(a[i].Length)
}

// Sort interface: by weight - length, descending
type ByDifference []Job

func (a ByDifference) Len() int           { return len(a) }
func (a ByDifference) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDifference) Less(i, j int) bool {
	iDiff := a[i].Weight - a[i].Length
	jDiff := a[j].Weight - a[j].Length
	var result bool
	if (iDiff == jDiff) {
		result = a[j].Weight < a[i].Weight
	} else {
		result = jDiff < iDiff
	}
	return result
}

func main() {
	if (len(os.Args) < 2) {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	jobs := ReadJobs(filename)

	CalculateCompletionTimes(jobs, ByDifference(jobs), "difference")
	CalculateCompletionTimes(jobs, ByRatio(jobs), "ratio")
}

func ReadJobs(filename string) []Job {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
	    os.Exit(1)
	}

	scanner := bufio.NewScanner(bufio.NewReader(f))
	count := Atoi(ReadLine(scanner))
	jobs := make([]Job, count)

	for i := 0; i < count ; i++ {
		jobParts := strings.Split(ReadLine(scanner), " ")
		jobs[i] = Job{ Weight: Atoi(jobParts[0]), Length: Atoi(jobParts[1]) }
	}

	return jobs
}

func CalculateCompletionTimes(jobs []Job, sortJobs sort.Interface, desc string) {
	sort.Sort(sortJobs)
	sum := SumCompletionTimes(jobs)
	fmt.Printf("Sum by %s: %d\n", desc, sum)
}

func SumCompletionTimes(jobs []Job) int {
	var sum, completionTime int = 0, 0
	for _, job := range jobs {
		completionTime += job.Length
		sum += job.Weight * completionTime
	}
	return sum
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