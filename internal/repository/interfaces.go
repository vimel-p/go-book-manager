package repository

import "sistema/internal/models"

type Repositorio interface {
	Guardar(libro *models.Libro) error
	BuscarPorID(id int) (*models.Libro, error)
	BuscarPorTitulo(titulo string) ([]*models.Libro, error)
	ListarTodos() ([]*models.Libro, error)
	Actualizar(libro *models.Libro) error
	ObtenerEstadisticas() (int, float64, int, error)
	Eliminar(id int) error // ← NUEVO MÉTODO AGREGADO
}
