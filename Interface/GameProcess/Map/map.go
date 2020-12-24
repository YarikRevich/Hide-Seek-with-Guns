package Map

import (
	"math"
	"github.com/faiface/pixel"
)

type HeroBorder interface{
	Right() int
	Top() int
	Bottom() int
	Left() int
}

type HB struct{}

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
	Init(FullMap *pixel.Sprite)
	Right() float64
	Top() float64
	Bottom() float64
	Left() float64
}

type CB struct{
	FullMap *pixel.Sprite
}

func (c *CB)Init(FullMap *pixel.Sprite){
	c.FullMap = FullMap
}

func (c CB)Right()float64{
	return 1244/1.65
}

func (c CB)Left()float64{
	return c.FullMap.Picture().Bounds().Center().X/4.3
}

func (c CB)Top()float64{
	return 805/1.55
}

func (c CB)Bottom()float64{
	return c.FullMap.Picture().Bounds().Center().Y/17
}

type Collisions interface{
	Init()
	Beetwen(x float64, cx float64)bool
	IsCollision(vector pixel.Vec) bool
	IsDoor(vector pixel.Vec)(pixel.Vec, string, bool)
	DeleteDoor(vector pixel.Vec)
	DrawDoors(drawHor func(pixel.Vec), drawVer func(pixel.Vec))
}

type C struct{
	Colls map[string][][]pixel.Vec
	Doors map[string][]pixel.Vec
}

func (c *C)Init(){
	//Inits all the collisions collected in a specially
	//sorted map.

	c.Colls = map[string][][]pixel.Vec{
			"ver": [][]pixel.Vec{
				[]pixel.Vec{pixel.V(-84, 768), pixel.V(-84, 672)},
				[]pixel.Vec{pixel.V(-49, 768), pixel.V(-49, 672)},
				[]pixel.Vec{pixel.V(260, 757), pixel.V(260, 448)},
				[]pixel.Vec{pixel.V(299, 760), pixel.V(299, 483)},
				[]pixel.Vec{pixel.V(110, 670), pixel.V(110, 457)},
				[]pixel.Vec{pixel.V(-145, 472), pixel.V(-145, 211)},
				[]pixel.Vec{pixel.V(617, -92), pixel.V(617, -227)},
				[]pixel.Vec{pixel.V(617, 13), pixel.V(617, -32)},
				[]pixel.Vec{pixel.V(572, -92), pixel.V(572, -227)},
				[]pixel.Vec{pixel.V(572, 31), pixel.V(572, -32)},
				[]pixel.Vec{pixel.V(275, 319), pixel.V(275, 79)},
				[]pixel.Vec{pixel.V(-107, 513), pixel.V(-107, 211)},
				
			},
			"hor": [][]pixel.Vec{
				[]pixel.Vec{pixel.V(-79, 775), pixel.V(-55, 775)},
				[]pixel.Vec{pixel.V(-79, 676), pixel.V(-55, 676)},
				[]pixel.Vec{pixel.V(268, 772), pixel.V(292, 676)},
				[]pixel.Vec{pixel.V(268, 565), pixel.V(292, 565)},
				[]pixel.Vec{pixel.V(107, 565), pixel.V(251, 565)},
				[]pixel.Vec{pixel.V(107, 526), pixel.V(251, 526)},
				[]pixel.Vec{pixel.V(301, 670), pixel.V(563.5, 670)},
				[]pixel.Vec{pixel.V(300, 408), pixel.V(622, 408)},
				[]pixel.Vec{pixel.V(610, 301), pixel.V(842.5, 301)},
				[]pixel.Vec{pixel.V(607, 49), pixel.V(1000, 49)},
				[]pixel.Vec{pixel.V(934, -54), pixel.V(1088, -54)},
				[]pixel.Vec{pixel.V(1147, -56), pixel.V(1229, -56)},
				[]pixel.Vec{pixel.V(38, 685), pixel.V(100, 685)},
				[]pixel.Vec{pixel.V(-244, 472), pixel.V(-139, 472)},
				[]pixel.Vec{pixel.V(617, 10), pixel.V(895, 10)},
				[]pixel.Vec{pixel.V(491, 40), pixel.V(581, 40)},
				[]pixel.Vec{pixel.V(332, 37), pixel.V(437, 37)},
				[]pixel.Vec{pixel.V(275, 79), pixel.V(437, 79)},
				[]pixel.Vec{pixel.V(495, 79), pixel.V(570, 79)},
				[]pixel.Vec{pixel.V(357, 205), pixel.V(568, 205)},
				[]pixel.Vec{pixel.V(275, 205), pixel.V(297, 205)},
				[]pixel.Vec{pixel.V(275, 241), pixel.V(297, 241)},
				[]pixel.Vec{pixel.V(357, 241), pixel.V(568, 241)},
			},
	}

	c.Doors = map[string][]pixel.Vec{
		"hor": []pixel.Vec{
			pixel.V(6.5, 641),
		},
		"ver": []pixel.Vec{
			pixel.V(-68, 780),
			pixel.V(278, 784),
		},
	}
}

func (c C)Beetwen(x float64, cx float64)bool{
	//Checks whether current coords are in the
	//available range.

	if ((x >= (cx-8)) && (x <= cx)) || ((x >= (cx+8)) && (x <= cx)){
		return true
	}
	return false
}

func (c C)IsCollision(vector pixel.Vec)bool{
	//Checks whether next position is a collision.

	for _, vec := range c.Colls["ver"]{
		if c.Beetwen(vector.X, vec[0].X) && ((vec[0].Y >= vector.Y) && (vec[1].Y <= vector.Y)){
			return true
		} 
	}

	for _, vec := range c.Colls["hor"]{
		if c.Beetwen(vector.Y, vec[0].Y) && ((vec[0].X <= vector.X) && (vec[1].X >= vector.X)){
			return true
		}
	}
	return false
}

func (c C)IsDoor(vector pixel.Vec)(pixel.Vec, string, bool){
	//Checks whether next position is a door 

	for key, values := range c.Doors{
		for _, value := range values{
			if math.Abs(value.X - vector.X) <= 30 && math.Abs(value.Y - vector.Y) <= 60{
				return value, key, true
			}
		}
	}
	return vector, "-", false
}

func (c *C)DeleteDoor(vector pixel.Vec){
	for variant, values := range c.Doors{
		for index, value := range values{
			if value.Eq(vector){
				c.Doors[variant] = append(values[:index], values[index+1:]...)
			}
		}
	}
}

func (c C)DrawDoors(drawHor func(pixel.Vec), drawVer func(pixel.Vec)){
	hor := c.Doors["hor"]
	ver := c.Doors["ver"]
	for _, value := range hor{
		drawHor(value)
	}
	for _, value := range ver{
		drawVer(value)
	}
}
