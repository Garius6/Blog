package storage

type Storage interface {
	Article() ArticleRepository
}
