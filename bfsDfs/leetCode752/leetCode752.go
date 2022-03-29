package leetcode752

func openLock(deadends []string, target string) int {
    if target == "0000" {
        return 0
    }
	m := make(map[string]bool,0)
	for i:=0; i<len(deadends); i++ {
        if deadends[i] == "0000" {
            return -1
        }
		m[deadends[i]] = true
	}

	q := make([]string,1)
	q[0] = "0000"
	step := 1

	for len(q) > 0 {
		p := make([]string,0)
		for i:=0; i<len(q); i++ {
			s := q[i]
            
			for j:=0; j<4; j++ {
				x := s[j]
				b := []byte(s)
				if int(x) -'0' >= 1 && int(x) - '0' <= 8 {
					b[j] = x + 1
                    
					if _, ok := m[string(b)]; !ok {
						if string(b) == target {
							return step
						}
						p = append(p, string(b))
                        
						m[string(b)] = true
					}
					b[j] = x - 1
					if _, ok := m[string(b)]; !ok {
						if string(b) == target {
							return step
						}
						p = append(p, string(b))
						m[string(b)] = true
					}
				} else if int(x) - '0' == 0 {
					b[j] = '1'
					if _, ok := m[string(b)]; !ok {
						if string(b) == target {
							return step
						}
						p = append(p, string(b))
						m[string(b)] = true
					}
					b[j] = '9'
					if _, ok := m[string(b)]; !ok {
						if string(b) == target {
							return step
						}
						p = append(p, string(b))
						m[string(b)] = true
					}
				} else {
					b[j] = '0'
					if _, ok := m[string(b)]; !ok {
						if string(b) == target {
							return step
						}
						p = append(p, string(b))
						m[string(b)] = true
					}
					b[j] = '8'
					if _, ok := m[string(b)]; !ok {
						if string(b) == target {
							return step
						}
						p = append(p, string(b))
						m[string(b)] = true
					}
				}
			}
		}
		step++
		q = p
	}
	return -1
}
