package either

func (self *Either[L, R]) init() {
	if self.left != nil {
		return
	}
	left := true
	self.left = &left
	var l L
	self.data = l
}

func (self Either[L, R]) Left() (res L, ok bool) {
	self.init()

	if !*self.left {
		return
	}
	return self.data.(L), true
}

func (self Either[L, R]) Right() (res R, ok bool) {
	self.init()

	if *self.left {
		return
	}
	return self.data.(R), true
}

func (self Either[L, R]) IsLeft() bool {
	self.init()
	return *self.left
}

func (self Either[L, R]) IsRight() bool {
	self.init()
	return !*self.left
}

func (self *Either[L, R]) SetLeft(v L) {
	left := true
	self.left = &left
	self.data = v
}

func (self *Either[L, R]) SetRight(v R) {
	left := false
	self.left = &left
	self.data = v
}
