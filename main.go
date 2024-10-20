package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Задача 1: Проверка на простоту
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("Число не является простым, делится на %d\n", i)
			return false
		}
	}
	return true
}

// Задача 2: НОД
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Задача 3: Сортировка пузырьком
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println("Промежуточный результат:", arr)
	}
}

// Задача 4: Таблица умножения
func multiplicationTable() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

// Задача 5: Фибоначчи с мемоизацией
var memo = map[int]int{}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

// Задача 6: Обратные числа
func reverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

// Задача 7: Треугольник Паскаля
func generatePascalTriangle(levels int) {
	triangle := make([][]int, levels)
	for i := 0; i < levels; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		fmt.Println(triangle[i])
	}
}

// Задача 8: Число палиндром
func isPalindrome(n int) bool {
	original, reversed := n, 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}

// Задача 9: Нахождение максимума и минимума
func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("Массив не должен быть пустым")
	}
	min, max := arr[0], arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// Задача 10: Игра "Угадай число"
func guessTheNumber() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess, attempts int

	for attempts < 10 {
		fmt.Print("Введите ваше предположение: ")
		fmt.Scan(&guess)
		attempts++

		if guess < target {
			fmt.Println("Загаданное число больше")
		} else if guess > target {
			fmt.Println("Загаданное число меньше")
		} else {
			fmt.Printf("Поздравляем! Вы угадали число %d за %d попыток.\n", target, attempts)
			return
		}
	}
	fmt.Printf("Вы исчерпали все попытки. Загаданное число было: %d\n", target)
}

// Задача 11: Число Армстронга
func isArmstrong(n int) bool {
	original := n
	var sum int
	digits := int(math.Log10(float64(n))) + 1

	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		n /= 10
	}
	return sum == original
}

// Задача 12: Подсчет слов в строке
func countWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

// Задача 13: Игра "Жизнь"
const (
	rows = 5
	cols = 5
)

func printGrid(grid [][]bool) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] {
				fmt.Print("* ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countNeighbors(grid [][]bool, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && ny >= 0 && nx < rows && ny < cols && grid[nx][ny] {
				count++
			}
		}
	}
	return count
}

func nextGeneration(grid [][]bool) [][]bool {
	next := make([][]bool, rows)
	for i := range next {
		next[i] = make([]bool, cols)
		for j := range next[i] {
			aliveNeighbors := countNeighbors(grid, i, j)
			if grid[i][j] {
				next[i][j] = aliveNeighbors == 2 || aliveNeighbors == 3
			} else {
				next[i][j] = aliveNeighbors == 3
			}
		}
	}
	return next
}

// Задача 14: Цифровой корень
func digitalRoot(n int) int {
	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}
	return n
}

// Задача 15: Римские цифры
func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			roman += symbols[i]
		}
	}
	return roman
}

func main() {
	for {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1. Проверка на простоту")
		fmt.Println("2. НОД")
		fmt.Println("3. Сортировка пузырьком")
		fmt.Println("4. Таблица умножения")
		fmt.Println("5. Фибоначчи с мемоизацией")
		fmt.Println("6. Обратные числа")
		fmt.Println("7. Треугольник Паскаля")
		fmt.Println("8. Число палиндром")
		fmt.Println("9. Нахождение максимума и минимума")
		fmt.Println("10. Игра 'Угадай число'")
		fmt.Println("11. Число Армстронга")
		fmt.Println("12. Подсчет слов в строке")
		fmt.Println("13. Игра 'Жизнь'")
		fmt.Println("14. Цифровой корень")
		fmt.Println("15. Римские цифры")
		fmt.Println("0. Выйти")

		var choice int
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			if isPrime(n) {
				fmt.Println("Число является простым")
			} else {
				fmt.Println("Число не является простым")
			}
		case 2:
			var a, b int
			fmt.Print("Введите два числа: ")
			fmt.Scan(&a, &b)
			fmt.Println("НОД:", gcd(a, b))
		case 3:
			var size int
			fmt.Print("Введите размер массива: ")
			fmt.Scan(&size)
			arr := make([]int, size)
			fmt.Println("Введите элементы массива:")
			for i := range arr {
				fmt.Scan(&arr[i])
			}
			bubbleSort(arr)
			fmt.Println("Отсортированный массив:", arr)
		case 4:
			multiplicationTable()
		case 5:
			var n int
			fmt.Print("Введите номер числа Фибоначчи: ")
			fmt.Scan(&n)
			fmt.Println("Число Фибоначчи:", fibonacci(n))
		case 6:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			fmt.Println("Обратное число:", reverseNumber(n))
		case 7:
			var levels int
			fmt.Print("Введите количество уровней: ")
			fmt.Scan(&levels)
			generatePascalTriangle(levels)
		case 8:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			if isPalindrome(n) {
				fmt.Println("Число является палиндромом")
			} else {
				fmt.Println("Число не является палиндромом")
			}
		case 9:
			var size int
			fmt.Print("Введите размер массива: ")
			fmt.Scan(&size)
			arr := make([]int, size)
			fmt.Println("Введите элементы массива:")
			for i := range arr {
				fmt.Scan(&arr[i])
			}
			min, max := findMinMax(arr)
			fmt.Println("Минимум:", min, "Максимум:", max)
		case 10:
			guessTheNumber()
		case 11:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			if isArmstrong(n) {
				fmt.Println("Число является числом Армстронга")
			} else {
				fmt.Println("Число не является числом Армстронга")
			}
		case 12:
			var text string
			fmt.Print("Введите строку: ")
			fmt.Scan(&text)
			wordCount := countWords(text)
			fmt.Println("Количество уникальных слов:", wordCount)
		case 13:
			grid := [][]bool{
				{false, true, false, false, false},
				{false, true, false, false, false},
				{false, true, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
			}
			printGrid(grid)
			for i := 0; i < 5; i++ {
				grid = nextGeneration(grid)
				fmt.Printf("Поколение %d:\n", i+1)
				printGrid(grid)
			}
		case 14:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			fmt.Println("Цифровой корень:", digitalRoot(n))
		case 15:
			var n int
			fmt.Print("Введите арабское число: ")
			fmt.Scan(&n)
			fmt.Println("Римское число:", intToRoman(n))
		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
