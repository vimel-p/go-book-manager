package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sistema/internal/database"
	"sistema/internal/handlers"
	"sistema/internal/repository"
	"sistema/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// 1. Conectar a la base de datos
	log.Println("📦 Conectando a base de datos...")
	db, err := database.NuevaConexion()
	if err != nil {
		log.Fatal("❌ Error conectando a BD:", err)
	}
	defer db.Close()
	log.Println("✅ Conectado a SQLite")

	// 2. Inicializar las capas
	repo := repository.NuevoRepositorioSQLite(db)
	servicio := service.NuevoServicio(repo)
	libroHandler := handlers.NuevoLibroHandler(servicio)

	// 3. Configurar el enrutador
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	// 4. Definir los endpoints
	api.HandleFunc("/libros", libroHandler.ListarLibros).Methods("GET")
	api.HandleFunc("/libros/{id}", libroHandler.ObtenerLibro).Methods("GET")
	api.HandleFunc("/libros/buscar", libroHandler.BuscarPorTitulo).Methods("GET")
	api.HandleFunc("/libros", libroHandler.CrearLibro).Methods("POST")
	api.HandleFunc("/libros/{id}", libroHandler.EliminarLibro).Methods("DELETE")
	api.HandleFunc("/libros/{id}/comprar", libroHandler.ComprarLibro).Methods("POST")
	api.HandleFunc("/estadisticas", libroHandler.ObtenerEstadisticas).Methods("GET")

	// Endpoint de salud
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"estado":  "ok",
			"mensaje": "API funcionando correctamente",
		})
	}).Methods("GET")

	// 5. Iniciar el servidor
	log.Println("🚀 Servidor web iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
