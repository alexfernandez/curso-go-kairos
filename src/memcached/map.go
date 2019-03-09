package memcached

var mem map[string]string = make(map[string]string)

func Get(key string) string {
	return mem[key]
}

func Set(key string, value string) {
	mem[key] = value
}

func Del(key string) {
	delete(mem, key)
}
