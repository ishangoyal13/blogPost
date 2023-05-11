package models

var migrationModels = []interface{}{
	&User{},
	&Blog{},
}

func GetMigrationModels() []interface{} {
	return migrationModels
}
