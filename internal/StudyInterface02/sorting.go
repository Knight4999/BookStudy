package studyinterface02

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string        // 名称
	Artist string        // 歌手
	Album  string        // 专辑
	Year   int           // 年代
	Length time.Duration //时长
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "-----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type ByArtist []*Track

func (a ByArtist) Len() int           { return len(a) }
func (a ByArtist) Less(i, j int) bool { return a[i].Artist < a[j].Artist }
func (a ByArtist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
