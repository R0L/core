package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	switch os.Args[1] {
	case "message":
		message()
	}

}

func message() {
	times := make([]string, 0)
	t, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	for i := -120; i <= 0; i++ {
		day := time.Duration(i)
		ts := t.Add(day * 24 * time.Hour)
		ts = ts.Add(time.Minute * time.Duration(rand.Intn(24*60)))
		times = append(times, fmt.Sprintf("%s|%s", ts.Format("2006-01-02T15:04:05"), "feat(init):up"))
	}

	fmt.Print(strings.Join(times, " "))
}
