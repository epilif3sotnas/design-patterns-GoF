type
    Device* = concept a
        a.isEnabled() is bool
        a.enable() is bool
        a.disable() is bool
        a.getVolume() is uint8
        a.setVolume(uint8) is bool
        a.getChannel() is uint16
        a.setChannel(uint16) is bool
        a.getDeviceInfo() is string