package main

import (
	"fmt"
	"math/rand"
	"time"

	util "gitee.com/thoohv5/utils/file"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	util.WriteFile(fmt.Sprintf("doc/%s.go", RandString(5)), RandString(10000))
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
