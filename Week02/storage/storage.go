package storage

import (
	"database/sql"
)

// BlogStorage ...
type BlogStorage struct {
	blogDAO *BlogDAO
}

// BlogDAO ...
type BlogDAO struct {
}

// Blog ...
type Blog struct {
	Title string
	Text  string
}

// NewBlogStorage ...
func NewBlogStorage(dao *BlogDAO) *BlogStorage {
	return &BlogStorage{
		blogDAO: dao,
	}
}

// GetBlogsForUserID ...
func (storage *BlogStorage) GetBlogsForUserID(userID int64) ([]*Blog, error) {
	// storage layer can add additional logging and combine the data from the DAO layer
	// and compse the data for the downstream service handler
	return storage.blogDAO.GetBlogsForUserID(userID)
}

// GetBlogsForUserID ...
func (dao *BlogDAO) GetBlogsForUserID(userID int64) ([]*Blog, error) {
	return nil, sql.ErrNoRows
}
