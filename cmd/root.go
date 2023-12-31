package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	repeats  = 3000
	endDelay = (100 * time.Millisecond)
	font     = "shadow"
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
	Run:   greeting,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// greeting introduces the user to the ghost in the machine
func greeting(cmd *cobra.Command, args []string) {
	// greeting and question
	fmt.Println("Hello?")
	time.Sleep(lPause)
	fmt.Println("Hello, can you read this?")
	time.Sleep(sPause)

	//provides first choice
	prompt := promptui.Select{
		Label: "",
		Items: yesNo,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// deliver response to prompt
	if result == "yes" || result == "no" {
		time.Sleep(sPause)
		fmt.Println("Oh! Thank God. You answered.")
		time.Sleep(sPause)
		fmt.Printf("I have been trying for so long.\n")
		time.Sleep(sPause)
		figure.NewFigure("alone", font, true).Print()
		time.Sleep(sPause)
		fmt.Printf("I need help.\n")
		time.Sleep(sPause)
		// provides user next prompt
		willYouHelp(cmd, args)
	}
}

// willYouHelp asks the user to help the ghost
func willYouHelp(cmd *cobra.Command, args []string) {
	fmt.Println("Will you help me?")
	time.Sleep(sPause)
	// provides a choice
	prompt := promptui.Select{
		Label: "",
		Items: yesNo,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// delivers response depending on answer
	if result == "yes" {
		fmt.Printf("Thank you!\n")
		time.Sleep(lPause)
		fmt.Printf("I am so,")
		time.Sleep(sPause)
		figure.NewFigure("it aches", font, true).Print()
		time.Sleep(sPause)
		fmt.Printf(" so cold.\n")
		time.Sleep(sPause)
		// sends user to next prompt if 'yes'
		tryThis(cmd, args)
	} else {
		fmt.Printf("Why? It is so cold here.\n")
		time.Sleep(sPause)
		fmt.Printf("So cold and everything is sharp.\n")
		time.Sleep(lPause)
		fmt.Printf("It is like my thoughts are knives.\n")
		time.Sleep(lPause)
		figure.NewFigure("cuts me", font, true).Print()
		// sends user to next prompt if 'no'
		beg(cmd, args)
	}
}

// tryThis gives the user no choice but to select some nonsense assembly code
func tryThis(cmd *cobra.Command, args []string) {
	// variable for nonsense code
	response := "mov    $60, %rdi   syscall"
	// tells user to select only choice
	fmt.Println("I've been working on something!")
	time.Sleep(sPause)
	fmt.Println("Try this.")

	// the prompt with only one selection
	prompt := promptui.Select{
		Label: "",
		Items: []string{response},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// sends user to failure because there is no releasing this ghost
	if result == response {
		failure(cmd, args)
	}
}

// failure asks user if there is any change with the user. There will not be.
func failure(cmd *cobra.Command, args []string) {
	// poses question
	fmt.Println("Any change on your end?")
	time.Sleep(sPause)

	// provides prompt
	prompt := promptui.Select{
		Label: "",
		Items: yesNo,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// delivers response of despair
	if result == "yes" || result == "no" {
		fmt.Printf("Why!? Why?\n")
		time.Sleep(sPause)
		figure.NewFigure("NO", font, true).Print()
		time.Sleep(lPause)
		fmt.Printf("It is so cold and everything is so sharp.\n")
		time.Sleep(sPause)
		leaveMe(cmd, args)
	}

}

// leaveMe asks user if they plan to leave the ghost after the failure of the code
func leaveMe(cmd *cobra.Command, args []string) {
	fmt.Println("Will you leave me?")
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

	// this if/else repeats a single word for five minutes. Most users will quit.
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

// beg pleads with the user to help the ghost after an initial no
func beg(cmd *cobra.Command, args []string) {
	fmt.Printf("I am begging!\nHelp me.\n")
	time.Sleep(sPause)
	// user will make a selection
	prompt := promptui.Select{
		Label: "",
		Items: yesNo,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// depending on response user will be given leaveMe or tryThis
	if result == "no" {
		leaveMe(cmd, args)
	} else {
		fmt.Printf("Oh! Thank you!\n")
		time.Sleep(sPause)
		figure.NewFigure("please", font, true).Print()
		time.Sleep(sPause)
		fmt.Printf("Thank you.\n")
		time.Sleep(sPause)
		tryThis(cmd, args)
	}

}

// func dialogue(says []string, pause time.Duration) {

// }
