package models

import (
	"database/sql"
	"errors"
	"time"
)

type Article struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	// Author  string
}

type ArticleModel struct {
	DB *sql.DB
}

func (m *ArticleModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO articles (title, content, created)
	VALUES($1, $2);`

	result, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ArticleModel) Get(id int) (*Article, error) {

	stmt := `SELECT id, title, content, created_at FROM articles WHERE id = $1;`

	s := &Article{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *ArticleModel) Latest() ([]*Article, error) {

	stmt := `SELECT id, title, content, created_at FROM articles
	ORDER BY id DESC LIMIT 10;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []*Article{}

	for rows.Next() {
		s := &Article{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created)
		if err != nil {
			return nil, err
		}

		articles = append(articles, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
