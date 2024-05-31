package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate random passwords",
	Long: `Generate random passwords with customizable options.
	
	For example:
	passwordGen generate -l 12 -d -s`,

	Run: GeneratePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 8, "Length of generated password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special char in the generated password")
}

func GeneratePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "!@#$%^&*()_+{}[]|:;,.<>?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Println("Generating Password")
	fmt.Println(string(password))
}
