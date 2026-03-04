// Vianka Melina Paredes Rivas
// Curso: Programación Orientada a Objetos
// Descripción: 2do avance en la implementación de un sistema básico de gestión de libros en Go.

// habia un error por no usar package main, lo agregue para que funcione el programa
package main

//agregue import para usar fmt, errors, strings en el main
import (
	"errors"
	"fmt"
	"strings"
)

// ESTRUCTURAS BÁSICAS
type Libro struct {
	id     int
	titulo string
	autor  string
	precio float64
	anio   int
	stock  int // para controlar el stock de libros
}

// getters para acceder a los atributos
func (l *Libro) GetID() int {
	return l.id
}

func (l *Libro) GetTitulo() string {
	return l.titulo
}

func (l *Libro) GetAutor() string {
	return l.autor
}

func (l *Libro) GetAnio() int {
	return l.anio
}

func (l *Libro) GetPrecio() float64 {
	return l.precio
}

func (l *Libro) GetStock() int {
	return l.stock
}

// setters para modificar los artributos con validaciones
func (l *Libro) SetTitulo(nuevoTitulo string) error {
	if nuevoTitulo == "" {
		return errors.New("el título no puede estar vacío")
	}
	l.titulo = nuevoTitulo
	return nil
}

func (l *Libro) SetAutor(nuevoAutor string) error {
	if nuevoAutor == "" {
		return errors.New("el autor no puede estar vacío")
	}
	l.autor = nuevoAutor
	return nil
}

func (l *Libro) SetPrecio(nuevoPrecio float64) error {
	if nuevoPrecio <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	l.precio = nuevoPrecio
	return nil
}

func (l *Libro) SetStock(nuevoStock int) error {
	if nuevoStock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	l.stock = nuevoStock
	return nil
}

// aqui se añadio condicionales para verificar el stock de libros

func (l *Libro) tieneStock(cantidad int) bool {
	if l.stock <= 0 {
		return false
	}
	return l.stock >= cantidad
}

func (l *Libro) reducirStock(cantidad int) error {
	if !l.tieneStock(cantidad) {
		return errors.New("no hay suficiente stock")
	}
	l.stock -= cantidad
	return nil
}

// una funcion de constructor para crear un nuevo libro

func nuevoLibro(id int, titulo, autor string, anio int, precio float64, stock int) (*Libro, error) {
	if titulo == "" {
		return nil, errors.New("el título no puede estar vacio")
	}
	if autor == "" {
		return nil, errors.New("el autor no puede estar vacio")
	}
	if anio < 1900 || anio > 2026 {
		return nil, errors.New("el año debe estar entre 1900 y 2026")
	}
	if precio <= 0 {
		return nil, errors.New("el precio no puede ser menor a 0")
	}
	if stock < 0 {
		return nil, errors.New("el stock no puede ser negativo")
	}

	return &Libro{
		id:     id,
		titulo: titulo,
		autor:  autor,
		precio: precio,
		anio:   anio,
		stock:  stock,
	}, nil
}

// aqui hare una estrutura para manejar el inventario de libros y un map para guardarlos
type repositorioLibros struct {
	libros map[int]*Libro
	nextID int
}

func nuevoRepositorio() *repositorioLibros {
	return &repositorioLibros{
		libros: make(map[int]*Libro),
		nextID: 1,
	}
}

// aqui guardare los libros y sino tiene id se le asignara
func (r *repositorioLibros) guardar(libro *Libro) error {
	if libro == nil {
		return errors.New("el libro no puede ser nulo")
	}

	if libro.GetID() == 0 {
		libro.id = r.nextID
		r.nextID++
	}

	r.libros[libro.GetID()] = libro
	return nil
}

// buscar un libro por id
func (r *repositorioLibros) buscarPorID(id int) (*Libro, error) {
	libro, existe := r.libros[id]
	if !existe {
		return nil, fmt.Errorf("libro no encontrado con ID: %d", id)
	}
	return libro, nil
}

// buscar libro por titulo
func (r *repositorioLibros) buscarPorTitulo(titulo string) ([]*Libro, error) {
	if titulo == "" {
		return nil, errors.New("el título no puede estar vacío")
	}
	var resultados []*Libro
	titulobusqueda := strings.ToLower(titulo)
	for _, libro := range r.libros {
		if strings.Contains(strings.ToLower(libro.GetTitulo()), titulobusqueda) {
			resultados = append(resultados, libro)
		}
	}
	return resultados, nil
}

// listar todos los libros como un slice
func (r *repositorioLibros) listarTodos() []*Libro {
	lista := make([]*Libro, 0, len(r.libros))
	for _, libro := range r.libros {
		lista = append(lista, libro)
	}
	return lista
}

// esta es una estructura que manejara la logica
type servicioLibros struct {
	repositorio *repositorioLibros
}

func nuevoServicio(repositorio *repositorioLibros) *servicioLibros {
	return &servicioLibros{
		repositorio: repositorio,
	}
}

// aqui hare las funciones del servicio como agregar
func (s *servicioLibros) agregarLibro(titulo, autor string, anio int, precio float64, stock int) error {
	libro, err := nuevoLibro(0, titulo, autor, anio, precio, stock)
	if err != nil {
		return err
	}

	err = s.repositorio.guardar(libro)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Libro agregado: %s\n", titulo)
	return nil
}

