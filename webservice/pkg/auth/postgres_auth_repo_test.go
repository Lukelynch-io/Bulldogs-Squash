package auth_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"
	"webservice/pkg/auth"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// SHA256 hash of "password"
const testPasswordHash = "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"

func SetupPostgresContainer(t *testing.T, ctx context.Context) (*postgres.PostgresContainer, func(), bool) {

	dbName := "Auth"
	dbUser := "postgres"
	dbPassword := "password"
	seedDataFile, err := filepath.Abs("../../testdata/setup_auth_database.sql")
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
		postgres.WithInitScripts(seedDataFile),
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

func TestInsertNewUserIntoTestContainerDB(t *testing.T) {
	// Arrange
	ctx := context.Background()
	pqContainer, terminate, isSuccess := SetupPostgresContainer(t, ctx)
	defer terminate()
	if !isSuccess {
		t.Fatal("Something went wrong setting up the test container")
	}
	connStr, err := pqContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(connStr)
	db := new(auth.PostgresAuthDatabase)
	db.Connect(connStr)
	testUser := auth.User{
		UserId:       auth.UserId(uuid.New().String()),
		Username:     "hello",
		PasswordHash: testPasswordHash,
		Role:         auth.Viewer,
	}
	// Act
	db.InsertTestUserIntoAuthDatabase(testUser)
	actualUser := db.SelectUserFromAuthDatabase(testUser.Username)
	// Assert
	t.Log(actualUser)
	assert.Equal(t, actualUser.Username, testUser.Username)
}
