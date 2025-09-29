package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	db, err := pgxpool.New(ctx, "postgresql://user:password@localhost:5432/walletdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
