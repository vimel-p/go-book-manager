package main

import "fmt"

// 1. Estructura del libro
type Libro struct {
    id     int
    titulo string
    autor  string
    precio float64
}

// 2. Base de datos en memoria
var libros []Libro
var siguienteID = 1

// 3. FUNCIONES PRINCIPALES

// Agregar nuevo libro
func agregarLibro(t, a string, p float64) {
    nuevo := Libro{siguienteID, t, a, p}
    libros = append(libros, nuevo)
    siguienteID++
    fmt.Println("✅ Libro agregado:", t)
}

// Mostrar todos los libros
func mostrarLibros() {
    if len(libros) == 0 {
        fmt.Println("📭 No hay libros")
        return
    }
    fmt.Println("\n📚 LIBROS EN CATÁLOGO:")
    for _, l := range libros {
        fmt.Printf("%d. %s - %s ($%.2f)\n", l.id, l.titulo, l.autor, l.precio)
    }
}

// Buscar por título
func buscarTitulo(t string) {
    fmt.Println("\n🔍 Buscando:", t)
    encontrado := false
    for _, l := range libros {
        if l.titulo == t {
            fmt.Printf("   Encontrado: %s - %s\n", l.titulo, l.autor)
            encontrado = true
        }
    }
    if !encontrado {
        fmt.Println("   No encontrado")
    }
}

// Calcular el promedio
func promedioPrecios() float64 {
    if len(libros) == 0 {
        return 0
    }
    total := 0.0
    for _, l := range libros {
        total += l.precio
    }
    return total / float64(len(libros))
}

// 4. PROGRAMA PRINCIPAL
func main() {
    fmt.Println("=== SISTEMA DE LIBROS ELECTRÓNICOS ===")
    
    // Agregar algunos libros
    agregarLibro("El principito", "Antoine", 9.99)
    agregarLibro("1984", "George Orwell", 14.99)
    agregarLibro("Cien años de soledad", "Gabo", 19.99)
    
    // Mostrar catálogo
    mostrarLibros()
    
    // Buscar
    buscarTitulo("1984")
    buscarTitulo("Harry Potter")
    
    // Estadística
    fmt.Printf("\n💰 Precio promedio: $%.2f\n", promedioPrecios())
    fmt.Printf("📊 Total de libros: %d\n", len(libros))
}
