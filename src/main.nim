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
  patterns/creational/singleton/[
    carstorage,
  ]


proc main() =
  ################################
  # Design Pattern Abstract Factory => more information in https://en.wikipedia.org/wiki/Abstract_factory_pattern
  
  var cat = newCat("Stewie", "Cat")
  var abstractFactory = newAbstactFactory[Cat](cat)
  echo abstractFactory.getAnimalInfo()

  ################################


  ################################
  # Design Pattern Builder => more information in https://en.wikipedia.org/wiki/Builder_pattern

  let setup = newSetup()
  let admin = newAdmin[Setup](setup)
  var car: Car = admin.build()
  echo car.getCarInfo()
  
  ################################


  ################################
  # Design Pattern Factory Method => more information in https://en.wikipedia.org/wiki/Factory_method_pattern

  let creator = newCreator()
  var company: Company1 = create[Company1](creator)
  echo company.getCompanyInfo()
  
  ################################


  ################################
  # Design Pattern Prototype => more information in https://en.wikipedia.org/wiki/Prototype_pattern

  ### CODE - WIP ###
  
  ################################


  ################################
  # Design Pattern Singleton => more information in https://en.wikipedia.org/wiki/Singleton_pattern

  getInstance().addCar("VW Golf")
  getInstance().addCar("VW Passat")
  getInstance().addCar("Ford Fiesta")
  echo getInstance().getCars()
  
  ################################

when isMainModule:
  main()