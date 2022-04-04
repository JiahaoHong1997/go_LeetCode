package weekCamp

type Encrypter struct {
	mp  [26]string
	cnt map[string]int
}

func Constructor(keys []byte, values, dictionary []string) Encrypter {
	mp := [26]string{}
	for i, key := range keys {
		mp[key-'a'] = values[i]
	}
	e := Encrypter{mp, map[string]int{}}
	for _, s := range dictionary {
		e.cnt[e.Encrypt(s)]++
	}
	return e
}

func (e *Encrypter) Encrypt(word1 string) string {
	res := make([]byte, 0, len(word1)*2)
	for _, ch := range word1 {
		s := e.mp[ch-'a']
		if s == "" { return "" }
		res = append(res, s...)
	}
	return string(res)
}

func (e *Encrypter) Decrypt(word2 string) int { return e.cnt[word2] }


/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */
