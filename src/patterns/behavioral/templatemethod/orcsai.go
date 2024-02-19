package templatemethod


// std
import(
	"fmt"
)


type OrcsAI struct {
	GameAI
}


func NewOrcsAI() *OrcsAI {
	gameAI := &OrcsAI{}
	gameAI.GameAI = *NewGameAI(gameAI)

	return gameAI
}

func (self *OrcsAI) BuildStructures() {
	fmt.Println("Building Structures OrcsAI...")
}

func (self *OrcsAI) BuildUnits() {
	fmt.Println("Building Units OrcsAI...")
}

func (self *OrcsAI) SendScouts(position int) {
	fmt.Println("Sending Scouts OrcsAI...", position)
}

func (self *OrcsAI) SendWarriors(position int) {
	fmt.Println("Sending Warriors OrcsAI...", position)
}