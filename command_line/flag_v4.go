package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// flagSet：实现选项分组
	git := flag.NewFlagSet("git", flag.ContinueOnError)
	add := git.String("add", "", "add file")
	commit := git.String("commit", "", "commit file")
	err := git.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("err: ", err)
		return
	}
	fmt.Println("git: ", *add, *commit)

	golang := flag.NewFlagSet("go", flag.ContinueOnError)
	run := golang.String("run", "", "compile and run Go program")
	build := golang.String("build", "", "compile packages and dependencies")
	err = golang.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("err: ", err)
		return
	}
	fmt.Println("golang: ", *run, *build)
}
