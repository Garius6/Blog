package teststorage

import (
	"fmt"

	"github.com/Garius6/blog/internal/models"
)

type ArticleRepository struct {
	storage *Storage
	current int
}

func (a *ArticleRepository) FindByID(ID int) (*models.Article, error) {
	if _, ok := a.storage.db[ID]; !ok {
		return nil, fmt.Errorf("Not found")
	}

	return a.storage.db[ID], nil
}

func (a *ArticleRepository) Create(title, information string) (int, error) {
	a.storage.db[a.current] = &models.Article{Title: title, Information: information}
	a.current++
	return a.current, nil
}

func (a *ArticleRepository) GetAll() ([]models.Article, error) {
	res := make([]models.Article, 0)
	for _, value := range a.storage.db {
		res = append(res, *value)
	}

	return res, nil
}
