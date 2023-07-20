package time_test

import (
	"BookStudy/cmd/stdlib/time"
	"fmt"
	"testing"
)

func TestTimeDemo(t *testing.T) {
	time.TimeDemo()
	fmt.Println("[---------------------------]")
	time.TimeStampDemo()
	fmt.Println("[---------------------------]")
	time.TimeDruationDemo()
	fmt.Println("[---------------------------]")
	// time.TickDemo()
	time.FormatDemo()
}
