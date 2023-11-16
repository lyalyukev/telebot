package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var rootCmd = &cobra.Command{
	Use:   "kbot",
	Short: "Telebot",
	Long:  "This is a simple command-line application written in Go using Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		pref := telebot.Settings{
			Token:  os.Getenv("TELE_TOKEN"),
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}

		kbot, err := telebot.NewBot(pref)
		if err != nil {
			log.Fatal(err)
			return
		}

		kbot.Handle("/hello", func(c telebot.Context) error {
			return c.Send("Hello!")
		})

		kbot.Start()

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
