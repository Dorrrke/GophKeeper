package cmd

import (
	"fmt"

	"github.com/Dorrrke/GophKeeper/internal/config"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/Dorrrke/GophKeeper/internal/services"
	"github.com/Dorrrke/GophKeeper/internal/storage"
)

func saveAuthData(name, login, password string) (string, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return "", err
	}
	keepService := services.New(storage)
	authData := models.LoginModel{
		Name:     name,
		Login:    login,
		Password: password,
	}
	cID, err := keepService.SaveLogin(authData)
	if err != nil {
		return "failed to save user auth data", err
	}

	return fmt.Sprintf("User auth data has been saved; ID: %v", cID), nil
}
