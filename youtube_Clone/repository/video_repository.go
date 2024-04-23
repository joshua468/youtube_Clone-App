package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/joshua468/youtube_Clone/models"
)

// VideoRepository handles database operations related to videos.
type VideoRepository struct {
	db *sql.DB
}

// NewVideoRepository creates a new instance of VideoRepository.
func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{
		db: db,
	}
}

// Create inserts a new video record into the database.
func (r *VideoRepository) Create(video *models.Video) error {
	// Prepare the SQL statement
	query := "INSERT INTO videos (id, user_id, title, content) VALUES (?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with video data
	_, err = stmt.Exec(video.ID, video.UserID, video.Title, video.Content)
	if err != nil {
		return err
	}

	return nil
}

// GetByID retrieves a video record from the database by its ID.
func (r *VideoRepository) GetByID(id uuid.UUID) (*models.Video, error) {
	query := "SELECT id, user_id, title, content FROM videos WHERE id = ?"
	row := r.db.QueryRow(query, id)

	video := &models.Video{}
	err := row.Scan(&video.ID, &video.UserID, &video.Title, &video.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("video not found")
		}
		return nil, err
	}

	return video, nil
}

// Update modifies an existing video record in the database.
func (r *VideoRepository) Update(video *models.Video) error {
	query := "UPDATE videos SET title=?, content=? WHERE id=?"

	_, err := r.db.Exec(query, video.Title, video.Content, video.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a video record from the database by its ID.
func (r *VideoRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM videos WHERE id = ?"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("video not found")
	}

	return nil
}

// List retrieves all video records from the database.
func (r *VideoRepository) List() ([]*models.Video, error) {
	query := "SELECT id, user_id, title, content FROM videos"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []*models.Video
	for rows.Next() {
		video := &models.Video{}
		if err := rows.Scan(&video.ID, &video.UserID, &video.Title, &video.Content); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return videos, nil
}
