package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/anton-ag/javacode/internal/http"
	"github.com/anton-ag/javacode/internal/repo"
	"github.com/anton-ag/javacode/internal/service"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
)

var total = 1000
var connString = "postgresql://user:secret@localhost:5432/wallet"

type APITestSuite struct {
	suite.Suite

	uuid    string
	db      *sql.DB
	repo    *repo.Repo
	service *service.Service
	handler *http.Handler
}

func (s *APITestSuite) SetupSuite() {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		s.FailNow("Ошибка подключения к БД", err)
	}
	s.db = db

	s.initDependencies()
	err = s.populateDB(total)
	if err != nil {
		s.FailNow(err.Error())
	}
}

func (s *APITestSuite) TearDownSuite() {
	s.db.Close()
}

func (s *APITestSuite) initDependencies() {
	s.repo = repo.InitRepo(s.db)
	s.service = service.InitService(s.repo)
	s.handler = http.NewHandler(s.service)
}

func (s *APITestSuite) populateDB(data int) error {
	query := `CREATE TABLE IF NOT EXISTS wallet (
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	total integer DEFAULT 0
	);`
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка инициализации таблицы: %v", err)
	}

	var id string
	err = s.db.QueryRow("INSERT INTO wallet(total) VALUES ($1) RETURNING id", data).Scan(&id)
	if err != nil {
		return fmt.Errorf("ошибка заполнения таблицы: %v", err)
	}

	s.uuid = id
	log.Printf("uuid: %s", s.uuid)
	return nil
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}
