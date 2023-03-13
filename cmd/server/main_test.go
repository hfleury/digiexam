package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"runtime"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/suite"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MainTestSuite struct {
	suite.Suite
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
	DbConn   *gorm.DB
}

func (suite *MainTestSuite) SetupSuite() {
	suite.Pool, suite.Resource = suite.initDB()
}

func (suite *MainTestSuite) TearDownSuite() {
	closeDB(suite.Pool, suite.Resource)
}

func initMigrations(dbConn *gorm.DB) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrate, err := migrate.NewWithDatabaseInstance(
		"file://../../../database/postgres/migration",
		"horsedbtest", driver)
	if err != nil {
		log.Fatalf("NewWithDatabaseInstance %v", err)
	}

	err = migrate.Up()
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *MainTestSuite) initDB() (*dockertest.Pool, *dockertest.Resource) {
	pgURL := initPostgres()
	pgPass, _ := pgURL.User.Password()

	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.1-alpine",
		Env: []string{
			"POSTGRES_USER=" + pgURL.User.Username(),
			"POSTGRES_PASSWORD=" + pgPass,
			"POSTGRES_DB=" + pgURL.Path,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "localhost", HostPort: "5433"},
			},
		},
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal("Could not connect to docker")
	}

	resource, err := pool.RunWithOptions(&runOpts)
	if err != nil {
		log.Fatal("Could not start postgres container")
	}

	pgURL.Host = resource.Container.NetworkSettings.IPAddress

	// Docker layer network is different on Mac
	if runtime.GOOS == "darwin" {
		pgURL.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		suite.DbConn, err = gorm.Open(psql.Open(pgURL.String()), &gorm.Config{})
		if err != nil {
			return err
		}
		return err
	}); err != nil {
		phrase := fmt.Sprintf("Could not connect to docker: %s", err)
		log.Fatal(phrase)
	}

	suite.initMigrations()

	return pool, resource
}

func closeDB(pool *dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		phrase := fmt.Sprintf("Could not purge resource: %s", err)
		log.Fatal(phrase)
	}
}

func (suite *MainTestSuite) initMigrations() {
	sqlDB, err := suite.DbConn.DB()
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrate, err := migrate.NewWithDatabaseInstance(
		"file://../../../database/postgres/migration",
		"horsedbtest", driver)
	if err != nil {
		log.Fatalf("NewWithDatabaseInstance %v", err)
	}

	err = migrate.Up()
	if err != nil {
		log.Fatal(err)
	}
}

func initPostgres() *url.URL {
	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("rootuser", "nosecret"),
		Path:   "horsedbtest",
		Host:   "localhost:5432",
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	return pgURL
}
