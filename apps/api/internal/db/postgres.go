package db

import (
	"api/internal/config"
	"api/internal/logger"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// uppercase pulbic variable, lowcase private variable
// It's a connection pool manager
var DB *sql.DB;

func InitDB() {
	// postgres://username:password@host:port/dbname

	str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",  
		config.Appcfg.DBUser,
		config.Appcfg.DBPassword,
		config.Appcfg.DBHost,
		config.Appcfg.DBPort,
		config.Appcfg.DBName,
		config.Appcfg.DBSSLMode,
	)
	var err error
	DB, err = sql.Open("postgres", str)

	if err != nil {
		logger.Log.Error("Error in opening DB", zap.Error(err))
	}

	err = DB.Ping()

	if err != nil {
		logger.Log.Error("Error in connecting DB...", zap.Error(err))
	}

	logger.Log.Info("Connected to DB");
}

