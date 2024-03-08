package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Dorrrke/GophKeeper/internal/domain/models"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, user models.UserModel) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO users(login, hash) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, user.Login, user.Hash)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("user alredy exists: %w", err)
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) GetUserHash(ctx context.Context, login string) (string, error) {
	stmt, err := s.db.Prepare("SELECT hash FROM users WHERE login = ?")
	if err != nil {
		return "", err
	}
	row := stmt.QueryRowContext(ctx, login)
	var hash string
	err = row.Scan(&hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found: %w", err)
		}
		return "", err
	}
	return hash, nil
}

func (s *Storage) SaveCard(ctx context.Context, card models.CardModel) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO cards(name, number, date, cvv) VALUES(?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, card.Name, card.Number, card.Date, card.CVVCode)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("card alredy exists: %w", err)
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) SaveLogin(ctx context.Context, loginData models.LoginModel) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO logins(name, login, password) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, loginData.Name, loginData.Login, loginData.Password)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("login alredy exists: %w", err)
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) SaveText(ctx context.Context, textData models.TextDataModel) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO text_data(name, data) VALUES(?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, textData.Name, textData.Data)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("text alredy exists: %w", err)
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) GetAllCards(ctx context.Context) ([]models.CardModel, error) {
	stmt, err := s.db.Prepare("SELECT name, number, date, cvv FROM cards")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []models.CardModel
	for rows.Next() {
		var card models.CardModel
		err := rows.Scan(&card.Name, &card.Number, &card.Date, &card.CVVCode)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (s *Storage) GetAllLogins(ctx context.Context) ([]models.LoginModel, error) {
	stmt, err := s.db.Prepare("SELECT name, login, password FROM logins")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var logins []models.LoginModel
	for rows.Next() {
		var login models.LoginModel
		err := rows.Scan(&login.Name, &login.Login, &login.Password)
		if err != nil {
			return nil, err
		}
		logins = append(logins, login)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return logins, nil
}

func (s *Storage) GetAllTextData(ctx context.Context) ([]models.TextDataModel, error) {
	stmt, err := s.db.Prepare("SELECT name, data FROM text_data")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tData []models.TextDataModel
	for rows.Next() {
		var data models.TextDataModel
		err := rows.Scan(&data.Name, &data.Data)
		if err != nil {
			return nil, err
		}
		tData = append(tData, data)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tData, nil
}

func (s *Storage) GetCardByName(ctx context.Context, name string) (models.CardModel, error) {
	stmt, err := s.db.Prepare("SELECT name, number, date, cvv FROM cards WHERE name = ?")
	if err != nil {
		return models.CardModel{}, err
	}

	row := stmt.QueryRowContext(ctx, name)
	var card models.CardModel
	err = row.Scan(&card.Name, &card.Number, &card.Date, &card.CVVCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.CardModel{}, fmt.Errorf("card not found") // TODO: Вынести ошибки в отдельный пакет
		}
		return models.CardModel{}, err
	}
	return card, nil
}

func (s *Storage) GetLoginByName(ctx context.Context, name string) (models.LoginModel, error) {
	stmt, err := s.db.Prepare("SELECT name, login, password FROM logins WHERE name = ?")
	if err != nil {
		return models.LoginModel{}, err
	}

	row := stmt.QueryRowContext(ctx, name)
	var login models.LoginModel
	err = row.Scan(&login.Name, &login.Login, &login.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.LoginModel{}, fmt.Errorf("login not found") // TODO: Вынести ошибки в отдельный пакет
		}
		return models.LoginModel{}, err
	}
	return login, nil
}

func (s *Storage) GetTextDataByName(ctx context.Context, name string) (models.TextDataModel, error) {
	stmt, err := s.db.Prepare("SELECT name, data FROM text_data WHERE name = ?")
	if err != nil {
		return models.TextDataModel{}, err
	}

	row := stmt.QueryRowContext(ctx, name)
	var data models.TextDataModel
	err = row.Scan(&data.Name, &data.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.TextDataModel{}, fmt.Errorf("text not found") // TODO: Вынести ошибки в отдельный пакет
		}
		return models.TextDataModel{}, err
	}
	return data, nil
}
