package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)


var DbConn *pgx.Conn;

func ConnectDb() {
	var err error;
	DbConn, err = pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"));
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err);
		os.Exit(1);
	}
	// defer DbConn.Close(context.Background());

	// Verify the connection
    if err := DbConn.Ping(context.Background()); err != nil {
        fmt.Println("Unable to ping database:", err);
    }

	fmt.Println("Connected successfully to database");
}