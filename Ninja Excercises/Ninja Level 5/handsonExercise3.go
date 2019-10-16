package main

import (
	"fmt"
)

func main() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	pingao := truck{
		vehicle: vehicle{
			doors: 4,
			color: "maroon",
		},
		fourWheel: true,
	}

	hauwao := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "racing yellow",
		},
		luxury: false,
	}

	fmt.Println(pingao)
	fmt.Println(hauwao)

	fmt.Println(pingao.doors)
	fmt.Println(pingao.color)
	fmt.Println(pingao.fourWheel)

	fmt.Println(hauwao.doors)
	fmt.Println(hauwao.color)
	fmt.Println(hauwao.luxury)
}
