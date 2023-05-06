package postgres

import (
	"embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed sql
var sqlFolder embed.FS

// LoadSQLFile loads a SQL file containing the given file name. This returns the contents of the SQL file, or an error
// indicating why the SQL file could not be loaded. An error will also be returned if the given fileName is empty or
// blank.
func LoadSQLFile(fileName string) (string, error) {
	if fileName = strings.TrimSpace(fileName); fileName == "" {
		return "", errors.New("fileName cannot be empty or blank")
	}

	filePath := fmt.Sprintf("sql/%s.sql", fileName)
	contents, err := sqlFolder.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}
