package com.level.oj.wordsearch;

class Solution {

    private char[][] board;

    private boolean[][] marked;

    private int rowNum;

    private int columnNum;

    public boolean exist(char[][] board, String word) {
        this.board = board;
        rowNum = board.length;
        columnNum = board[0].length;
        this.marked = new boolean[rowNum][columnNum];

        char[] wordArray = word.toCharArray();

        for (int i = 0; i < rowNum; i++) {
            for (int j = 0; j < columnNum; j++) {
                if (found(wordArray, 0, i, j)) {
                    return true;
                }
            }
        }

        return false;
    }

    private boolean found(char[] word, int offset, int startRow, int startColumn) {
        if (offset == word.length) {
            return true;
        }
        char pre = word[offset];
        if (startRow >= 0 && startRow < rowNum && startColumn >= 0 && startColumn < columnNum
            && marked[startRow][startColumn] == false && pre == board[startRow][startColumn]) {
            marked[startRow][startColumn] = true;
            offset += 1;
            if (found(word, offset, startRow + 1, startColumn) || found(word, offset, startRow - 1, startColumn)
                || found(word, offset, startRow, startColumn + 1) || found(word, offset, startRow, startColumn - 1)) {
                marked[startRow][startColumn] = false;
                return true;
            } else {
                marked[startRow][startColumn] = false;
                return false;
            }
        } else {
            return false;
        }
    }

    public static void main(String[] args) {
        char[][] board = new char[][] {
            {'##'}
        }
    }
}
