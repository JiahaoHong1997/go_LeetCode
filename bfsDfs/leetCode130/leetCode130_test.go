package leetcode130

import (
	"testing"
	"reflect"
)

func TestSolve(t *testing.T) {
	board1 := [][]byte{
		{'X','X','X','X'},
		{'X','O','O','X'},
		{'X','X','O','X'},
		{'X','O','X','X'},
	}

	board2 := [][]byte{
		{'O','X','X','O','X'},
		{'X','O','O','X','O'},
		{'X','O','X','O','X'},
		{'O','X','O','O','O'},
		{'X','X','O','X','O'},
	}

	solve(board1)
	solve(board2)
	want1 := [][]byte{
		{'X','X','X','X'},
		{'X','X','X','X'},
		{'X','X','X','X'},
		{'X','O','X','X'},
	}

	want2 := [][]byte{
		{'O','X','X','O','X'},
		{'X','X','X','X','O'},
		{'X','X','X','O','X'},
		{'O','X','O','O','O'},
		{'X','X','O','X','O'},
	}

	if !reflect.DeepEqual(board1,want1) {
		t.Errorf("expected:%v, got:%v", want1, board1)
	}

	if !reflect.DeepEqual(board2,want2) {
		t.Errorf("expected:%v, got:%v", want2, board2)
	}
}