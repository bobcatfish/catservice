package cat

import (
	"html/template"
)

// TektonCat holds information about a cat that has been associated with a Tekton release
type TektonCat struct {
	// Name is the name of this cat
	Name string

	// DisplayName is the fancy display name of this cat
	DisplayName template.HTML

	// Image is the path to an image on disk that represents this cat
	Image string

	// Release is the release of Tekton that this cat was associated with
	Release string
}

func GetCatsOfTekton() []TektonCat {
	rainbowCats := []TektonCat{}
	for _, cat := range catsOfTekton {
		rainbowCats = append(rainbowCats, TektonCat{
			DisplayName: template.HTML(cat.Name),
			Name:        cat.Name,
			Image:       cat.Image,
			Release:     cat.Release,
		})
	}
	return rainbowCats
}

// catsOfTekton are some wonderful cats
var catsOfTekton = []TektonCat{{
	Name:    "Acadia",
	Image:   "acadia_0.3.1.jpeg",
	Release: "v0.4.0",
}, {
	Name:    "Annabelle",
	Image:   "annabelle_0.6.0.jpeg",
	Release: "v0.6.0",
}, {
	Name:    "Chip",
	Image:   "chip_0.5.0.jpeg",
	Release: "v0.5.0",
}, {
	Name:    "Coco",
	Image:   "coco_0.20.0.jpeg",
	Release: "v0.20.0",
}, {
	Name:    "Dip",
	Image:   "dip_0.7.0.jpeg",
	Release: "v0.7.0",
}, {
	Name:    "Meatball",
	Image:   "meatball_0.16.3.jpeg",
	Release: "v0.16.3",
}, {
	Name:    "Yoshimi",
	Image:   "yoshimi_0.4.0.jpeg",
	Release: "v0.4.0",
}}
