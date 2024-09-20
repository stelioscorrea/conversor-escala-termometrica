package converter

import "fmt"

/*
 * Para transformar Celsius para Kelvin
 * Tk = Tc + 273
 */
func FromCelsiusToKelvin(degrees float32) float32 {
	fmt.Println(":= Convertendo da Escala de Celsius(°C) para Fahrenheit(F)                    =:")
	fmt.Println(":= -------------------------------------------------------------------------- =:")
	Celsius := degrees + 273
	return Celsius
}

/*
 * Para transformar Kelvin para Celsius
 * Tc = Tk - 273
 */
func FromKelvinToCelsius(degrees float32) float32 {
	fmt.Println(":= Convertendo da Escala de Kelvin(K) para Celsius(°C)                        =:")
	fmt.Println(":= -------------------------------------------------------------------------- =:")
	Kelvin := degrees - 273
	return Kelvin
}

/*
 * Para transformar Fahrenheit para Celsius
 * Tc / 5 = (Tf - 32)/9
 * 9Tc = 5(Tf - 32)
 * Tc = (5(Tf - 32))/9
 */
func FromFahrenheitToCelsius(degrees float32) float32 {
	fmt.Println(":= Convertendo da Escala de Fahrenheit(F) para Celsius(°C)                    =:")
	fmt.Println(":= -------------------------------------------------------------------------- =:")
	Celsius := (5 * (degrees - 32)) / 9
	return Celsius
}

/*
 * Para transformar Celsius para Fahrenheit
 * Tc / 5 = (Tf - 32)/9
 * 9Tc = 5(Tf - 32)
 * Tf - 32 = 9Tc /  5
 * Tf = (9Tc / 5) + 32
 */
func FromCelsiusToFahrenheit(degrees float32) float32 {
	fmt.Println(":= Convertendo da Escala de Celsius(°C) para Fahrenheit(F)                    =:")
	fmt.Println(":= -------------------------------------------------------------------------- =:")
	Fahrenheit := ((9 * degrees) / 5) + 32
	return Fahrenheit
}
