package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (s *Storage) SaveCard(ctx context.Context, name, number, date string, cvv int) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO cards(name, number, date, cvv) VALUES(?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, name, number, date, cvv)
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

func (s *Storage) SaveLogin(ctx context.Context, name, login, password string) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO caloginsrds(name, login, password) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, name, login, password)
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

func (s *Storage) SaveText(ctx context.Context, name, data string) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO text_data(name, data) VALUES(?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, name, data)
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
