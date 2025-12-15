package main

import "fmt"

// 1. Interfaz Observer
type Observer interface {
	Actualizar(mensaje string)
}

// 2. Sujeto (Observable)
type Notificador struct {
	observers []Observer
}

// Registrar observadores
func (n *Notificador) Suscribir(o Observer) {
	n.observers = append(n.observers, o)
}

// Notificar a todos
func (n *Notificador) Notificar(mensaje string) {
	for _, o := range n.observers {
		o.Actualizar(mensaje)
	}
}

// 3. Observadores concretos
type Usuario struct {
	Nombre string
}

func (u *Usuario) Actualizar(mensaje string) {
	fmt.Println(u.Nombre, "recibió:", mensaje)
}

// 4. Cliente
func main() {
	notificador := &Notificador{}

	u1 := &Usuario{Nombre: "Ana"}
	u2 := &Usuario{Nombre: "Carlos"}

	notificador.Suscribir(u1)
	notificador.Suscribir(u2)

	notificador.Notificar("Nuevo mensaje disponible")
}
