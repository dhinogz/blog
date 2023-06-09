package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"blog.dhinojosa.io/internal/models"

	_ "github.com/lib/pq"
)

type application struct {
	config   config
	errorLog *log.Logger
	infoLog  *log.Logger
	articles *models.ArticleModel
}

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func main() {

	// Network and DB config, input comes from command line flags
	var cfg config

	// Port and environment configuration
	flag.IntVar(&cfg.port, "port", 4000, "Web server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Postgres DB configuration
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("BLOG_DB_DSN"), "PostgreSQL DSN")

	// Postgres connection rules
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	// Create system log
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	infoLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	// Connect to db
	db, err := openDB(cfg)
	if err != nil {
		errorLog.Fatal(err)
	}
	// Defer closes the database at the end of the main function
	defer db.Close()
	infoLog.Printf("database connection pool established")

	// TODO: Change articles to models pointer
	app := application{
		config:   cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
		articles: &models.ArticleModel{DB: db},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infoLog.Printf("Starting %s server on %d", cfg.env, cfg.port)
	infoLog.Printf("Connected to database %s", cfg.db.dsn)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
