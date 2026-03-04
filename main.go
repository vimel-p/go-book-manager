package main

import (
	"fmt"
	"sistema/internal/database"
	"sistema/internal/repository"
	"sistema/internal/service"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("📚 SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS")
	fmt.Println(strings.Repeat("=", 50))

	// Conectar a base de datos
	fmt.Println("\n📦 Conectando a base de datos...")
	db, err := database.NuevaConexion()
	if err != nil {
		fmt.Printf("❌ Error conectando a BD: %v\n", err)
		return
	}
	defer db.Close()
	fmt.Println("✅ Conectado a SQLite (data/libros.db)")

	// Inicializar sistema
	fmt.Println("\n📦 Inicializando sistema...")
	repo := repository.NuevoRepositorioSQLite(db)
	servicio := service.NuevoServicio(repo)

	// Agregar libros de ejemplo
	fmt.Println("\n📝 Agregando libros al catálogo...")

	var err2 error

	err2 = servicio.AgregarLibro("El principito", "Antoine de Saint-Exupéry", 1943, 9.99, 10)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	err2 = servicio.AgregarLibro("1984", "George Orwell", 1949, 14.99, 5)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	err2 = servicio.AgregarLibro("Cien años de soledad", "Gabriel García Márquez", 1967, 19.99, 3)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	err2 = servicio.AgregarLibro("El Hobbit", "J.R.R. Tolkien", 1937, 24.99, 7)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	err2 = servicio.AgregarLibro("Fahrenheit 451", "Ray Bradbury", 1953, 12.99, 4)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	// Mostrar catálogo
	servicio.MostrarCatalogo()

	// Búsquedas
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🔍 DEMOSTRACIÓN DE BÚSQUEDAS")
	fmt.Println(strings.Repeat("=", 50))

	servicio.Buscar("1984")
	servicio.Buscar("hobbit")
	servicio.Buscar("Harry Potter")

	// Compras
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🛒 DEMOSTRACIÓN DE COMPRAS")
	fmt.Println(strings.Repeat("=", 50))

	servicio.ComprarLibro(2, 2)

	err2 = servicio.ComprarLibro(3, 10)
	if err2 != nil {
		fmt.Printf("❌ Error controlado: %v\n", err2)
	}

	err2 = servicio.ComprarLibro(99, 1)
	if err2 != nil {
		fmt.Printf("❌ Error controlado: %v\n", err2)
	}

	// Estadísticas
	servicio.MostrarEstadisticas()

	// Catálogo actualizado
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📚 CATÁLOGO ACTUALIZADO")
	fmt.Println(strings.Repeat("=", 50))
	servicio.MostrarCatalogo()

	// Validaciones
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("⚠️ DEMOSTRACIÓN DE VALIDACIONES")
	fmt.Println(strings.Repeat("=", 50))

	err2 = servicio.AgregarLibro("", "Autor", 2000, 10.99, 5)
	if err2 != nil {
		fmt.Printf("❌ Error al crear libro sin título: %v\n", err2)
	}

	err2 = servicio.AgregarLibro("Libro", "Autor", 1800, 10.99, 5)
	if err2 != nil {
		fmt.Printf("❌ Error al crear libro con año inválido: %v\n", err2)
	}

	err2 = servicio.AgregarLibro("Libro", "Autor", 2000, -5, 5)
	if err2 != nil {
		fmt.Printf("❌ Error al crear libro con precio negativo: %v\n", err2)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("✅ PROGRAMA FINALIZADO CORRECTAMENTE")
	fmt.Println(strings.Repeat("=", 50))
}
