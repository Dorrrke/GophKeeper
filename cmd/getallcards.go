package cmd

import (
	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func getCards() ([]models.CardModel, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return nil, err
	}
	keepService := services.New(storage)
	cards, err := keepService.GetCards()
	if err != nil {
		return nil, err
	}

	return cards, nil
}
