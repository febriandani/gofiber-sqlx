package repository

import (
	mt "gofiber-sqlx/model/todolist"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// Repository interface with CRUD operations
type UserRepository interface {
	GetUserByID(id int) (*mt.User, error)
	CreateUser(user *mt.User) error
	UpdateUser(user *mt.User) error
	DeleteUser(id int) error
}

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) UserRepository {
	return &TodoListRepository{db: db}
}

// GetUserByID retrieves a user by ID
func (r *TodoListRepository) GetUserByID(id int) (*mt.User, error) {
	user := &mt.User{}
	err := r.db.Get(user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new user
func (r *TodoListRepository) CreateUser(user *mt.User) error {
	_, err := r.db.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", user)
	return err
}

// UpdateUser updates an existing user
func (r *TodoListRepository) UpdateUser(user *mt.User) error {
	_, err := r.db.NamedExec("UPDATE users SET name=:name, email=:email WHERE id=:id", user)
	return err
}

// DeleteUser deletes a user by ID
func (r *TodoListRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
