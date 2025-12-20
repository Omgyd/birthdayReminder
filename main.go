package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
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

func calculateAge(birthdate time.Time) int {
	now := time.Now()
	age := now.Year() - birthdate.Year()
	if now.YearDay() < birthdate.YearDay() {
		age--
	}
	return age
}

func sendEmail(birthdays []Birthday) {
	if len(birthdays) == 0 {
		return // No birthdays today, no need to send email
	}

	password := os.Getenv("EMAIL_PASSWORD")
	fromEmail := os.Getenv("FROM_EMAIL")
	toEmail := []string{os.Getenv("TO_EMAIL")}

	var subject, body string
	if len(birthdays) > 1 {
		names := make([]string, len(birthdays))
		for i, b := range birthdays {
			names[i] = b.Name
		}
		nameList := strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]
		subject = fmt.Sprintf("Happy Birthday to %s!", nameList)

		bodyParts := []string{}
		for _, b := range birthdays {
			age := calculateAge(b.Birthdate)
			bodyParts = append(bodyParts, fmt.Sprintf("%s (%d)", b.Name, age))
		}
		body = fmt.Sprintf("Happy Birthday to %s! Hope you all have a wonderful day!", strings.Join(bodyParts, ", "))
	} else {
		b := birthdays[0]
		age := calculateAge(b.Birthdate)
		subject = fmt.Sprintf("Happy %dth Birthday to %s!", age, b.Name)
		body = fmt.Sprintf("Happy Birthday to %s! Hope you have a wonderful day!", b.Name)
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", strings.Join(toEmail, ","), subject, body))
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, toEmail, msg)
	if err != nil {
		log.Fatal("Failed to send email:", err)
	}
	log.Println("Email sent successfully!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	birthdays := readCSVFile("./data/birthdays.csv")
	todaysBirthdays := checkTodaysBirthdays(birthdays)
	if len(todaysBirthdays) > 0 {
		sendEmail(todaysBirthdays)
	} else {
		log.Println("No birthdays today")
	}

}
