package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/votoznotna/go-rest-api/internal/comment"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

// CommentRow - models how our comments look in the database
type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

// GetComment - retrieves a comment from the database by ID
func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {

	_, err := d.Client.ExecContext(ctx, "SELECT pg_sleep(16)")
	if err != nil {
		return comment.Comment{}, err
	}
	// fetch CommentRow from the database and then convert to comment.Comment
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author 
		FROM comments 
		WHERE id = $1`,
		uuid,
	)
	err = row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("an error occurred fetching a comment by uuid: %w", err)
	}
	// sqlx with context to ensure context cancelation is honoured
	return convertCommentRowToComment(cmtRow), nil
}
