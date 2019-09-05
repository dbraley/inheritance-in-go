package model

type Credential interface {
}

type credentialData struct {
	id        int
	accountId int
	credType  string
	status    string
}
