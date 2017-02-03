package c4

import (
	"strings"
)

// Id of nil (i.e. '')
var (
	NIL_ID *ID
)

func init() {
	NIL_ID, _ = Identify(strings.NewReader(``))
}
