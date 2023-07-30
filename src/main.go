package main

// std
import(
	"fmt"
)

// internal
import(
	"learn-design-patterns-GoF/patterns/creational/abstractfactory"
	"learn-design-patterns-GoF/patterns/creational/builder"
	"learn-design-patterns-GoF/patterns/creational/factorymethod"
	"learn-design-patterns-GoF/patterns/creational/prototype"
)


func main() {
	// ################################
	//  Design Pattern Abstract Factory => more information in https://refactoring.guru/design-patterns/abstract-factory
	fmt.Println("\n\nDesign Pattern Abstract Factory")
	
	cat := abstractfactory.NewCat("Stewie", "Cat")
	abstractFactory_ := abstractfactory.NewAbstactFactory(cat)
	fmt.Println("\n" + abstractFactory_.GetAnimalInfo())
  
	// ################################
  
  
	// ################################
	// Design Pattern Builder => more information in https://refactoring.guru/design-patterns/builder
	fmt.Println("\n\nDesign Pattern Builder")
  
	setup := builder.NewSetup()
	admin := builder.NewAdmin(setup)
	car := admin.Build()
	fmt.Println("\n" + car.GetCarInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Factory Method => more information in https://refactoring.guru/design-patterns/factory-method
	fmt.Println("\n\nDesign Pattern Factory Method")
  
	creator := factorymethod.NewCreator()
	company := creator.Create()
	fmt.Println("\n" + company.GetCompanyInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Prototype => more information in https://refactoring.guru/design-patterns/prototype
	fmt.Println("\n\nDesign Pattern Prototype")
  
	zaraClothing := prototype.NewZara("Zara", "Jeans", 19.99, "Mafia Guy")
	fmt.Println("\n" + zaraClothing.GetClothingInfo())
  
	zaraClothing2 := zaraClothing.Copy()
	fmt.Println(zaraClothing2.GetClothingInfo())
  
	zaraClothing2.SetModel("T-Shirt")
	zaraClothing2.SetPrice(15.99)
  
	fmt.Println(zaraClothing.GetClothingInfo())
	fmt.Println(zaraClothing2.GetClothingInfo())
	
	// ################################
  
  
	// // ################################
	// // Design Pattern Singleton => more information in https://refactoring.guru/design-patterns/singleton
	// echo "\n\nDesign Pattern Singleton\n"
  
	// getInstance().addCar("VW Golf")
	// getInstance().addCar("VW Passat")
	// getInstance().addCar("Ford Fiesta")
	// echo getInstance().getCars()
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Adapter => more information in https://refactoring.guru/design-patterns/adapter
	// echo "\n\nDesign Pattern Adapter\n"
  
	// echo "Object Adapter\n"
  
	// let government = newGovernment()
	// var governmentImpl = newGovernmentAdapterImpl(government)
	// var broker = newBroker[GovernmentAdapterImpl](governmentImpl)
  
	// echo broker.sendGains(parseJson("""{"taxPayerId":1234567890, "taxPayerGains":10000}"""))
  
  
	// echo "\nClass Adapter\n"
  
	// let classAdapter = newClassAdapter()
	// var broker2 = newBroker[ClassAdapter](classAdapter)
  
	// echo broker2.sendGains(parseJson("""{"taxPayerId":1234562490, "taxPayerGains":156000}"""))
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Bridge => more information in https://refactoring.guru/design-patterns/bridge
	// echo "\n\nDesign Pattern Bridge\n"
  
	// var tv = newTv()
	// var radio = newRadio()
  
	// var remoteRadio = newRemote[Radio](radio)
	// echo "Radio"
	// echo radio.getDeviceInfo()
	
	// discard remoteRadio.volumeUp()
	// discard remoteRadio.tooglePower()
  
	// echo radio.getDeviceInfo()
  
	// var remoteTv = newRemote[Tv](tv)
	// echo "\nTV"
	// echo tv.getDeviceInfo()
  
	// discard remoteTv.volumeDown()
	// discard remoteTv.channelUp()
  
	// echo tv.getDeviceInfo()
  
  
	// var advancedRemoteTv = newAdvancedRemote[Tv](tv)
  
	// echo "\nTV with Advanced Remote"
	// echo tv.getDeviceInfo()
  
	// discard advancedRemoteTv.mute()
  
	// echo tv.getDeviceInfo()
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Composite => more information in https://refactoring.guru/design-patterns/composite
	// echo "\n\nDesign Pattern Composite\n"
  
	// var circle = newCircle(0, 0, PI)
	// var dot = newDot(0, 0)
	
	// var coumpoundGraphic = newCoumpoundGraphic()
	// coumpoundGraphic.add(circle)
	// coumpoundGraphic.add(dot)
  
	// discard coumpoundGraphic.draw()
	// discard coumpoundGraphic.move(1,1)
	// discard coumpoundGraphic.draw()
  
	// // ################################
}