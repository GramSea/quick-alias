package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)
	zshrc := dirname + "/.zshrc"
	file, err := os.Open(zshrc)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var aliases []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "alias") {
			cmd := strings.Split(text, "alias")
			aliases = append(aliases, cmd[1])
		}
	}
	i := 1
	for alias := range aliases {
		fmt.Println(strconv.Itoa(i) + ")" + aliases[alias])
		i++
	}
	fmt.Println("Enter the number of the alias you want to use: ")
	var input int
	fmt.Scanln(&input)

	cmd_str := strings.Split(aliases[input-1], "=")[1]
	len := len(cmd_str)
	cmd_str = cmd_str[1 : len-1]
	fmt.Println(cmd_str)
	cmd_args := strings.Split(cmd_str, " ")
	cmd_main := cmd_args[0]
	cmd_args = cmd_args[1:]

	cmd := exec.Command(cmd_main, cmd_args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")

}
