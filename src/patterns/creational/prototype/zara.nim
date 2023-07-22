# internal
import
    patterns/creational/prototype/[
        clothing
    ]


type
    Zara* = ref object of Clothing


proc newZara*(brand: string, model: string, price: float, owner: string): Zara =
    return Zara(
        brand: brand,
        model: model,
        price: price,
        owner: owner
    )

proc copy*(self: Zara): Zara =
    return Zara(
        brand: self.getBrand(),
        model: self.getModel(),
        price: self.getPrice(),
        owner: self.getOwner()
    )