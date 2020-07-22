# What is this ?

Just another CSV loader, which load csv and dumps into different Database.

The program use worker pub/sub pattern internally to speed up the data loading.

This may gain lots of time if your ETL process involves loading a huge csv files

# How does it work ?

1. Read raw CSV input
2. Distribute the work load into different worker, each worker is an independent goroutine. This helps speed up the data loading process

For more [information](https://medium.com/life-of-a-senior-data-engineer/worker-pattern-in-golang-for-data-etl-ebf8a52da636)

# Supported loading

As of now, only Postgres interface is implemented, thus you can load CSV to POSTGRES_HOST

# Add more supported Database
Either create an issue for follow the guidelines to implement it yourself.

# Guidelines
1. Create your own type of worker (Refer to `workerpool\postgresworker.go`).
2. Your new worker must satisfy the interface

```golang

// Worker is the work horse
type Worker interface {
	ExecuteTask([]string, *sync.WaitGroup, ...interface{}) error
}

```

3. Using your own worker, initiate your own loader, refer to `loader.go`
For example, you may create a `MongoDBWorker` struct, then your loader function may look like

```golang

// LoadRecordToDatabase take records and dump to Database
func LoadRecordToDatabase(records []string, numberOfGoroutine int, args ...interface{}) error {
	var wg sync.WaitGroup

	// Initiate worker pool
	// Use the corresponding worker type
	workerArray := make([]pool.Worker, 0)
	for i := 0; i < numberOfGoroutine; i++ {

    // Initiate MongoDBWorker
		w := &pool.MongoDBWorker{ID: strconv.FormatInt(int64(i), 10)}
		workerArray = append(workerArray, w)
	}
	workerPool := &pool.WorkerPool{Wg: &wg, Pool: workerArray}
	workerPool.ExecuteJob(records, args...)
	wg.Wait()
	return nil
}

```
### Build and Execution

Build with Docker

```bash
docker build . -t csvloader
# run the image in the directory where you can locate the csv and mount it to the container
docker run exec -it csvloader /bin/bash --mount src=`pwd`,target=/csvloader
# inside your container
cd /csvloader

POSTGRES_HOST="localhost" POSTGRES_PORT=5432 POSTGRES_USER="postgres" POSTGRES_PASS="admin" POSTGRES_DBNAME="ronin" go run *.go -filePath=sample.csv -query="INSERT INTO samples VALUES (\$1, \$2) ON CONFLICT (id) DO UPDATE SET value = \$2 RETURNING id" -nbrgoroutines=5

```
