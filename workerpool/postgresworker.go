package workerpool

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "ronin"
)

// PostgresWorker is just a type of worker
type PostgresWorker struct {
	ID string
}

func (w *PostgresWorker) inserttoDB(records []string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO samples (id, value)
	VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET value = $2
	RETURNING id`
	id := 0

	// Parse array of records to argument
	// interface{} has a different memory layout from a string
	// as db.QueryRow below takes []interface{} as argument,
	// the conversion must be done manually
	recordArgs := make([]interface{}, len(records))
	for i, val := range records {
		recordArgs[i] = val
	}
	err = db.QueryRow(sqlStatement, recordArgs...).Scan(&id)
	if err != nil {
		panic(err)
	}
}

//ExecuteTask load the record to Postgres
func (w *PostgresWorker) ExecuteTask(records []string, wg *sync.WaitGroup) error {
	defer wg.Done()
	// TODO: load records to database
	w.inserttoDB(records)
	fmt.Printf("Worker's id %s , executing task, message is %s \n", w.ID, records)
	return nil
}
