package services

import (
	"context"
	"fmt"

	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	SaveUser(ctx context.Context, user models.UserModel) (int64, error)
	GetUserHash(ctx context.Context, login string) (string, error)
	SaveCard(ctx context.Context, card models.CardModel) (int64, error)
	SaveLogin(ctx context.Context, login models.LoginModel) (int64, error)
	SaveText(ctx context.Context, text models.TextDataModel) (int64, error)
	GetAllCards(ctx context.Context) ([]models.CardModel, error)
	GetAllLogins(ctx context.Context) ([]models.LoginModel, error)
	GetAllTextData(ctx context.Context) ([]models.TextDataModel, error)
	GetCardByName(ctx context.Context, name string) (models.CardModel, error)
	GetLoginByName(ctx context.Context, name string) (models.LoginModel, error)
	GetTextDataByName(ctx context.Context, name string) (models.TextDataModel, error)
}

type KeepService struct {
	stor Storage
}

func New(stor Storage) *KeepService {
	return &KeepService{
		stor: stor,
	}
}

func (kp *KeepService) RegisterUser(login string, pass string) (int64, error) {
	hash, err := hashPass(pass)
	if err != nil {
		return -1, err
	}
	uID, err := kp.stor.SaveUser(context.Background(), models.UserModel{
		Login: login,
		Hash:  hash,
	})
	if err != nil {
		// TODO: Обработка ошибки гадичия такого пользователя.
		return -1, err
	}
	return uID, nil
}

func (kp *KeepService) LoginUser(login string, pass string) error {
	hashFromDB, err := kp.stor.GetUserHash(context.Background(), login)
	if err != nil {
		// TODO: Обработка ошибки отсутствия пользователя.
		return err
	}
	if !matchPass(pass, hashFromDB) {
		return fmt.Errorf("invalid login or password")
	}
	return nil
}

func (kp *KeepService) SaveCard(card models.CardModel) (int64, error) {
	cID, err := kp.stor.SaveCard(context.Background(), card)
	if err != nil {
		//TODO: Проверка ошибки на наличие такой карты в бд
		return -1, err
	}
	return cID, nil
}

func (kp *KeepService) SaveLogin(loginData models.LoginModel) (int64, error) {
	lID, err := kp.stor.SaveLogin(context.Background(), loginData)
	if err != nil {
		//TODO: Проверка ошибки на наличие таких данных в бд
		return -1, err
	}
	return lID, nil
}

func (kp *KeepService) SaveTextData(textData models.TextDataModel) (int64, error) {
	tID, err := kp.stor.SaveText(context.Background(), textData)
	if err != nil {
		//TODO: Проверка ошибки на наличие таких данных в бд
		return -1, err
	}
	return tID, nil
}

func (kp *KeepService) SaveBinaryData(card models.BinaryDataModel) error {

	return nil
}

func (kp *KeepService) GetCards() ([]models.CardModel, error) {
	cards, err := kp.stor.GetAllCards(context.Background())
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (kp *KeepService) GetLogins() ([]models.LoginModel, error) {
	logins, err := kp.stor.GetAllLogins(context.Background())
	if err != nil {
		return nil, err
	}
	return logins, nil
}

func (kp *KeepService) GetTextData() ([]models.TextDataModel, error) {
	tData, err := kp.stor.GetAllTextData(context.Background())
	if err != nil {
		return nil, err
	}
	return tData, nil
}

func (kp *KeepService) GetCardByName(cName string) (models.CardModel, error) {
	card, err := kp.stor.GetCardByName(context.Background(), cName)
	if err != nil {
		return models.CardModel{}, err
	}
	return card, nil
}

func (kp *KeepService) GetLoginByName(lName string) (models.LoginModel, error) {
	login, err := kp.stor.GetLoginByName(context.Background(), lName)
	if err != nil {
		return models.LoginModel{}, err
	}
	return login, nil
}

func (kp *KeepService) GetTextDataByName(tName string) (models.TextDataModel, error) {
	tData, err := kp.stor.GetTextDataByName(context.Background(), tName)
	if err != nil {
		return models.TextDataModel{}, err
	}
	return tData, nil
}

func hashPass(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func matchPass(pass string, hashFromDB string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashFromDB), []byte(pass))
	return err == nil
}
