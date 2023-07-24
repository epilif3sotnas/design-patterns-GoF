type
    Radio* = ref object
        isEnabled: bool
        volume: uint8
        channel: uint16
    

proc newRadio*(): Radio =
    return Radio(
        isEnabled: false,
        volume: 0,
        channel: 1
    )

proc isEnabled*(self: Radio): bool =
    return self.isEnabled

proc enable*(self: Radio): bool =
    self.isEnabled = true
    return true

proc disable*(self: Radio): bool =
    self.isEnabled = false
    return true

proc getVolume*(self: Radio): uint8 =
    return self.volume

proc setVolume*(self: Radio, volume: uint8): bool =
    self.volume = volume
    return true

proc getChannel*(self: Radio): uint16 =
    return self.channel

proc setChannel*(self: Radio, channel: uint16): bool =
    self.channel = channel
    return true

proc getDeviceInfo*(self: Radio): string =
    return "\nIs Enabled: " & $self.isEnabled &
            "\nVolume: " & $self.volume &
            "\nChannel: " & $self.channel