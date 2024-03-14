package cmd

import (
	"fmt"

	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func saveTextData(name, data string) (string, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return "", err
	}
	keepService := services.New(storage)
	textData := models.TextDataModel{
		Name: name,
		Data: data,
	}
	cID, err := keepService.SaveTextData(textData)
	if err != nil {
		return "failed to save user card", err
	}

	return fmt.Sprintf("User card has been saved; ID: %v", cID), nil
}
