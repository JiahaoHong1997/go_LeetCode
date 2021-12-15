package leetcode340

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	s := []string{
		"eceba",
		"aa",
		"dsanjivbsibsnofdeihfbihsbfewikbfuywvfiusbfvciuewvfuywbfdoqwifbewyuidsadsaiudbweiuyfbewiognoiuwfghdaiuwdbfweifnowehbfwiuygfyufvwuiyb",
		"nbdicsbfuvybdwiuehnfwoiufhiufnoireghbiyufguyfbdiufhbaisuhhhjhdhfvhjfhheqwdneowfhgqwyiufrgerubqwodbhawiuybasidbewyufgve wjdbqwu",
		"dsabjkdbiewbquygdwqiudbwiuqdbhwqoidheryufguafdnewonfregfbejdnalkscnvkjbuyfvwuuehbdfehfbvewufvqhdbqndqwuidgqwiukjbcfdshjcvajkdbqwdbwqiufdgewuyfvwuj",
		"dsanodbewyuigdudvqwibdnwolqnfoiwehfbiubdiuqwbdqwkjdbnwqiujdbqwyuifdtydgcqofnpewofhweofghqwdbpoqwicbhduikfbvureyhfvbewuydfvwqudbwqjhdvqwhgjdjqhw",
	}

	k := []int{2,1,4,6,3,10}
	want := []int{3,2,8,15,4,30}
	got := []int{}

	for i:=0; i<len(s); i++ {
		got =append(got, lengthOfLongestSubstringKDistinct(s[i],k[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want,got)
	}
}