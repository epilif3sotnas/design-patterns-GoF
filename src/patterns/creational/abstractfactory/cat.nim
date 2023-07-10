type
    Cat* = ref object
        name: string
        specieName: string


proc newCat*(name: string, specieName: string): Cat =
    return Cat(
        name: name,
        specieName: specieName
    )

proc getName*(self: Cat): string =
    return self.name

proc setName*(self: Cat, name: string) =
    self.name = name

proc getSpecieName*(self: Cat): string =
    return self.specieName

proc setSpecieName*(self: Cat, specieName: string) =
    self.specieName = specieName