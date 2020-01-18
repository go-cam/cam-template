package migrations

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/models"
)

func init() {
	m := new(m200118_092535_init_user)
	cam.App.AddMigration(m)
}

type m200118_092535_init_user struct {
	models.Migration
}

// up
func (migration *m200118_092535_init_user) Up() {
	migration.Exec(`
		CREATE TABLE user (
			id INTEGER NOT NULL AUTO_INCREMENT,
			name VARCHAR(64) NOT NULL COMMENT 'username',
			delete_at TIMESTAMP COMMENT 'delete timestamp',
			update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update timestamp',
			create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create timestamp',
			PRIMARY KEY(id)
		) ENGINE innodb CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
`)
}

// down
func (migration *m200118_092535_init_user) Down() {
}
