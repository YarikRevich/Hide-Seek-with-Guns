package Map

import (
	"math"
	"Game/Window"
	"Game/Heroes/Users"
	"github.com/faiface/pixel"
)



//!!! IT IS A PART OF HERO BORDERS' OPTIONS !!!



type HeroBorder interface {
	//It has all the important methods
	//for hero border tracking 

	//Returns a right border of map
	Right() int

	//Returns a top border of map
	Top() int

	//Returns a bottom border of map
	Bottom() int

	//Returns a left border of map
	Left() int
}

type HB struct{}

func (b HB) Right() int {
	return 1223
}

func (b HB) Left() int {
	return -295
}

func (b HB) Top() int {
	return 805
}

func (b HB) Bottom() int {
	return -225
}

//!!! IT IS A CAM BORDERS' OPTIONS

type CamBorder interface {
	//It has all the important methods
	//for cam border tracking 

	//Inits fullmap for the futher
	//border tracking
	Init(FullMap *pixel.Sprite)

	//Returns a right border for camera
	Right()                    float64

	//Returns a top border for camera
	Top()                      float64

	//Returns a bottom border for camera
	Bottom()                   float64

	//Returns a left border for camera
	Left()                     float64
}

type CB struct {
	FullMap *pixel.Sprite
}

func (c *CB) Init(FullMap *pixel.Sprite) {
	c.FullMap = FullMap
}

func (c CB) Right() float64 {
	return 1244 / 1.65
}

func (c CB) Left() float64 {
	return c.FullMap.Picture().Bounds().Center().X / 4.3
}

func (c CB) Top() float64 {
	return 805 / 1.55
}

func (c CB) Bottom() float64 {
	return c.FullMap.Picture().Bounds().Center().Y / 17
}



//!!! IT IS A PART OF GENERAL COLLISIONS !!!



type UtilCollisions interface{
	//Main utilible methods for collisions checking ...

	Beetwen(float64, float64, float64, float64) bool
}

type UC struct{}


func (c UC) Beetwen(x float64, cx float64, coefp float64, coefn float64) bool {
	//Checks whether current coords are in the
	//available range ...

	if ((x >= (cx - coefp)) && (x <= cx)) || ((x >= (cx + coefn)) && (x <= cx)) {
		return true
	}
	return false
}


type MapCollisions interface{
	//Configures map collisions ...

	Init()
	IsColumn(pixel.Vec, UtilCollisions) bool
}

type MC struct{
	Colls        map[string][][]pixel.Vec
}

func (m *MC) Init(){
	m.Colls = map[string][][]pixel.Vec{
		"ver": {
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
			[]pixel.Vec{pixel.V(-33, -92), pixel.V(-33, -227)},
			[]pixel.Vec{pixel.V(81, 37), pixel.V(81, -34)},
			[]pixel.Vec{pixel.V(107, 37), pixel.V(107, -34)},
			[]pixel.Vec{pixel.V(302, -168), pixel.V(302, -227)},
			[]pixel.Vec{pixel.V(302, 37), pixel.V(302, -116)},
			[]pixel.Vec{pixel.V(329, -168), pixel.V(329, -227)},
			[]pixel.Vec{pixel.V(329, 37), pixel.V(329, -116)},
		},
		"hor": {
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
			[]pixel.Vec{pixel.V(-296, -119), pixel.V(-86, -119)},
			[]pixel.Vec{pixel.V(-184, 43), pixel.V(74, 43)},
			[]pixel.Vec{pixel.V(-28, -87), pixel.V(101, -87)},
			[]pixel.Vec{pixel.V(74, -35), pixel.V(104, -35)},
			[]pixel.Vec{pixel.V(110, 46), pixel.V(300, 46)},
		},
	}
}

func (c MC) IsColumn(vector pixel.Vec, util UtilCollisions) bool {
	//Checks whether next position is a collision.

	for _, vec := range c.Colls["ver"] {
		if util.Beetwen(vector.X, vec[0].X, 6.5, 6.5) && ((vec[0].Y >= vector.Y) && (vec[1].Y <= vector.Y)) {
			return true
		}
	}

	for _, vec := range c.Colls["hor"] {
		if util.Beetwen(vector.Y, vec[0].Y, 6.5, 6.5) && ((vec[0].X <= vector.X) && (vec[1].X >= vector.X)) {
			return true
		}
	}
	return false
}

