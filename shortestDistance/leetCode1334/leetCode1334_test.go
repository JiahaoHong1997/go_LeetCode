package leetcode1334

import "testing"

var (
	n1 	[]int
	edges 	[][][]int
	distanceThreshold []int
)

func init() {
	n1 = []int{4,5}
	edges = [][][]int{
		{{0,1,3},{1,2,1},{1,3,4},{2,3,1}},
		{{0,1,2},{0,4,8},{1,2,3},{1,4,2},{2,3,1},{3,4,1}},
	}
	distanceThreshold = []int{4,2}
}

func BenchmarkDj(b *testing.B) {
	benchmarkDj(b)
}

func BenchmarkFl(b *testing.B) {
	benchmarkFl(b)
}

func benchmarkDj(b *testing.B) {
	for n:=0; n<b.N; n++ {
		for i:=0; i<len(n1); i++ {
			findTheCityDijkstra(n1[i],edges[i],distanceThreshold[i])
		}
	}
}

func benchmarkFl(b *testing.B) {
	for n:=0; n<b.N; n++ {
		for i:=0; i<len(n1); i++ {
			findTheCityFloyd(n1[i],edges[i],distanceThreshold[i])
		}
	}
}