package postgres_test

import (
	"context"
	"testing"
	"webservice/infra/postgres"
	"webservice/pkg/auth"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

// SHA256 hash of "password"
const testPasswordHash = "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"

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
	db := new(postgres.PostgresAuthDatabase)
	db.Connect(connStr)
	testUser := postgres.User{
		UserId:       uuid.New().String(),
		Username:     "hello",
		PasswordHash: testPasswordHash,
		Role:         string(auth.Viewer),
	}
	// Act
	db.InsertTestUserIntoAuthDatabase(testUser)
	actualUser := db.SelectUserFromAuthDatabase(testUser.Username)
	// Assert
	t.Log(actualUser)
	assert.Equal(t, actualUser.Username, testUser.Username)
}
