package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

type Row struct {
	id any
}

func dbConnect() (*pgxpool.Pool, *pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	databaseUrl := os.Getenv("DATABASE_URL")

	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, nil, err
	}
	fmt.Println("Connected to database")

	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to acquire connection from pool: %v\n", err)
		return nil, nil, err
	}

	return dbPool, conn.Conn(), nil
}

func execQuery(query string, Conn *pgx.Conn) ([]Row, int, string) {

	startQuery := time.Now()
	rows, err := Conn.Query(context.Background(), query)
	execTime := time.Since(startQuery)

	if err != nil {
		fmt.Println("Error executing query:", err)
	}
	defer rows.Close()

	var rowSlice []Row
	for rows.Next() {
		var r Row
		err := rows.Scan(&r.id)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error scanning rows:", err)
	}

	for rows.Next() {
		var r Row
		err := rows.Scan(&r.id)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return rowSlice, len(rowSlice), execTime.String()
}

func eQuery(query string, dbPool *pgxpool.Pool) (int64, string) {
	startQuery := time.Now()
	rows, err := dbPool.Exec(context.Background(), query)
	execTime := time.Since(startQuery)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, "0"
	}

	return rows.RowsAffected(), execTime.String()
}
