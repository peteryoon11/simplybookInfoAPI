package main

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	fmt.Println(time.Now().Format("2018-09-11"))
	t := time.Now()

	fmt.Printf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())
	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err = uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s", u2)

}
