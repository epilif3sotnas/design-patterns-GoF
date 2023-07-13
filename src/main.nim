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
  ]


proc main() =
  ################################
  # Design Pattern Abstract Factory => more information in https://en.wikipedia.org/wiki/Abstract_factory_pattern
  
  var cat = newCat("Stewie", "Cat")
  var abstractFactory = newAbstactFactory[Cat](cat)
  echo abstractFactory.getAnimalInfo()

  ################################


  ################################
  # Desing Pattern Builder => more information in https://en.wikipedia.org/wiki/Builder_pattern

  let setup = newSetup()
  let admin = newAdmin[Setup](setup)
  var car: Car = admin.build()
  echo car.getCarInfo()
  
  ################################

when isMainModule:
  main()