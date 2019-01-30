package main

type Movie struct {
		Title	string
		Year	int `json:"released"`
		Color 	bool `json:"color, omitempty"`
		Actors	[]string
}

var movies = []Movie{
		{Title: "The Big Lebowski", Year: 1998, Color: true, Actors: []string{"Jeff Bridges"}}
		// ...
}
