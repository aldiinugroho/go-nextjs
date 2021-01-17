package seeds

import (
	"fmt"
	error "rst/errors"
)

func Seeder() {
	error.Block{
		Try: func() {
			tagseeder()
			contentseeder()
		},
		Catch: func(e error.Exception) {
			fmt.Println("msg error from seeder")
		},
	}.Do()
}
