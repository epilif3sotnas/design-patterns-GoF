type
    Graphic* = concept a
        a.move(float, float) is bool
        a.draw() is bool

    GraphicImpl* = ref object of RootObj


proc move*(self: GraphicImpl, x,y: float): bool =
    raise newException(Exception, "Method not implemented")


proc draw*(self: GraphicImpl): bool =
    raise newException(Exception, "Method not implemented")