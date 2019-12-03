package mysql_repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hurtuh/indriver/domain"
	"time"
)

type MysqlRepo struct {
	mysqlConn *sql.DB
}

func NewMysqlRepo(host, port, user, pass, dbName string) (*MysqlRepo, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return &MysqlRepo{mysqlConn:db}, nil
}

func (m *MysqlRepo) SelectAll() ([]*domain.Candidate, error) {
	query := "SELECT id, created, name, lastname, interview, description FROM candidate"
	rows, err := m.mysqlConn.Query(query)
	if err != nil {
		return nil, err
	}

	candidates := make([]*domain.Candidate, 0)
	for rows.Next() {
		candidate := new(domain.Candidate)
		err = rows.Scan(&candidate.ID, &candidate.Created, &candidate.Name, &candidate.LastName, &candidate.Interview, &candidate.Description)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, candidate)
	}

	return candidates, nil
}


func (m *MysqlRepo) SelectByID(id uint) (*domain.Candidate, error) {
	query := "SELECT id, created, name, lastname, interview, description FROM candidate WHERE id = ?"
	candidate := new(domain.Candidate)

	err := m.mysqlConn.QueryRow(query, id).Scan(&candidate.ID, &candidate.Created, &candidate.Name, &candidate.LastName, &candidate.Interview, &candidate.Description)
	if err != nil {
		return nil, err
	}

	return candidate, nil
}

func (m *MysqlRepo) EditDescription(candidateID uint, desc string) error {
	query := "UPDATE candidate SET description = ? WHERE id = ?"
	stmt, err := m.mysqlConn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(query, desc, candidateID)
	return err
}

func (m *MysqlRepo) EditInterview(candidateID uint, newTime time.Time) error {
	query := "UPDATE candidate SET interview = ? WHERE id = ?"
	stmt, err := m.mysqlConn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(query, newTime, candidateID)
	return err
}

func (m *MysqlRepo) NewCandidate(candidate *domain.Candidate) error {
	query := "INSERT INTO candidate (name, lastname, interview, description) VALUES (?, ?, ?, ?)"

	stmt, err := m.mysqlConn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(candidate.Name, candidate.LastName, candidate.Interview, candidate.Description)
	return err
}

func (m *MysqlRepo) DeleteCandidate(candidateID uint) error {
	query := "DELETE FROM candidates WHERE id = ?"

	stmt, err := m.mysqlConn.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(candidateID)
	return err
}
