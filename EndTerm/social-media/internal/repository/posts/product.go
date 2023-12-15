package posts

import (
	"context"
	"database/sql"
	"fmt"
	interfaces "github.com/olzhas-b/social-media/internal/interfaces/repository"
	"github.com/olzhas-b/social-media/internal/models"
	"github.com/olzhas-b/social-media/modules/logger"
	"go.uber.org/zap"
	"log"
	"strings"
)

type product struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) interfaces.IPosts {
	return &product{
		db:  db,
		ctx: ctx,
	}
}

func (r *product) List(searchText string, userID uint64) (posts []models.PostDTO, err error) {
	query := `
		SELECT
			id,
			user_id,
			author,
			title,
			text,
			image,
			date_created
		FROM
			posts`

	addWhereClause := func(condition string, value interface{}) {
		if value != nil {
			if strings.Contains(query, "WHERE") {
				query += " AND"
			} else {
				query += " WHERE"
			}

			query += fmt.Sprintf(condition, value)
		}
	}
	if userID != 0 {
		addWhereClause(" user_id = %d", userID)
	}

	log.Println(query)
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post models.PostDTO
		err = rows.Scan(&post.ID, &post.UserID, &post.Author, &post.Title, &post.Text, &post.Image, &post.DateCreatedAt)
		if err != nil {
			return
		}
		log.Println(post)

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return
	}

	log.Println(posts)
	return
}

func (r *product) ByID(id uint64) (post models.Posts, err error) {
	l := logger.WorkLogger.Named("repo.posts.List")

	query := `
		SELECT
			id,
			user_id,
			author,
			title,
			text,
			image,
			date_created
		FROM
			posts
			WHERE id = $1`

	row := r.db.QueryRow(query, id)

	err = row.Scan(&post.ID, &post.UserID, &post.Author, &post.Title, &post.Text, &post.Image, &post.DateCreatedAt)
	if err != nil {
		l.Error("couldn't scan posts", zap.Error(err))
		return
	}

	return
}

func (r *product) Add(post models.Posts) (err error) {
	l := logger.WorkLogger.Named("repo.posts.Add").With(zap.Any("post", post))

	_, err = r.db.Exec("INSERT INTO posts (user_id, author, title, text, image) VALUES ($1, $2, $3, $4, $5)",
		post.UserID,
		post.Author,
		post.Title,
		post.Text,
		post.Image)
	if err != nil {
		l.Error("couldn't insert posts")
		return
	}

	return
}

func (r *product) Remove(id uint64) (err error) {
	l := logger.WorkLogger.Named("repo.posts.Remove").With(zap.Uint64("id", id))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		l.Error("couldn't delete posts")
		return
	}

	return
}
