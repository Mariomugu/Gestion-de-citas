package citas

import "time"

type Servicio struct {
	nombre       string
	duracion     time.Duration
	equipamiento Equipamiento
}
