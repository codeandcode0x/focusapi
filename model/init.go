package model

import (
	"encoding/json"
	"focusapi/db"
	"log"
	"os"
)

// init
func init() {
	if runMode := os.Getenv("RUN_MODE"); runMode == "testing" {
		// TO-DO
	} else {
		AutoMigrateAll()
	}
}

// Migrate Model
func AutoMigrateAll() {
	m := map[string]interface{}{
		"id":        "uint64",
		"name":      "string",
		"CreatedAt": "0001-01-01T00:00:00Z",
		"UpdatedAt": "0001-01-01T00:00:00Z",
		"DeletedAt": nil,
	}

	js, _ := json.Marshal(m)

	js1, _ := json.Marshal(&User{})
	log.Println(string(js1))
	log.Println(string(js))

	_ = db.Conn.Table("users").AutoMigrate(&User{})
	// _ = db.Conn.Table("ts").AutoMigrate(string("ts"))
	_ = db.Conn.Table("instances").AutoMigrate(&Instance{})
	db.Conn.Exec("INSERT INTO gin_scaffold.users (`id`,`name`,`password`,`email`,`age`,`birthday`,`member_number`,`created_at`,`updated_at`,`deleted_at`,`role`) VALUES (1,'admin','$2a$10$BjYFeoOSaD8Xzs2KumA7Z.duVszjK8lB1ZaZDkdlc5bTzvPSbGify','admin@admin.com',0,'2021-05-24 23:39:01.278','1','2021-05-24 23:39:01.280','2021-05-24 23:39:01.280',NULL,'管理员'); ")
}
