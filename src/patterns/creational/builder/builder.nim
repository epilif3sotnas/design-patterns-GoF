# internal
import
    patterns/creational/builder/[
        car
    ]


type
    Builder* = concept a
        a.buildCar() is Car