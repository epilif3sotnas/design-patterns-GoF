package templatemethod


// std
import(
	"fmt"
)


type IGameAI interface {
	BuildStructures()
	BuildUnits()
	SendScouts(position int)
	SendWarriors(position int)
}

type GameAI struct {
	iGameAI IGameAI
}


func NewGameAI(iGameAI IGameAI) *GameAI {
	return &GameAI{
		iGameAI: iGameAI,
	}
}

func (self *GameAI) Turn() {
	self.CollectResources()
	self.iGameAI.BuildStructures()
	self.iGameAI.BuildUnits()
	self.Attack(false)
}

func (self *GameAI) CollectResources() {
	fmt.Println("Collecting resources")
}

func (self *GameAI) Attack(isEnemyClose bool) {
	fmt.Println("Attacking")

	if isEnemyClose == true {
		self.iGameAI.SendWarriors(1)
	}

	self.iGameAI.SendScouts(1)
}

func (self *GameAI) SetIGameAI(iGameAI IGameAI) {
	self.iGameAI = iGameAI
}

func (self *GameAI) GetIGameAI() IGameAI {
	return self.iGameAI
}