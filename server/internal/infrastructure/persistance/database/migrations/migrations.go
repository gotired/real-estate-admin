package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run(db *sql.DB) {
	dirPath := "/app/migrations"

	migrationFiles, err := listFiles(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range migrationFiles {
		if err := migrate(db, fmt.Sprintf("%s/%s", dirPath, file)); err != nil {
			log.Fatalf("Failed to apply migration %s: %v", file, err)
		}
	}
}

func migrate(db *sql.DB, filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	return err
}

func listFiles(dirPath string) ([]string, error) {
	var sqlFiles []string

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.Type().IsRegular() && strings.HasSuffix(entry.Name(), ".sql") {
			sqlFiles = append(sqlFiles, entry.Name())
		}
	}

	return sqlFiles, nil
}
