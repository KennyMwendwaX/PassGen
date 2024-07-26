package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes  = "0123456789"
	specialBytes = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

var rng *rand.Rand

func init() {
	// Create a new random generator with a random seed
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randString(n int, charset string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}

func generatePassword(length int, useLetters, useNumbers, useSpecials bool) string {
	charset := ""
	if useLetters {
		charset += letterBytes
	}
	if useNumbers {
		charset += numberBytes
	}
	if useSpecials {
		charset += specialBytes
	}
	return randString(length, charset)
}

func main() {
	var length int
	var useLetters, useNumbers, useSpecials bool

	var rootCmd = &cobra.Command{
		Use:   "password",
		Short: "PasswordGen is a simple password generator",
		Long:  `A simple command line tool to generate secure passwords.`,
		Run: func(cmd *cobra.Command, args []string) {
			password := generatePassword(length, useLetters, useNumbers, useSpecials)
			fmt.Printf("Generated Password: %s\n", password)
		},
	}

	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of the password")
	rootCmd.Flags().BoolVarP(&useLetters, "letters", "t", true, "Include letters")
	rootCmd.Flags().BoolVarP(&useNumbers, "numbers", "n", true, "Include numbers")
	rootCmd.Flags().BoolVarP(&useSpecials, "specials", "s", true, "Include special characters")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
