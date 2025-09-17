package app

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	db, err := pgxpool.New(ctx, "postgresql://user:password@localhost:5432/walletdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
