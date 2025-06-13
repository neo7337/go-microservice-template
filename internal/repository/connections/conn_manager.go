package connections

import (
	"errors"
	"fmt"
	"io"

	"github.com/neo7337/go-microservice-template/internal/config"
	"github.com/neo7337/go-microservice-template/internal/repository"
	postgres_repo "github.com/neo7337/go-microservice-template/internal/repository/postgres"

	"oss.nandlabs.io/golly/collections"
	"oss.nandlabs.io/golly/l3"
	"oss.nandlabs.io/golly/lifecycle"
)

var connectionClosers = collections.NewArrayList[io.Closer]()
var logger = l3.Get()

// GetRepoConnection creates and returns a SimpleComponent responsible for managing
// repository connections within the application's lifecycle. It initializes the
// connection manager with start and stop functions to open and close connections,
// and handles errors during startup by logging and triggering a shutdown of all
// components if connections fail to open.
func GetRepoConnection(manager lifecycle.ComponentManager) *lifecycle.SimpleComponent {
	connectionManager := &lifecycle.SimpleComponent{
		CompId: "connection-manager",
		StartFunc: func() error {
			return OpenConnections()
		},
		AfterStart: func(err error) {
			if err != nil {
				logger.Error("Failed to open connections", "error", err)
				manager.StopAll()
			} else {
				logger.Info("Connections opened successfully")
			}
		},
		StopFunc: func() error {
			return CloseConnections()
		},
	}
	return connectionManager
}

// OpenConnections initializes and opens connections for all enabled repository providers
// as specified in the application configuration. It iterates through each provider,
// checks if it is enabled, and attempts to open a connection based on the provider's name.
// If a connection fails to open, it logs the error and returns it. On success, it registers
// the corresponding repositories for the provider's modules. Returns an error if any connection
// fails to open, otherwise returns nil.
func OpenConnections() (err error) {
	configDetails := config.GetConfig()
	for _, provider := range configDetails.Repository.Providers {
		if !provider.Enabled {
			logger.Info("Skipping disabled provider", "name", provider.Name)
			continue
		}

		logger.Info("Opening connection for provider", "name", provider.Name)

		switch provider.Name {
		case "postgres":
			conn, err := OpenPostgresConnection(provider.Connection)
			if err != nil {
				logger.Error("Failed to open Postgres connection", "error", err, "provider", provider.Name)
				return err
			}
			RegisterPostgresRepos(conn, provider.Modules)
		}
	}
	return
}

// OpenPostgresConnection establishes a new connection to a PostgreSQL database using the provided
// connection configuration. It returns a pointer to a PostgresRepo instance and an error if the
// connection fails. The function also registers the connection for later closure and logs a
// success message upon successful connection.
//
// Parameters:
//   - connection: config.Connection containing the database connection details.
//
// Returns:
//   - db: *postgres_repo.PostgresRepo, the connected Postgres repository instance.
//   - err: error, non-nil if the connection could not be established.
func OpenPostgresConnection(connection config.Connection) (db *postgres_repo.PostgresRepo, err error) {
	database := postgres_repo.NewPostgresRepo()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", connection.Host, connection.Username, connection.Password, connection.DbName, connection.Port)
	db, err = database.Connect(dsn)
	if err != nil {
		err = errors.New("Failed to connect to Postgres: " + err.Error())
		return
	}
	connectionClosers.Add(db)
	logger.Info("Postgres connection opened successfully")
	return
}

// RegisterPostgresRepos registers repository implementations for the specified modules using the provided PostgresRepo instance.
// For each module name in the modules slice, it creates the corresponding repository and registers it with the repository manager.
// Supported modules include "users", "travellers", "preferences", and "loyalty_programs".
//
// Parameters:
//   - db:        Pointer to a PostgresRepo instance used to initialize the repositories.
//   - modules:   Slice of strings specifying which modules' repositories to register.
func RegisterPostgresRepos(db *postgres_repo.PostgresRepo, modules []string) {
	for _, module := range modules {
		switch module {
		case "users":
			repo := postgres_repo.NewPostgresUsersRepo(db)
			repository.Manager.Register(repository.UsersRepo, repo)
		}
	}
}

// CloseConnections iterates over all registered connection closers and closes each connection.
// It returns an error if any occurs during the closing process.
func CloseConnections() (err error) {
	if !connectionClosers.IsEmpty() {
		for it := connectionClosers.Iterator(); it.HasNext(); {
			it.Next().Close()
		}
	}
	return
}
