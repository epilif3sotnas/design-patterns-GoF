package templatemethod


// std
import(
	"fmt"
)


type MonstersAI struct {
	GameAI
}


func NewMonstersAI() *MonstersAI {
	gameAI := &MonstersAI{}
	gameAI.GameAI = *NewGameAI(gameAI)

	return gameAI
}

func (self *MonstersAI) BuildStructures() {
	fmt.Println("Building Structures MonstersAI...")
}

func (self *MonstersAI) BuildUnits() {
	fmt.Println("Building Units MonstersAI...")
}

func (self *MonstersAI) CollectResources() {
	fmt.Println("Collecting Resources Not Available...")
}

func (self *MonstersAI) SendScouts(position int) {
	fmt.Println("Sending Scouts MonstersAI...", position)
}

func (self *MonstersAI) SendWarriors(position int) {
	fmt.Println("Sending Warriors MonstersAI...", position)
}