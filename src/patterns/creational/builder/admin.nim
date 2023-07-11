# internal
import
    patterns/creational/builder/[
        builder,
        car
    ]


type
    Admin* = ref object

# type
#     Admin* = ref object
#         builder: Buidler


proc newAdmin*(): Admin =
    return Admin()

# proc newAdmin*(builder: Builder): Admin =
#     return Admin(builder: Builder)

proc build*(self: Admin, builder: Builder): Car =
    return builder.buildCar()

# proc create*(self: Admin): Car =
#     return self.builder.buildCar()