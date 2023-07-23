# std
import
    std/[
        json
    ]


type
    GovernmentAdapter* = concept a
        a.sendGainsInfo(JsonNode) is bool