package main

import (
	"database/sql"

	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/logger"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// TODO: Инициализация конфига
	cfg := config.ReadConfig()

	// TODO: Инициализация логгера
	zlog := logger.SetupLogger(cfg.DebugFlag)

	zlog.Debug().Msg("Sucsess init logger and cfg")
	zlog.Info().Msg("Sucsess init logger and cfg")
	zlog.Error().Msg("Sucsess init logger and cfg")

	// TODO: Инициализация сервисов

	// TODO: Инициализация базы данных
	database, err := sql.Open("sqlite3", "./gophkeeper.db")
	if err != nil {
		panic(err)
	}

	st, err := database.Prepare("CREATE TABLE IF NOT EXISTS cards (id INTEGER PRIMARY KEY, name TEXT, number TEXT, date TEXT, cvv INTEGER)")
	if err != nil {
		panic(err)
	}
	st.Exec()

	// TODO: Инициализация сервера

	// TODO: Запуск серевер

	// TODO: ShutDown
}
