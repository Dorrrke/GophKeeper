package services

import (
	"context"
	"errors"

	errText "github.com/Dorrrke/GophKeeper/internal/domain/errors"
	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidPassword = errors.New(errText.InvalidPasswordError)

type Storage interface {
	SaveUser(ctx context.Context, user models.UserModel) (int64, error)
	GetUserHash(ctx context.Context, login string) (int64, string, error)
	SaveCard(ctx context.Context, card models.CardModel, uID int64) (int64, error)
	SaveLogin(ctx context.Context, login models.LoginModel, uID int64) (int64, error)
	SaveText(ctx context.Context, text models.TextDataModel, uID int64) (int64, error)
	SaveBin(ctx context.Context, binData models.BinaryDataModel, uID int64) (int64, error)
	GetAllCards(ctx context.Context, uID int64) ([]models.CardModel, error)
	GetAllBin(ctx context.Context, uID int64) ([]models.BinaryDataModel, error)
	GetAllLogins(ctx context.Context, uID int64) ([]models.LoginModel, error)
	GetAllTextData(ctx context.Context, uID int64) ([]models.TextDataModel, error)
	GetCardByName(ctx context.Context, name string, uID int64) (models.CardModel, error)
	GetLoginByName(ctx context.Context, name string, uID int64) (models.LoginModel, error)
	GetTextDataByName(ctx context.Context, name string, uID int64) (models.TextDataModel, error)
	GetBinByName(ctx context.Context, name string, uID int64) (models.BinaryDataModel, error)
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
		return -1, err
	}
	return uID, nil
}

func (kp *KeepService) LoginUser(login string, pass string) (models.UserModel, error) {
	uID, hashFromDB, err := kp.stor.GetUserHash(context.Background(), login)
	if err != nil {
		return models.UserModel{}, err
	}
	if !matchPass(pass, hashFromDB) {
		return models.UserModel{}, ErrInvalidPassword
	}
	return models.UserModel{
		UserID: uID,
		Login:  login,
		Hash:   pass,
	}, nil
}

func (kp *KeepService) SaveCard(card models.CardModel, uID int64) (int64, error) {
	cID, err := kp.stor.SaveCard(context.Background(), card, uID)
	if err != nil {
		return -1, err
	}
	return cID, nil
}

func (kp *KeepService) SaveLogin(loginData models.LoginModel, uID int64) (int64, error) {
	lID, err := kp.stor.SaveLogin(context.Background(), loginData, uID)
	if err != nil {
		return -1, err
	}
	return lID, nil
}

func (kp *KeepService) SaveTextData(textData models.TextDataModel, uID int64) (int64, error) {
	tID, err := kp.stor.SaveText(context.Background(), textData, uID)
	if err != nil {
		return -1, err
	}
	return tID, nil
}

func (kp *KeepService) SaveBinaryData(binData models.BinaryDataModel, uID int64) (int64, error) {
	bID, err := kp.stor.SaveBin(context.Background(), binData, uID)
	if err != nil {
		return -1, err
	}
	return bID, nil
}

func (kp *KeepService) GetBins(uID int64) ([]models.BinaryDataModel, error) {
	bins, err := kp.stor.GetAllBin(context.Background(), uID)
	if err != nil {
		return nil, err
	}
	return bins, nil
}

func (kp *KeepService) GetCards(uID int64) ([]models.CardModel, error) {
	cards, err := kp.stor.GetAllCards(context.Background(), uID)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (kp *KeepService) GetLogins(uID int64) ([]models.LoginModel, error) {
	logins, err := kp.stor.GetAllLogins(context.Background(), uID)
	if err != nil {
		return nil, err
	}
	return logins, nil
}

func (kp *KeepService) GetTextData(uID int64) ([]models.TextDataModel, error) {
	tData, err := kp.stor.GetAllTextData(context.Background(), uID)
	if err != nil {
		return nil, err
	}
	return tData, nil
}

func (kp *KeepService) GetCardByName(cName string, uID int64) (models.CardModel, error) {
	card, err := kp.stor.GetCardByName(context.Background(), cName, uID)
	if err != nil {
		return models.CardModel{}, err
	}
	return card, nil
}

func (kp *KeepService) GetLoginByName(lName string, uID int64) (models.LoginModel, error) {
	login, err := kp.stor.GetLoginByName(context.Background(), lName, uID)
	if err != nil {
		return models.LoginModel{}, err
	}
	return login, nil
}

func (kp *KeepService) GetTextDataByName(tName string, uID int64) (models.TextDataModel, error) {
	tData, err := kp.stor.GetTextDataByName(context.Background(), tName, uID)
	if err != nil {
		return models.TextDataModel{}, err
	}
	return tData, nil
}

func (kp *KeepService) GetBinByName(tName string, uID int64) (models.BinaryDataModel, error) {
	tData, err := kp.stor.GetBinByName(context.Background(), tName, uID)
	if err != nil {
		return models.BinaryDataModel{}, err
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
