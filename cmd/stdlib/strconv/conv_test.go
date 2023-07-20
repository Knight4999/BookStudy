package strconv_test

import (
	"BookStudy/cmd/stdlib/strconv"
	"fmt"
	"testing"
)

func TestConv(t *testing.T) {
	/* scan := bufio.NewScanner(os.Stdin)
	s := scan.Text()
	switch s {
	case "1":
		strconv.StoI()
	case "2":
		strconv.ItoS()
	case "3":
		strconv.ParseTest()
	} */

	strconv.ParseTest()
	fmt.Println("-----------------------------------")
	strconv.FormatTest()
	fmt.Println("-----------------------------------")
	strconv.IsPrint()
	fmt.Println("-----------------------------------")
	strconv.AppendTest()
}
