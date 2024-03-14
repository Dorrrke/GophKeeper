package services

import "github.com/Dorrrke/GophKeeper/internal/domain/models"

type Storage interface{}

type KeepService struct {
	stor Storage
}

func New(stor Storage) *KeepService {
	return &KeepService{
		stor: stor,
	}
}

func RegisterUser(login string, pass string) error {

	return nil
}

func LoginUser(login string, pass string) error {

	return nil
}

func SaveCard(card models.CardModel) error {

	return nil
}

func SaveLogin(card models.LoginModel) error {

	return nil
}

func SaveTextData(card models.TextDataModel) error {

	return nil
}

func SaveBinaryData(card models.BinaryDataModel) error {

	return nil
}
