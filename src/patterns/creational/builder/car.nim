type
    Car* = ref object
        engine: string
        horsepower: uint16


proc newCar*(engine: string, horsepower: uint16): Car =
    return Car(
        engine: engine,
        horsepower: horsepower
    )

proc getCarInfo*(self: Car): string =
    return "Engine: " & self.engine & "\nHorsepower: " & $self.horsepower

proc getEngine*(self: Car): string =
    return self.engine

proc setEngine*(self: Car, engine: string) =
    self.engine = engine

proc getHorsepower*(self: Car): uint16 =
    return self.horsepower

proc setHorsepower*(self: Car, horsepower: uint16) =
    self.horsepower = horsepower