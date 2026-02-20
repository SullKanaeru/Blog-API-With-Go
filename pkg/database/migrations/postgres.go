package database

import (
	"log"
	"blog_api/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := config.GetEnv("DATABASE_URL", "")

	db, err := gorm.Open(gorm_postgres.New(gorm_postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	log.Println("Database terhubung! Menjalankan migrasi SQL...")

	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gagal mengambil instance sql.DB:", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal("Gagal membuat driver migrasi:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./pkg/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("Gagal inisialisasi migrasi:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Gagal menjalankan file migrasi up:", err)
	}

	log.Println("Migrasi database berhasil/up-to-date!")
}