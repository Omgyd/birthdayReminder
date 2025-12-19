package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

//TODO: Read csv file and print the content

type Birthday struct {
	Name      string
	Birthdate time.Time
}

func readCSVFile(filePath string) []Birthday {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	birthdays := []Birthday{}
	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	for _, record := range records[1:] {
		name := record[0]
		birthdate, err := time.Parse("2006-01-02", record[1])
		if err != nil {
			log.Fatal("Unable to parse date for "+name, err)
		}
		birthdays = append(birthdays, Birthday{
			Name:      name,
			Birthdate: birthdate,
		})
	}
	return birthdays
}

func checkTodaysBirthdays(birthdays []Birthday) []Birthday {
	todaysBirthdays := []Birthday{}
	today := time.Now()
	for _, birthday := range birthdays {
		if birthday.Birthdate.Month() == today.Month() && birthday.Birthdate.Day() == today.Day() {
			todaysBirthdays = append(todaysBirthdays, birthday)
		}
	}
	return todaysBirthdays
}

func sendEmail(name []string, age int) {
	password := os.Getenv("EMAIL_PASSWORD")
	fromEmail := os.Getenv("FROM_EMAIL")
	toEmail := []string{os.Getenv("TO_EMAIL")}

	if len(name) > 1 {
		// Figure out how to create message with multiple people
	} else if len(name) == 1 {
		//Set message for one person
	}
	// Finish logic to send email

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
