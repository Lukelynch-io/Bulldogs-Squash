package post_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"
	"webservice/pkg/post"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupPostgresContainer(t *testing.T, ctx context.Context) (*postgres.PostgresContainer, func(), bool) {

	dbName := "Posts"
	dbUser := "postgres"
	dbPassword := "password"
	seedDataFile, err := filepath.Abs("../../testdata/setup_post_database.sql")
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

func TestInsertPost(t *testing.T) {
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
	var db post.PostStorage
	pgdb := new(post.PostgresPostDatabase)
	pgdb.Connect(connStr)
	db = pgdb
	testPost := post.Post{
		ID:          uuid.NewString(),
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "Test URL",
	}
	// Act
	_, insertError := db.InsertPost(testPost)
	if insertError != nil {
		t.Fatal(insertError)
	}
}
