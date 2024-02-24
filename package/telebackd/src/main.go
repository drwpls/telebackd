package main

import (
	"asdas/command"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func serveCommand(timeout int, userCommand string, args string) (result string, err error) {
	switch userCommand {
	case "run":
		var runBinary string
		var runArgs []string
		parts := strings.Split(args, " ")
		switch len(parts) {
		case 0:
			runArgs = []string{}
		case 1:
			runBinary = parts[0]
			runArgs = []string{}
		default:
			runBinary = parts[0]
			runArgs = parts[1:]
		}
		output, err := command.Exec(timeout, runBinary, runArgs)
		if err != nil {
			return "", fmt.Errorf("exec error: %s", err)
		} else {
			if len(output) == 0 {
				output = "Success"
			}
			return output, nil

		}
	default:
		return "", fmt.Errorf("unknown command: %s", userCommand)
	}
}

func serveMessage(bot *tgbotapi.BotAPI, adminID int64, updates tgbotapi.UpdatesChannel) {
	var timeout = 5

	for update := range updates {
		if update.Message != nil { // If we got a message
			// check if message from admin
			if update.Message.From.ID != adminID {
				log.Printf("command from non-admin: %s", update.Message.From.UserName)
				message := tgbotapi.NewMessage(update.Message.Chat.ID, "You are not admin")
				bot.Send(message)
				continue
			}

			if update.Message.IsCommand() {
				log.Printf("[%s] command: %s %s", update.Message.From.UserName, update.Message.Command(), update.Message.CommandArguments())
				out, err := serveCommand(timeout, update.Message.Command(), update.Message.CommandArguments())

				if err != nil {
					log.Printf("command error: %s", err)
					message := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("command error: %s", err))
					_, err := bot.Send(message)
					if err != nil {
						log.Printf("send error: %s", err)
					}
				} else {
					split4k := func(s string) []string {
						var split []string
						for len(s) > 0 {
							if len(s) > 4096 {
								split = append(split, s[:4096])
								s = s[4096:]
							} else {
								split = append(split, s)
								break
							}
						}
						return split
					}

					if len(out) > 2*4096 {
						file := tgbotapi.FileBytes{Name: "output.txt", Bytes: []byte(out)}
						message := tgbotapi.NewDocument(update.Message.Chat.ID, file)
						_, err := bot.Send(message)
						if err != nil {
							log.Printf("send error: %s", err)
						}
					} else {
						for _, pout := range split4k(out) {
							message := tgbotapi.NewMessage(update.Message.Chat.ID, pout)
							_, err := bot.Send(message)
							if err != nil {
								log.Printf("send error: %s", err)
							}
						}
					}

				}
			}

		}
	}
}

func main() {
	var c config
	if err := c.Load(); err != nil {
		log.Fatalf("config load error: %s", err)
	}

	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("authorized on bot %s", bot.Self.UserName)
	bot.Send(tgbotapi.NewMessage(int64(c.AdminID), "bot started"))

	updateOptions := tgbotapi.NewUpdate(0)
	updateOptions.Timeout = 60
	updates := bot.GetUpdatesChan(updateOptions)

	serveMessage(bot, c.AdminID, updates)
}
