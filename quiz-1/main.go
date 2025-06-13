package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Problem struct {
	question    string
	answer      string
	user_answer string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File required")
		return
	}

	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Err reading file ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Err reading all", err)
	}

	quiz := make(chan Problem, len(records))

	correct := 0
	go func() {
		for _, record := range records {
			fmt.Printf("%s = ", record[0])
			var u_answer string
			fmt.Scanln(&u_answer)
			quiz <- Problem{
				question:    record[0],
				answer:      record[1],
				user_answer: u_answer,
			}
		}
		close(quiz)
	}()

	for q := range quiz {
		if q.answer == q.user_answer {
			correct++
		}
	}

	fmt.Printf("\n\nYou got %d/%d correct", correct, len(records))
}
