package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"webapp/pkg/data"
	"webapp/pkg/repository"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB
var testRepo repository.DatabaseRepo

func TestMain(m *testing.M) {
	// connect to docker; fail if docker not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker; is it running? %s", err)
	}

	// ping docker to make sure it's working
	err = p.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	pool = p

	// set up our docker options, specifying the image and so forth
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	hostConfig := func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	}

	// get a resource (docker image)
	resource, err = pool.RunWithOptions(&opts, hostConfig)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource: %s", err)
	}

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// start the image and wait until it's ready
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			return err
		}
		return testDB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// populate the database with empty tables
	err = createTestTables()
	if err != nil {
		log.Fatalf("could not create test tables: %s", err)
	}

	testRepo = &PostgresDBRepo{DB: testDB}

	// run the tests
	code := m.Run()

	// purge resource when done
	if purgeErr := pool.Purge(resource); purgeErr != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func createTestTables() error {
	tableSQL, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("can't ping database")
	}
}

func TestPostgresDBRepoInsertUser(t *testing.T) {
	testUser := data.User{
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin2@example.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("insert user return an error: %s", err)
	}
	if id != 3 {
		t.Errorf("insert user return wrong id; expected 3, but got %d", id)
	}
}

func TestPostgresDBrepoAllUsers(t *testing.T) {
	users, err := testRepo.AllUsers()
	if err != nil {
		t.Errorf("can't get all users: %s", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, but got %d", len(users))
	}

	testUser := data.User{
		FirstName: "Jack",
		LastName:  "Smith",
		Email:     "admin3@example.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = testRepo.InsertUser(testUser)

	users, err = testRepo.AllUsers()
	if err != nil {
		t.Errorf("can't get all users: %s", err)
	}
	if len(users) != 3 {
		t.Errorf("expected 3 users, but got %d", len(users))
	}
}

func TestPostgresDBRepoGetUserById(t *testing.T) {
	user, err := testRepo.GetUser(3)
	if err != nil {
		t.Errorf("can't get user by id: %s", err)
	}
	if user.Email != "admin2@example.com" {
		t.Errorf("expected user email, but got %s", user.Email)
	}
	// test not found
	_, err = testRepo.GetUser(5)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}

func TestPostgresDBRepoGetUserByEmail(t *testing.T) {
	user, err := testRepo.GetUserByEmail("admin2@example.com")
	if err != nil {
		t.Errorf("can't get user by email: %s", err)
	}
	if user.ID != 3 {
		t.Errorf("expected user id 3, but got %d", user.ID)
	}
}

func TestPostgresDBRepoUpdateUser(t *testing.T) {
	user, err := testRepo.GetUser(3)
	if err != nil {
		t.Errorf("can't get user by id: %s", err)
	}
	user.FirstName = "Jane"
	user.Email = "jane@smith.com"
	err = testRepo.UpdateUser(*user)
	if err != nil {
		t.Errorf("can't update user: %s", err)
	}
	user, err = testRepo.GetUser(3)
	if err != nil {
		t.Errorf("can't get user by id: %s", err)
	}
	if user.FirstName != "Jane" {
		t.Errorf("expected user first name Jane, but got %s", user.FirstName)
	}
	if user.Email != "jane@smith.com" {
		t.Errorf("expected user email jane@smith.com but got %s", user.Email)
	}
}

func TestPostgresDBRepoDeleteUser(t *testing.T) {
	err := testRepo.DeleteUser(3)
	if err != nil {
		t.Errorf("can't delete user: %s", err)
	}
	_, err = testRepo.GetUser(3)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}

func TestPostgresDBRepoResetPassword(t *testing.T) {
	err := testRepo.ResetPassword(2, "newpassword")
	if err != nil {
		t.Errorf("can't reset password: %s", err)
	}
	user, err := testRepo.GetUser(2)
	if err != nil {
		t.Errorf("can't get user by id: %s", err)
	}
	matches, err := user.PasswordMatches("newpassword")
	if err != nil {
		t.Errorf("can't compare password: %s", err)
	}
	if !matches {
		t.Error("password doesn't match")
	}
}

func TestPostgresDBrepoInsertUserImage(t *testing.T) {
	var image data.UserImage
	image.UserID = 2
	image.FileName = "test.jpg"
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now()

	newId, err := testRepo.InsertUserImage(image)
	if err != nil {
		t.Errorf("can't insert user image: %s", err)
	}
	if newId != 2 {
		t.Errorf("expected new id 1, but got %d", newId)
	}
	// user id not found
	image.UserID = 100
	_, err = testRepo.InsertUserImage(image)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}
