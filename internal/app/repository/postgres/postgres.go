package postgres

import (
	"database/sql"
	"fmt"
	"music-player-api/internal/app/repository"
	"sync"

	_ "github.com/lib/pq"
)

type PGPool struct {
	pool    *sql.DB
	maxCons uint
	numCons uint
	mutex   *sync.Mutex
}

func NewPGPool(dsn string, maxCons uint) (repository.Connection, error) {
	pool, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error while creating a connection pool to Postgres: %v", err)
	}

	pool.SetMaxOpenConns(int(maxCons))

	if err := pool.Ping(); err != nil {
		return nil, fmt.Errorf("pg ping check error: %v", err)
	}

	return &PGPool{
		pool:    pool,
		maxCons: maxCons,
		numCons: 0,
		mutex:   &sync.Mutex{},
	}, nil
}

func (p *PGPool) GetConnection() (*sql.DB, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.numCons == p.maxCons {
		return nil, fmt.Errorf("pg connection pool is full")
	}

	p.numCons++

	return p.pool, nil
}

func (p *PGPool) ReleaseConnection(conn *sql.DB) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.numCons--
}
