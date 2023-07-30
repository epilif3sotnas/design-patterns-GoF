# internal
import
    patterns/structural/composite/[
        circle,
        graphic,
        dot
    ]


type
    CoumpoundGraphic* = ref object
        children: seq[GraphicImpl]


proc newCoumpoundGraphic*(): CoumpoundGraphic =
    return CoumpoundGraphic()

proc add*(self: CoumpoundGraphic, child: Graphic) =
    self.children.add(child)

proc move*(self: CoumpoundGraphic, x,y: float): bool =
    for child in self.children:
        discard child.move(x,y)

proc draw*(self: CoumpoundGraphic): bool =
    for child in self.children:
        discard child.draw()