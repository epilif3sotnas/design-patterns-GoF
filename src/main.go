package main


// std
import (
	"encoding/json"
	"fmt"
	"math"
)

// internal
import (
	"learn-design-patterns-GoF/patterns/creational/abstractfactory"
	"learn-design-patterns-GoF/patterns/creational/builder"
	"learn-design-patterns-GoF/patterns/creational/factorymethod"
	"learn-design-patterns-GoF/patterns/creational/prototype"
	"learn-design-patterns-GoF/patterns/creational/singleton"
	"learn-design-patterns-GoF/patterns/structural/adapter"
	"learn-design-patterns-GoF/patterns/structural/bridge"
	"learn-design-patterns-GoF/patterns/structural/composite"
	"learn-design-patterns-GoF/patterns/structural/decorator"
	"learn-design-patterns-GoF/patterns/structural/facade"
	"learn-design-patterns-GoF/patterns/structural/flyweight"
	"learn-design-patterns-GoF/patterns/structural/proxy"
	"learn-design-patterns-GoF/patterns/behavioral/chainofresponsability"
	"learn-design-patterns-GoF/patterns/behavioral/command"
	"learn-design-patterns-GoF/patterns/behavioral/interpreter"
	"learn-design-patterns-GoF/patterns/behavioral/iterator"
	"learn-design-patterns-GoF/patterns/behavioral/mediator"
	"learn-design-patterns-GoF/patterns/behavioral/memento"
	"learn-design-patterns-GoF/patterns/behavioral/observer"
	"learn-design-patterns-GoF/patterns/behavioral/state"
	"learn-design-patterns-GoF/patterns/behavioral/strategy"
)


