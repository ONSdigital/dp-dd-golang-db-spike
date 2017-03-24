# dp-dd-golang-db-spike

## Running the basic SQL version.

Update `run.sh` update 

`export DB_USER=  export DB_PWD=`  `export DB_NAME=`

To match your database configuration, then `./run.sh`

## Running the go-pg version.

Update `run.sh` update 

`export DB_USER=  export DB_PWD=`  `export DB_NAME=`

To match your database configuration, then `./run.sh`

HTTP GET `:8000/dataset/{DATASET_ID}` will return the requested dataset.<br/>
HTTP GET `:8000/dataResource/{DATARESOURCE_ID}` will return the requested data resource.
