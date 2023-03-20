package _2325_decodeMessage

func decodeMessage(key string, message string) string {
	mp := make(map[byte]byte)
	mp[' '] = ' '
	for i, j := 'a', 0; i <= 'z'; j++ {
		if _, exists := mp[key[j]]; !exists {
			mp[key[j]] = byte(i)
			i++
		}
	}

	cs := []byte(message)
	for i := 0; i < len(message); i++ {
		cs[i] = mp[cs[i]]
	}
	return string(cs)
}
