package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"sort"
	"time"
)

var (
	dsn = flag.String("dsn", "root:password@tcp(db:3306)/benchdb", "Data Source Name")
)

func queryTable(db *sql.DB, tableName string) error {
	rows, err := db.Query("SELECT * FROM " + tableName)
	if err != nil {
		panic(err.Error())
		return err
	}

	for rows.Next() {
		var id int
		var firstName, lastName, email, gender, ipAddress string
		err = rows.Scan(&id, &firstName, &lastName, &email, &gender, &ipAddress)
		if err != nil {
			fmt.Printf("error=%v\n", err)
			return err
		}
	}
	return nil
}

func runBenchmark(try int, f func() error) {
	var min, max, avg int64
	min = math.MaxInt64
	max = math.MinInt64
	elapsedTimes := make([]int, try)
	for i := 0; i < try; i++ {
		begin := time.Now()
		err := f()
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
		elapsedTimes[i] = int(elapsed)
	}
	sort.Sort(sort.IntSlice(elapsedTimes))
	fmt.Printf("min=%f, max=%f, avg=%f, median=%f\n", float64(min)/1000000.0, float64(max)/1000000.0,
		float64(avg)/1000000.0/float64(try), float64(elapsedTimes[len(elapsedTimes)/2])/1000000.0)
}

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	tryCount := 100
	fmt.Println("bulk select from mem_tbl:")
	runBenchmark(tryCount, func() error {
		return queryTable(db, "mem_tbl")
	})

	fmt.Println("bulk select from disk_tbl")
	runBenchmark(tryCount, func() error {
		return queryTable(db, "disk_tbl")
	})
}
