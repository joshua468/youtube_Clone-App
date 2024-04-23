package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/joshua468/youtube_Clone/errors"
	"github.com/joshua468/youtube_Clone/models"
	"github.com/joshua468/youtube_Clone/services"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) Create(user *models.User) error {
	// Prepare the SQL statement
	query := "INSERT INTO users (id, username, email, password, is_admin) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with user data
	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.Password, user.IsAdmin)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	query := "SELECT id, username, email, password, is_admin FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	// Construct the SQL query to update the user's information
	query := "UPDATE users SET username=?, email=?, password=?, is_admin=? WHERE id=?"

	// Execute the SQL query with the user's updated information
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.IsAdmin, user.ID)
	if err != nil {
		// Return any errors encountered during the database operation
		return err
	}

	// Check if any rows were affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, return an error indicating that the user was not found
	if rowsAffected == 0 {
		return errors.ErrorUserNotFound
	}

	// Return nil to indicate that the update operation was successful
	return nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	// Construct the SQL query to delete the user by ID
	query := "DELETE FROM users WHERE id = ?"

	// Execute the SQL query to delete the user
	result, err := r.db.Exec(query, id)
	if err != nil {
		// Return any errors encountered during the database operation
		return err
	}

	// Check if any rows were affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, return an error indicating that the user was not found
	if rowsAffected == 0 {
		return errors.ErrorUserNotFound
	}

	// Return nil to indicate that the delete operation was successful
	return nil
}

func (r *UserRepository) List() ([]*models.User, error) {
	query := "SELECT id, username, email, password, is_admin FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsAdmin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
