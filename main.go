package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var args = os.Args
	if len(args) > 1 {
		switch args[1] {
			case "--version":
				fmt.Println("1.0.0")
			case "--help":
				fmt.Println("./to-do new: creates new todo with left over todos from the previous day")
			case "new":
				checkPrevDailyNote()
			default:
				fmt.Println("I don't know that command, sorry")
		}
	}
}

func checkPrevDailyNote() {
	curr_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	now := time.Now()
	mon, day, yer := now.Month(), now.Day(), now.Year()
	today_file := fmt.Sprintf("%04d-%02d-%02d.md", yer, mon, day)
	today_filepath := filepath.Join(curr_dir, today_file)

	today_bytedata, err := os.ReadFile(today_filepath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("bytes: %v\n", string(today_bytedata))

}
