package domain

import (
	"errors"
	"net/http"
	"time"
)

type Handlers interface {
	AddRecord(w http.ResponseWriter, r *http.Request)
	DelRecord(w http.ResponseWriter, r *http.Request)
	GetRecord(w http.ResponseWriter, r *http.Request)
	EditRecord(w http.ResponseWriter, r *http.Request)
}

type Service interface {
	NewCandidate(r *http.Request) (code int64, msg error, result interface{})
	EditCandidate(r *http.Request) (code int64, msg error, result interface{})
	DeleteCandidate(r *http.Request) (code int64, msg error, result interface{})
	GetCandidate(r *http.Request) (code int64, msg error, result interface{})
}

type Repository interface {
	SelectAll() ([]*Candidate, error)
	SelectByID(id uint) (*Candidate, error)
	EditDescription(candidateID uint, desc string) error
	EditInterview(candidateID uint, newTime time.Time) error
	NewCandidate(candidate *Candidate) error
	DeleteCandidate(candidateID uint) error
}

type Candidate struct {
	ID          uint      `json:"id"`
	Created     time.Time `json:"created"`
	Name        string    `json:"name"`
	LastName    string    `json:"lastname"`
	Interview   time.Time `json:"interview"`
	Description string    `json:"description"`
}

type EditCandidateReq struct {
	CandidateID uint      `json:"candidate_id"`
	NewTime     time.Time `json:"new_time"`
	Description string    `json:"description"`
}

type DeleteCandidateReq struct {
	CandidateID uint `json:"candidate_id"`
}

type GetCandidateReq struct {
	CandidateID   uint      `json:"candidate_id"`
	InterviewTime time.Time `json:"interview_time"`
}

const (
	CodeSuccess = iota
	CodeBadRequest
	CodeDatabaseError
)

var (
	ErrorWithUnmarshal = errors.New("error with unmarshal, please check request")
	ErrorWithDatabase  = errors.New("error with database, please contact developers")
)
