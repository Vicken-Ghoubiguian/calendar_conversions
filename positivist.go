package calendar_conversions

import (
	"fmt"
	"time"
)

type positivistMonth int

 //
 const (
 	Moses positivistMonth = iota
 	Homer
 	Aristotle
 	Archimedes
 	Caesar
 	Saint_Paul
 	Charlemagne
 	Dante
 	Gutenberg
 	Shakespeare
 	Descartes
 	Frederic
 	Bichat
 )

//
var days_from_months = map[positivistMonth][]string{
	Moses: {
		"Prometheus",
		"Hercules",
		"Orpheus",
		"Ulysses",
		"Lycurgus",
		"Romulus",
		"Numa",
		"Belus",
		"Sesostris",
		"Menu",
		"Cyrus",
		"Zoroaster",
		"The Druids",
		"Buddha",
		"Fo-Hi",
		"Lao-Tzu",
		"Meng-Tzu",
		"The Priests of Tibet",
		"The Priests of Japan",
		"Manco Capac",
		"Confucius",
		"Abraham",
		"Joseph",
		"Samuel",
		"Solomon",
		"Isaac",
		"St. John the Baptist",
		"Muhammad",
	},
	Homer: {
		"a",
		"b",
		"c",
	},
	Aristotle: {
		"a",
		"b",
		"c",
	},
	Archimedes: {
		"a",
		"b",
		"c",
	},
	Caesar: {
		"a",
		"b",
		"c",
	},
	Saint_Paul: {
		"a",
		"b",
		"c",
	},
	Charlemagne: {
		"a",
		"b",
		"c",
	},
	Dante: {
		"a",
		"b",
		"c",
	},
	Gutenberg: {
		"a",
		"b",
		"c",
	},
	Shakespeare: {
		"a",
		"b",
		"c",
	},
	Descartes: {
		"a",
		"b",
		"c",
	},
	Frederic: {
		"a",
		"b",
		"c",
	},
	Bichat: {
		"Copernicus",
		"Kepler",
		"Huygens",
		"Jacques Bernoulli",
		"Bradley",
		"Volta",
		"Galileo",
		"Viète",
		"Wallis",
		"Clairaut",
		"Euler",
		"D'Alembert",
		"Lagrange",
		"Newton",
		"Bergmann",
		"Priestley",
		"Cavendish",
		"Guyton Morveau",
		"Berthollet",
		"Berzelius",
		"Lavoisier",
		"Harvey",
		"Boerhaave",
		"Linnaeus",
		"Haller",
		"Lamarck",
		"Broussais",
		"Gall",
	},
}

//
type Positivist_date struct {
}

//
func (current_positivist_date *Positivist_date) Initialize_positivist_date(dt time.Time) {
}

func main() {

	fmt.Println("hello world")
}
