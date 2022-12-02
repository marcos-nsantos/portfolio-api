package database

import "github.com/marcos-nsantos/portfolio-api/internal/entity"

func (c *Connection) CreateTables() error {
	return c.Client.AutoMigrate(&entity.Project{})
}
