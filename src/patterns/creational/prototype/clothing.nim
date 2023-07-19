type
    Clothing* = ref object of RootObj
        brand*: string
        model*: string
        price*: float
        owner*: string


proc getClothingInfo*(self: Clothing): string =
    return "Brand: " & self.brand & "\nModel: " & self.model & "\nPrice: " & $self.price & "\nOwner: " & self.owner

proc copy*(self: Clothing): Clothing =
    raise newException(Exception, "Method `copy` not implemented")

proc getBrand*(self: Clothing): string =
    return self.brand

proc setBrand*(self: Clothing, brand: string): void =
    self.brand = brand

proc getModel*(self: Clothing): string =
    return self.model

proc setModel*(self: Clothing, model: string): void =
    self.model = model

proc getPrice*(self: Clothing): float =
    return self.price

proc setPrice*(self: Clothing, price: float): void =
    self.price = price

proc getOwner*(self: Clothing): string =
    return self.owner

proc setOwner*(self: Clothing, owner: string): void =
    self.owner = owner