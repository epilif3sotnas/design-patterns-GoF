# internal
import
    patterns/creational/factorymethod/[
        company,
        company1
    ]


type
    Creator* = ref object


proc newCreator*(): Creator =
    return Creator()

proc create*[TCompany: Company](self: Creator): TCompany =
    return newCompany1("Klarna", "SE4550000000058398257466")