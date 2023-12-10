package database

var database string

func init() {
	database = "pgsql"
}

func GetDatabase() string {
	return database
}