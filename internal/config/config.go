package config

type Config struct {
	Database database `json:"database"`
	Server   server   `json:"server"`
}

/*
server := "localhost"
    port := 1433
    user := "kursUser"
    password := "kursPswd"
    database := "TEST"

    connectionString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;", server, port, user, password, database)

*/

func NewConfig() *Config {
	return &Config{}
}
