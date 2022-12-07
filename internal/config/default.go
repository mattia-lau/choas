package config

var defaultConfig = map[string]interface{}{
	"db.migrate.enable": false,
	"db.migrate.dir":    "",
	"db.path":           "choas.db",
	"server.port":       8080,
	"env":               "dev",
}