// buscar por id
func (s *servicioLibros) buscarPorID(id int) (*Libro, error) {
	return s.repositorio.buscarPorID(id)
}

// aqui pondre un mostrar catalogo
func (s *servicioLibros) mostrarCatalogo() {
	libros := s.repositorio.listarTodos()

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

// buscar libros por titulo
func (s *servicioLibros) buscar(titulo string) {
	fmt.Printf("\n🔍 Buscando: \"%s\"\n", titulo)
	resultados, err := s.repositorio.buscarPorTitulo(titulo)
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

// aqui hare una funcion para comprar un libro
func (s *servicioLibros) comprarLibro(id int, cantidad int) error {
	fmt.Printf("\n🛒 Comprando libro ID: %d, Cantidad: %d\n", id, cantidad)

	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	libro, err := s.repositorio.buscarPorID(id)
	if err != nil {
		return err
	}

	err = libro.reducirStock(cantidad)
	if err != nil {
		return err
	}

	fmt.Printf("Compra exitosa: %d x %s\n", cantidad, libro.GetTitulo())
	fmt.Printf("Stock restante: %d\n", libro.GetStock())
	return nil
}

// aqui muestra información resumida
func (s *servicioLibros) mostrarEstadisticas() {
	libros := s.repositorio.listarTodos()

	if len(libros) == 0 {
		fmt.Println("\n📊 No hay datos para estadísticas")
		return
	}

	// Calcular promedio de precios
	totalPrecios := 0.0
	totalStock := 0

	for _, libro := range libros {
		totalPrecios += libro.GetPrecio()
		totalStock += libro.GetStock()
	}

	promedio := totalPrecios / float64(len(libros))

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📊 ESTADÍSTICAS")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total de libros: %d\n", len(libros))
	fmt.Printf("Precio promedio: $%.2f\n", promedio)
	fmt.Printf("Stock total: %d unidades\n", totalStock)
}

// el programa principal muestra todas las funcionalidades del sistema
func main() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("📚 SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS")
	fmt.Println(strings.Repeat("=", 50))

	// iniciamos el sistema
	fmt.Println("\n📦 Inicializando sistema...")
	repositorio := nuevoRepositorio()
	servicio := nuevoServicio(repositorio)

	// añadire algunos libros de ejemplo
	fmt.Println("\n📝 Agregando libros al catálogo...")

	var err error

	err = servicio.agregarLibro("El principito", "Antoine de Saint-Exupéry", 1943, 9.99, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = servicio.agregarLibro("1984", "George Orwell", 1949, 14.99, 5)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = servicio.agregarLibro("Cien años de soledad", "Gabriel García Márquez", 1967, 19.99, 3)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = servicio.agregarLibro("El Hobbit", "J.R.R. Tolkien", 1937, 24.99, 7)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = servicio.agregarLibro("Fahrenheit 451", "Ray Bradbury", 1953, 12.99, 4)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// veremos el caalogo completo
	servicio.mostrarCatalogo()

	// veremos las busquedas
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🔍 DEMOSTRACIÓN DE BÚSQUEDAS")
	fmt.Println(strings.Repeat("=", 50))

	// haremos busquedas para mostrar el manejo de errores y resultados
	servicio.buscar("1984")
	servicio.buscar("hobbit")
	servicio.buscar("Harry Potter")

	// veremos el manejo de stock
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🛒 DEMOSTRACIÓN DE COMPRAS")
	fmt.Println(strings.Repeat("=", 50))

	// Compra exitosa
	servicio.comprarLibro(2, 2) // Compra 2 unidades del libro ID 2

	// Intento de compra con stock insuficiente
	err = servicio.comprarLibro(3, 10) // Quiere 10, pero hay 3
	if err != nil {
		fmt.Printf("❌ Error controlado: %v\n", err)
	}

	// Intento de compra de libro inexistente
	err = servicio.comprarLibro(99, 1)
	if err != nil {
		fmt.Printf("❌ Error controlado: %v\n", err)
	}

	// mostrar estadísticas
	servicio.mostrarEstadisticas()

	// mostrar catálogo actualizado
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📚 CATÁLOGO ACTUALIZADO")
	fmt.Println(strings.Repeat("=", 50))
	servicio.mostrarCatalogo()

	// veremos validaciones
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("⚠️ DEMOSTRACIÓN DE VALIDACIONES")
	fmt.Println(strings.Repeat("=", 50))

	err = servicio.agregarLibro("", "Autor", 2000, 10.99, 5)
	if err != nil {
		fmt.Printf("❌ Error al crear libro sin título: %v\n", err)
	}

	err = servicio.agregarLibro("Libro", "Autor", 1800, 10.99, 5)
	if err != nil {
		fmt.Printf("❌ Error al crear libro con año inválido: %v\n", err)
	}

	err = servicio.agregarLibro("Libro", "Autor", 2000, -5, 5)
	if err != nil {
		fmt.Printf("❌ Error al crear libro con precio negativo: %v\n", err)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("✅ PROGRAMA FINALIZADO CORRECTAMENTE")
	fmt.Println(strings.Repeat("=", 50))
}
