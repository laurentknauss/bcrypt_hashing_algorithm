package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	reader := bufio.NewReader(os.Stdin) 

	fmt.Print("Please enter some words , integers or a sentence of your choice: " ) 

	name, err := reader.ReadString('\n') 
	if err != nil { 
		log.Fatalf("Error handling your input: %v", err )
	}

// Trim any leading/trailing whitespaces and newlines from the input
name = strings.TrimSpace(name)

// Convert the name to a byte slice since bcrpyt uses byte slice 
nameBytes :=[]byte(name) 

// Generate a Bcrypt hash of the name 
hashedResult, err  := bcrypt.GenerateFromPassword(nameBytes, bcrypt.DefaultCost) 
if err != nil { 
	fmt.Println("Error", err) 
	return 
}
fmt.Printf("The Brcypt hash of your input '%s' is %s\n", name, string(hashedResult))
}