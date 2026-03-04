# 📚 Sistema de Gestión de Libros Electrónicos

**Autor:** Vianka Melina Paredes Rivas  
**Repositorio:** [github.com/vimel-p/go-book-manager](https://github.com/vimel-p/go-book-manager)

## 📋 DESCRIPCIÓN
Sistema completo de gestión de libros electrónicos con API REST desarrollado en Go. Implementa conceptos de programación funcional, estructuras de datos, encapsulación, manejo de errores y concurrencia.

## 🚀 TECNOLOGÍAS
- Go 1.x
- SQLite
- Gorilla Mux
- Concurrencia con goroutines

## 🌐 ENDPOINTS DE LA API

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/health` | Verificar estado del servidor |
| GET | `/api/libros` | Listar todos los libros |
| GET | `/api/libros/{id}` | Obtener libro por ID |
| GET | `/api/libros/buscar?titulo=` | Buscar por título |
| POST | `/api/libros` | Crear nuevo libro |
| DELETE | `/api/libros/{id}` | Eliminar libro |
| POST | `/api/libros/{id}/comprar` | Comprar libro (reduce stock) |
| GET | `/api/estadisticas` | Ver estadísticas del catálogo |

## 🔧 INSTALACIÓN Y EJECUCIÓN

```bash
# Clonar el repositorio
git clone https://github.com/vimel-p/go-book-manager.git
cd go-book-manager

# Instalar dependencias
go mod tidy

# Ejecutar servidor
go run cmd/api/main.go
# Ejecutar servidor web
go run cmd/api/main.go

# Ejecutar pruebas
go test ./tests/...
