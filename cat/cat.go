package cat

// Cat holds information about a cat, such as my adorable little monster Yoshimi!
type Cat struct {
	// Name is the name of this cat.
	Name string

	// Mood is the main mood this cat displays.
	Mood string

	// Image is the path to an image on disk that represents this cat.
	Image string

	// Age is the number of cat years old this cat is.
	Age int

	// Pronoun is the personal pronoun to use for this cat.
	Pronoun string
}

// AgeInCatYears will look at how old c is in human years and convert that to the
// cat years equivalent.
func AgeInCatYears(hoomanYears int) int {
	// One cat year is the equivalent of about 7 cat years, right?
	// But I've heard it's a bit more complicated...
	return hoomanYears * 7
}

// Yoshimi is my ancient cat.
func Yoshimi() *Cat {
	return &Cat{
		Name:    "Yoshimi",
		Mood:    "grouchy",
		Image:   "yoshimi2.jpg",
		Age:     AgeInCatYears(17),
		Pronoun: "She",
	}
}
