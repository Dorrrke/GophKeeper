package models

type CardModel struct {
	Name    string
	Number  string
	Date    string
	CVVCode int
}

type LoginModel struct {
	Name     string
	Login    string
	Password string
}

type TextDataModel struct {
	Name string
	Data string
}

type BinaryDataModel struct {
	Name string
	Data []byte
}

type UserModel struct {
	UserID int64  `json:"u_id"`
	Login  string `json:"login"`
	Hash   string `json:"hash"`
}
