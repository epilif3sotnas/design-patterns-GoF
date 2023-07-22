# std
import
    std/[
        json
    ]


type
    GovernmentAdapter* = concept a
        a.sendGainsInformation(JsonNode) is bool