package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/manasky/cryptomeme/api"
	"github.com/manasky/cryptomeme/database"
	"github.com/manasky/cryptomeme/meme"
	"github.com/manasky/cryptomeme/service"
	"github.com/manasky/cryptomeme/syncer"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	flags := pflag.NewFlagSet(path.Base(os.Args[0]), pflag.ContinueOnError)

	flags.String("host", "127.0.0.1:8080", "host:port for the HTTP server")
	flags.String("db-driver", "sqlite3", "database driver")
	flags.String("db-host", "localhost", "database host")
	flags.String("db-port", "3306", "database host")
	flags.String("db-name", "crypromeme", "database name")
	flags.String("db-username", "", "database username")
	flags.String("db-password", "", "database password")
	flags.String("market-delay", "180", "market sync delay in seconds")
	flags.String("market-chart-delay", "21600", "chart sync delay in seconds")
	flags.String("api-endpoint", "", "api endpoint")
	flags.String("cache-driver", "memory", "cache driver. supported drivers: ent (mysql, sqlite3, postgres), memory")
	flags.String("meme-dataset-path", "dataset.json", "path of meme dataset json file")
	flags.String("meme-cache-duration", "3600", "cache duration in seconds")
	flags.String("allowed-origins", "*", "allowed origins for the API")

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
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}

func main() {
	cg := api.NewCoinGeco(viper.GetString("api-endpoint"))

	m, err := database.NewManager(context.Background(), viper.GetString("cache-driver"), getCacheDriverConfig(viper.GetString("cache-driver")))
	if err != nil {
		panic(err)
	}

	syncr := syncer.New()
	defer syncr.StopAllTasks()

	srv := service.NewService(cg, m)
	err = syncr.RegisterTask("markets", time.Duration(viper.GetInt("market-delay"))*time.Second, srv.SyncMarkets)
	if err != nil {
		panic(err)
	}
	err = syncr.RegisterTask("markets-chart", time.Duration(viper.GetInt("market-chart-delay"))*time.Second, srv.SyncMarketsChart)
	if err != nil {
		panic(err)
	}

	mm, err := meme.New(viper.GetString("meme-dataset-path"), time.Duration(viper.GetInt("meme-cache-duration"))*time.Second)
	if err != nil {
		panic(err)
	}

	s := api.NewServer(m, mm)

	r := mux.NewRouter()
	r.HandleFunc("/v1/markets", s.Markets)
	http.Handle("/", r)

	go func() {
		err = http.ListenAndServe(viper.GetString("host"), handlers.CORS(
			handlers.AllowedOrigins([]string{viper.GetString("allowed-origins")}),
			handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
		)(r))
		if err != nil {
			panic(err)
		}
	}()

	log.Printf("Server started on %s", viper.GetString("host"))
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}

func getCacheDriverConfig(driver string) interface{} {
	switch driver {
	case "ent":
		return &database.EntConfig{
			Driver:   viper.GetString("db-driver"),
			Username: viper.GetString("db-username"),
			Password: viper.GetString("db-password"),
			Host:     viper.GetString("db-host"),
			Port:     viper.GetString("db-port"),
			Database: viper.GetString("db-name"),
		}
	default:
		return nil
	}
}
