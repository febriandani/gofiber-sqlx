package repository

import (
	mt "gofiber-sqlx/model/user"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// Repository interface with CRUD operations
type UserRepository interface {
	GetUserByID(id int) (*mt.User, error)
	CreateUser(user *mt.User) (int64, error)
	UpdateUser(user *mt.User) error
	DeleteUser(id int) error
	GetUsersRepo(offset, limit int) ([]mt.User, error)
}

type UsertRepository struct {
	db *sqlx.DB
}

func NewUsertRepository(db *sqlx.DB) UserRepository {
	return &UsertRepository{db: db}
}

// GetUserByID retrieves a user by ID
func (r *UsertRepository) GetUserByID(id int) (*mt.User, error) {
	user := &mt.User{}
	err := r.db.Get(user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new user
func (r *UsertRepository) CreateUser(user *mt.User) (int64, error) {
	var (
		param = make([]interface{}, 0)
	)

	param = append(param, user.Name)
	param = append(param, user.Email)

	query := "INSERT INTO users (name, email) VALUES (?, ?) returning id"

	query, args, err := sqlx.In(query, param...)
	if err != nil {
		return 0, err
	}

	query = r.db.Rebind(query)

	res := r.db.QueryRow(query, args...)

	err = res.Err()
	if err != nil {
		return 0, err
	}

	var id int64
	err = res.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

// UpdateUser updates an existing user
func (r *UsertRepository) UpdateUser(user *mt.User) error {
	_, err := r.db.NamedExec("UPDATE users SET name=:name, email=:email WHERE id=:id", user)
	return err
}

// DeleteUser deletes a user by ID
func (r *UsertRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func (r *UsertRepository) GetUsersRepo(offset, limit int) ([]mt.User, error) {
	var result []mt.User

	query := `select id, name, email from users OFFSET ((?-1)*?) ROWS
	FETCH NEXT ? ROWS ONLY `

	query, args, err := sqlx.In(query, offset, limit, limit)

	query = r.db.Rebind(query)

	err = r.db.Select(&result, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
