package sqlstorage

import (
	"log"

	"github.com/Garius6/blog/internal/models"
)

func TestClearDatabase() {
	s, err := New("test.db")
	if err != nil {
		log.Fatalf("Database connection error: %s", err.Error())
	}

	s.Db.Migrator().DropTable(&models.Article{})
}
