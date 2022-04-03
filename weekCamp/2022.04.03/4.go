package weekCamp

type Encrypter struct {
	keys       map[byte]int
	k          []byte
	values     map[string][]int
	v          []string
	dictionary map[string]bool
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {

	m := make(map[byte]int)
	for i := 0; i < len(keys); i++ {
		m[keys[i]] = i
	}

	t := make(map[string][]int)
	for i:=0; i<len(values); i++ {
		t[values[i]] = append(t[values[i]], i)
	}

	d := make(map[string]bool)
	for i:=0; i<len(dictionary); i++ {
		d[dictionary[i]] = true
	}
	return Encrypter{
		keys:       m,
		k: keys,
		values: t,
		v:     values,
		dictionary: d,
	}
}

func (this *Encrypter) Encrypt(word1 string) string {

	b := ""
	for i := 0; i < len(word1); i++ {
		c := word1[i]
		index := this.keys[c]
		b += this.v[index]
	}

	return b
}

func (this *Encrypter) Decrypt(word2 string) int {

	count := 0
	sub := ""
	dec := []string{}
	for i:=0; i<len(word2); i++ {
		count++
		sub += word2[i:i+1]
		if count == 2 {
			count = 0
			sub = ""
		}

		list := this.values[sub]
		times := len(list)
		tList := make([]string, len(dec))
		copy(tList, dec)
		for i:=0; i<times-1; i++ {
			dec = append(dec, tList...)
		}

		for j:=0; j<len(list); j++ {
			index := list[j]
			c := this.k[index]
			for k:=0; k<len(dec)/times; k++ {
				ff := j*len(dec)/times + k
				dec[ff] += string(c)
			}
		}
	}

	ret := 0
	for i:=0; i<len(dec); i++ {
		s := dec[i]
		if this.dictionary[s] {
			ret++
		}
	}

	return ret
}

/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */
