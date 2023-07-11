# internal
import
    patterns/creational/builder/[
        car
    ]


type
    Setup* = ref object


proc newSetup*(): Setup =
    return Setup()

proc buildCar*(self: Setup): Car =
    return newCar("Toyota", 184)