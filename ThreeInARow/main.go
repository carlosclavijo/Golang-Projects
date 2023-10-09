package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Size: ", size)
	arr := make([][]string, size)
	for i := range arr {
		arr[i] = make([]string, size)
	}
	fmt.Printf("Len[0]: %d\n", len(arr))
	fmt.Println(mapGame(arr))
	b, q := verification(arr)
	turno := true
	for !b {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		if turno {
			arr[x-1][y-1] = "X"
			turno = false
		} else {
			arr[x-1][y-1] = "O"
			turno = true
		}
		fmt.Println(mapGame(arr))
		b, q = verification(arr)
	}
	fmt.Println("Gano el jugador", q)
}

func mapGame(a [][]string) string {
	var m string
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			m += " "
			if a[i][j] == "" {
				m += " "
			} else {
				m += a[i][j]
			}
			m += " "
			if j < len(a)-1 {
				m += "|"
			} else if j == len(a)-1 {
				m += "\n"
				if i != len(a)-1 {
					for k := 0; k < len(a); k++ {
						m += "---"
					}
					for k := 0; k < len(a)-1; k++ {
						m += "-"
					}
					m += "\n"
				}
			}
		}
	}
	return m
}

func verification(a [][]string) (bool, int) {
	var (
		v      bool
		winner int
	)
	for i := 0; i < len(a); i++ {
		var aux string
		if a[i][0] != "" {
			aux = a[i][0]
			cont := 1
			for j := 1; j < len(a); j++ {
				if a[i][j] == aux {
					cont++
				}
			}
			v, winner = countPlays(cont, aux, len(a))
			if v {
				return v, winner
			}
		}
	}
	for i := 0; i < len(a); i++ {
		var aux string
		if a[0][i] != "" {
			aux = a[0][i]
			cont := 1
			for j := 1; j < len(a); j++ {
				if a[j][i] == aux {
					cont++
				}
			}
			v, winner = countPlays(cont, aux, len(a))
			if v {
				return v, winner
			}
		}
	}
	if a[0][0] != "" {
		aux := a[0][0]
		cont := 1
		for i := 1; i < len(a); i++ {
			if a[i][i] == aux {
				cont++
			}
		}
		v, winner = countPlays(cont, aux, len(a))
		if v {
			return v, winner
		}
	}
	if a[len(a)-1][0] != "" {
		aux := a[len(a)-1][0]
		cont := 1
		for i := 1; i < len(a); i++ {
			if a[len(a)-1-i][i] == aux {
				cont++
			}
		}
		v, winner = countPlays(cont, aux, len(a))
		if v {
			return v, winner
		}
	}
	return v, winner
}

func countPlays(count int, str string, length int) (bool, int) {
	if count == length {
		if str == "X" {
			return true, 1
		} else if str == "O" {
			return true, 2
		}
	}
	return false, 0
}
