package client

import (
	"context"
	"strconv"

	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	gophkeeperv1 "github.com/Dorrrke/goph-keeper-proto/gen/go/gophkeeper"
	"google.golang.org/grpc"
)

type KeeperClient struct {
	client gophkeeperv1.GophKeeperClient
	conn   *grpc.ClientConn
}

func New(ctx context.Context, addr string) (*KeeperClient, error) {
	conn, err := grpc.DialContext(ctx, addr)
	if err != nil {
		return nil, err
	}
	client := gophkeeperv1.NewGophKeeperClient(conn)
	return &KeeperClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *KeeperClient) Register(ctx context.Context, login, password string) error {
	_, err := c.client.SignUp(ctx, &gophkeeperv1.SignUpRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *KeeperClient) Login(ctx context.Context, login, password string) error {
	_, err := c.client.SignIn(ctx, &gophkeeperv1.SingInRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *KeeperClient) Sync(ctx context.Context, model models.SyncModel) (models.SyncModel, error) {
	protoModel := modelToProtoModel(model)
	res, err := c.client.SyncDB(ctx, &gophkeeperv1.SyncDBRequest{
		Auth:  protoModel.Auth,
		Bins:  protoModel.Bins,
		Cards: protoModel.Cards,
		Texts: protoModel.Texts,
	})
	if err != nil {
		return models.SyncModel{}, err
	}
	resModel, err := protoModelToModel(models.ProtoSyncModel{
		Cards: res.Cards,
		Auth:  res.Auth,
		Texts: res.Texts,
		Bins:  res.Bins,
	})
	if err != nil {
		return models.SyncModel{}, err
	}
	return resModel, nil
}

func modelToProtoModel(model models.SyncModel) models.ProtoSyncModel {
	var pModel models.ProtoSyncModel
	for _, data := range model.Bins {
		bin := &gophkeeperv1.SyncBinData{
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		pModel.Bins = append(pModel.Bins, bin)
	}
	for _, data := range model.Auth {
		auth := &gophkeeperv1.SyncAuth{
			Name:     data.Name,
			Login:    data.Login,
			Password: data.Password,
			Deleted:  data.Deleted,
			Updated:  data.Updated,
		}
		pModel.Auth = append(pModel.Auth, auth)
	}
	for _, data := range model.Cards {
		card := &gophkeeperv1.SyncCard{
			Name:    data.Name,
			Number:  data.Number,
			Date:    data.Date,
			Cvv:     strconv.Itoa(data.CVVCode),
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		pModel.Cards = append(pModel.Cards, card)
	}
	for _, data := range model.Texts {
		text := &gophkeeperv1.SyncText{
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		pModel.Texts = append(pModel.Texts, text)
	}

	return pModel
}

func protoModelToModel(model models.ProtoSyncModel) (models.SyncModel, error) {
	var sModel models.SyncModel
	for _, data := range model.Bins {
		bin := models.SyncBinaryDataModel{
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		sModel.Bins = append(sModel.Bins, bin)
	}
	for _, data := range model.Auth {
		auth := models.SyncLoginModel{
			Name:     data.Name,
			Login:    data.Login,
			Password: data.Password,
			Deleted:  data.Deleted,
			Updated:  data.Updated,
		}
		sModel.Auth = append(sModel.Auth, auth)
	}
	for _, data := range model.Cards {
		cvv, err := strconv.Atoi(data.Cvv)
		if err != nil {
			return models.SyncModel{}, err
		}
		card := models.SyncCardModel{
			Name:    data.Name,
			Number:  data.Number,
			Date:    data.Date,
			CVVCode: cvv,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		sModel.Cards = append(sModel.Cards, card)
	}
	for _, data := range model.Texts {
		text := models.SyncTextDataModel{
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		sModel.Texts = append(sModel.Texts, text)
	}

	return sModel, nil
}
