package handlers

import (
	"encoding/json"
	"net/http"
	"sistema/internal/models"
	"sistema/internal/service"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type LibroHandler struct {
	servicio *service.ServicioLibros
}

func NuevoLibroHandler(servicio *service.ServicioLibros) *LibroHandler {
	return &LibroHandler{servicio: servicio}
}

// GET /api/libros
func (h *LibroHandler) ListarLibros(w http.ResponseWriter, r *http.Request) {
	libros, err := h.servicio.Repo().ListarTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}

// GET /api/libros/{id}
func (h *LibroHandler) ObtenerLibro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	libro, err := h.servicio.Repo().BuscarPorID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libro)
}

// POST /api/libros
func (h *LibroHandler) CrearLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro

	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if libro.Titulo == "" {
		http.Error(w, "El título es obligatorio", http.StatusBadRequest)
		return
	}

	err = h.servicio.AgregarLibro(
		libro.Titulo,
		libro.Autor,
		libro.Anio,
		libro.Precio,
		libro.Stock,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Libro creado exitosamente",
	})
}

// GET /api/libros/buscar?titulo=...
func (h *LibroHandler) BuscarPorTitulo(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Query().Get("titulo")
	if titulo == "" {
		http.Error(w, "El parámetro 'titulo' es requerido", http.StatusBadRequest)
		return
	}

	resultados, err := h.servicio.Repo().BuscarPorTitulo(titulo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}

// GET /api/estadisticas
func (h *LibroHandler) ObtenerEstadisticas(w http.ResponseWriter, r *http.Request) {
	total, promedio, stockTotal, err := h.servicio.Repo().ObtenerEstadisticas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	estadisticas := map[string]interface{}{
		"total_libros":    total,
		"precio_promedio": promedio,
		"stock_total":     stockTotal,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estadisticas)
}

// DELETE /api/libros/{id}
func (h *LibroHandler) EliminarLibro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.servicio.Repo().Eliminar(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Libro eliminado exitosamente",
	})
}

// POST /api/libros/{id}/comprar (CON GOROUTINE)
func (h *LibroHandler) ComprarLibro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Obtener cantidad del query string (default: 1)
	cantidad := 1
	cantidadStr := r.URL.Query().Get("cantidad")
	if cantidadStr != "" {
		cantidad, _ = strconv.Atoi(cantidadStr)
	}

	// Canal para recibir el resultado de la goroutine
	resultado := make(chan error)

	// --- GOROUTINE (CONCURRENCIA) ---
	go func() {
		// Simulamos que la compra toma tiempo
		time.Sleep(100 * time.Millisecond)

		err := h.servicio.ComprarLibro(id, cantidad)
		resultado <- err
	}()
	// ---------------------------------

	// Esperar el resultado con timeout
	select {
	case err := <-resultado:
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"mensaje":  "Compra exitosa",
			"id_libro": id,
			"cantidad": cantidad,
		})

	case <-time.After(3 * time.Second):
		http.Error(w, "Tiempo de espera agotado", http.StatusRequestTimeout)
	}
}
