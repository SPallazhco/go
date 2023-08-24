package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p("Date: ", now)
	then := time.Date(2019, 11, 17, 20, 34, 58, 0030471, time.UTC)
	p(then)
	then = then.Add(time.Hour * 1) // Si multiplico por 24 seria un dia
	p(then)
	then = then.Add(-time.Hour * 1)
	p(then)
	p()
	p(then.Year())
	p(then.Month())
	p(then.Day()) // Lo mismo para segundos minutos horas location weekday

	//Comparaciones de fechas
	p()
	p("Then is before: ", then.Before(now))
	p("Then is after: ", then.After(now))
	p("Then is equal: ", then.Equal(now))

	// Diferencia de fechas
	p()
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

}
