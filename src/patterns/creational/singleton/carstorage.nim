type
    CarStorage* = ref object
        cars: seq[string]


proc newCarStorage(): CarStorage =
    return CarStorage()


var instance = newCarStorage()


proc getInstance*(): CarStorage =
    return instance


proc getCars*(self: CarStorage): seq[string] =
    return self.cars

proc addCar*(self: CarStorage, car: string) =
    self.cars.add(car)