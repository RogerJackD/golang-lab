package main

import (
	"fmt"
	"sync"
)

// Estructura del Singleton
type Config struct {
	Nombre string
}

var (
	instancia *Config
	once      sync.Once
)

// Función que devuelve la única instancia
func GetConfig() *Config {
	once.Do(func() {
		instancia = &Config{
			Nombre: "Configuración Global",
		}
	})
	return instancia
}

func main() {
	// Ambas variables apuntan al mismo objeto
	c1 := GetConfig()
	c2 := GetConfig()

	fmt.Println("Nombre:", c1.Nombre)
	fmt.Println("¿Es la misma instancia?", c1 == c2) // true

	// Si modificas una, la otra también cambia
	c1.Nombre = "Nueva Config"
	fmt.Println("Nombre en c2:", c2.Nombre)
}
