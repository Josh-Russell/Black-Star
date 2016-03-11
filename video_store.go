package main

import (
	"database/sql"
	"fmt"
)

const pageSize = 4

var globalVideoStore VideoStore

type VideoStore interface {
	Save(video *Video) error
	Find(id string) (*Video, error)
	FindAll(offset int) ([]Video, error)
	FindAllByUser(user *User, offset int) ([]Video, error)
}

type DBVideoLocationStore struct {
	db *sql.DB
}

func NewDBVideoLocationStore() VideoStore {
	return &DBVideoLocationStore{
		db: globalMySQLDB,
	}
}

func (store *DBVideoLocationStore) Save(video *Video) error {
	fmt.Println("words")
	_, err := store.db.Exec(
		`
		REPLACE INTO videos
			(videoID, title, username, description, filepath, upvotes, downvotes, mature)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)
		`,
		video.ID,
		video.title,
		video.Username,
		video.Description,
		video.Location,
		video.upvotes,
		video.downvotes,
		video.mature,
	)
	return err
}

func (store *DBVideoLocationStore) Find(id string) (*Video, error) {
	//id = strings.TrimPrefix(id, ":")
	row := store.db.QueryRow(
		`
		SELECT videoID, title, username, description, filepath, upvotes, downvotes, mature
		FROM videos
		WHERE videoID = ?`,
		id,
	)

	video := Video{}
	err := row.Scan(
		&video.ID,
		&video.title,
		&video.Username,
		&video.Description,
		&video.Location,
		&video.upvotes,
		&video.downvotes,
		&video.mature,
	)
	return &video, err
}

func (store *DBVideoLocationStore) FindAll(offset int) ([]Video, error) {
	rows, err := store.db.Query(
		`
		SELECT videoID, title, username, description, filepath, upvotes, downvotes, mature
		FROM videos
		ORDER BY title DESC
		LIMIT ?
		OFFSET ?
		`,
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	videos := []Video{}
	for rows.Next() {
		video := Video{}
		err := rows.Scan(
			&video.ID,
			&video.title,
			&video.Username,
			&video.Description,
			&video.Location,
			&video.upvotes,
			&video.downvotes,
			&video.mature,
		)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil
}

func (store *DBVideoLocationStore) FindAllByUser(user *User, offset int) ([]Video, error) {
	rows, err := store.db.Query(
		`
		SELECT videoID, title, username, description, filepath, upvotes, downvotes, mature
		FROM videos
		WHERE username = ?
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?`,
		user.ID,
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	videos := []Video{}
	for rows.Next() {
		video := Video{}
		err := rows.Scan(
			&video.ID,
			&video.title,
			&video.Username,
			&video.Description,
			&video.Location,
			&video.upvotes,
			&video.downvotes,
			&video.mature,
		)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil
}
