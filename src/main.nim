# std
import
  std/[
    json
  ]


# internal
import
  patterns/creational/abstractfactory/[
    abstractfactory,
    cat
  ],
  patterns/creational/builder/[
    admin,
    car,
    setup
  ],
  patterns/creational/factorymethod/[
    creator,
    company1
  ],
  patterns/creational/prototype/[
    clothing,
    zara
  ],
  patterns/creational/singleton/[
    carstorage
  ],
  patterns/structural/adapter/[
    broker,
    classadapter,
    government,
    governmentadapter,
    governmentadapterimpl
  ],
  patterns/structural/bridge/[
    tv,
    radio,
    remote
  ]


proc main() =
  ################################
  # Design Pattern Abstract Factory => more information in https://refactoring.guru/design-patterns/abstract-factory
  echo "\n\nDesign Pattern Abstract Factory\n"
  
  var cat = newCat("Stewie", "Cat")
  var abstractFactory = newAbstactFactory[Cat](cat)
  echo abstractFactory.getAnimalInfo()

  ################################


  ################################
  # Design Pattern Builder => more information in https://refactoring.guru/design-patterns/builder
  echo "\n\nDesign Pattern Builder\n"

  let setup = newSetup()
  let admin = newAdmin[Setup](setup)
  var car: Car = admin.build()
  echo car.getCarInfo()
  
  ################################


  ################################
  # Design Pattern Factory Method => more information in https://refactoring.guru/design-patterns/factory-method
  echo "\n\nDesign Pattern Factory Method\n"

  let creator = newCreator()
  var company: Company1 = create[Company1](creator)
  echo company.getCompanyInfo()
  
  ################################


  ################################
  # Design Pattern Prototype => more information in https://refactoring.guru/design-patterns/prototype
  echo "\n\nDesign Pattern Prototype\n"

  var zaraClothing = newZara("Zara", "Jeans", 19.99, "Mafia Guy")
  echo zaraClothing.getClothingInfo()

  var zaraClothing2 = zaraClothing.copy()
  echo zaraClothing2.getClothingInfo()

  zaraClothing2.setModel("T-Shirt")
  zaraClothing2.setPrice(15.99)

  echo zaraClothing.getClothingInfo()
  echo zaraClothing2.getClothingInfo()
  
  ################################


  ################################
  # Design Pattern Singleton => more information in https://refactoring.guru/design-patterns/singleton
  echo "\n\nDesign Pattern Singleton\n"

  getInstance().addCar("VW Golf")
  getInstance().addCar("VW Passat")
  getInstance().addCar("Ford Fiesta")
  echo getInstance().getCars()
  
  ################################


  ################################
  # Design Pattern Adapter => more information in https://refactoring.guru/design-patterns/adapter
  echo "\n\nDesign Pattern Adapter\n"

  echo "Object Adapter\n"

  let government = newGovernment()
  var governmentImpl = newGovernmentAdapterImpl(government)
  var broker = newBroker[GovernmentAdapterImpl](governmentImpl)

  echo broker.sendGains(parseJson("""{"taxPayerId":1234567890, "taxPayerGains":10000}"""))


  echo "\nClass Adapter\n"

  let classAdapter = newClassAdapter()
  var broker2 = newBroker[ClassAdapter](classAdapter)

  echo broker2.sendGains(parseJson("""{"taxPayerId":1234562490, "taxPayerGains":156000}"""))
  
  ################################


  ################################
  # Design Pattern Bridge => more information in https://refactoring.guru/design-patterns/bridge
  echo "\n\nDesign Pattern Bridge\n"

  var tv = newTv()
  var radio = newRadio()

  var remoteRadio = newRemote[Radio](radio)
  echo "Radio"
  echo radio.getDeviceInfo()
  
  discard remoteRadio.volumeUp()
  discard remoteRadio.tooglePower()

  echo radio.getDeviceInfo()

  var remoteTv = newRemote[Tv](tv)
  echo "\nTV"
  echo tv.getDeviceInfo()

  discard remoteTv.volumeDown()
  discard remoteTv.channelUp()

  echo tv.getDeviceInfo()


  # var advancedRemoteTv = newAdvancedRemote(tv)
  
  ################################

when isMainModule:
  main()