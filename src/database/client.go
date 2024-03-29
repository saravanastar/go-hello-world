package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbClient interface {
	Ready() bool
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DbClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		5432,
		"disable",
	)
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		return nil, error
	}
	client := Client{DB: db}

	return client, nil
}

func (client Client) Ready() bool {
	var ready string
	transaction := client.DB.Raw("select 1 as ready").Scan(&ready)
	if transaction.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
