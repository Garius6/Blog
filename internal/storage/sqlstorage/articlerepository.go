package sqlstorage

import (
	"github.com/Garius6/blog/internal/models"
)

type ArticleRepository struct {
	storage *Storage
}

func (a *ArticleRepository) FindByID(ID int) (*models.Article, error) {
	article := &models.Article{}
	i := a.storage.Db.First(article, ID)
	if i.Error != nil {
		return nil, i.Error
	}

	return article, nil
}

func (a *ArticleRepository) Create(title, information string) (int, error) {
	article := &models.Article{Title: title, Information: information}
	result := a.storage.Db.Select("Title", "Information").Create(article)
	if result.Error != nil {
		return -1, result.Error
	}

	return article.ID, nil
}

func (a *ArticleRepository) GetAll() ([]models.Article, error) {
	articles := make([]models.Article, 0)
	if res := a.storage.Db.Find(&articles); res.Error != nil {
		return nil, res.Error
	}

	return articles, nil
}
