package infra_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"
	infra "webservice/infra/postgres"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupPostgresContainer(t *testing.T, ctx context.Context) (*postgres.PostgresContainer, func(), bool) {

	dbName := "docker"
	dbUser := "user"
	dbPassword := "password"
	setupFile, err := filepath.Abs("../../testdata/setup_testcontainer.sql")
	if err != nil {
		t.Log(err)
		t.Fail()
		return nil, func() {}, false
	}

	postgresContainer, startContainerError := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.WithInitScripts(setupFile),
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
	return postgresContainer, terminateContainer, true
}

func TestTestContainers(t *testing.T) {
	ctx := context.Background()
	postgresContainer, terminateContainerFunction, isSuccess := SetupPostgresContainer(t, ctx)
	defer terminateContainerFunction()
	if isSuccess {
	}
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Log(err)
	}
	infra.SelectFromPostgresTable(connStr)
}
