package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var addTodoMigration_0002 = &Migration{
	Number: 2,
	Name:   "Add todo",
	Forwards: func(db *gorm.DB) error {
		const addUserSQL = `
			CREATE TABLE todos(
 				id serial PRIMARY KEY,
 				name text NOT NULL,
 				done boolean NOT NULL,
 				user_id int not null,
 				created_at TIMESTAMP NOT NULL,
 				updated_at TIMESTAMP NOT NULL,
 				deleted_at TIMESTAMP,
 				
 				CONSTRAINT todos_user_id_fkey FOREIGN KEY (user_id)
      				REFERENCES users (id) MATCH SIMPLE
      				ON UPDATE NO ACTION ON DELETE CASCADE
			);
		`

		err := db.Exec(addUserSQL).Error
		return errors.Wrap(err, "unable to create todos table")
	},
}

func init() {
	Migrations = append(Migrations, addTodoMigration_0002)
}
