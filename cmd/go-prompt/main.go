package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func PrintIntro() {
	c := color.New(color.FgBlue)

	c.Println("   ______         ____                             __ ")
	c.Println("  / ____/___     / __ \\_________  ____ ___  ____  / /_")
	c.Println(" / / __/ __ \\   / /_/ / ___/ __ \\/ __ `__ \\/ __ \\/ __/")
	c.Println("/ /_/ / /_/ /  / ____/ /  / /_/ / / / / / / /_/ / /_  ")
	c.Println("\\____/\\____/  /_/   /_/   \\____/_/ /_/ /_/ .___/\\__/  ")
	c.Println("					/_/              ")

	color.Unset()
}

type gitRepo struct {
	url string
}

func (r *gitRepo) Clone() error {
	fmt.Printf("Cloning at %q\n", r.url)

	var out bytes.Buffer
	cmd := exec.Command("git", []string{"clone", r.url}...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Failed to clone repo %q\n", out)
		return err
	}

	fmt.Printf("Successfuly clonned repo\n")
	return nil
}

const (
	python_repo = "https://github.com/python/cpython.git"
	go_repo     = "https://github.com/golang/go.git"
	ruby_repo   = "https://github.com/ruby/ruby.git"
)

var repoMap = map[string]gitRepo{
	"Python": {url: python_repo},
	"Go":     {url: go_repo},
	"Ruby":   {url: ruby_repo},
}

func main() {
	PrintIntro()
	prompt := promptui.Select{
		Label: "Select Repo to clone",
		Items: []string{"Python", "Go", "Ruby"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Cloning the github repo for %q\n", result)

	repo := repoMap[result]
	repo.Clone()
}
