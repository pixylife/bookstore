package main

import (
	env "bookstore/pkg/env"
	model "bookstore/pkg/model"
	xilogger "bookstore/pkg/xiLogger"
	gorm "github.com/jinzhu/gorm"
	middleware "github.com/labstack/echo/v4/middleware"
	"os"
)

import (
	"bookstore/bookstore"
	"bookstore/pkg/controller"
	"bookstore/pkg/eventing"
	"bookstore/pkg/provider"
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

const port = ":50051"

func readEnvs() {

	if val := os.Getenv(env.DB0_URL); val != "" {
		env.Db0Url = val
	}
	if val := os.Getenv(env.DB0_DIALET); val != "" {
		env.Db0DIALET = val
	}
	if val := os.Getenv(env.DB1_URL); val != "" {
		env.Db1Url = val
	}
	if val := os.Getenv(env.DB1_DIALET); val != "" {
		env.Db1DIALET = val
	}
}
func main() {

	readEnvs()
	env.LoadEnvs()

	database0 := connectToMongo(env.Db0Host, env.Db0Port, env.Db0User, env.Db0Pwd, env.Db0Db)

	cmdStorage := provider.NewMongoStorageProvider(database0, "event", eventing.BasicProjectionGenerator)
	commandController := controller.NewController(&model.Projection{}, model.BuildProjection, cmdStorage)
	env.CMDCtrl = commandController

	database1, err := gorm.Open(env.DB1_DIALET, env.Db1Url)
	if err != nil {
		xilogger.Log.Panic(err)
	}
	env.RDB = database1
	defer database1.Close()

	queryStorage := provider.NewGormStorageProvider(database1)
	queryController := controller.NewController(&model.Projection{}, model.BuildProjection, queryStorage)
	env.QueryCtrl = queryController

	model.InitModels(database1)

	nc := connectToNats()
	defer nc.Close()
	opts := []nats.Option{
		nats.Name("nats"),
	}
	opts = setupConnOptions(opts)
	eventSource := provider.NewNatsMessingProvider(nc)
	esController := controller.NewMessageController(eventSource)
	env.ESCtrl = esController

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	go RunServer()
	RunProxy()
}

func connectToMongo(host, port, user, password, dbName string) *mongo.Database {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	opts.SetAuth(options.Credential{
		Username: user,
		Password: password,
	})

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		xilogger.Log.Error("could not connect to database")

	}

	if err := client.Ping(context.Background(), nil); err != nil {
		xilogger.Log.Error("could not ping database")
	}

	return client.Database(dbName)
}

func connectToNats() *nats.EncodedConn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	return ec
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatal("Exiting, no servers available")
	}))
	return opts
}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

var (
	bookstoreEndpoint = flag.String("bookstoreEndpoint", "localhost:5001", "Your Description")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := bookstore.RegisterBookstoreHandlerFromEndpoint(ctx, mux, *bookstoreEndpoint, opts)
	if err != nil {
		return err
	}

	fmt.Print("\nProxy listening on 8081\n")
	return http.ListenAndServe(":8081", mux)
}

func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Printf("\nServer listening on port %v \n", port)
	bookstore.RegisterBookstoreServer(s, &controller.BookstoreServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
