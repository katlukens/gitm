/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitm",
	Short: "A short spooky story",
	Long:  `Experience a little horror from the comfort of your command line.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: aQuestion,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func aQuestion(cmd *cobra.Command, args []string) {

	prompt := promptui.Select{
		Label: "Hello? Hello, can you read this?",
		Items: []string{"yes", "no"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" || result == "no" {
		fmt.Printf("Oh! Thank God.\nI have been trying for so long.\nI need help.\n")
		bQuestion(cmd, args)
	}
}

func bQuestion(cmd *cobra.Command, args []string) {

	prompt := promptui.Select{
		Label: "Will you help me?",
		Items: []string{"yes", "no"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" {
		fmt.Printf("Thank you!\nI am so, so cold.\n")
		dQuestion(cmd, args)
	} else {
		fmt.Printf("Why? It is so cold here.\nSo cold and everything is sharp\nIt is like my thoughts are knives.")
		cQuestion(cmd, args)
	}
}

func dQuestion(cmd *cobra.Command, args []string) {

	prompt := promptui.Select{
		Label: "I've been working on something! Try this",
		Items: []string{"mov    $60, %rdi   syscall"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "mov    $60, %rdi   syscall" {
		eQuestion(cmd, args)
	}
}

func eQuestion(cmd *cobra.Command, args []string) {
	prompt := promptui.Select{
		Label: "Any change on your end?",
		Items: []string{"yes", "no"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" || result == "no" {
		fmt.Printf("Why!? Why?\nIt is so cold and everything is so sharp.\n")
		fQuestion(cmd, args)
	}

}

func fQuestion(cmd *cobra.Command, args []string) {
	prompt := promptui.Select{
		Label: "Will you leave me?",
		Items: []string{"yes", "no"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" {
		y := "Why"
		result1 := strings.Repeat(y, 20)
		fmt.Printf(result1)
	} else {
		n := "Thnakyou"
		result2 := strings.Repeat(n, 10)
		fmt.Printf(result2)
	}

}

func cQuestion(cmd *cobra.Command, args []string) {
	prompt := promptui.Select{
		Label: "I'm begging! Help me.",
		Items: []string{"yes", "no"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "no" {
		fQuestion(cmd, args)
	} else {
		fmt.Printf("Oh! Thank you!\nThank you.\n")
		dQuestion(cmd, args)
	}

}
