package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Azizbek-Qodirov/hr_platform_evaluation_service/config"
	st "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	Evaluations 	st.EvaluationStorage
	Guides st.GuideStorage
	notifications st.NotificationStorage
}

func NewPostgresStorage() (st.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, nil

}

func (s *Storage) Evaluation() st.EvaluationStorage{
	if s.Evaluations == nil {
		s.Evaluations = &EvaluationStorage{s.Db}
	}
	return s.Evaluations
}


func (s *Storage) Guide() st.GuideStorage{
	if s.Guides == nil {
		s.Guides = &GuideStorage{s.Db}
	}
	return s.Guides
}


func (s *Storage) Notification() st.NotificationStorage{
	if s.notifications == nil {
		s.notifications = &NotificationStorage{s.Db}
	}
	return s.notifications
}




