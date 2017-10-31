package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"time"
)

const (
	tryCount = 10
)

var (
	dsn = flag.String("dsn", "mysql:password@tcp(db:3306)/benchdb", "Data Source Name")
)

func queryTable(db *sql.DB, tableName string) (int, error) {
	rows, err := db.Query("SELECT * FROM " + tableName)
	if err != nil {
		panic(err.Error())
		return 0, err
	}

	numRecords := 0
	for rows.Next() {
		numRecords += 1
		var id int
		var firstName, lastName, email, gender, ipAddress string
		err = rows.Scan(&id, &firstName, &lastName, &email, &gender, &ipAddress)
		if err != nil {
			fmt.Printf("error=%v\n", err)
			return 0, err
		}
	}
	return numRecords, nil
}

func runBenchmark(db *sql.DB, tableName string) {
	var min, max, avg int64
	min = math.MaxInt64
	max = math.MinInt64
	for i := 0; i < tryCount; i++ {
		begin := time.Now()
		_, err := queryTable(db, tableName)
		if err != nil {
			fmt.Printf("error=%v\n", err)
			continue
		}

		elapsed := time.Since(begin).Nanoseconds()
		avg += elapsed
		if max < elapsed {
			max = elapsed
		}
		if elapsed < min {
			min = elapsed
		}
	}
	fmt.Printf("from %s: min=%f, max=%f, avg=%f\n", tableName, float64(min)/1000000.0, float64(max)/1000000.0, float64(avg)/1000000.0/float64(tryCount))
}

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for i := 1; i <= 10; i++ {
		fmt.Printf("#%d\n", i)
		runBenchmark(db, "mem_tbl")
		runBenchmark(db, "disk_tbl")
	}
}
