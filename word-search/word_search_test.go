package word_search

import (
	"testing"
)

func TestWordSearch_Case1(t *testing.T) {
	board := [][]byte{{'a', 'b'}, {'d', 'c'}}
	if !exist(board, "abcd") {
		t.Errorf("failed.")
	}

}

func TestWordSearch_Case2(t *testing.T) {
	board := [][]byte{{'b'}, {'a'}, {'b'}, {'b'}, {'a'}}
	if exist(board, "baa") {
		t.Errorf("failed.")
	}

}