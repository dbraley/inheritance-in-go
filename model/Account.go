package model

type Account interface {
	getId() int
	getEmail() string
	isAdmin() bool
	getCreds() []Credential
}

type accountData struct {
	id    int
	email string
	admin bool
}

func (account accountData) getId() int {
	return account.id
}

func (account accountData) getEmail() string {
	return account.email
}

func (account accountData) isAdmin() bool {
	return account.admin
}

func (account accountData) getCreds() []Credential {
	return nil
}
