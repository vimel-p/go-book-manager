package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NuevaConexion abre la base de datos y crea la tabla si no existe
func NuevaConexion() (*sql.DB, error) {
	// Abrir conexión a SQLite (crea el archivo si no existe)
	db, err := sql.Open("sqlite3", "./data/libros.db")
	if err != nil {
		return nil, err
	}

	// Crear la tabla si no existe
	err = crearTabla(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// crearTabla ejecuta el SQL para crear la tabla libros
func crearTabla(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS libros (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        titulo TEXT NOT NULL,
        autor TEXT NOT NULL,
        anio INTEGER NOT NULL,
        precio REAL NOT NULL,
        stock INTEGER NOT NULL
    );`

	_, err := db.Exec(query)
	return err
}
