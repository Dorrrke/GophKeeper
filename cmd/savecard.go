package cmd

import (
	"fmt"
	"strconv"

	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func saveUserCard(cardName, number, date string, cvv string) (string, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return "", err
	}
	cvvCode, err := strconv.Atoi(cvv)
	if err != nil {
		return "", err
	}
	keepService := services.New(storage)
	card := models.CardModel{
		Name:    cardName,
		Number:  number,
		Date:    date,
		CVVCode: cvvCode,
	}
	cID, err := keepService.SaveCard(card)
	if err != nil {
		return "failed to save user card", err
	}

	return fmt.Sprintf("User card has been saved; ID: %v", cID), nil
}
