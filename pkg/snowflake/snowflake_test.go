package snowflake

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	// 129945560944640
	// 129989966041088
	// 130080097439744
	// 130140864516096
	// 134216767508480
	// 157056447287296
	// 382897567567872
	Init("2021-07-01", 1)
	fmt.Println(GetID())
}
