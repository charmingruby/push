package container

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/charmingruby/push/internal/infra/database/mongo_repository"
	"github.com/charmingruby/push/pkg/mongodb"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DBName = "test_db"
	DBUser = "test_user"
	DBPass = "test_password"
)

type MongoTestDatabase struct {
	DB        *mongo.Database
	DBAddr    string
	container testcontainers.Container
}

func NewMongoTestDatabase() *MongoTestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	container, db, dbAddr, err := createContainer(ctx)
	if err != nil {
		log.Fatal("failed to setup container", err)
	}
	cancel()

	return &MongoTestDatabase{
		container: container,
		DB:        db,
		DBAddr:    dbAddr,
	}
}

func (tdb *MongoTestDatabase) TearDown() {
	_ = tdb.container.Terminate(context.Background())
}

func (tdb *MongoTestDatabase) CreateCollections() error {
	if err := tdb.DB.CreateCollection(context.Background(), mongo_repository.NOTIFICATION_COLLECTION); err != nil {
		return err
	}

	if err := tdb.DB.CreateCollection(context.Background(), mongo_repository.COMMUNICATION_CHANNEL_COLLECTION); err != nil {
		return err
	}

	return nil
}

func (tdb *MongoTestDatabase) DropCollections() error {
	if err := tdb.DB.Drop(context.Background()); err != nil {
		return err
	}

	return nil
}

func createContainer(ctx context.Context) (testcontainers.Container, *mongo.Database, string, error) {
	var env = map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": DBUser,
		"MONGO_INITDB_ROOT_PASSWORD": DBPass,
		"MONGO_INITDB_DATABASE":      DBName,
	}
	var port = "27017/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo",
			ExposedPorts: []string{port},
			Env:          env,
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to get container external port: %v", err)
	}

	log.Println("mongo container ready and running at port: ", p.Port())

	uri := fmt.Sprintf("mongodb://%s:%s@localhost:%s", DBUser, DBPass, p.Port())
	db, err := mongodb.NewMongoConnection(uri, "testdb")
	if err != nil {
		return container, db, uri, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, uri, nil
}
