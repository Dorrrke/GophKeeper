package cmd

import (
	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func getTexts() ([]models.TextDataModel, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return nil, err
	}
	keepService := services.New(storage)
	texts, err := keepService.GetTextData()
	if err != nil {
		return nil, err
	}

	return texts, nil
}