func main() {
	// ################################
	//  Design Pattern Abstract Factory => more information in https://refactoring.guru/design-patterns/abstract-factory
	fmt.Print("\n\nDesign Pattern Abstract Factory\n\n")
	
	cat := abstractfactory.NewCat("Stewie", "Cat")
	abstractFactory_ := abstractfactory.NewAbstactFactory(cat)
	fmt.Println(abstractFactory_.GetAnimalInfo())
  
	// ################################
  
  
	// ################################
	// Design Pattern Builder => more information in https://refactoring.guru/design-patterns/builder
	fmt.Print("\n\nDesign Pattern Builder\n\n")
  
	setup := builder.NewSetup()
	admin := builder.NewAdmin(setup)
	car := admin.Build()
	fmt.Println(car.GetCarInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Factory Method => more information in https://refactoring.guru/design-patterns/factory-method
	fmt.Print("\n\nDesign Pattern Factory Method\n\n")
  
	creator := factorymethod.NewCreator()
	company := creator.Create()
	fmt.Println(company.GetCompanyInfo())
	
	// ################################
  
  
	// ################################
	// Design Pattern Prototype => more information in https://refactoring.guru/design-patterns/prototype
	fmt.Print("\n\nDesign Pattern Prototype\n\n")
  
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
	fmt.Print("\n\nDesign Pattern Singleton\n\n")
  
	singleton.GetInstance().AddCar("VW-Golf")
	singleton.GetInstance().AddCar("VW-Passat")
	singleton.GetInstance().AddCar("Ford-Fiesta")

	fmt.Println(singleton.GetInstance().GetCars())
	
	// ################################
  
  
	// ################################
	// Design Pattern Adapter => more information in https://refactoring.guru/design-patterns/adapter
	fmt.Print("\n\nDesign Pattern Adapter\n\n")
  
	fmt.Print("Object Adapter\n\n")
  
	government := adapter.NewGovernment()
	governmentImpl := adapter.NewGovernmentAdapterImpl(government)
	broker := adapter.NewBroker(governmentImpl)

	value, _ := json.Marshal(map[string]any{
		"tax_payer_id": "1234567890",
		"tax_payer_gains": 10000,
	})

	broker.SendGains(value)
  
  
	fmt.Print("\nClass Adapter\n\n")
  
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
	fmt.Print("\n\nDesign Pattern Bridge\n\n")
  
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
	fmt.Print("\n\nDesign Pattern Composite\n\n")
  
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
	// Design Pattern Decorator => more information in https://refactoring.guru/design-patterns/decorator
	fmt.Print("\n\nDesign Pattern Decorator\n\n")
  
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


	// ################################
	// Design Pattern Facade => more information in https://refactoring.guru/design-patterns/facade
	fmt.Print("\n\nDesign Pattern Facade\n\n")
  
	computer := facade.NewComputer()
	computer.Start()
  
	// ################################


	// ################################
	// Design Pattern Flyweight => more information in https://refactoring.guru/design-patterns/flyweight
	fmt.Print("\n\nDesign Pattern Flyweight\n\n")
  
	forest := flyweight.NewForest()
	forest.PlantTree(13.12, 12.12, "Ash", "green", "wood")
	forest.PlantTree(12.12, 12.2, "Beech", "red", "wood")
	forest.PlantTree(1.12, 11.12, "Cherry", "white", "plant")
	forest.PlantTree(12.12, 2.12, "Aspen", "green", "wood")

	canvas := flyweight.NewCanvas(100, 100)
	forest.Draw(canvas)

	// ################################


	// ################################
	// Design Pattern Proxy => more information in https://refactoring.guru/design-patterns/proxy
	fmt.Print("\n\nDesign Pattern Proxy\n\n")
  
	aYoutubeService := proxy.NewThirdPartyYoutubeClass()
	aYoutubeProxy := proxy.NewCachedYoutubeClass(aYoutubeService)
	youtubeManager := proxy.NewYoutubeManager(aYoutubeProxy)

	youtubeManager.RenderListPanel()
	youtubeManager.RenderVideoPage("AAA")
	youtubeManager.RenderVideo("AAA")

	// ################################


	// ################################
	// Design Pattern Chain of Responsability => more information in https://refactoring.guru/design-patterns/chain-of-responsibility
	fmt.Print("\n\nDesign Pattern Chain of Responsability\n\n")
  
	dialog := chainofresponsability.NewDialog(
		nil,
		"Dialog: Tool Tip Text",
		[]*chainofresponsability.Component{},
		"Dialog: Wiki Page URL",
	)
	panel := chainofresponsability.NewPanel(
		nil,
		"Panel: Tool Tip Text",
		[]*chainofresponsability.Component{},
		"Panel: Model Help Text",
	)
	button := chainofresponsability.NewButton(
		nil,
		"Button: Tool Tip Text",
	)

	panel.Add(&button.Component)
	dialog.Add(&panel.Component)

	dialog.ShowHelp()
	panel.ShowHelp()
	button.ShowHelp()
	
	// ################################


	// ################################
	// Design Pattern Command => more information in https://refactoring.guru/design-patterns/command
	fmt.Print("\n\nDesign Pattern Command\n\n")
  
	editor := command.NewEditor("Hello, World!")
	commandHistory := command.NewCommandHistory([]command.Command{})

	app := command.NewApplication(
		"I AM HERE",
		[]*command.Editor{editor},
		editor,
		commandHistory,
	)
	app.InitCommands()

	fmt.Println("1 (No Command):", app.GetActiveEditor().GetSelection())

	app.ExecuteCommand(command.PASTE)
	fmt.Println("2 (PASTE Command):", app.GetActiveEditor().GetSelection())
	
	app.ExecuteCommand(command.UNDO)
	fmt.Println("3 (UNDO Command):", app.GetActiveEditor().GetSelection())

	app.ExecuteCommand(command.CUT)
	fmt.Println("4 (CUT Command):", app.GetActiveEditor().GetSelection())

	app.ExecuteCommand(command.PASTE)
	fmt.Println("5 (PASTE Command):", app.GetActiveEditor().GetSelection())

	app.SetClipboard(app.GetActiveEditor().GetSelection() + "\tA New Line")
	app.ExecuteCommand(command.PASTE)
	fmt.Println("6 (PASTE Command):", app.GetActiveEditor().GetSelection())

	app.ExecuteCommand(command.COPY)
	fmt.Println("7 (COPY Command):", app.GetActiveEditor().GetSelection())

	app.ExecuteCommand(command.PASTE)
	fmt.Println("8 (PASTE Command):", app.GetActiveEditor().GetSelection())
	
	// ################################


	// ################################
	// Design Pattern Interpreter => more information in https://refactoring.guru/design-patterns/interpreter
	fmt.Print("\n\nDesign Pattern Interpreter\n\n")
  
	person1 := interpreter.NewTerminalExpression("Filipe")
	person2 := interpreter.NewTerminalExpression("Santos")
	isSingle := interpreter.NewOrExpression(person1, person2)

	vikram := interpreter.NewTerminalExpression("Vikram")
	committed := interpreter.NewTerminalExpression("Committed")
	isCommitted := interpreter.NewAndExpression(vikram, committed)

	fmt.Println(isSingle.Interpreter("Kushagra"))
	fmt.Println(isSingle.Interpreter("Lokesh"))
	fmt.Println(isSingle.Interpreter("Filipe"))
	fmt.Println(isSingle.Interpreter("Achint"))
	fmt.Println(isSingle.Interpreter("Santos"))
		
	fmt.Println(isCommitted.Interpreter("Committed, Vikram"))
	fmt.Println(isCommitted.Interpreter("Single, Vikram"))
	
	// ################################


	// ################################
	// Design Pattern Iterator => more information in https://refactoring.guru/design-patterns/iterator
	fmt.Print("\n\nDesign Pattern Iterator\n\n")

	socialSpammer := iterator.NewSocialSpammer()
	facebook := iterator.NewFacebook()

	profile1 := iterator.NewProfile("1", "email@example.com")

	app2 := iterator.NewApplication(socialSpammer, facebook)
	app2.SendSpamToFriends(profile1)
	app2.SendSpamToCoworkers(profile1)

	// ################################


	// ################################
	// Design Pattern Mediator => more information in https://refactoring.guru/design-patterns/mediator
	fmt.Print("\n\nDesign Pattern Mediator\n\n")

	authenticationDialog := mediator.NewAuthenticationDialog("")


	checkboxMediator := mediator.NewCheckBox(authenticationDialog)
	buttonMediator := mediator.NewButton(authenticationDialog)
	textboxMediator := mediator.NewTextBox(authenticationDialog)

	authenticationDialog.SetLoginOrRegisterChkBx(checkboxMediator)

	checkboxMediator.Check()
	fmt.Println("Status: ", authenticationDialog.GetTitle())
	buttonMediator.Click()
	fmt.Println("Status: ", authenticationDialog.GetTitle())
	textboxMediator.KeyPress()
	fmt.Println("Status: ", authenticationDialog.GetTitle())

	authenticationDialog.GetLoginOrRegisterChkBx().SetChecked(false)
	buttonMediator.Click()
	fmt.Println("Status: ", authenticationDialog.GetTitle())

	authenticationDialog.GetLoginOrRegisterChkBx().SetChecked(true)

	for i := 0; i < 5; i++ {
		buttonMediator.Click()
		fmt.Println("Status: ", authenticationDialog.GetTitle())
	}

	// ################################


	// ################################
	// Design Pattern Memento => more information in https://refactoring.guru/design-patterns/memento
	fmt.Print("\n\nDesign Pattern Memento\n\n")

	editor_ := memento.NewEditor()
	command_ := memento.NewCommand()

	editor_.SetText("Test 1")
	command_.MakeBackup(editor_)

	editor_.SetText("Test 2")
	editor_.SetText("Test 3")
	editor_.SetCursor(1,32)
	editor_.SetSelectionWidth(1)

	fmt.Print(
		"Editor Information:",
		"\nText: ", editor_.GetText(),
		"\nX: ", editor_.GetCurX(),
		"\nY: ", editor_.GetCurY(),
		"\nSelection Width: ", editor_.GetSelectionWidth(),
	)

	fmt.Println("\n\nRestoring editor...")

	command_.Undo()
	
	fmt.Print(
		"\nEditor Information:",
		"\nText: ", editor_.GetText(),
		"\nX: ", editor_.GetCurX(),
		"\nY: ", editor_.GetCurY(),
		"\nSelection Width: ", editor_.GetSelectionWidth(),
		"\n",
	)

	// ################################


	// ################################
	// Design Pattern Observer => more information in https://refactoring.guru/design-patterns/observer
	fmt.Print("\n\nDesign Pattern Observer\n\n")

	eventManager := observer.NewEventManager()
	editor1_ := observer.NewEditor(eventManager)
	app_ := observer.NewApplication(editor1_)

	app_.Config()
	app_.Run()

	// ################################


	// ################################
	// Design Pattern State => more information in https://refactoring.guru/design-patterns/state
	fmt.Print("\n\nDesign Pattern State\n\n")

	player := state.NewAudioPlayer(
		50,
		[]string{"music1", "music2", "music3", "music4", "music5", "music6", "music7"},
		"music1",
		0,
	)
	
	player.SetState(state.NewReadyState(player))

	for i := 0; i < 3; i++ {
		player.ClickPlay()
		player.ClickPlay()
		player.ClickNext()
		player.ClickNext()
		player.ClickPrevious()
		player.ClickLock()
		player.ClickPlay()
		player.ClickLock()
		player.ClickPlay()
	}

	// ################################


	// ################################
	// Design Pattern Strategy => more information in https://refactoring.guru/design-patterns/strategy
	fmt.Print("\n\nDesign Pattern Strategy\n\n")

	add := strategy.NewAdd()
	subtract := strategy.NewSubtract()
	multiply := strategy.NewMultiply()
	divide := strategy.NewDivide()
	
	context_123_ := strategy.NewContext(nil)


	fmt.Println("--- Add ---")
	context_123_.SetStrategy(add)
	context_123_.ExecuteOperation(1, 2)

	fmt.Println("--- Subtract ---")
	context_123_.SetStrategy(subtract)
	context_123_.ExecuteOperation(1, 2)

	fmt.Println("--- Multiply ---")
	context_123_.SetStrategy(multiply)
	context_123_.ExecuteOperation(1, 2)

	fmt.Println("--- Divide ---")
	context_123_.SetStrategy(divide)
	context_123_.ExecuteOperation(1, 2)

	// ################################


	// ################################
	// Design Pattern Template Method => more information in https://refactoring.guru/design-patterns/template-method
	fmt.Print("\n\nDesign Strategy Template Method\n\n")

	

	// ################################
}