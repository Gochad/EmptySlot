package database

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{db: &sql.DB{}}
}

func (p *Postgres) Database() *sql.DB {
	return p.db
}

func (p *Postgres) Connect(driver string, dataSourceName string) error {
	var err error
	p.db, err = sql.Open(driver, dataSourceName)
	if err != nil {
		return err
	}

	err = p.db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to database")
	return nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) InsertData(data string) error {
	_, err := p.db.Exec("INSERT INTO table_name (column_name) VALUES (?)", data)
	return err
}

func (p *Postgres) GetData() ([]string, error) {
	rows, err := p.db.Query("SELECT column_name FROM table_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var data string
		err := rows.Scan(&data)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}
