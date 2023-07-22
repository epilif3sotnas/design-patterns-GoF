type
    Government* = ref object


proc newGovernment*(): Government =
    return Government()

proc sendGainsInformation*(self: Government, data: tuple[taxPayerId: uint, taxPayerGains: float]): bool =
    echo "Data successfully received"
    echo "Tax Payer ID: " & $data.taxPayerId & "\nTax Payer Gains: " & $data.taxPayerGains

    return true