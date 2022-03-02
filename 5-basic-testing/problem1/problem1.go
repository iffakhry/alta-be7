package problem1

var (
	// nama = "jerry"
	Nama = "Abbi"
)

func Add(a, b int) int {
	return a + b
}

// func multiply(a, b int) int {
// 	return a * b
// }

func Substract(a, b int) int {
	return a - b
}

func GanjilGenap(x int) string {
	if x%2 == 0 {
		return "Genap"
	}

	return "Ganjil"
}

// func main() {
// 	fmt.Println("Hello World")

// 	fmt.Println("Hasil 1+2 = ", Add(1, 2))
// 	fmt.Println("Hasil 3-1 = ", Substract(3, 1))
// 	fmt.Println()
// }
