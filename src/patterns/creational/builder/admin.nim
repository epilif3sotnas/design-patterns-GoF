# internal
import
    patterns/creational/builder/[
        builder,
        car
    ]


type
    Admin*[TBuilder: Builder] = ref object
        builder: TBuilder


proc newAdmin*[T: Builder](builder: T): Admin[T] =
    return Admin[T](builder: builder)

proc build*(self: Admin): Car =
    return self.builder.buildCar()