package repository

import (
	"database/sql"
	"fmt"
	"sistema/internal/models"
	"strings"
)

type RepositorioSQLite struct {
	db *sql.DB
}

func NuevoRepositorioSQLite(db *sql.DB) *RepositorioSQLite {
	return &RepositorioSQLite{db: db}
}

// Guardar inserta un nuevo libro
func (r *RepositorioSQLite) Guardar(libro *models.Libro) error {
	query := `INSERT INTO libros (titulo, autor, anio, precio, stock) 
              VALUES (?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
		libro.GetTitulo(),
		libro.GetAutor(),
		libro.GetAnio(),
		libro.GetPrecio(),
		libro.GetStock())

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	libro.SetID(int(id))

	return nil
}

// BuscarPorID obtiene un libro
func (r *RepositorioSQLite) BuscarPorID(id int) (*models.Libro, error) {
	query := `SELECT id, titulo, autor, anio, precio, stock 
              FROM libros WHERE id = ?`

	row := r.db.QueryRow(query, id)

	var libro models.Libro
	var idLibro, anio, stock int
	var titulo, autor string
	var precio float64

	err := row.Scan(&idLibro, &titulo, &autor, &anio, &precio, &stock)
	if err != nil {
		return nil, fmt.Errorf("libro no encontrado con ID: %d", id)
	}

	libro.SetID(idLibro)
	libro.SetTitulo(titulo)
	libro.SetAutor(autor)
	libro.SetAnio(anio)
	libro.SetPrecio(precio)
	libro.SetStock(stock)

	return &libro, nil
}

// BuscarPorTitulo busca libros
func (r *RepositorioSQLite) BuscarPorTitulo(titulo string) ([]*models.Libro, error) {
	query := `SELECT id, titulo, autor, anio, precio, stock 
              FROM libros WHERE LOWER(titulo) LIKE ?`

	rows, err := r.db.Query(query, "%"+strings.ToLower(titulo)+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Libro

	for rows.Next() {
		var libro models.Libro
		var id, anio, stock int
		var tit, autor string
		var precio float64

		err := rows.Scan(&id, &tit, &autor, &anio, &precio, &stock)
		if err != nil {
			continue
		}

		libro.SetID(id)
		libro.SetTitulo(tit)
		libro.SetAutor(autor)
		libro.SetAnio(anio)
		libro.SetPrecio(precio)
		libro.SetStock(stock)

		resultados = append(resultados, &libro)
	}

	return resultados, nil
}

// ListarTodos devuelve todos los libros
func (r *RepositorioSQLite) ListarTodos() ([]*models.Libro, error) {
	query := `SELECT id, titulo, autor, anio, precio, stock FROM libros`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*models.Libro

	for rows.Next() {
		var libro models.Libro
		var id, anio, stock int
		var tit, autor string
		var precio float64

		err := rows.Scan(&id, &tit, &autor, &anio, &precio, &stock)
		if err != nil {
			continue
		}

		libro.SetID(id)
		libro.SetTitulo(tit)
		libro.SetAutor(autor)
		libro.SetAnio(anio)
		libro.SetPrecio(precio)
		libro.SetStock(stock)

		lista = append(lista, &libro)
	}

	return lista, nil
}

// Actualizar modifica un libro
func (r *RepositorioSQLite) Actualizar(libro *models.Libro) error {
	query := `UPDATE libros SET titulo=?, autor=?, anio=?, precio=?, stock=? 
              WHERE id=?`

	_, err := r.db.Exec(query,
		libro.GetTitulo(),
		libro.GetAutor(),
		libro.GetAnio(),
		libro.GetPrecio(),
		libro.GetStock(),
		libro.GetID())

	return err
}

// ObtenerEstadisticas calcula totales
func (r *RepositorioSQLite) ObtenerEstadisticas() (int, float64, int, error) {
	query := `SELECT 
                COUNT(*) as total,
                AVG(precio) as promedio,
                SUM(stock) as stock_total
              FROM libros`

	row := r.db.QueryRow(query)

	var total int
	var promedio float64
	var stockTotal int

	err := row.Scan(&total, &promedio, &stockTotal)
	if err != nil {
		return 0, 0, 0, err
	}

	return total, promedio, stockTotal, nil
}

// Eliminar borra un libro
func (r *RepositorioSQLite) Eliminar(id int) error {
	query := `DELETE FROM libros WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}
