package sqlstorage

import (
	"github.com/Garius6/blog/internal/models"
	"github.com/Garius6/blog/internal/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	Db                *gorm.DB
	articleRepository *ArticleRepository
}

func New(databaseURL string) (*Storage, error) {
	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Article{
		Title:       "",
		Information: "",
	})

	return &Storage{
		Db: db,
	}, nil
}

func (s Storage) Article() storage.ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}

	s.articleRepository = &ArticleRepository{
		storage: &s,
	}

	return s.articleRepository
}
