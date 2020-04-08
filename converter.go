package main

import (
	"fmt"
)

//Converter ...
type Converter struct{}

//Centimeter ..
type Centimeter float64

//Feet ...
type Feet float64

//Minutes ...
type Minutes float64

//Seconds ...
type Seconds float64

//Milliseconds ...
type Milliseconds float64

//Celcius ...
type Celcius float64

//Fahrenheit ...
type Fahrenheit float64

//Radians ...
type Radians float64

//Degree ...
type Degree float64

//Kilogram ...
type Kilogram float64

//Pounds ...
type Pounds float64

//CentToFeet Coverter
func (cvr Converter) CentToFeet(c Centimeter) Feet {
	var res = Feet(c / 30.48)
	return res
}

//FeetToCent Converter
func (cvr Converter) FeetToCent(f Feet) Centimeter {
	var res = Centimeter(f * 30.48)
	return res
}

//MinutesToSecs Converter
func (cvr Converter) MinutesToSecs(m Minutes) Seconds {
	var res = Seconds(m * 60)
	return res
}

//SecsToMinutes Converter
func (cvr Converter) SecsToMinutes(s Seconds) Minutes {
	var res = Minutes(s / 60)
	return res
}

//SecstoMillisecs Converter
func (cvr Converter) SecstoMillisecs(s Seconds) Milliseconds {
	var res = Milliseconds(s * 1000)
	return res
}

//CelciustoFahrenheit Converter
func (cvr Converter) CelciustoFahrenheit(c Celcius) Fahrenheit {
	var res = Fahrenheit((c * 1.8) + 32)
	return res
}

//FahrenheittoCelcius Converter
func (cvr Converter) FahrenheittoCelcius(f Fahrenheit) Celcius {
	var res = Celcius((f - 32) * 1.8)
	return res
}

//RadianToDegree Converter
func (cvr Converter) RadianToDegree(r Radians) Degree {
	var res = Degree(r * (3.142 / 180))
	return res
}

// DegreeToradian Converter
func (cvr Converter) DegreeToradian(d Degree) Radians {
	var res = Radians(d * (180 / 3.142))
	return res
}

//KilogramsToPounds Converter
func (cvr Converter) KilogramsToPounds(k Kilogram) Pounds {
	var res = Pounds(k * 2.20462)
	return res
}

//PoundsToKilograms Converter
func (cvr Converter) PoundsToKilograms(p Pounds) Kilogram {
	var res = Kilogram(p / 2.20462)
	return res
}

func main() {
	// to access the methods
	// we need an instance to the Converter sruct
	cvr := Converter{}
	// calling our method and printing the result at the same time
	fmt.Println(cvr.CentToFeet(304.8))
	fmt.Println(cvr.FeetToCent(10))
	fmt.Println(cvr.MinutesToSecs(10))
	fmt.Println(cvr.SecsToMinutes(10))
	fmt.Println(cvr.SecstoMillisecs(20))
	fmt.Println(cvr.CelciustoFahrenheit(5))
	fmt.Println(cvr.FahrenheittoCelcius(33))
	fmt.Println(cvr.RadianToDegree(33))
	fmt.Println(cvr.DegreeToradian(23))
	fmt.Println(cvr.KilogramsToPounds(33))
	fmt.Println(cvr.PoundsToKilograms(44))
}
