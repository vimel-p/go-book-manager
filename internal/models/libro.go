package models

import "errors"

type Libro struct {
	ID     int     `json:"id"`
	Titulo string  `json:"titulo"`
	Autor  string  `json:"autor"`
	Precio float64 `json:"precio"`
	Anio   int     `json:"anio"`
	Stock  int     `json:"stock"`
}

// Getters
func (l *Libro) GetID() int {
	return l.ID
}

func (l *Libro) GetTitulo() string {
	return l.Titulo
}

func (l *Libro) GetAutor() string {
	return l.Autor
}

func (l *Libro) GetAnio() int {
	return l.Anio
}

func (l *Libro) GetPrecio() float64 {
	return l.Precio
}

func (l *Libro) GetStock() int {
	return l.Stock
}

// Setters
func (l *Libro) SetID(id int) {
	l.ID = id
}

func (l *Libro) SetTitulo(titulo string) error {
	if titulo == "" {
		return errors.New("el título no puede estar vacío")
	}
	l.Titulo = titulo
	return nil
}

func (l *Libro) SetAutor(autor string) error {
	if autor == "" {
		return errors.New("el autor no puede estar vacío")
	}
	l.Autor = autor
	return nil
}

func (l *Libro) SetAnio(anio int) error {
	if anio < 1900 || anio > 2026 {
		return errors.New("año inválido")
	}
	l.Anio = anio
	return nil
}

func (l *Libro) SetPrecio(precio float64) error {
	if precio <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	l.Precio = precio
	return nil
}

func (l *Libro) SetStock(stock int) error {
	if stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	l.Stock = stock
	return nil
}

// Métodos de stock
func (l *Libro) TieneStock(cantidad int) bool {
	if l.Stock <= 0 {
		return false
	}
	return l.Stock >= cantidad
}

func (l *Libro) ReducirStock(cantidad int) error {
	if !l.TieneStock(cantidad) {
		return errors.New("no hay suficiente stock")
	}
	l.Stock -= cantidad
	return nil
}

// Constructor
func NuevoLibro(id int, titulo, autor string, anio int, precio float64, stock int) (*Libro, error) {
	if titulo == "" {
		return nil, errors.New("el título no puede estar vacío")
	}
	if autor == "" {
		return nil, errors.New("el autor no puede estar vacío")
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
		ID:     id,
		Titulo: titulo,
		Autor:  autor,
		Precio: precio,
		Anio:   anio,
		Stock:  stock,
	}, nil
}
