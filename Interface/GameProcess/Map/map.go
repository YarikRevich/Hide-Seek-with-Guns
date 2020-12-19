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
	IsCollision(vector pixel.Vec) bool
	React(vector pixel.Vec)
}

type C struct{
	Colls []pixel.Vec
}

func (c *C)Init(){
	c.Colls = []pixel.Vec{pixel.V(853, 434)}
}

func (c C)IsCollision(vector pixel.Vec)bool{
	for _, vec := range c.Colls{
		if (((vector.X-20) <= vec.X) && ((vector.X+20) >= vec.X)) && (((vector.Y-20) <= vec.Y) && ((vector.Y+20) >= vec.Y)){ 
			return true
		}
	}
	return false
}

func (c C)React(vector pixel.Vec){
	fmt.Println("You are at collision")
}
