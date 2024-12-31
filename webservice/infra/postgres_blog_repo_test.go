package infra_test

import (
	"context"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupPostgresContainer(t *testing.T) (*postgres.PostgresContainer, func(), bool) {
	ctx := context.Background()

	dbName := "users"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, startContainerError := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	terminateContainer := func() {
		t.Log("Shutting down container...")
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			t.Logf("failed to terminate container: %s", err)
			t.Fail()
		}
	}
	if startContainerError != nil {
		t.Logf("failed to start container: %s", startContainerError)
		t.Fail()
		return nil, func() {}, false
	}
	connStr, err := postgresContainer.ConnectionString(ctx)
	if err != nil {
		t.Log(err)
	}
	t.Log("Connection String: ", connStr)
	return postgresContainer, terminateContainer, true
}

func TestTestContainers(t *testing.T) {
	_, terminateContainerFunction, isSuccess := SetupPostgresContainer(t)
	if isSuccess {
	}
	terminateContainerFunction()
}
