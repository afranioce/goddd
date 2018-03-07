package config

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	os.Setenv("API_PATH", path.Dir(filename)+"/..")

	var err error
	if currentEnvironment, ok := os.LookupEnv("ENV"); ok == true {
		err = godotenv.Load(os.ExpandEnv("$API_PATH/.env." + currentEnvironment))
	} else {
		err = godotenv.Load(os.ExpandEnv("$API_PATH/.env"))
	}

	if err != nil {
		panic(err)
	}
}

// MustEnv Return value of environment variable
func MustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Unset environment variable %v", k)
	}
	return v
}
