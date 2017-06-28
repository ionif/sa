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
	houses []*House var
}

var housenum = flag.Int("h", 10, "number of houses")
var salesmennum = flag.Int("s", 1, "number of salesmen")
var temp = flag.Float64("t", 10, "initial temperature")
var coolingRate = flag.Float64("c", 0.01, "cooling rate")
var initd = 0
var currentdist = 0

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

func (s *Salesman) distanceto(h *House) float64 {
	xdist := math.Abs(float64(s.x - h.x))
	ydist := math.Abs(float64(s.y - h.y))
	dist := math.Sqrt((xdist*xdist) + (ydist*ydist))

	return dist
}

func NewSalesman() *Salesman {
	s := new(Salesman)
	r := rand.Intn(len(houses))
	h := houses[r]
	s.h = append(houses[:r])
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

func Route() *Salesman {
	for i:=1; i < len(s.h) {
		r := rand.Intn(len(s.h))
		dist := s.distanceto(s.h[r])
		initd += dist
		s.x = s.h[r].x
		s.y = s.h[r].y
		s.h = append(houses[:r])
	}
	return initd
}

func SA() {
	currentdist = Route()
	for i := 1; i <= temp; temp*(1-coolingRate) {
		d := Route()
		if acceptanceProbability(d, currentdist, temp) > rand.Float64 {
			currentdist = d
		}
	}
	fmt.Printf(currentdist)
}

func main() {
	flag.Parse()
	setupHouses()
	SA()

}
