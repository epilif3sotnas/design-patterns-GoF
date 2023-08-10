package main


// std
import (
	"encoding/json"
	"fmt"
	"math"
)


// internal
import(
	"learn-design-patterns-GoF/patterns/creational/abstractfactory"
	"learn-design-patterns-GoF/patterns/creational/builder"
	"learn-design-patterns-GoF/patterns/creational/factorymethod"
	"learn-design-patterns-GoF/patterns/creational/prototype"
	"learn-design-patterns-GoF/patterns/creational/singleton"
	"learn-design-patterns-GoF/patterns/structural/adapter"
	"learn-design-patterns-GoF/patterns/structural/bridge"
	"learn-design-patterns-GoF/patterns/structural/composite"
	"learn-design-patterns-GoF/patterns/structural/decorator"
)


func main() {
	// ################################
	//  Design Pattern Abstract Factory => more information in https://refactoring.guru/design-patterns/abstract-factory
	fmt.Println("\n\nDesign Pattern Abstract Factory\n")
	
	cat := abstractfactory.NewCat("Stewie", "Cat")
	abstractFactory_ := abstractfactory.NewAbstactFactory(cat)
	fmt.Println(abstractFactory_.GetAnimalInfo())
  
	// ################################
  
  
	// ################################
	// Design Pattern Builder => more information in https://refactoring.guru/design-patterns/builder
	fmt.Println("\n\nDesign Pattern Builder\n")
  
	setup := builder.NewSetup()
	admin := builder.NewAdmin(setup)
	car := admin.Build()
	fmt.Println(car.GetCarInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Factory Method => more information in https://refactoring.guru/design-patterns/factory-method
	fmt.Println("\n\nDesign Pattern Factory Method\n")
  
	creator := factorymethod.NewCreator()
	company := creator.Create()
	fmt.Println(company.GetCompanyInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Prototype => more information in https://refactoring.guru/design-patterns/prototype
	fmt.Println("\n\nDesign Pattern Prototype\n")
  
	zaraClothing := prototype.NewZara("Zara", "Jeans", 19.99, "Mafia Guy")
	fmt.Println(zaraClothing.GetClothingInfo())
  
	zaraClothing2 := zaraClothing.Copy()
	fmt.Println(zaraClothing2.GetClothingInfo())
  
	zaraClothing2.SetModel("T-Shirt")
	zaraClothing2.SetPrice(15.99)
  
	fmt.Println(zaraClothing.GetClothingInfo())
	fmt.Println(zaraClothing2.GetClothingInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Singleton => more information in https://refactoring.guru/design-patterns/singleton
	fmt.Println("\n\nDesign Pattern Singleton\n")
  
	singleton.GetInstance().AddCar("VW-Golf")
	singleton.GetInstance().AddCar("VW-Passat")
	singleton.GetInstance().AddCar("Ford-Fiesta")

	fmt.Println(singleton.GetInstance().GetCars())
	
	// ################################
  
  
	// ################################
	// Design Pattern Adapter => more information in https://refactoring.guru/design-patterns/adapter
	fmt.Println("\n\nDesign Pattern Adapter\n")
  
	fmt.Println("Object Adapter\n")
  
	government := adapter.NewGovernment()
	governmentImpl := adapter.NewGovernmentAdapterImpl(government)
	broker := adapter.NewBroker(governmentImpl)

	value, _ := json.Marshal(map[string]any{
		"tax_payer_id": "1234567890",
		"tax_payer_gains": 10000,
	})

	broker.SendGains(value)
  
  
	fmt.Println("\nClass Adapter\n")
  
	classAdapter := adapter.NewClassAdapter()
	broker2 := adapter.NewBroker(classAdapter)

	value2, _ := json.Marshal(map[string]any{
		"tax_payer_id": "1234567890",
		"tax_payer_gains": 10000,
	})

	broker2.SendGains(value2)
	
	// ################################
  
  
	// ################################
	// Design Pattern Bridge => more information in https://refactoring.guru/design-patterns/bridge
	fmt.Println("\n\nDesign Pattern Bridge\n")
  
	tv := bridge.NewTv()
	radio := bridge.NewRadio()
  
	remoteRadio := bridge.NewRemote(radio)
	fmt.Println("Radio")
	fmt.Println(radio.GetDeviceInfo())
	
	remoteRadio.VolumeUp()
	remoteRadio.TooglePower()
  
	fmt.Println(radio.GetDeviceInfo())
  
	var remoteTv = bridge.NewRemote(tv)
	fmt.Println("\nTV")
	fmt.Println(tv.GetDeviceInfo())
  
	remoteTv.VolumeDown()
	remoteTv.ChannelUp()
  
	fmt.Println(tv.GetDeviceInfo())
  
  
	advancedRemoteTv := bridge.NewAdvancedRemote(tv)
  
	fmt.Println("\nTV with Advanced Remote")
	fmt.Println(tv.GetDeviceInfo())
  
	advancedRemoteTv.Mute()
  
	fmt.Println(tv.GetDeviceInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Composite => more information in https://refactoring.guru/design-patterns/composite
	fmt.Println("\n\nDesign Pattern Composite\n")
  
	circle := composite.NewCircle(0, 0, math.Pi)
	dot := composite.NewDot(0, 0)
	
	coumpoundGraphic := composite.NewCoumpoundGraphic()
	coumpoundGraphic.AddChild(circle)
	coumpoundGraphic.AddChild(dot)
  
	coumpoundGraphic.Draw()
	coumpoundGraphic.Move(1,1)
	coumpoundGraphic.Draw()
  
	// ################################


	// ################################
	// Design Pattern Composite => more information in https://refactoring.guru/design-patterns/composite
	fmt.Println("\n\nDesign Pattern Decorator\n")
  
	file := decorator.NewFileDataSource("test")
	file.WriteData([]byte("Data to test! Eheheh"))

	encryption := decorator.NewEncryptionDecorator(file)

	fmt.Println("\nFile raw")
	fmt.Println(string(file.ReadData()))

	fmt.Println("\nFile Encrypted")
	encryption.WriteData(file.ReadData())
	fmt.Println(string(file.ReadData()))
	fmt.Println(string(encryption.ReadData()))
  
	// ################################
}