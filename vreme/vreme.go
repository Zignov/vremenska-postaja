// Package vreme omogoča upravljanje vremenskih postaj,
// dodajanje temperaturnih meritev in izračun statistike.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z imenom,
// krajem in temperaturnimi meritvami.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo vremensko postajo v slovar postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}

// DodajMeritev doda temperaturno meritev, če postaja obstaja
// in je temperatura znotraj dovoljenega območja.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	e, ok := postaje[oznaka]

	if !ok {
		fmt.Println("Postaja z oznako", oznaka, "ne obstaja.")
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Meritev temperature", temperatura, "je izven dovoljenega območja.")
		return
	}

	e.meritve = append(e.meritve, temperatura)
	postaje[oznaka] = e
	fmt.Println("Meritev temperature", temperatura, "je bila uspešno dodana za postajo", oznaka)
}

// statistika izračuna najnižjo, najvišjo in povprečno
// temperaturo, če ima postaja dovolj meritev.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	e, ok := postaje[oznaka]
	st := len(e.meritve)
	if st < stMeritev || st == 0 {
		ok = false
	}

	if !ok {
		fmt.Println("Postaja z oznako", oznaka, "ne obstaja ali nima dovolj meritev.")
		return
	}

	najnizja = e.meritve[0]
	najvisja = e.meritve[0]
	sestevek := 0.0

	for _, temperatura := range e.meritve {
		if temperatura < najnizja {
			najnizja = temperatura
		}
		if temperatura > najvisja {
			najvisja = temperatura
		}
		sestevek += temperatura
	}

	povprecna = sestevek / float64(st)
	return najnizja, najvisja, povprecna, ok
}

// IzpisMeritev izpiše vse vremenske postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERTIVE:")

	for postaja := range postaje {
		fmt.Println("Postaja:", postaja)
		fmt.Println("Ime:", postaje[postaja].ime)
		fmt.Println("Kraj:", postaje[postaja].kraj)
		fmt.Println("Meritve:", postaje[postaja].meritve)
		fmt.Println()
	}

}

// IzpisPorocila izpiše statistiko postaj, ki imajo
// zadostno število meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	fmt.Println("POROČILO:")

	for postaja := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, postaja, stMeritev)
		if !ok {
			continue
		}

		fmt.Print("Ime: ", postaje[postaja].ime, ", ")
		fmt.Print("Min: ", najnizja, ", ")
		fmt.Print("Max: ", najvisja, ", ")
		fmt.Print("Povprečna: ", povprecna, " => ")

		switch {
		case povprecna < 5.0:
			fmt.Print("Mrzlo")
		case povprecna < 15.0:
			fmt.Print("Hladno")
		case povprecna < 25.0:
			fmt.Print("Toplo")
		default:
			fmt.Print("Vroče")
		}

		fmt.Println()
	}
}
