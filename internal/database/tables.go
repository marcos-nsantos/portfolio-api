package database

import "github.com/marcos-nsantos/portfolio-api/internal/entity"

func (c *Connection) CreateTables() error {
	return c.Client.AutoMigrate(
		&entity.Project{},
		&entity.User{},
	)
}

func (c *Connection) DropTables() error {
	return c.Client.Migrator().DropTable(
		&entity.Project{}, &entity.User{},
	)
}
