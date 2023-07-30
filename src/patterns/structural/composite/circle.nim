# internal
import
    patterns/structural/composite/[
        dot
    ]


type
    Circle* = ref object of Dot
        radius: float


proc newCircle*(x,y,radius: float): Circle =
    return Circle(
        x: x,
        y: y,
        radius: radius
    )

proc draw*(self: Circle): bool =
    echo "\n\nCircle Drawn\nx: " & $self.x & "\ny: " & $self.y & "\nradius: " & $self.radius
    
    return true

proc move*(self: Circle, x,y: float): bool =
    echo "\n\nCircle Moved\nFrom\nx: " & $self.x & "\ny: " & $self.y & "\nTo\nx: " & $x & "\ny: " & $y

    self.x = x
    self.y = y
    
    return true

proc radius*(self: Circle): float =
    return self.radius

proc radius*(self: Circle, radius: float) =
    self.radius = radius