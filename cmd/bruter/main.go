package main

import (
	"flag"
	"log"
	"os"

	"github.com/MrWebUzb/bruter/internal/bruter"
	"golang.org/x/crypto/bcrypt"
)

var (
	wordlist []string
	err      error
	hashes   []bruter.HashFile
)

func main() {
	hashFile := flag.String("h", "", `
		Path to hash file.
		Hash file contains with login:password_hash format
		or password_hash format
	`)

	wordlistFile := flag.String("w", "", `
		Wordlist for comparing password hashes
	`)

	verbose := flag.Bool("v", false, "Verbose mode")

	flag.Parse()

	if *hashFile == "" {
		log.Printf("please enter hash file\n")
		os.Exit(1)
	}

	if *wordlistFile == "" {
		log.Printf("please enter wordlist file\n")
		os.Exit(1)
	}

	hashes, err = bruter.ReadHashFile(*hashFile)

	if err != nil {
		log.Printf("error while reading hash file: %v\n", err)
		os.Exit(1)
	}

	log.Printf("%d hashes loaded\n", len(hashes))

	wordlist, err = bruter.ReadWordlistFile(*wordlistFile)
	if err != nil {
		log.Printf("error while reading wordlist file: %v\n", err)
		os.Exit(1)
	}
	log.Printf("%d wordlist generated\n", len(wordlist))

	c := make(chan bool, len(hashes))
	for _, hash := range hashes {
		go func(hash bruter.HashFile) {
			checkHash(hash, *verbose)
			c <- true
		}(hash)
	}

	for i := 0; i < len(hashes); i++ {
		<-c
	}
}

func checkHash(hash bruter.HashFile, verbose bool) {
	for _, word := range wordlist {
		if verbose {
			log.Printf("(%s):(%s) checking...\n", hash.Username, word)
		}
		err := bcrypt.CompareHashAndPassword([]byte(hash.PasswordHash), []byte(word))

		if err != nil {
			if verbose {
				log.Printf("(%s):(%s) not matched\n", hash.Username, word)
			}
			continue
		}
		log.Printf("(%s):(%s)\n", hash.Username, word)
		break
	}
}
