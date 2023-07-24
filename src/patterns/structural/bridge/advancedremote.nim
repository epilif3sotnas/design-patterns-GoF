# internal
import
    patterns/structural/bridge/[
        device,
        remote
    ]


type
    AdvancedRemote* = ref object of Remote[TDevice: Device]


proc newAdvancedRemote*(device: TDevice): AdvancedRemote =
    return AdvancedRemote(device: device)

proc mute*(self: AdvancedRemote): bool =
    return self.device.setVolume(0)