package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//func Connect(databseUrl string) (*pgxpool.Pool, error) {
//
//	// context = objek yang membawa informasi request (timeout, cancel operation)
//	var ctx context.Context = context.Background()
//	var config *pgxpool.Config
//
//	var err error
//
//	config, err = pgxpool.ParseConfig(databseUrl)
//	if err != nil {
//		log.Printf("Error parse DATABASE URL %s\n", err)
//		return nil, err
//	}
//
//	var pool *pgxpool.Pool
//	pool, err = pgxpool.NewWithConfig(ctx, config)
//	if err != nil {
//		log.Printf("Unable to create connection pool %s\n", err)
//		return nil, err
//	}
//
//	// test koneksi = select 1 or datasource.getConnection()
//	err = pool.Ping(ctx)
//	if err != nil {
//		log.Printf("Unable to ping DATABASE URL %s\n", err)
//		pool.Close()
//		return nil, err
//	}
//
//	log.Println("Successfully connected to DATABASE URL")
//	return pool, nil
//}

func Connect(databaseUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Printf("Unable to connect to database: %s\n", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Unable to get underlying sql.DB: %s\n", err)
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Printf("Unable to ping database: %s\n", err)
		return nil, err
	}

	log.Println("Successfully connected to DATABASE URL")
	return db, nil
}
