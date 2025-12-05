package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)


var DbConn *pgx.Conn;

func ConnectDb() {
	var err error;
	DbConn, err = pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"));
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err);
	}
	// defer DbConn.Close(context.Background());

	// Verify the connection
    if err := DbConn.Ping(context.Background()); err != nil {
        log.Fatalf("Unable to ping database: %v\n", err);
    }

	log.Println("Connected successfully to database");
}