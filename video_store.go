package main

import "database/sql"

const pageSize = 25

var globalImageStore ImageStore

type ImageStore interface {
	Save(vid *video) error
	Find(id string) (*video, error)
	FindAll(offset int) ([]video, error)
	FindAllByUser(user *User, offset int) ([]video, error)
}

type DBImageStore struct {
	db *sql.DB
}

func NewDBImageStore() ImageStore {
	return &DBImageStore{
		db: globalMySQLDB,
	}
}

func (store *DBImageStore) Save(vid *video) error {
	_, err := store.db.Exec(
		`
		REPLACE INTO images
			(id, user_id, name, location, description, size, created_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?)
		`,
		vid.ID,
		vid.UserID,
		vid.Name,
		vid.Location,
		vid.Description,
		vid.Size,
		vid.CreatedAt,
	)
	return err
}

func (store *DBImageStore) Find(id string) (*video, error) {
	row := store.db.QueryRow(
		`
		SELECT id, user_id, name, location, description, size, created_at
		FROM images
		WHERE id = ?`,
		id,
	)

	vid := video{}
	err := row.Scan(
		&vid.ID,
		&vid.UserID,
		&vid.Name,
		&vid.Location,
		&vid.Description,
		&vid.Size,
		&vid.CreatedAt,
	)
	return &vid, err
}

func (store *DBImageStore) FindAll(offset int) ([]video, error) {
	rows, err := store.db.Query(
		`
		SELECT id, user_id, name, location, description, size, created_at
		FROM images
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?
		`,
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	vids := []video{}
	for rows.Next() {
		vid := video{}
		err := rows.Scan(
			&vid.ID,
			&vid.UserID,
			&vid.Name,
			&vid.Location,
			&vid.Description,
			&vid.Size,
			&vid.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		vids = append(vids, vid)
	}

	return vids, nil
}

func (store *DBImageStore) FindAllByUser(user *User, offset int) ([]video, error) {
	rows, err := store.db.Query(
		`
		SELECT id, user_id, name, location, description, size, created_at
		FROM images
		WHERE user_id = ?
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

	vids := []video{}
	for rows.Next() {
		vid := video{}
		err := rows.Scan(
			&vid.ID,
			&vid.UserID,
			&vid.Name,
			&vid.Location,
			&vid.Description,
			&vid.Size,
			&vid.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		vids = append(vids, vid)
	}

	return vids, nil
}
