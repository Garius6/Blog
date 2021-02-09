package teststorage

import (
	"github.com/Garius6/blog/internal/models"
	"github.com/Garius6/blog/internal/storage"
)

type Storage struct {
	db                map[int]*models.Article
	articleRepository *ArticleRepository
}

func New() *Storage {
	return &Storage{
		db: make(map[int]*models.Article),
	}
}

func (s *Storage) Article() storage.ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}

	s.articleRepository = &ArticleRepository{
		storage: s,
		current: 1,
	}

	return s.articleRepository
}
