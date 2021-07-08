package main

import (
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `localhost`,
	User:     `postgres`,
	Password: ``,
}

type Person struct {
	ID        uint   `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func main() {

	// Use Open to access the database.
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	fmt.Printf("Connected to %q with DSN:\n\t%q", sess.Name(), settings)

	fmt.Println()

	collections, err := sess.Collections()
	if err != nil {
		log.Fatal("Collections: ", err)
	}

	for i := range collections {
		// Name returns the name of the collection.
		fmt.Printf("-> %q\n", collections[i].Name())
	}

	person := sess.Collection("persons")
	ok, err := person.Exists()
	fmt.Printf("Q: Does table %q exists?\n", person.Name())
	fmt.Printf("R: %v (%v)", ok, err)

	fmt.Println()

	persons := []Person{}
	err = person.Find().All(&persons)
	if err != nil {
		log.Fatal("Persons.Find: ", err)
	}

	// Print the queried information.
	for i := range persons {
		fmt.Printf("record #%d: %#v\n", i, persons[i])
	}

	fmt.Println()

	// find 1 with id
	var p2 Person
	q := sess.SQL().
		SelectFrom("persons").
		Where("id = ?", 2)
	err = q.One(&p2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p2: %#v\n", p2)
}
