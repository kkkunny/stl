package stlmaps

// Reverse 反转键值对
func Reverse[K, V comparable, KV ~map[K]V, VK ~map[V]K](m KV) VK {
	res := make(map[V]K, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

func Empty[K comparable, V any, KV ~map[K]V](hmap KV) bool {
	return len(hmap) == 0
}