type DoorsCollisions interface{
	Init()
	DoorTraker(pixel.Vec)
	IsDoor(pixel.Vec)                          (pixel.Vec, string, bool)
	DeleteDoor(pixel.Vec)
	IsNearDeletedDoor(pixel.Vec)                bool
	RecreateDeletedDoors()
	DrawDoors(func(pixel.Vec), func(pixel.Vec))
}

type DC struct{
	Doors        map[string][]pixel.Vec
	DeletedDoors map[string][]pixel.Vec
}

func (d *DC) Init(){

	d.Doors = map[string][]pixel.Vec{
		"hor": {
			pixel.V(6.5, 641),
		},
		"ver": {
			pixel.V(-68, 780),
			pixel.V(278, 784),
			pixel.V(88.5, 410),
			pixel.V(587.5, 318.7),
			pixel.V(865, 509),
			pixel.V(1026.5, 720),
			pixel.V(640, 537.5),
			pixel.V(989, 48),
			pixel.V(914, -192),
			pixel.V(597.5, -86.5),
			pixel.V(313.5, -164.4),
			pixel.V(91, -82),
			pixel.V(89, 239.5),
		},
	}

	d.DeletedDoors = map[string][]pixel.Vec{}
}

func (c *DC) DoorTraker(vector pixel.Vec){
	vector, _, ok := c.IsDoor(vector)

	if ok{
		c.DeleteDoor(vector)
	}
	if !c.IsNearDeletedDoor(vector){
	 	c.RecreateDeletedDoors()
	}
}

func (c *DC) IsDoor(vector pixel.Vec) (pixel.Vec, string, bool) {
	//Checks whether next position is a door

	for key, values := range c.Doors {
		for _, value := range values {
			if math.Abs(value.X-vector.X) <= 30 && math.Abs(value.Y-vector.Y) <= 60 {
				return value, key, true
			}
		}
	}
	return vector, "-", false
}

func (c *DC) DeleteDoor(vector pixel.Vec) {
	for variant, values := range c.Doors {
		for index, value := range values {
			if value.Eq(vector) {
				c.DeletedDoors[variant] = append(c.DeletedDoors[variant], value)
				c.Doors[variant] = append(values[:index], values[index+1:]...)
			}
		}
	}
}

func (c *DC) IsNearDeletedDoor(vector pixel.Vec) bool{
	for _, values := range c.DeletedDoors {
		for _, value := range values {
			if math.Abs(value.X-vector.X) <= 30 && math.Abs(value.Y-vector.Y) <= 60 {
				return true
			}
		}
	}
	return false
}

func (c *DC) RecreateDeletedDoors() {
	//Recrates deleted doors.

	for variant, values := range c.DeletedDoors{
		for index, value := range values{
			c.Doors[variant] = append(c.Doors[variant], value)
			c.DeletedDoors[variant] = append(values[:index], values[index+1:]...)
		}
	}
}

func (c DC) DrawDoors(drawHor func(pixel.Vec), drawVer func(pixel.Vec)) {
	hor := c.Doors["hor"]
	ver := c.Doors["ver"]
	for _, value := range hor {
		drawHor(value)
	}
	for _, value := range ver {
		drawVer(value)
	}
}


//Main configurator for hero icon collisions ...

type HeroCollisions interface{
	IsHero(pixel.Vec, []*Users.User, UtilCollisions) heroPosControllers
}

type HC struct{}


type heroPosControllers interface{
	Top()    bool
	Bottom() bool
	Left()   bool
	Right()  bool
}

type hP struct{
	vector        pixel.Vec
	othervectors []*Users.User
	util          UtilCollisions
}

func (h *hP) Top()bool{
	for _, value := range h.othervectors{
		if math.Abs(float64(value.X)-h.vector.X) <= 10 && math.Abs(float64(value.Y)-h.vector.Y) <= 17{
			return true
		}
	}
	return false
}

func (h *hP) Bottom()bool{
	for _, value := range h.othervectors{
		if math.Abs(float64(value.X)-h.vector.X) <= 10 && math.Abs(float64(value.Y)-h.vector.Y) <= 17{
			return true
		}
	}
	return false
}

func (h *hP) Left()bool{
	for _, value := range h.othervectors{
		if math.Abs(float64(value.X)-h.vector.X) <= 10 && math.Abs(float64(value.Y)-h.vector.Y) <= 17{
			return true
		}
	}
	return false
}

func (h *hP) Right()bool{
	for _, value := range h.othervectors{
		if math.Abs(float64(value.X)-h.vector.X) <= 10 && math.Abs(float64(value.Y)-h.vector.Y) <= 17{
			return true
		}
	}
	return false
}


