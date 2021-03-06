package main

import (
	"io/ioutil"
	"strings"
	"bytes"
	"bufio"
)

func main() {
	// Read Gone file
	file, _ := ioutil.ReadFile("hello.gone")
	boxCode := string(file)

	// Compile Gone to Go
	boxCode = strings.Replace(boxCode, "fun()", "func() {", -1)
	boxCode = strings.Replace(boxCode, "done", ")", -1)
	boxCode = strings.Replace(boxCode, "end", "}", -1)

	reader := bufio.NewReader(bytes.NewBuffer([]byte(boxCode)))

	var lines string

	for {
		ss, _, err := reader.ReadLine()
		ss_s := string(ss)

		if err == nil {
			if strings.TrimSpace(ss_s) == "import" {
				ss_s = strings.Replace(ss_s, "import", "import (", 1)
			} else if strings.TrimSpace(ss_s) == "var" {
				ss_s = strings.Replace(ss_s, "var", "var (", 1)
			} else if strings.HasPrefix(ss_s, "interface ") {
				ss_plit := strings.Split(ss_s, " ")
				ss_s = "type " + ss_plit[1] + " interface {"
			} else if strings.HasPrefix(ss_s, "struct ") {
				ss_s = "type " + strings.Split(ss_s, " ")[1] + " struct {"
			} else if strings.HasPrefix(ss_s, "fun ") {
				ss_s = strings.Replace(ss_s, "fun ", "func ", 1) + " {"
			} else if strings.HasPrefix(strings.TrimSpace(ss_s), "if ") {
				ss_s = ss_s + " {"
			} else if strings.TrimSpace(ss_s) == "else" {
				ss_s = "} " + ss_s + " {"
			} else if strings.HasPrefix(strings.TrimSpace(ss_s), "else ") {
				ss_s = "} " + ss_s + " {"
			} else if strings.HasPrefix(strings.TrimSpace(ss_s), "for ") {
				ss_s = ss_s + " {"
			} else if strings.TrimSpace(ss_s) == "for" {
				ss_s = ss_s + " {"
			} else if strings.HasPrefix(strings.TrimSpace(ss_s), "switch ") {
				ss_s = ss_s + " {"
			}

			lines += ss_s + "\n"

		} else {
			break
		}		
	}

	ioutil.WriteFile("hello.go", []byte(lines), 0664)
}
