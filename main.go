package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes  = "0123456789"
	specialBytes = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

var (
	length      int
	useLetters  bool
	useNumbers  bool
	useSpecials bool
	copyToClip  bool
	verbose     bool
)

func init() {
	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of the password (8-128)")
	rootCmd.Flags().BoolVarP(&useLetters, "letters", "t", true, "Include letters")
	rootCmd.Flags().BoolVarP(&useNumbers, "numbers", "n", true, "Include numbers")
	rootCmd.Flags().BoolVarP(&useSpecials, "specials", "s", false, "Include special characters")
	rootCmd.Flags().BoolVarP(&copyToClip, "copy", "c", false, "Copy to clipboard")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}

func secureRandomString(n int, charset string) (string, error) {
	b := make([]byte, n)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}

func generatePassword(length int, useLetters, useNumbers, useSpecials bool) (string, error) {
	if length < 8 || length > 128 {
		return "", fmt.Errorf("password length must be between 8 and 128")
	}

	if !useLetters && !useNumbers && !useSpecials {
		return "", fmt.Errorf("at least one character type must be selected")
	}

	var charset string
	var requiredChars []byte

	if useLetters {
		charset += letterBytes
		requiredChars = append(requiredChars, letterBytes[0])
	}
	if useNumbers {
		charset += numberBytes
		requiredChars = append(requiredChars, numberBytes[0])
	}
	if useSpecials {
		charset += specialBytes
		requiredChars = append(requiredChars, specialBytes[0])
	}

	password, err := secureRandomString(length-len(requiredChars), charset)
	if err != nil {
		return "", err
	}

	for _, char := range requiredChars {
		position, err := rand.Int(rand.Reader, big.NewInt(int64(len(password)+1)))
		if err != nil {
			return "", err
		}
		password = password[:position.Int64()] + string(char) + password[position.Int64():]
	}

	return password, nil
}

func calculateEntropy(password string) float64 {
	charset := float64(len(letterBytes) + len(numberBytes) + len(specialBytes))
	return float64(len(password)) * (float64(len(password)) / charset)
}

var rootCmd = &cobra.Command{
	Use:   "password",
	Short: "PasswordGen is a simple password generator",
	Long:  `A simple command line tool to generate secure passwords.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, err := generatePassword(length, useLetters, useNumbers, useSpecials)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Generated Password: %s\n", password)

		if verbose {
			entropy := calculateEntropy(password)
			fmt.Printf("Password entropy: %.2f bits\n", entropy)
		}

		if copyToClip {
			err := clipboard.WriteAll(password)
			if err != nil {
				fmt.Println("Error copying to clipboard:", err)
			} else {
				fmt.Println("Password copied to clipboard.")
			}
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
