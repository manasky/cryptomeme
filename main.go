package main

import (
	"context"
	"cryptolist/api"
	"cryptolist/database"
	"cryptolist/syncer"
	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	flags := pflag.NewFlagSet(path.Base(os.Args[0]), pflag.ContinueOnError)

	flags.String("host", "127.0.0.1:8080", "host:port for the HTTP server")
	flags.String("db-driver", "sqlite3", "database driver")
	flags.String("db-host", "localhost", "database host")
	flags.String("db-port", "3306", "database host")
	flags.String("db-name", "cryptolist", "database name")
	flags.String("db-username", "", "database username")
	flags.String("db-password", "", "database password")
	flags.String("delay", "10", "sync delay in seconds")
	flags.String("api-endpoint", "", "api endpoint")

	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Panic("failed to parse arguments")
	}

	err = viper.BindPFlags(flags)
	if err != nil {
		log.Panic("failed to bind flags")
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Print(err)
	}

	viper.AutomaticEnv()
}

func main() {
	cg := api.NewCoinGeco(viper.GetString("api-endpoint"))

	db, err := database.Open(viper.GetString("db-driver"), viper.GetString("db-username"), viper.GetString("db-password"), viper.GetString("db-host"), viper.GetString("db-port"), viper.GetString("db-name"))
	if err != nil {
		panic(err)
	}

	err = db.Migrate(context.Background())
	if err != nil {
		panic(err)
	}

	syncr, _ := syncer.New(cg, db, &syncer.Config{
		Delay: time.Duration(viper.GetInt("delay")) * time.Second,
	})
	defer syncr.Stop()

	err = syncr.Start()
	if err != nil {
		panic(err)
	}

	s := api.NewServer(db)

	r := mux.NewRouter()
	r.HandleFunc("/v1/markets", s.Markets)
	http.Handle("/", r)

	go func() {
		err = http.ListenAndServe(viper.GetString("host"), nil)
		if err != nil {
			panic(err)
		}
	}()

	log.Printf("Server started on %s", viper.GetString("host"))
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
