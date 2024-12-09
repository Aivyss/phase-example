package repository

type AccountRepository interface {
	Insert(userID string, password string) error
}
