package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsUniq(t *testing.T) {
	assert.True(t, IsUniq("abcde"))
	assert.True(t, IsUniq("aAbBcCdDeE"))
	assert.False(t, IsUniq("abcded"))
	assert.False(t, IsUniq("abcdefghijklmnopqrstuvwxyzz"))

	defer ExecTime(time.Now(), "main")
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{2, 1}, TwoSum([]int{9, 2, 5, 6}, 7))
	assert.Equal(t, []int{3, 2}, TwoSum([]int{6, 9, 5, 2}, 7))
	assert.Nil(t, TwoSum([]int{9, 5, 2, 6}, 100))

	defer ExecTime(time.Now(), "main")
}

func TestAnagrams(t *testing.T) {
	assert.Equal(t, [][]string{{"saco", "caso", "cosa"}, {"arresto", "rastreo"}, {"programa"}}, Anagrams([]string{"saco", "arresto", "programa", "rastreo", "caso", "cosa"}))
}

func TestZeroMatrix(t *testing.T) {
	_ = [][]int{
		{2, 1, 3, 0, 2},
		{7, 4, 1, 3, 8},
		{4, 0, 1, 2, 1},
		{9, 3, 4, 1, 9},
	}
}
