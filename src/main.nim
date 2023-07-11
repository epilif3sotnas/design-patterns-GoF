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
  # In Nim, this pattern does not work yet (2023-07-10), since concepts can't be saved in a type class
  
  var cat = newCat("Stewie", "Cat")
  var abstractFactory = newAbstactFactory()
  echo abstractFactory.getAnimalInfo(cat)

  # var abstractFactory = newAbstactFactory(cat)
  # echo abstractFactory.getAnimalInfo()
  ################################


  ################################
  # Desing Pattern Builder => more information in https://en.wikipedia.org/wiki/Builder_pattern
  # In Nim, this pattern does not work yet (2023-07-11), since concepts can't be saved in a type class

  let setup = newSetup()
  let admin = newAdmin()
  var car: Car = admin.build(setup)
  echo car.getCarInfo()

  # let admin = newAdmin(setup)
  # var car: Car = admin.create()
  # echo car.getCarInfo()
  ################################

when isMainModule:
  main()