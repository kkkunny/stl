package either

import "fmt"

func (e Either[L, R]) String() string {
	if e.IsLeft() {
		return fmt.Sprintf("%v", e.l)
	} else {
		return fmt.Sprintf("%v", e.r)
	}
}
