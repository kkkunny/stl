package either

import "fmt"

func (self Either[L, R]) String() string {
	self.init()

	if self.IsLeft() {
		return fmt.Sprintf("%v", self.data.(L))
	} else {
		return fmt.Sprintf("%v", self.data.(R))
	}
}
