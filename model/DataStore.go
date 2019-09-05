package model

type DataStore struct {
	Accounts      []Account
	nextAccountId int
}

func (db *DataStore) init() {
	if db.Accounts == nil {
		db.Accounts = make([]Account, 0)
	}
}

func (db *DataStore) addAccount() Account {
	id := db.nextAccountId
	db.nextAccountId += 1
	newAccount := accountData{
		id: id,
	}
	return newAccount
}
