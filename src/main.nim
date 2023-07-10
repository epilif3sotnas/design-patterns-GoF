# internal
import
  patterns/creational/abstractfactory/[
    abstractfactory,
    cat
  ]


proc main() =
  # Design Pattern Abstract Factory => more information in https://en.wikipedia.org/wiki/Abstract_factory_pattern
  # In Nim this pattern does not work yet (2023-07-10), since concepts can't be saved in a type class
  var cat = newCat("Stewie", "Cat")
  var abstractFactory = newAbstactFactory()
  echo abstractFactory.getAnimalInfo(cat)

  # var abstractFactory = newAbstactFactory(cat)
  # echo abstractFactory.getAnimalInfo()

when isMainModule:
  main()