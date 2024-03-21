package models

import (
	"context"
	"errors"
	"fmt"

	"strconv"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	// DB *sql.DB
	DB *pgxpool.Pool
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) error {

	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES (
		$1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + $3::INTERVAL);`

	expiresToString := strconv.Itoa(expires) + "DAYS"

	commTag, err := m.DB.Exec(context.Background(), stmt, title, content, expiresToString)

	if err != nil {
		return err
	}

	fmt.Println(commTag)

	return nil
}

// This will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE id=$1`

	row := m.DB.QueryRow(context.Background(), stmt, id)

	// fmt.Println(row)

	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrorNoRecord
		} else {
			fmt.Println("No row error", err)
			return nil, err
		}
	}

	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
