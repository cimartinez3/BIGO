package main

import (
	"log"
	"strings"
	"time"
)

func main() {
	println(strings.Contains("mar", "ram"))
}

// 1. COMPROBAR SI LOS CARACTERES SON UNICOS
// isUniq("abcde") = true
// isUniq("abcded") = false

func IsUniq(str string) bool {
	// O(n^2) - MALA SOLUCION
	//for i := range str {
	//	for j := 0; j < len(str); j++ {
	//		if j == i {
	//			continue
	//		}
	//		//println(fmt.Sprintf("%c --- %c", str[i], str[j]))
	//		if str[i] == str[j] {
	//			return false
	//		}
	//	}
	//}
	//return true

	// O(n) - MEJOR SOLUCION
	if len(str) > 128 {
		// si es mayor que la cantidad de caracteres posibles fijo se va a repetir
		return false
	}

	myMap := make(map[string]string)
	for i := range str {
		car := string(str[i])
		_, exists := myMap[car]

		if exists {
			return false
		}

		myMap[car] = car
	}
	return true
}

// 2. SUMA DE DOS VALORES DADO UN INPUT REQUERIDO, RETORNA LAS POSICIONES.
// nums = [9,2,5,6]  -- target = 7
// res = [1,2] porque 5 + 2 = 7

func TwoSum(nums []int, trg int) []int {
	// O(n^2) - MALA SOLUCION
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] > trg {
	//		continue
	//	}
	//
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == trg {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//
	//return nil

	// O(N) - MEJOR SOLUCION
	n := make(map[int]int)
	for i, num := range nums {
		_, exist := n[num]

		if exist {
			return []int{i, n[num]}
		}

		n[trg-num] = i
	}

	return nil
}

// 3. Obtener los que sean anagramas
// in - ["saco", "arresto","programa","rastreo","caso"]
// out - [["saco","caso"],["arresto","rastrero"],["programa"]]

func Anagrams(strs []string) [][]string {
	myMap := make(map[int32][]string)
	var res [][]string

	for _, str := range strs {
		cont := getHashASCII(str)
		myMap[cont] = append(myMap[cont], str)
	}

	for _, m := range myMap {
		res = append(res, m)
	}

	return res
}

// 4. Encerar filas y columnas de una matriz
// In:  2 1 3 0 2
//		7 4 1 3 8
//		4 0 1 2 1
//		9 3 4 1 9
//
// Out: 0 0 0 0 0
//		7 0 1 0 8
//		0 0 0 0 0
//		9 0 4 0 9
// todo se cambia los valores de la matriz original.

func ZeroMatrix(matrix [][]int) [][]int {
	var aux = make([][]int, len(matrix))

	copy(aux, matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				continue
			}

			aux[i] = createZeros(len(matrix[i]))

			for k := 0; k < len(matrix); k++ {
				aux[k][j] = 0
			}
		}
	}
	return aux
}

func createZeros(l int) []int {
	var a []int

	for i := 0; i < l; i++ {
		a = append(a, 0)
	}
	return a
}

func getHashASCII(str string) int32 {
	cont := int32(0)
	for _, i := range str {
		cont += i
	}

	return cont
}

func stringBuilders() {
	var sb strings.Builder

	ns := []string{"aaa", "bbb", "ccc"}

	for _, n := range ns {
		sb.WriteString(n)
	}

	println(sb.String())
}

func ExecTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
