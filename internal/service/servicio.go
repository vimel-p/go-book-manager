package service

import (
	"fmt"
	"sistema/internal/models"
	"sistema/internal/repository"
	"strings"
)

type ServicioLibros struct {
	repo repository.Repositorio
}

func NuevoServicio(repo repository.Repositorio) *ServicioLibros {
	return &ServicioLibros{
		repo: repo,
	}
}

func (s *ServicioLibros) AgregarLibro(titulo, autor string, anio int, precio float64, stock int) error {
	libro, err := models.NuevoLibro(0, titulo, autor, anio, precio, stock)
	if err != nil {
		return err
	}

	err = s.repo.Guardar(libro)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Libro agregado: %s\n", titulo)
	return nil
}

func (s *ServicioLibros) BuscarPorID(id int) (*models.Libro, error) {
	return s.repo.BuscarPorID(id)
}

func (s *ServicioLibros) MostrarCatalogo() {
	libros, err := s.repo.ListarTodos()
	if err != nil {
		fmt.Println("Error al leer catálogo:", err)
		return
	}

	if len(libros) == 0 {
		fmt.Println("No hay libros en el catálogo.")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Catálogo de Libros:")
	fmt.Println(strings.Repeat("=", 50))

	for i, libro := range libros {
		fmt.Printf("%d. ID: %d | %s - %s (%d) | $%.2f | Stock: %d\n",
			i+1,
			libro.GetID(),
			libro.GetTitulo(),
			libro.GetAutor(),
			libro.GetAnio(),
			libro.GetPrecio(),
			libro.GetStock())
	}

	fmt.Printf("\n📊 Total de libros: %d\n", len(libros))
}

func (s *ServicioLibros) Buscar(titulo string) {
	fmt.Printf("\n🔍 Buscando: \"%s\"\n", titulo)
	resultados, err := s.repo.BuscarPorTitulo(titulo)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if len(resultados) == 0 {
		fmt.Println("No se encontraron libros.")
		return
	}

	fmt.Printf("Resultados encontrados: %d\n", len(resultados))
	for i, libro := range resultados {
		fmt.Printf("%d. ID: %d | %s - %s | $%.2f | Stock: %d\n",
			i+1,
			libro.GetID(),
			libro.GetTitulo(),
			libro.GetAutor(),
			libro.GetPrecio(),
			libro.GetStock())
	}
}

func (s *ServicioLibros) ComprarLibro(id int, cantidad int) error {
	fmt.Printf("\n🛒 Comprando libro ID: %d, Cantidad: %d\n", id, cantidad)

	if cantidad <= 0 {
		return fmt.Errorf("la cantidad debe ser mayor a cero")
	}

	libro, err := s.repo.BuscarPorID(id)
	if err != nil {
		return err
	}

	err = libro.ReducirStock(cantidad)
	if err != nil {
		return err
	}

	s.repo.Actualizar(libro)

	fmt.Printf("Compra exitosa: %d x %s\n", cantidad, libro.GetTitulo())
	fmt.Printf("Stock restante: %d\n", libro.GetStock())
	return nil
}

func (s *ServicioLibros) MostrarEstadisticas() {
	total, promedio, stockTotal, err := s.repo.ObtenerEstadisticas()
	if err != nil {
		fmt.Println("\n📊 Error al calcular estadísticas:", err)
		return
	}

	if total == 0 {
		fmt.Println("\n📊 No hay datos para estadísticas")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📊 ESTADÍSTICAS")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total de libros: %d\n", total)
	fmt.Printf("Precio promedio: $%.2f\n", promedio)
	fmt.Printf("Stock total: %d unidades\n", stockTotal)
}

func (s *ServicioLibros) Repo() repository.Repositorio {
	return s.repo
}