func (h *HC) IsHero(vector pixel.Vec, othervectors []*Users.User, util UtilCollisions) heroPosControllers{

	return heroPosControllers(&hP{
		vector:  vector,
		othervectors: othervectors,
		util: util,
	})
}


//Inits all the types of collisions
type Collisions interface {
	Init()
	IsCritical(pixel.Vec, []*Users.User, string) bool
	GetHeroCollisions() HeroCollisions
	GetDoorsCollisions() DoorsCollisions
	GetMapCollisions() MapCollisions
	GetUtilCollisions() UtilCollisions
}

type C struct {
	HeroCollisions HeroCollisions
	DoorsCollisions DoorsCollisions
	MapCollisions MapCollisions
	UtilCollisions UtilCollisions
}

func (c *C) Init() {
	c.HeroCollisions = HeroCollisions(new(HC))
	doorscoll := DoorsCollisions(new(DC))
	doorscoll.Init()
	c.DoorsCollisions = doorscoll
	mapcoll := MapCollisions(new(MC))
	mapcoll.Init()
	c.MapCollisions = mapcoll
	c.UtilCollisions = UtilCollisions(new(UC))
	
}

func (c *C) IsCritical(vector pixel.Vec, otherusers []*Users.User, types string)bool{
	//Checks whether each of beneath probable issues are existing
	//If though only one does it returns true ...
	
	if c.GetMapCollisions().IsColumn(vector, c.GetUtilCollisions()){
		return true
	}
	switch types{
	case "top": 
		return c.GetHeroCollisions().IsHero(vector, otherusers, c.GetUtilCollisions()).Top()
	case "bottom":
		return c.GetHeroCollisions().IsHero(vector, otherusers, c.GetUtilCollisions()).Bottom()
	case "left":
		return c.GetHeroCollisions().IsHero(vector, otherusers, c.GetUtilCollisions()).Left()
	case "right":
		return c.GetHeroCollisions().IsHero(vector, otherusers, c.GetUtilCollisions()).Right()
	}
	return false
}

func (c *C) GetHeroCollisions() HeroCollisions{
	return c.HeroCollisions
}

func (c *C) GetDoorsCollisions() DoorsCollisions{
	return c.DoorsCollisions
}

func (c *C) GetMapCollisions() MapCollisions{
	return c.MapCollisions
}

func (c *C) GetUtilCollisions() UtilCollisions{
	return c.UtilCollisions
}



///!!! IT IS A PART OF MAP ANALIZER !!!



type Analizer interface {
	Init(float64, float64, []int, Collisions)
	AnalizeAvailablePlaces()[]pixel.Vec
	ChangeToBottom()
	ChangeToBottomException()
	ChangeToTop()
	ChangeToTopException()
	//ChangeToLeft()
	//ChangeToRight()
}

type Analises struct {
	Borders     []int
	Base        pixel.Vec
	Regime      string
	BeingPlaces []pixel.Vec
	Collisions  Collisions
}

func (A *Analises) Init(x float64, y float64, borders []int, collisions Collisions) {
	A.Base = pixel.V(x, y)
	A.Borders = borders
	A.Collisions = collisions
	A.Regime = "top"
}

func (A *Analises)ChangeToBottom(){
	A.Regime = "bottom"
	A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X, A.Base.Y+20))
	A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X+20, A.Base.Y))
	A.Base = pixel.V(A.Base.X+20, A.Base.Y+20)
}

func (A *Analises)ChangeToBottomException(){
	A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X, A.Base.Y+20))
	A.Base = pixel.V(A.Base.X, A.Base.Y+20)
}

func (A *Analises)ChangeToTop(){
	A.Regime = "top"
	A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X, A.Base.Y-20))
	A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X+20, A.Base.Y))
	A.Base = pixel.V(A.Base.X+20, A.Base.Y-20)
}

func (A *Analises)ChangeToTopException(){
	A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X, A.Base.Y-20))
	A.Base = pixel.V(A.Base.X, A.Base.Y-20)
}

