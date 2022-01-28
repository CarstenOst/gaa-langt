package hello

import (
	"rsc.io/quote"
)

func pithySay(arrIndex int) string {

	pithySayings := [4]string{quote.Glass(), quote.Go(), quote.Hello(), quote.Opt()}

	return pithySayings[arrIndex]

}
