type
    Tv* = ref object
        isEnabled: bool
        volume: uint8
        channel: uint16
    

proc newTv*(): Tv =
    return Tv(
        isEnabled: false,
        volume: 0,
        channel: 1
    )

proc isEnabled*(self: Tv): bool =
    return self.isEnabled

proc enable*(self: Tv): bool =
    self.isEnabled = true
    return true

proc disable*(self: Tv): bool =
    self.isEnabled = false
    return true

proc getVolume*(self: Tv): uint8 =
    return self.volume

proc setVolume*(self: Tv, volume: uint8): bool =
    self.volume = volume
    return true

proc getChannel*(self: Tv): uint16 =
    return self.channel

proc setChannel*(self: Tv, channel: uint16): bool =
    self.channel = channel
    return true

proc getDeviceInfo*(self: Tv): string =
    return "\nIs Enabled: " & $self.isEnabled &
            "\nVolume: " & $self.volume &
            "\nChannel: " & $self.channel