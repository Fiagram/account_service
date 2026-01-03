package logic

type Role uint8

const (
	None Role = iota
	Admin
	member
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

type DeleteAccountParams struct {
	Username string
}
