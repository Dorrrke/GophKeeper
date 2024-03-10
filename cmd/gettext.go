package cmd

import (
	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func getText(name string) (models.TextDataModel, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return models.TextDataModel{}, err
	}
	keepService := services.New(storage)
	tData, err := keepService.GetTextDataByName(name)
	if err != nil {
		return models.TextDataModel{}, err
	}

	return tData, nil
}
