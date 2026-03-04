# 📚 Sistema de Gestión de Libros Electrónicos en Go

**Asignatura:** Programación Orientada a Objetos 

## 🎯 Objetivo
Desarrollar un sistema completo de gestión de libros electrónicos que evoluciona desde una aplicación de consola hasta una API REST, integrando todos los conceptos de las 4 unidades: programación funcional, estructuras de datos, encapsulación, manejo de errores y servicios web con concurrencia.

## 🚀 Características
- Agregar, mostrar y buscar libros
- Calcular estadísticas básicas
- Basado en lo aprendido
- comprar basado en el stock libros


## 📦 FUNCIONALIDADES PRINCIPALES

### **Unidad 1: Fundamentos**
- ✅ Estructura Libro con campos básicos
- ✅ Funciones para agregar y listar libros
- ✅ Condicionales para validaciones
- ✅ Iteraciones para recorrer catálogo

### **Unidad 2: Estructuras de Datos**
- ✅ Slices para resultados de búsqueda
- ✅ Map para repositorio de libros (clave: ID)
- ✅ Búsquedas por ID y título parcial

### **Unidad 3: POO y Manejo de Errores**
- ✅ Encapsulación (campos privados)
- ✅ Getters y Setters con validaciones
- ✅ Manejo de errores con tipo `error`
- ✅ Interfaces para repositorio

### **Unidad 4: Servicios Web**
- ✅ 8 endpoints REST
- ✅ Serialización JSON
- ✅ Concurrencia con goroutines
- ✅ Pruebas unitarias y de integración

## 🌐 ENDPOINTS DE LA API (8 servicios web)

| Método | Endpoint | Descripción | Unidad |
|--------|----------|-------------|--------|
| GET | `/api/libros` | Listar todos los libros | 2 |
| GET | `/api/libros/{id}` | Buscar libro por ID | 2 |
| GET | `/api/libros/buscar?titulo={titulo}` | Buscar por título | 2 |
| POST | `/api/libros` | Agregar nuevo libro | 1 |
| PUT | `/api/libros/{id}` | Actualizar libro completo | 3 |
| PATCH | `/api/libros/{id}/stock` | Actualizar solo stock | 3 |
| DELETE | `/api/libros/{id}` | Eliminar libro | 1 |
| POST | `/api/libros/{id}/comprar` | Comprar libro (reduce stock) | 3 |
| GET | `/api/estadisticas` | Ver estadísticas del catálogo | 2 |
| GET | `/api/health` | Verificar estado del servidor | 4 |

## 🔧 INSTALACIÓN Y EJECUCIÓN

```bash
# Clonar repositorio
git clone https://github.com/tuusuario/gestion-libros.git
cd gestion-libros

# Ejecutar servidor web
go run cmd/api/main.go

# Ejecutar pruebas
go test ./tests/...
