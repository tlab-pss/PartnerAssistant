package infrastructures

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// InitEnvironment : .env file in Root path
func InitEnvironment() {
	InitEnvWithPath("")
}

// InitEnvWithPath : テスト実行時のenvfileのpathを統一するための関数
func InitEnvWithPath(path string) {
	err := godotenv.Load(fmt.Sprintf(path+"envfiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
