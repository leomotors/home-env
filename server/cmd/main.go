package main

import "github.com/leomotors/home-env/services"

func main() {
	services.RegisterSensor("main_room", "Office Room")
	services.RegisterSensor("living_room", "Living Room")
}
