# internal
import
    patterns/structural/composite/[
        graphic
    ]



type
    Dot* = ref object of GraphicImpl
        x*: float
        y*: float


proc newDot*(x,y: float): Dot =
    return Dot(
        x: x,
        y: y
    )

proc move*(self: Dot, x,y: float): bool =
    echo "\n\nDot Moved\nFrom\nx: " & $self.x & "\ny: " & $self.y & "\nTo\nx: " & $x & "\ny: " & $y

    self.x = x
    self.y = y
    
    return true

proc draw*(self: Dot): bool =
    echo "\n\nDot Drawn\nx: " & $self.x & "\ny: " & $self.y
    
    return true

proc x*(self: Dot): float =
    return self.x

proc x*(self: Dot, x: float) =
    self.x = x

proc y*(self: Dot): float =
    return self.y

proc y*(self: Dot, y: float) =
    self.y = y