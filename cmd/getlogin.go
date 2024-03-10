package cmd

import (
	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func getLogin(name string) (models.LoginModel, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return models.LoginModel{}, err
	}
	keepService := services.New(storage)
	auth, err := keepService.GetLoginByName(name)
	if err != nil {
		return models.LoginModel{}, err
	}

	return auth, nil
}
