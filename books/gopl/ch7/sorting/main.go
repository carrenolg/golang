package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track

// byArtist implements sort.Interface
func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byYear []*Track

// byYear implements sort.Interface
func (x byYear) Len() int {
	return len(x)
}

func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	fmt.Fprint(tw, "\n")
	tw.Flush() // calculate column widths and print table
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// customSort implements sort.Interface
func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func main() {
	// sorting
	printTracks(tracks) // tracks: type []*Track

	// sorting to tracks by the Artist field
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	// sorting to tracks by the Year field
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	// custom sortig
	// "f" as value for field less (struct: customSort)
	f := func(x, y *Track) bool {
		/*if x.Year != y.Year {
			return x.Year < y.Year
		}*/
		// by Title
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		// by Album
		if x.Album != y.Album {
			return x.Album < y.Album
		}
		return false
		//return x.Title < y.Title
	}
	custom := customSort{
		t:    tracks,
		less: f,
	}

	// sorting to tracks by Title,Album using customing function "f"
	sort.Sort(custom)
	printTracks(tracks)
}