func (A *Analises) AnalizeAvailablePlaces()[]pixel.Vec{
	//Analizes all the available places due to
	//beneath algorithm. It has two regimes.
	//At the first one it checks whether there is
	//any available place goes to top if there is 
	//any block it tries to go to the right and does all
	//the places actions before. If there is regime 'bottom'
	//it does everything vice verse. At last it returns 
	//massive with all the available places.

	for i:=0; i<=150; i++{
		switch A.Regime {
		case "top":
			var found bool = false
			var exc bool = false
			if A.Collisions.GetMapCollisions().IsColumn(pixel.V(A.Base.X, A.Base.Y+46), A.Collisions.GetUtilCollisions()){
				if !A.Collisions.GetMapCollisions().IsColumn(pixel.V(A.Base.X+20, A.Base.Y), A.Collisions.GetUtilCollisions()){
					A.ChangeToBottom()
					found = true
				}else{
					A.ChangeToBottomException()
					exc = true
				}
			}
			for _, value := range A.BeingPlaces {
				if pixel.V(A.Base.X, A.Base.Y + 21).Eq(value) {
					if !pixel.V(A.Base.X+20, A.Base.Y).Eq(value){
						A.ChangeToBottom()
						found = true
					}else{
						A.ChangeToBottomException()
						exc = true
					}
				}
			}
			for _, value := range A.Borders {
				if (A.Base.Y + 20) == float64(value) {
					if !(A.Base.X+20 == float64(value)){
						A.ChangeToBottom()
						found = true
					}else{
						A.ChangeToBottomException()
						exc = true
					}
				}
			}
			if !found && !exc{
				A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X, A.Base.Y+20))
				A.Base = pixel.V(A.Base.X, A.Base.Y+20)
			}
		case "bottom":
			var found bool = false
			var exc bool = false
			if A.Collisions.GetMapCollisions().IsColumn(pixel.V(A.Base.X, A.Base.Y-46), A.Collisions.GetUtilCollisions()){
				if !A.Collisions.GetMapCollisions().IsColumn(pixel.V(A.Base.X+20, A.Base.Y), A.Collisions.GetUtilCollisions()){
					A.ChangeToTop()
					found = true
				}else{
				 	A.ChangeToTopException()
				 	exc = true
				}
			}
			for _, value := range A.BeingPlaces {
				if pixel.V(A.Base.X, A.Base.Y - 21).Eq(value){
					if !pixel.V(A.Base.X+20, A.Base.Y).Eq(value){
						A.ChangeToTop()
						found = true
					}else{
					 	A.ChangeToTopException()
					 	exc = true
					}
				}
			}
			for _, value := range A.Borders {
				if (A.Base.Y - 20) == float64(value) {
					if !(A.Base.X+20 == float64(value)){
						A.ChangeToTop()
						found = true
					}else{
						A.ChangeToTopException()
						exc = true
					}
				}
			}
			if !found && !exc{
				A.BeingPlaces = append(A.BeingPlaces, pixel.V(A.Base.X, A.Base.Y-20))
				A.Base = pixel.V(A.Base.X, A.Base.Y-20)
			}
		}
	}
	return A.BeingPlaces
}



//!!! IT IS A PART OF GENERAL CAMERA SETTINGS !!!



type Cam interface{
	//Special interface to get into the functionality of camera

	//Inits all the important dependences
	Init(winconf *Window.WindowConfig, userconfig Users.User, borders CamBorder)

	//Collibrates camera to the bottom
	collibrateBottom() float64

	//Collibrates camera to the top
	collibrateTop()    float64

	//Collibrates camera to the left
	collibrateLeft()   float64

	//Collibrates camera to the right
	collibrateRight()  float64

	//Generally collibrates camera to its start position
	collibrate()       pixel.Vec

	//Sets collibration and cam position
	SetCam()

	//Sets camera to window matrix scaling it before
	UpdateCam()
}

type CM struct{
	//Structure for camera. Conatains 
	//all the important dependences

	winconf    *Window.WindowConfig
	userconfig Users.User
	borders    CamBorder
}

func (c *CM)Init(winconf *Window.WindowConfig, userconfig Users.User, borders CamBorder){
	//Inits important dependences

	c.winconf    = winconf
	c.userconfig = userconfig
	c.borders    = borders
}

func (c *CM)collibrateBottom()float64{
	//Collibrates camera due to bottom position

	bottom := c.borders.Bottom()
	Y := c.userconfig.Y
	for{
		if float64(Y) >= bottom{
			return float64(Y)
		}
		Y++
	}
}

func (c *CM)collibrateTop()float64{
	//Collibrates camera due to top position

	top := c.borders.Top()
	Y := c.userconfig.Y
	for{
		if float64(Y) <= top{
			return float64(Y)
		}
		Y--
	}
}

func (c *CM)collibrateLeft()float64{
	//Collibrates camera due to left position

	left := c.borders.Left()
	X := c.userconfig.X
	for{
		if float64(X) >= left{
			return float64(X)
		}  
		X++
	}
} 

