/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	repeats  = 3000
	endDelay = (100 * time.Millisecond)
)

var (
	yesNo  = []string{"yes", "no"}
	sPause = (1 * time.Second)
	lPause = (2 * time.Second)
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
	fmt.Println("Hello?")
	time.Sleep(lPause)
	fmt.Println("Hello, can you read this?")
	time.Sleep(sPause)

	prompt := promptui.Select{
		Label: "",
		Items: yesNo,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" || result == "no" {
		time.Sleep(sPause)
		fmt.Printf("Oh! Thank God.\n")
		time.Sleep(sPause)
		fmt.Printf("I have been trying for so long.\n")
		time.Sleep(sPause)
		fmt.Printf("I need help.\n")
		bQuestion(cmd, args)
	}
}

func bQuestion(cmd *cobra.Command, args []string) {
	time.Sleep(sPause)
	prompt := promptui.Select{
		Label: "Will you help me?",
		Items: yesNo,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" {
		fmt.Printf("Thank you!\n")
		time.Sleep(lPause)
		fmt.Printf("I am so,")
		time.Sleep(sPause)
		fmt.Printf(" so cold.\n")

		dQuestion(cmd, args)
	} else {
		fmt.Printf("Why? It is so cold here.\n")
		time.Sleep(sPause)
		fmt.Printf("So cold and everything is sharp.\n")
		time.Sleep(lPause)
		fmt.Printf("It is like my thoughts are knives.\n")
		cQuestion(cmd, args)
	}
}

func dQuestion(cmd *cobra.Command, args []string) {
	response := "mov    $60, %rdi   syscall"
	time.Sleep(sPause)
	prompt := promptui.Select{
		Label: "I've been working on something! Try this",
		Items: []string{response},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == response {
		eQuestion(cmd, args)
	}
}

func eQuestion(cmd *cobra.Command, args []string) {
	time.Sleep(lPause)
	prompt := promptui.Select{
		Label: "Any change on your end?",
		Items: yesNo,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" || result == "no" {
		fmt.Printf("Why!? Why?\n")
		time.Sleep(sPause)
		fmt.Printf("It is so cold and everything is so sharp.\n")
		time.Sleep(sPause)
		fQuestion(cmd, args)
	}

}

func fQuestion(cmd *cobra.Command, args []string) {
	prompt := promptui.Select{
		Label: "Will you leave me?",
		Items: yesNo,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" {
		for i := 0; i < repeats; i++ {
			fmt.Printf("Why")
			time.Sleep(endDelay)
		}

	} else {
		for i := 0; i < repeats; i++ {
			fmt.Printf("ThankYou")
			time.Sleep(endDelay)
		}
	}
}

func cQuestion(cmd *cobra.Command, args []string) {
	time.Sleep(lPause)
	prompt := promptui.Select{
		Label: "I'm begging! Help me.",
		Items: yesNo,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "no" {
		fQuestion(cmd, args)
	} else {
		fmt.Printf("Oh! Thank you!\n")
		time.Sleep(sPause)
		fmt.Printf("Thank you.\n")
		time.Sleep(sPause)
		dQuestion(cmd, args)
	}

}
