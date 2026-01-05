package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	strFLoats := strings.Fields(input)
	var floatValues []float64
	for _, s := range strFLoats {
		f, _ := strconv.ParseFloat(s, 64)
		floatValues = append(floatValues, f)
	}

	d := 2 * floatValues[0]

	if d == 0 {
		fmt.Println("Impossivel calcular")
		return
	}

	delta := math.Pow(floatValues[1], 2) - (4 * floatValues[0] * floatValues[2])
	if delta < 0 {
		fmt.Println("Impossivel calcular")
		return
	}
	x1 := ((-floatValues[1]) + math.Sqrt(delta)) / (2 * floatValues[0])
	x2 := ((-floatValues[1]) - math.Sqrt(delta)) / (2 * floatValues[0])
	fmt.Printf("R1 = %.5f\n", x1)
	fmt.Printf("R2 = %.5f\n", x2)
}
