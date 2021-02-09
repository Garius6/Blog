package storage

import (
	"github.com/Garius6/blog/internal/models"
)

type ArticleRepository interface {
	FindByID(ID int) (*models.Article, error)
	Create(title, information string) (int, error)
	GetAll() ([]models.Article, error)
}
