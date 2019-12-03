package mysql_repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hurtuh/indriver/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMysqlRepo_NewCandidate(t *testing.T) {
	conn := "inDriver:IsCool@tcp(0.0.0.0:3306)/interviews_test"
	db, err := sql.Open("mysql", conn)
	assert.NoError(t, err)

	repo := MysqlRepo{mysqlConn: db}

	candidate := &domain.Candidate{
		Name:        "test",
		LastName:    "testov",
		Interview:   time.Now(),
		Description: "test_desc",
	}

	err = repo.NewCandidate(candidate)
	assert.NoError(t, err)
}

func TestMysqlRepo_SelectAll(t *testing.T) {
	conn := "inDriver:IsCool@tcp(0.0.0.0:3306)/interviews_test"
	db, err := sql.Open("mysql", conn)
	assert.NoError(t, err)

	repo := MysqlRepo{mysqlConn: db}

	results, err := repo.SelectAll()
	assert.NoError(t, err)
	assert.NotNil(t, results)
}

func TestMysqlRepo_SelectByID(t *testing.T) {
	conn := "inDriver:IsCool@tcp(0.0.0.0:3306)/interviews_test"
	db, err := sql.Open("mysql", conn)
	assert.NoError(t, err)

	repo := MysqlRepo{mysqlConn: db}

	result, err := repo.SelectByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestMysqlRepo_DeleteCandidate(t *testing.T) {
	conn := "inDriver:IsCool@tcp(0.0.0.0:3306)/interviews_test"
	db, err := sql.Open("mysql", conn)
	assert.NoError(t, err)

	repo := MysqlRepo{mysqlConn: db}

	err = repo.DeleteCandidate(1)
	assert.NoError(t, err)
}

func TestMysqlRepo_EditDescription(t *testing.T) {
	conn := "inDriver:IsCool@tcp(0.0.0.0:3306)/interviews_test"
	db, err := sql.Open("mysql", conn)
	assert.NoError(t, err)

	repo := MysqlRepo{mysqlConn: db}

	err = repo.EditDescription(1, "test")
	assert.NoError(t, err)
}

func TestMysqlRepo_EditInterview(t *testing.T) {
	conn := "inDriver:IsCool@tcp(0.0.0.0:3306)/interviews_test"
	db, err := sql.Open("mysql", conn)
	assert.NoError(t, err)

	repo := MysqlRepo{mysqlConn: db}

	err = repo.EditInterview(1, time.Now())
	assert.NoError(t, err)
}
