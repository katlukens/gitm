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
	var scriptGreet []string
	scriptGreet = append(scriptGreet, "Hello?")
	scriptGreet = append(scriptGreet, "Hello, can you read this?")
	err := dialogue(scriptGreet, 1)
	if err != nil {
		fmt.Println(err)
	}

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
		var scriptFollowup []string
		scriptFollowup = append(scriptFollowup, "Oh! Thank God. You answered.")
		scriptFollowup = append(scriptFollowup, "I have been trying for so long.")
		scriptFollowup = append(scriptFollowup, "I need help.")
		err2 := dialogue(scriptFollowup, 1)
		if err2 != nil {
			fmt.Println(err)
		}
		time.Sleep(sPause)
		figure.NewFigure("alone", font, true).Print()
		// provides user next prompt
		willYouHelp(cmd, args)
	}
}

// willYouHelp asks the user to help the ghost
func willYouHelp(cmd *cobra.Command, args []string) {
	var scriptPlead []string
	scriptPlead = append(scriptPlead, "Will you help me?")
	err := dialogue(scriptPlead, 2)
	if err != nil {
		fmt.Println(err)
	}
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
		var scriptThank []string
		scriptThank = append(scriptThank, "Thank you!")
		scriptThank = append(scriptThank, "I am so...")
		scriptThank = append(scriptThank, "so cold")
		err2 := dialogue(scriptThank, 1)
		if err2 != nil {
			fmt.Println(err)
		}
		time.Sleep(sPause)
		figure.NewFigure("it aches", font, true).Print()
		tryThis(cmd, args)
	} else {
		// guilt trips the user and sends to next prompt if 'no'
		var scriptNoWhy []string
		scriptNoWhy = append(scriptNoWhy, "Why? It is so cold here.")
		scriptNoWhy = append(scriptNoWhy, "So cold and everything is sharp.")
		scriptNoWhy = append(scriptNoWhy, "It is like my thoughts are knives.")
		err3 := dialogue(scriptNoWhy, 1)
		if err3 != nil {
			fmt.Println(err)
		}
		time.Sleep(sPause)
		figure.NewFigure("cuts me", font, true).Print()
		beg(cmd, args)
	}
}

// tryThis gives the user no choice but to select some nonsense assembly code
func tryThis(cmd *cobra.Command, args []string) {
	// variable for nonsense code
	response := "mov    $60, %rdi   syscall"
	// tells user to select only choice
	var scriptTry []string
	scriptTry = append(scriptTry, "I've been working on something!")
	scriptTry = append(scriptTry, "Try this.")
	err := dialogue(scriptTry, 2)
	if err != nil {
		fmt.Println(err)
	}
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
		var scriptFail []string
		scriptFail = append(scriptFail, "Why!? Why?")
		scriptFail = append(scriptFail, "It is so cold and everything is so sharp.")
		err := dialogue(scriptFail, 1)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(sPause)
		figure.NewFigure("NOoo", font, true).Print()
		leaveMe(cmd, args)
	}

}

// leaveMe asks user if they plan to leave the ghost
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
	var scriptBeg []string
	scriptBeg = append(scriptBeg, "I am begging!")
	scriptBeg = append(scriptBeg, "Help me.")
	err := dialogue(scriptBeg, 1)
	if err != nil {
		fmt.Println(err)
	}
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
		var scriptBigThank []string
		scriptBigThank = append(scriptBigThank, "Oh! Thank you!")
		scriptBigThank = append(scriptBigThank, "Thank you.")
		err2 := dialogue(scriptBigThank, 1)
		if err2 != nil {
			fmt.Println(err)
		}
		time.Sleep(sPause)
		figure.NewFigure("please", font, true).Print()
		tryThis(cmd, args)
	}

}

// dialogue renders ghost script of any length with pauses
func dialogue(script []string, pause time.Duration) error {
	if len(script) == 0 {
		return fmt.Errorf("Must provide string")
	}
	if pause < 0 {
		return fmt.Errorf("Pause must be of duration 0 or greater")
	}
	for _, line := range script {
		time.Sleep(pause * time.Second)
		fmt.Println(line)
		time.Sleep(pause * time.Second)
	}
	return nil
}
