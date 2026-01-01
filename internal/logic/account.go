package logic

import "context"

type Role string

const (
	NoneRole   Role = "none"
	AdminRole  Role = "admin"
	NormalRole Role = "normal"
)

type AccountInfo struct {
	Username    string
	Fullname    string
	Email       string
	PhoneNumber string
	Role        Role
}

type CreateAccountParams struct {
	AccountInfo AccountInfo
	Password    string
}

type CreateAccountOutput struct {
	AccountId uint64
}

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error)
	// CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionOutput, error)
}

type account struct {
	hashLogic Hash
}

func NewAccount(
	hashLogic Hash,
) Account {
	return &account{
		hashLogic}
}

func (a *account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error) {
	// tx, err :=
	return CreateAccountOutput{}, nil
}

// func (a *account) CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionOutput, error) {
// 	return CreateSessionOutput{}, nil
// }
