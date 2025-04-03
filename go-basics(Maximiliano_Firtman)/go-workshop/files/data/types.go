package data

import "fmt"



type location string


func (origin location) DistanceTo(destination location) float64{
	//TODO caluclation ...
	fmt.Printf("Origin: %v \nDestination: %v\n", origin, destination)
	return 10
}

func LocationTest(){
	nyc := location("33.456334, 34.234")
	

	nyc.DistanceTo(location("-23,-44"))

	print(nyc)
}

//method
type distance float64 //miles
type distanceKm float64


func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}
func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}


func Test(){
	d := distance(34.5)
	km := d.ToKm()
	miles := km.ToMiles()
	print(d)
	print(miles)
}