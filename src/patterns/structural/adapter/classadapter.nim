# std
import
    std/[
        json
    ]

# internal
import
    patterns/structural/adapter/[
        government
    ]


type
    ClassAdapter* = ref object of Government


proc newClassAdapter*(): ClassAdapter =
    return ClassAdapter()

proc sendGainsInfo*(self: ClassAdapter, data: JsonNode): bool =
    return self.sendGainsInformation((data["taxPayerId"].getInt().uint(), data["taxPayerGains"].getFloat()))