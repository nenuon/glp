package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"../github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var issuesLessOneMonth []string
	var issuesLessOneYear []string
	var issuesOverOneYear []string
	fmt.Println(result.TotalCount, "issues:")
	now := time.Now()
	for _, item := range result.Items {
		result := fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		diff_hour := now.Sub(item.CreatedAt).Hours()
		if diff_hour < 24*30 {
			issuesLessOneMonth = append(issuesLessOneMonth, result)
		} else if diff_hour < 24*365 {
			issuesLessOneYear = append(issuesLessOneYear, result)
		} else {
			issuesOverOneYear = append(issuesOverOneYear, result)
		}
	}
	fmt.Println(issuesLessOneMonth)
	fmt.Println(issuesLessOneYear)
	fmt.Println(issuesOverOneYear)
}