func (c *CM)collibrateRight()float64{
	//Collibrates camera due to	right position

	right := c.borders.Right()
	X := c.userconfig.X
	for{
		if float64(X) <= right{
			return float64(X)
		}
		X--
	}
}

func (c *CM)collibrate()pixel.Vec{
	//Generally collibrates camera due to different positions

	var X float64
	var Y float64
	Y = c.collibrateBottom()
	NewY := c.collibrateTop()
	if NewY != float64(c.userconfig.Y){
		Y = NewY
	}
	X = c.collibrateLeft()
	NewX := c.collibrateRight()
	if NewX != float64(c.userconfig.X){
		X = NewX
	}
	return pixel.V(X, Y)

}

func (c *CM)SetCam(){
	//Sets all the camera's configurations

	coords := c.collibrate()
	c.winconf.Cam.CamPos = pixel.V(coords.X, coords.Y)
	c.winconf.Cam.CamZoom = 1.0
}

func (c *CM)UpdateCam(){
	//Sets camera as matrix scaling it

	cam := pixel.IM.Scaled(c.winconf.Cam.CamPos, c.winconf.Cam.CamZoom).Moved(c.winconf.Win.Bounds().Center().Sub(c.winconf.Cam.CamPos))
	c.winconf.Win.SetMatrix(cam)
}



//!!! IT IS A PART OF GENERAL MAP CONFIGURATOR !!!



type MapConf interface{
	//It is a main interface for map configuration

	//Inits all the interfaces
	Init()

	//Low level configurator for hero borders
	ConfHeroBorder()

	//Low level configurator for cam border
	ConfCamBorder(*Window.WindowConfig)

	//Low level configurator for map analizer
	ConfAnalizer()

	//Low level configurator for collisions
	ConfCollisions()

	//Low level configurator for camera
	ConfCam(*Window.WindowConfig, Users.User)

	//High level configurator for all beneath listed configurators
	ConfAll(*Window.WindowConfig, Users.User)

	//Returns hero border interface for the futher using
	GetHeroBorder()                           HeroBorder

	//Returns cam border interface for the futher using
	GetCamBorder()                            CamBorder

	//Returns collisions interface for the futher using
	GetCollisions()                           Collisions

	//Returns map analizer interface for the futher using
	GetAnailizer()                            Analizer

	//Returns camera interface for the futher using
	GetCam()                                  Cam
}

type MapC struct{
	HeroBorder  HeroBorder
	CamBorder   CamBorder
	Analizer    Analizer
	Collisions  Collisions
	Cam         Cam   
}

func (MC *MapC)Init(){
	MC.HeroBorder  = HeroBorder(new(HB))
	MC.CamBorder   = CamBorder(new(CB))
	MC.Analizer    = Analizer(new(Analises))
	MC.Collisions  = Collisions(new(C))
	MC.Cam         = Cam(new(CM))
}

func (MC *MapC)ConfHeroBorder(){
	//Here is nothing to configure, it is already done while initiating
}

func (MC *MapC)ConfCamBorder(winConf *Window.WindowConfig){
	MC.CamBorder.Init(winConf.BGImages.Game)
}

func (MC *MapC)ConfCollisions(){
	MC.Collisions.Init()
}

func (MC *MapC)ConfAnalizer(){
	MC.Analizer.Init(
		float64(MC.HeroBorder.Left()), 
		float64(MC.HeroBorder.Bottom()), 
		[]int{
			MC.HeroBorder.Left(), 
			MC.HeroBorder.Bottom(), 
			MC.HeroBorder.Right(), 
			MC.HeroBorder.Top(),
		}, 
		MC.Collisions,
		)
}

func (MC *MapC)ConfCam(winConf *Window.WindowConfig, userConfig Users.User){
	MC.Cam.Init(winConf, userConfig, MC.GetCamBorder())
	MC.Cam.SetCam()
}

func (MC *MapC)ConfAll(winConf *Window.WindowConfig, userConfig Users.User){
	//Configures all the map components

	MC.ConfHeroBorder()
	MC.ConfCamBorder(winConf)
	MC.ConfCollisions()
	MC.ConfAnalizer()
	MC.ConfCam(winConf, userConfig)
}

func (MC *MapC)GetHeroBorder()HeroBorder{
	return MC.HeroBorder
}

func (MC *MapC)GetCamBorder()CamBorder{
	return MC.CamBorder
}

func (MC *MapC)GetCollisions()Collisions{
	return MC.Collisions
}

func (MC *MapC)GetAnailizer()Analizer{
	return MC.Analizer
}

func (MC *MapC)GetCam()Cam{
	return MC.Cam
}
