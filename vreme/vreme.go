// Package vreme supports creating simulated weather stations, adding measurements to a chosen station,
// printing out all station measurements to command line, and creating and printing out a report of all existing stations
// to command line
//
// A station is defined by a unique identifier "oznaka", station name "ime", and location name "kraj"
//
// Usage example:
//
//	new_station = vreme.DodajPostajo(oznaka, ime, kraj)
//	vreme.DodajMeritev(oznaka, temperatura)
//
//	IzpisMeritev()
//	IzpisPorocila(min_st_meritev, t_min, t_max)
package vreme

import (
	"fmt"
	"math"
)

// Postaja is a struct which represents a station with the following attributes:
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// postaje is a global map structure which holds pairs of <k, v> = <oznaka, postaja>
var postaje = make(map[string]Postaja)

// DodajPostajo creates a station with given parameters
func DodajPostajo(oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{ime, kraj, make([]float64, 0)}
}

// DodajMeritev adds a measurement with value 'temperatura' to station with id 'oznaka'
func DodajMeritev(oznaka string, temperatura float64) {
	p_struct, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Println("Napaka: postaja z oznako", oznaka, "ne obstaja.")
		return
	}
	if temperatura < -50 || temperatura > 50 {
		fmt.Println("Napaka: temperatura", temperatura, "°C ni v dovoljenem območju.")
		return
	}
	temperatura = float64(math.Round(temperatura*100)) / 100
	p_struct.meritve = append(p_struct.meritve, temperatura)
	postaje[oznaka] = p_struct
}

// IzpisMeritev prints out all measurements of all existing stations
func IzpisMeritev() {
	fmt.Println("MERITVE:")
	for oznaka, p_struct := range postaje {
		s_kraj := "(" + p_struct.kraj + "):"
		fmt.Println(oznaka, "-", p_struct.ime, s_kraj, p_struct.meritve)
	}
}

// IzpisPorocila creates a report of all measurements recorded by all stations. This is more comprehensive
// than IzpisMeritev, as it prints out minimum, maximum and average temperatures as well as
// a description for each existing station
func IzpisPorocila(min_st_meritev int, t_min float64, t_max float64) {
	for oznaka, p_struct := range postaje {
		var opis string = ""
		min, max, avg, ok := statistika(oznaka, min_st_meritev, t_min, t_max)
		if !ok {
			fmt.Println(p_struct.ime, ": premalo meritev")
			continue
		}
		switch {
		case avg < 5:
			opis = "Mrzlo"
		case 5 <= avg && avg < 15:
			opis = "Hladno"
		case 15 <= avg && avg < 25:
			opis = "Toplo"
		case 25 <= avg:
			opis = "Vroče"
		}
		fmt.Println(p_struct.ime, ": min", min, "°C, max", max, "°C, povprečje", avg, "°C ->", opis)
	}
}

// statistika creates a statistic for a given station with id 'oznaka', used by IzpisPorocila to create a report
func statistika(oznaka string, min_st_meritev int, t_min float64, t_max float64) (najnizja, najvisja, povprecna float64, ok bool) {
	p_struct, obstaja := postaje[oznaka]
	ok = true
	if !obstaja {
		ok = false
		return
	}
	st_meritev := len(p_struct.meritve)
	if st_meritev < min_st_meritev {
		ok = false
		return
	}

	najnizja = t_max
	najvisja = t_min
	povprecna = 0
	for _, t := range p_struct.meritve {
		povprecna += t
		if t < najnizja {
			najnizja = t
		}
		if t > najvisja {
			najvisja = t
		}
	}
	povprecna /= float64(st_meritev)
	povprecna = float64(math.Round(povprecna*100)) / 100
	return
}
