package main

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch strings.ToLower(update.Message.Text) {
			case "привет!", "привет", "ghbdtn", "ку", "добрый", "доровки", "здарова", "здрасьте", "хай", "ку-ку", "приветуси", "хэй", "дороу":
				msg.Text = "Здравствуйте"

				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}
			case "аптека открыта?":
				msg.Text = "Да, мы ведь работаем 24/7"
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}

			case "получить скидку":
				msg.Text = "Пока что нет"
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}
				photo := tgbotapi.FilePath("photo") // could be `FileBytes`, `FileID`, `FileURL`, etc

				photoRequest := tgbotapi.NewPhoto(update.Message.Chat.ID, photo)
				m, err := bot.Send(photoRequest)
				if err != nil {
					fmt.Printf("err: %s\n", err)
				}

				for _, msgWithPhoto := range m.Photo {
					// that ID could be reused (no need to send the same photo to a server:
					// it already stores it), could be useful to store it somewhere
					fmt.Printf("i uploaded a photo, got its ID:%s\n", msgWithPhoto.FileID)
				}

			case "/start":
				msg.Text = "Здесь можно узнать наличие и стоимость товара"
				msg.ReplyMarkup = defBoard
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}

			case "наличие товара 💊":

				msg.Text = "Напишите название товара"
				msg.ReplyMarkup = button
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}

			case "назад 🚫":
				msg.Text = "Что-нибудь закажете?"
				msg.ReplyMarkup = defBoard
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}

			case "заказать":
				msg.Text = "Товар забронирован"
				msg.ReplyMarkup = deepShopBoard
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}

			case "да 👍":
				msg.Text = "Напишите что еще надо"
				msg.ReplyMarkup = button
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}
			case "нет 👎":
				msg.Text = "Приятных дальнейших заказов"
				msg.ReplyMarkup = defBoard
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}

			default:
				drugs = update.Message.Text
				sbis()
				if answToClient == "" {
					answToClient = "Такого препарата нет в базе"
				}
				msg.Text = answToClient
				_, err = bot.Send(msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}
				if answToClient != "Такого препарата нет в базе" {
					msg.ReplyMarkup = deepShopBoard
				}
				answToClient = ""

			}
		}

	}
}
