package cmd

import (
	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func getCard(name string) (models.CardModel, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return models.CardModel{}, err
	}
	keepService := services.New(storage)
	card, err := keepService.GetCardByName(name)
	if err != nil {
		return models.CardModel{}, err
	}

	return card, nil
}
