package iron

import (
	"image"
	"sync"
)

var loadIconOnce sync.Once

var icons map[string]image.Image

func loadIron(name string) image.Image
func loadIrons() {
	icons = map[string]image.Image{
		"xxx.png": loadIron("xxx.png"),
		"yyy.jpg": loadIron("yyy.png"),
		"zzz.jpg": loadIron("zzz.jpg"),
	}

}

// 并发安全
func Icon(name string) image.Image {
	loadIconOnce.Do(loadIrons)
	return icons[name]
}
