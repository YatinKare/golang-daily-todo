package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"strings"
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
				unchecked_move_to_tommorrow := checkPrevDailyNote()
				if err := write_new_file(unchecked_move_to_tommorrow); err != nil {
					fmt.Println("[ERROR]:", err)
				}
			default:
				fmt.Println("I don't know that command, sorry")
		}
	} else {
		fmt.Println("Please use a command. See '--help'")
	}
}

func checkPrevDailyNote() []string {
	curr_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	now := time.Now().AddDate(0, 0, -1)
	mon, day, yer := now.Month(), now.Day(), now.Year()
	today_file := fmt.Sprintf("%04d-%02d-%02d.md", yer, mon, day)
	today_filepath := filepath.Join(curr_dir, today_file)

	today_bytedata, err := os.ReadFile(today_filepath)
	today_data := string(today_bytedata)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(today_data, "\n")

	// loop through all lines of today's file
	var unchecked []string

	for _, val := range lines {
		if strings.Contains(val, "- [ ]") {
			// start := strings.Index(val, "- [ ]")
			// fmt.Printf("FOUND UNCHECKED [%d]:%s\n", index, val)
			// fmt.Printf("start %d\n", start)

			// unchecked_value := val[start + 6:len(val)]
			unchecked = append(unchecked, val)
		}
	}
	return unchecked
}

func write_new_file(unchecked_values []string) error {
	curr_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	now := time.Now()
	mon, day, yer := now.Month(), now.Day(), now.Year()
	tomorrow_file := fmt.Sprintf("%04d-%02d-%02d.md", yer, mon, day)
	tomorrow_filepath := filepath.Join(curr_dir, tomorrow_file)

	// buffer to write to file
	var buff string
	buff = buff + "# " + tomorrow_file[:len(tomorrow_file) - 3] + "\n"
	for _, unchecked_val := range unchecked_values {
		buff = buff + unchecked_val + "\n"
	}

	// os.CREATE: creates if does not exist, otherwise opens the existing one
	// os.EXCL: **only when os.CREATE is set**: returns err if there is an existing one
	f, err := os.OpenFile(tomorrow_filepath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("File exists in the current director, please check: \n%s\n", tomorrow_filepath)
		}
		return err
	}
	defer f.Close()

	_, err = f.WriteString(buff)

	if err != nil {
		return err
	}

	return nil
}
