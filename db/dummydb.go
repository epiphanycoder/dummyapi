package db

import "dummyapi/storage"

type DummyDb struct {
	authUserDb *AuthUserDb
	storeDb    *storage.StoreDb
}
