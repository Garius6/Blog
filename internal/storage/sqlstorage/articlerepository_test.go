package sqlstorage_test

import (
	"testing"

	"github.com/Garius6/blog/internal/storage"
	"github.com/Garius6/blog/internal/storage/sqlstorage"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	s, err := sqlstorage.New("test.db")
	if err != nil {
		t.Fatalf("Database connection error: %s", err.Error())
	}
	a := s.Article()
	defer sqlstorage.TestClearDatabase()

	testCases := []struct {
		name           string
		title          string
		information    string
		expectedResult int
	}{
		{
			name:           "valid data",
			title:          "Valid",
			information:    "Valid",
			expectedResult: 1,
		},
	}
	for _, tc := range testCases {
		res, _ := a.Create(tc.title, tc.information)
		assert.Equal(t, tc.expectedResult, res)
	}
}

func TestFindByID(t *testing.T) {
	s, err := sqlstorage.New("test.db")
	if err != nil {
		t.Fatalf("Database connection error: %s", err.Error())
	}
	a := s.Article()
	defer sqlstorage.TestClearDatabase()

	_, err = a.FindByID(1)
	assert.Equal(t, storage.ErrNotFound, err)

	a.Create("valid", "valid")

	_, err = a.FindByID(1)
	assert.Nil(t, err)

}
