type
    Company1* = ref object
        name: string
        iban: string


proc newCompany1*(name: string, iban: string): Company1 =
    return Company1(
        name: name,
        iban: iban
    )

proc sendMoney*(self: Company1, ibanToSend: string): bool =
    return true

proc getCompanyInfo*(self: Company1): string =
    return "Company Name: " & self.name & "\nIBAN: " & self.iban

proc getName*(self: Company1): string =
    return self.name

proc setName*(self: Company1, name: string) =
    self.name = name

proc getIban*(self: Company1): string =
    return self.iban

proc setIban*(self: Company1, iban: string) =
    self.iban = iban