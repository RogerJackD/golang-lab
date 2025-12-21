package main

import "fmt"

// 1. Interfaz Command
type Comando interface {
	Ejecutar()
}

// 2. Receptor
type Luz struct{}

func (l *Luz) Encender() {
	fmt.Println("Luz encendida")
}

func (l *Luz) Apagar() {
	fmt.Println("Luz apagada")
}

// 3. Comandos concretos
type EncenderLuz struct {
	luz *Luz
}

func (c *EncenderLuz) Ejecutar() {
	c.luz.Encender()
}

type ApagarLuz struct {
	luz *Luz
}

func (c *ApagarLuz) Ejecutar() {
	c.luz.Apagar()
}

// 4. Invocador
type ControlRemoto struct {
	comando Comando
}

func (c *ControlRemoto) PresionarBoton() {
	c.comando.Ejecutar()
}

// 5. Cliente
func main() {
	luz := &Luz{}

	encender := &EncenderLuz{luz: luz}
	apagar := &ApagarLuz{luz: luz}

	control := &ControlRemoto{}

	control.comando = encender
	control.PresionarBoton()

	control.comando = apagar
	control.PresionarBoton()
}
