package postgres

import "github.com/neo7337/go-microservice-template/pkg"

type PostgresUsersRepo struct {
	*PostgresRepo
}

func NewPostgresUsersRepo(repo *PostgresRepo) *PostgresUsersRepo {
	return &PostgresUsersRepo{
		PostgresRepo: repo,
	}
}

// GetUsers retrieves all users from the database.
func (r *PostgresUsersRepo) GetUsers() ([]pkg.User, error) {
	var users []pkg.User
	query := "SELECT id, name, email FROM users"
	rows, err := r.Database.Query(query)
	if err != nil {
		logger.Error("Failed to execute query", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user pkg.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			logger.Error("Failed to scan row", "error", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		logger.Error("Error encountered during row iteration", "error", err)
		return nil, err
	}
	logger.Info("Successfully retrieved users", "count", len(users))
	return users, nil
}
