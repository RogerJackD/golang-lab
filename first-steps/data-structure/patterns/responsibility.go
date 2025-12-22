package main

import "fmt"

// 1. Handler (interfaz)
type Manejador interface {
	SetSiguiente(Manejador)
	Manejar(nivel int)
}

// 2. Handler base
type BaseManejador struct {
	siguiente Manejador
}

func (b *BaseManejador) SetSiguiente(s Manejador) {
	b.siguiente = s
}

// 3. Manejadores concretos
type SoporteBasico struct {
	BaseManejador
}

func (s *SoporteBasico) Manejar(nivel int) {
	if nivel <= 1 {
		fmt.Println("Soporte básico resolvió el problema")
	} else if s.siguiente != nil {
		s.siguiente.Manejar(nivel)
	}
}

type SoporteAvanzado struct {
	BaseManejador
}

func (s *SoporteAvanzado) Manejar(nivel int) {
	if nivel <= 2 {
		fmt.Println("Soporte avanzado resolvió el problema")
	} else if s.siguiente != nil {
		s.siguiente.Manejar(nivel)
	}
}

type SoporteEspecialista struct {
	BaseManejador
}

func (s *SoporteEspecialista) Manejar(nivel int) {
	fmt.Println("Soporte especialista resolvió el problema")
}

// 4. Cliente
func main() {
	basico := &SoporteBasico{}
	avanzado := &SoporteAvanzado{}
	especialista := &SoporteEspecialista{}

	basico.SetSiguiente(avanzado)
	avanzado.SetSiguiente(especialista)

	basico.Manejar(3)
}
