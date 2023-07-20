package main

import (
	. "BookStudy/internal/StudyInterface01"
	. "BookStudy/internal/StudyInterface02"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func main() {
	var c ByteCounter
	c.Writer([]byte("hello"))
	fmt.Println(c)

	var w io.Writer
	fmt.Println(w == nil)
	// w.Write([]byte("hello")) 崩溃：对空指针取引用值
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w = new(bytes.Buffer)
	w.Write([]byte("hello"))
	fmt.Printf("%T\n", w)
	w = nil
	fmt.Println(w)
	fmt.Printf("%T\n", w)

	var x interface{} = time.Now()
	fmt.Println(x)

	var Tracks = []*Track{
		{"北京晚报", "In3", "《北京晚报》", 2010, Length("5m08s")},
		{"地下通道", "龙胆紫", "《W.T.F》", 2013, Length("3m33s")},
		{"立交桥", "龙胆紫", "《F.T.W》", 2018, Length("4m21s")},
		{"别伤我心", "黑豹乐队", "《黑豹》", 1997, Length("5m12s")},
	}
	sort.Sort(ByArtist(Tracks))
	PrintTracks(Tracks)
	sort.Reverse(ByArtist(Tracks))
	PrintTracks(Tracks)
}
