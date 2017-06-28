package main;

import (
	"flag"
	"fmt"
	"math/rand"
)

const Gridsize = 99; //100x100 grid

type House struct {
	x, y int
}

type Salesman struct {
	x, y int
}

var houses []*House
var housenum = flag.Int("h", 10, "number of houses")
var salesmennum = flag.Int("s", 1, "number of salesmen")
var temp = flag.Float64("t", 10, "initial temperature")
var coolingRate = flag.Float64("c", 0.01, "cooling rate")
 
func NewHouse() *House {
	h := new(House)
	h.x = rand.Intn(Gridsize)
	h.y = rand.Intn(Gridsize)

	return h
}

func setupHouses() {
	houses = make([]*House, *housenum)
	for i := 0; i < housenum; i++ {
		h := NewHouse()
		houses[i] =  h
	}
}

func (h *House) distanceto(h1 *House) float64 {
	xdist := math.Abs(float64(h.x - h1.x))
	ydist := math.Abs(float64(h.y - h1.y))
	dist := math.Sqrt((xdist*xdist) + (ydist*ydist))

	return dist
}

func NewSalesman() *Salesman {
	s := new(Salesman)
	r := rand.Intn(len(houses))
	h := houses[r]
	s.x = h.x
	s.y = h.y

	return s
}

func acceptanceProbability(newdist, currentdist, temp float64) {
	if newdist < currentdist {
		return 1.0
        } else {
		return math.exp((currentdist-newdist)/temp)
	}
}

func SA() {
	for i := 0; i <= temp; i++ {
		
	}
}

func main() {
	flag.Parse()
	setupHouses()
	
}
