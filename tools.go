//go:build tools

package tools

import (
	// Goa
	_ "goa.design/goa/v3/cmd/goa"

	// Gorm
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"

	// DotEnv
	_ "github.com/joho/godotenv"
)
