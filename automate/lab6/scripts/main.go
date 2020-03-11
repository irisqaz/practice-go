package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		oldfile := scanner.Text()
		newfile := strings.ReplaceAll(oldfile, "jane", "jdoe")

		_, err := os.Stat(oldfile)

		if !os.IsNotExist(err) && oldfile != newfile {
			command := exec.Command("mv", oldfile, newfile)
			result, err := command.Output()
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(string(result))
		}
	}
}
