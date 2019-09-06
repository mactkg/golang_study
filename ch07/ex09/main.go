package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
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

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

var trackTable = template.Must(template.New("tracktable").Parse(`
<table>
	<tr>
	  <th><a href="/?order=title">Title</a></th>
	  <th><a href="/?order=artist">Artist</a></th>
	  <th><a href="/?order=album">Album</a></th>
	  <th><a href="/?order=year">Year</a></th>
	  <th><a href="/?order=length">Length</a></th></tr>
	{{range . }}
	<tr>
		<td>{{.Title}}</td><td>{{.Artist}}</td><td>{{.Album}}</td><td>{{.Year}}</td><td>{{.Length}}</td>
	</tr>
	{{end}}
</table>
`))

func writeTracksToHTMLTable(w http.ResponseWriter, tracks []*Track) {
	if err := trackTable.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

type OrderType int

const (
	Title OrderType = iota
	Year
	Length
	Artist
	Album
)

type customSort struct {
	t       []*Track
	orderBy []OrderType
	less    func(x, y *Track, orderBy []OrderType) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j], x.orderBy) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var orderBy []OrderType
		order, ok := r.URL.Query()["order"]
		if ok {
			orders := strings.Split(order[0], ",")
			for _, o := range orders {
				switch o {
				case "title":
					orderBy = append(orderBy, Title)
				case "artist":
					orderBy = append(orderBy, Artist)
				case "length":
					orderBy = append(orderBy, Length)
				case "year":
					orderBy = append(orderBy, Year)
				case "album":
					orderBy = append(orderBy, Album)
				}
			}
		}

		sort.Sort(customSort{tracks, orderBy, func(x, y *Track, orderBy []OrderType) bool {
			for _, v := range orderBy {
				switch v {
				case Title:
					if x.Title != y.Title {
						return x.Title < y.Title
					}
				case Artist:
					if x.Artist != y.Artist {
						return x.Artist < y.Artist
					}
				case Year:
					if x.Year != y.Year {
						return x.Year < y.Year
					}
				case Length:
					if x.Length != y.Length {
						return x.Length < y.Length
					}
				case Album:
					if x.Album != y.Album {
						return x.Album < y.Album
					}
				default:
					return false
				}
			}
			return false
		}})
		writeTracksToHTMLTable(w, tracks)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
