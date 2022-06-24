package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"user/prisma/db"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Region    string `json:"region"`
	PostalZip string `json:"postalZip"`
	Country   string `json:"country"`
	GUID      string `json:"guid"`
}

func main() {
	client := db.NewClient()
	err := client.Prisma.Connect()
	if err != nil {
		log.Fatal("Unable to connect to postgresql")
	}

	defer client.Prisma.Disconnect()

	ctx := context.Background()

	jsonFile, err := os.Open("./prisma/seed/data.json")
	if err != nil {
		log.Fatal("Unable to open json file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []User
	_ = json.Unmarshal(byteValue, &users)

	for _, user := range users {
		_, err := client.User.UpsertOne(
			db.User.ID.Equals(user.GUID),
		).Create(
			db.User.ID.Set(user.GUID),
			db.User.Name.Set(user.Name),
			db.User.Email.Set(user.Email),
			db.User.PhoneNumber.Set(user.Phone),
			db.User.Address.Set(user.Address),
			db.User.City.Set(user.City),
			db.User.Region.Set(user.Region),
			db.User.PostalCode.Set(user.PostalZip),
			db.User.Country.Set(user.Country),
		).Update(
			db.User.Name.Set(user.Name),
			db.User.Email.Set(user.Email),
			db.User.PhoneNumber.Set(user.Phone),
			db.User.Address.Set(user.Address),
			db.User.City.Set(user.City),
			db.User.Region.Set(user.Region),
			db.User.PostalCode.Set(user.PostalZip),
			db.User.Country.Set(user.Country),
		).Exec(ctx)

		if err != nil {
			log.Println("Error seeding user data")
		}
	}

	log.Println("Successfully seeded user data")
}
