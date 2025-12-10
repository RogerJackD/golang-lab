package main

import "fmt"

type Reproductor interface {
	ReproducirMP3(nombre string)
}

type ReproductorVLC struct{}

func (r *ReproductorVLC) ReproducirVLC(nombre string) {
	fmt.Println("Reproduciendo archivo VLC:", nombre)
}

// Adaptador
type AdaptadorVLC struct {
	vlc *ReproductorVLC
}

func (a *AdaptadorVLC) ReproducirMP3(nombre string) {
	a.vlc.ReproducirVLC(nombre)
}

func main() {
	var reproductor Reproductor

	vlc := &ReproductorVLC{}

	reproductor = &AdaptadorVLC{vlc: vlc}

	reproductor.ReproducirMP3("video_concierto.vlc")
}
