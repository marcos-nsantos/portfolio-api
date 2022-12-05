package database

import "github.com/marcos-nsantos/portfolio-api/internal/entity"

func (c *Connection) CreateTables() error {
	return c.Client.AutoMigrate(
		&entity.User{},
		&entity.Project{},
	)
}

func (c *Connection) DropTables() error {
	return c.Client.Migrator().DropTable(
		&entity.User{},
		&entity.Project{},
	)
}
