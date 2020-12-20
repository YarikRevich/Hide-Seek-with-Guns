package Map

import (
	"fmt"
	"github.com/faiface/pixel"
)

type HeroBorder interface{
	Init(FullMap pixel.Picture)
	Right() int
	Top() int
	Bottom() int
	Left() int
}

type HB struct{
	FullMap pixel.Picture
}

func (b *HB)Init(FullMap pixel.Picture){
	b.FullMap = FullMap
}

func (b HB)Right()int{
	return 1244
}

func (b HB)Left()int{
	return -295
}

func (b HB)Top()int{
	return 805
}

func (b HB)Bottom()int{
	return -225
}

type CamBorder interface{
	Init(FullMap pixel.Picture)
	Right() float64
	Top() float64
	Bottom() float64
	Left() float64
}

type CB struct{
	FullMap pixel.Picture
}

func (c *CB)Init(FullMap pixel.Picture){
	c.FullMap = FullMap
}

func (c CB)Right()float64{
	return 1244/1.65
}

func (c CB)Left()float64{
	return c.FullMap.Bounds().Center().X/4.3
}

func (c CB)Top()float64{
	return 805/1.55
}

func (c CB)Bottom()float64{
	return c.FullMap.Bounds().Center().Y/17
}

type Collisions interface{
	Init()
	Beetwen(x float64, cx float64)bool
	IsCollision(vector pixel.Vec) bool
	React(vector pixel.Vec)
}

type C struct{
	Colls [][]pixel.Vec
}

func (c *C)Init(){
	c.Colls = [][]pixel.Vec{
		[]pixel.Vec{pixel.V(-84, 768), pixel.V(-84, 683)},
		[]pixel.Vec{pixel.V(-49, 768), pixel.V(-49, 683)},
		[]pixel.Vec{pixel.V(-79, 775), pixel.V(-55, 775)},
	}
}

func (c C)Beetwen(x float64, cx float64)bool{
	if ((x >= (cx-8)) && (x <= cx)) || ((x >= (cx+8)) && (x <= cx)){
		return true
	}
	return false
}

func (c C)IsCollision(vector pixel.Vec)bool{
	for _, vec := range c.Colls{
		if c.Beetwen(vector.X, vec[0].X) && ((vec[0].Y >= vector.Y) && (vec[1].Y <= vector.Y)){
			return true
		}  
		if c.Beetwen(vector.X, vec[0].X) && ((vec[0].X >= vector.X) && (vec[1].X <= vector.X)){
			return true
		}  
	}
	return false
}

func (c C)React(vector pixel.Vec){
	fmt.Println("You are at collision")
}
