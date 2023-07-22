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
    GovernmentAdapterImpl* = ref object
        government: Government


proc newGovernmentAdapterImpl*(government: Government): GovernmentAdapterImpl =
    return GovernmentAdapterImpl(
        government: government
    )

proc sendGainsInformation*(self: GovernmentAdapterImpl, data: JsonNode): bool =
    return self.government.sendGainsInformation((data["taxPayerId"].getInt().uint(), data["taxPayerGains"].getFloat()))