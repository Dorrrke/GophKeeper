package services

import (
	"context"
	"fmt"

	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	SaveUser(ctx context.Context, user models.UserModel) (int64, error)
	GetUserHash(ctx context.Context, login string) (string, error)
	SaveCard(ctx context.Context, card models.CardModel) (int64, error)
	SaveLogin(ctx context.Context, login models.LoginModel) (int64, error)
	SaveText(ctx context.Context, text models.TextDataModel) (int64, error)
}

type KeepService struct {
	stor Storage
	zlog *zerolog.Logger
}

func New(stor Storage, log *zerolog.Logger) *KeepService {
	return &KeepService{
		stor: stor,
		zlog: log,
	}
}

func (kp *KeepService) RegisterUser(login string, pass string) error {
	hash, err := hashPass(pass)
	if err != nil {
		return err
	}
	uID, err := kp.stor.SaveUser(context.Background(), models.UserModel{
		Login: login,
		Hash:  hash,
	})
	if err != nil {
		// TODO: Обработка ошибки гадичия такого пользователя.
		kp.zlog.Error().Err(err).Msg("save user error")
		return err
	}
	kp.zlog.Debug().Int64("uId", uID).Msg("User was saved")
	return nil
}

func (kp *KeepService) LoginUser(login string, pass string) error {
	hashFromDB, err := kp.stor.GetUserHash(context.Background(), login)
	if err != nil {
		// TODO: Обработка ошибки отсутствия пользователя.
		kp.zlog.Error().Err(err).Msg("get user password hash from db error")
		return err
	}
	if !matchPass(pass, hashFromDB) {
		kp.zlog.Error().Err(err).Msg("invalid login or password")
		return fmt.Errorf("invalid login or password")
	}
	return nil
}

func (kp *KeepService) SaveCard(card models.CardModel) error {
	cID, err := kp.stor.SaveCard(context.Background(), card)
	if err != nil {
		//TODO: Проверка ошибки на наличие такой карты в бд
		kp.zlog.Error().Err(err).Msg("save card error")
		return err
	}
	kp.zlog.Debug().Int64("cId", cID).Msg("saved card id")
	return nil
}

func (kp *KeepService) SaveLogin(loginData models.LoginModel) error {
	cID, err := kp.stor.SaveLogin(context.Background(), loginData)
	if err != nil {
		//TODO: Проверка ошибки на наличие таких данных в бд
		kp.zlog.Error().Err(err).Msg("save loginData error")
		return err
	}
	kp.zlog.Debug().Int64("lId", cID).Msg("saved login data id")
	return nil
}

func (kp *KeepService) SaveTextData(textData models.TextDataModel) error {
	cID, err := kp.stor.SaveText(context.Background(), textData)
	if err != nil {
		//TODO: Проверка ошибки на наличие таких данных в бд
		kp.zlog.Error().Err(err).Msg("save textData error")
		return err
	}
	kp.zlog.Debug().Int64("tId", cID).Msg("saved text data id")
	return nil
}

func (kp *KeepService) SaveBinaryData(card models.BinaryDataModel) error {

	return nil
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
