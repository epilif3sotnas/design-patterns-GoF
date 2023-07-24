# internal
import
    patterns/structural/bridge/[
        device
    ]


type
    Remote*[TDevice: Device] = ref object of RootObj
        device: TDevice


proc newRemote*[TDevice: Device](device: TDevice): Remote[TDevice] =
    return Remote[TDevice](device: device)

proc tooglePower*(self: Remote): bool =
    return if self.device.isEnabled(): self.device.disable() else: self.device.enable()

proc volumeDown*(self: Remote): bool =
    return self.device.setVolume(self.device.getVolume() - 1)

proc volumeUp*(self: Remote): bool =
    return self.device.setVolume(self.device.getVolume() + 1)

proc channelDown*(self: Remote): bool =
    return self.device.setChannel(self.device.getChannel() - 1)

proc channelUp*(self: Remote): bool =
    return self.device.setChannel(self.device.getChannel() + 1)

proc getDevice*(self: Remote): Device =
    return self.device

proc setDevice*(self: Remote, device: Device) =
    self.device = device