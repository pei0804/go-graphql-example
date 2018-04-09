package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Name string `json:"name"`
}

func (a *Account) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(a.Model.ID)
}

func (a *Account) NAME(ctx context.Context) *string {
	return &a.Name
}

func (db *DB) getAccount(ctx context.Context, id int32) (*Account, error) {
	var a Account
	err := db.DB.First(&a, id).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (db *DB) addAccount(ctx context.Context, input accountInput) (*Account, error) {
	a := Account{
		Name: input.Name,
	}
	err := db.DB.Create(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (db *DB) deleteAccount(ctx context.Context, accountID int32) (*bool, error) {
	var a Account
	err := db.DB.First(&a, accountID).Error
	if err != nil {
		return nil, err
	}
	err = db.DB.Delete(&a).Error
	if err != nil {
		return nil, err
	}
	return boolP(true), err
}

func (db *DB) updateAccount(ctx context.Context, args *accountInput) (*Account, error) {
	var a Account
	err := db.DB.First(&a, args.ID).Error
	if err != nil {
		return nil, err
	}
	updated := Account{
		Name: args.Name,
	}
	err = db.DB.Model(&a).Updates(updated).Error
	if err != nil {
		return nil, err
	}
	err = db.DB.First(&a, args.ID).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}
