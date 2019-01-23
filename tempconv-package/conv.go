package tempconv

// converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//converts Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
