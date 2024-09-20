package main

import (
	"fmt"
	converter "thermometric-scale/internal"
)

func main() {
	fmt.Println(":= -------------------------------------------------------------------------- =:")
	fmt.Println(":= Conversos de Escalas Termométricas                                               =:")
	fmt.Println(":= -------------------------------------------------------------------------- =:")

	var Degrees float32 = 350.0
	DegreesResult := converter.FromKelvinToCelsius(float32(Degrees))

	fmt.Printf(":= Ao converter %g K teremos %g °C\n\n\n\n\n", Degrees, DegreesResult)
}
