# internal
import
    patterns/structural/bridge/[
        device,
        remote
    ]


type
    AdvancedRemote*[TDevice: Device] = ref object of Remote[TDevice]


proc newAdvancedRemote*[TDevice: Device](device: TDevice): AdvancedRemote[TDevice] =
    return AdvancedRemote[TDevice](device: device)

proc mute*(self: AdvancedRemote): bool =
    return self.device.setVolume(0)