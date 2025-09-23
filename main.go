package main

import (
	"fmt"
	"sandbox/applications"
	"sandbox/repos"
)

var flightsTC = []repos.Flight{
	{Id: 1, Capacity: 45, Weight: 12},
	{Id: 2, Capacity: 45, Weight: 12},
}

var usersTC = []repos.User{
	{Id: 3, Name: "Bob", Age: 12},
	{Id: 4, Name: "John", Age: 12},
}

func main() {
	flights := repos.NewFlights(flightsTC)
	users := repos.NewUsers(usersTC)

	handler := applications.NewHandler(users, flights)

	ids := handler.GetAllIds()
	fmt.Println(ids)
}
