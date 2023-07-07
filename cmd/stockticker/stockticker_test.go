package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {

	build := exec.Command("go", "build", "-o", "stockticker_test")
	if err := build.Run(); err != nil {
		fmt.Println("Error Building Binary")
		os.Exit(1)

	}

	fmt.Println("Successful Build test")
	_, err := os.Stat("stockticker_test")
	if err != nil {
		fmt.Println("Error Packages Binary not found. Build might failed", err)
		os.Exit(1)
	}
	os.Remove("stockticker_test")
}
