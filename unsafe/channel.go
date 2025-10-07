package stlunsafe

func SendChannel[T any](ch chan<- T, v T) (success bool) {
	defer func() {
		if errObj := recover(); errObj != nil {
			success = false
		}
	}()
	ch <- v
	return true
}
