package interfaces

import "github.com/justpoypoy/go-base/domain"

// A UserRepository belong to the inteface layer
type UserRepository struct {
	SQLHandler SQLHandler
}

// FindAll is returns the number of entities.
func (ur *UserRepository) FindAll() (domain.Users, error) {
	const query = `
		SELECT
			id,
			fullname,
			email,
			is_active
		FROM 
			users
		`
	rows, err := ur.SQLHandler.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var users domain.Users
	for rows.Next() {
		var user = &domain.User{}
		if err = rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.IsActived); err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID is returns the entity identified by the given id.
func (ur *UserRepository) FindByID(userID uint) (domain.User, error) {
	const query = `
		SELECT
			id,
			fullname,
			email,
			is_active,
			created_at,
			updated_at,
			deleted_at
		FROM 
			users
		WHERE
			id = ?
	`
	row, err := ur.SQLHandler.Query(query, userID)
	defer row.Close()

	if err != nil {
		return domain.User{}, err
	}

	var user domain.User

	var id uint
	var fullname, email string
	var isActive bool
	row.Next()
	if err = row.Scan(&id, &fullname, &email, &isActive); err != nil {
		return domain.User{}, err
	}
	user.ID = id
	user.Fullname = fullname
	user.Email = email
	user.IsActived = isActive

	return user, nil
}
