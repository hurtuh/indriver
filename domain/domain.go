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
	NewCandidate(candidate *Candidate) (code int64, msg error, result interface{})
	EditCandidate(req *EditCandidateReq) (code int64, msg error, result interface{})
	DeleteCandidate(req *DeleteCandidateReq) (code int64, msg error, result interface{})
	GetCandidate(req *GetCandidateReq) (code int64, msg error, result interface{})
}

type Repository interface {
	SelectAll() ([]*Candidate, error)
	SelectByID(id uint64) (*Candidate, error)
	EditDescription(candidateID uint64, desc string) error
	EditInterview(candidateID uint64, newTime time.Time) error
	NewCandidate(candidate *Candidate) error
	DeleteCandidate(candidateID uint64) error
}

type Candidate struct {
	ID          uint64    `json:"candidate_id"`
	Created     time.Time `json:"created"`
	Name        string    `json:"name"`
	LastName    string    `json:"lastname"`
	Interview   time.Time `json:"interview"`
	Description string    `json:"description"`
}

type EditCandidateReq struct {
	CandidateID uint64    `json:"candidate_id"`
	NewTime     time.Time `json:"new_time"`
	Description string    `json:"description"`
}

type DeleteCandidateReq struct {
	CandidateID uint64 `json:"candidate_id"`
}

type GetCandidateReq struct {
	CandidateID   uint64    `json:"candidate_id"`
	InterviewTime time.Time `json:"interview_time"`
}

const (
	CodeSuccess = iota
	CodeBadRequest
	CodeDatabaseError
)

var (
	ErrorWithDecoding = errors.New("error with unmarshal, please check request")
	ErrorWithDatabase = errors.New("error with database, please contact developers")
)
