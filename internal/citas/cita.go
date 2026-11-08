package citas

import "time"

type Cita struct {
	clienta  string
	servicio Servicio
	inicio   time.Time
}
