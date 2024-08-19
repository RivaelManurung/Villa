package main

import (
	database "Backend/database"
	migrations "Backend/database/migrations"

)

func main() {
	database.Connection()
	migrations.Migration()
}
