package ci58

import "strings"


func reverseWords(s string) string {

    if s == "" || s == " " {
        return ""
    }
	str := strings.Split(s, " ")
	ret := []string{}
	for i:=len(str)-1; i>=0; i-- {
        if str[i] == "" {
			continue
		}
        // fmt.Println(str[i])
		ret = append(ret, str[i])
	}
    r := strings.Join(ret, " ")
	return r
}