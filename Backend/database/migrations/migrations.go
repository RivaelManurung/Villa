package migrations

import (
	"Backend/database"
	"fmt"
	"log"
	"Backend/models/entity"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&entity.Admin{},
	)

	if err != nil {
		// panic("gagal migrate")
		log.Println(err)
	}
	fmt.Println("Berhasil migrate")
}
