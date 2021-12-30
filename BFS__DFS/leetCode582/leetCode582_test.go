package leetcode582

import "testing"

var (
	pid  []int
	ppid []int
	kill int
)

func init() {
	for i := 1; i <= 50000; i++ {
		pid = append(pid, i)
	}

	for i := 1; i <= 50000; i++ {
		ppid = append(ppid, 1)
	}
	ppid[0] = 1
}

func BenchmarkBFS(b *testing.B) {
	benchmarkBFS(b, pid, ppid, kill)
}

func BenchmarkDFS(b *testing.B) {
	benchmarkDFS(b, pid, ppid, kill)
}

func benchmarkBFS(b *testing.B, pid []int, ppid []int, kiil int) {
	for n := 0; n < b.N; n++ {
		killProcessBFS(pid, ppid, kiil)
	}
}

func benchmarkDFS(b *testing.B, pid []int, ppid []int, kiil int) {
	for n := 0; n < b.N; n++ {
		killProcessDFS(pid, ppid, kiil)
	}
}
