package apptemp

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // conversion
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9) // conversion
}
