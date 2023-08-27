package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var defBoard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Наличие товара 💊"),
		tgbotapi.NewKeyboardButton("Получить скидку"),
		tgbotapi.NewKeyboardButton("Аптека открыта?"),
	),
)

var shopBoard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Заказать"),
	),
)

var deepShopBoard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Да 👍"),
		tgbotapi.NewKeyboardButton("Нет 👎"),
	),
)

var button = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад 🚫"),
	),
)
