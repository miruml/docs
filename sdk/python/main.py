from dataclasses import dataclass

from miru.client import Client, Config

def main():
    client: Client = Client.from_cert(
        ca_file="./path/to/ca.crt",
        cert_file="./path/to/client.crt",
        key_file="./path/to/client.key",
    )

    motion_control_config: Config[MotionControl] = client.init_config(
        slug="motion-control-config",
        schema_file="./path/to/development/file",
        data=MotionControl,
    )

    # Retrieve data from the config
    motion_control_config_data: MotionControl = motion_control_config.data()
    print("speed: ", motion_control_config_data.speed)
    print("spin enabled?: ", motion_control_config_data.features.spin)
    print("accelerometer offsets: ", motion_control_config_data.accelerometer.offsets)

    # Update the config data, this data is locally persisted until the next time the
    # data can be sent to the cloud.
    motion_control_config_data.features.spin = False
    motion_control_config_data.speed = 20
    motion_control_config.set(motion_control_config_data)


if __name__ == "__main__":
    main()  

# This the config structure generated by the JSON Schema definition defined for your
# config schema
@dataclass
class MotionControlFeatures:
    jump: bool
    spin: bool
    backflip: bool

@dataclass
class Offsets:
    x: float
    y: float
    z: float

@dataclass
class ScalingFactor:
    x: float
    y: float
    z: float

@dataclass
class Accelerometer:
    id: str
    offsets: Offsets
    scaling_factor: ScalingFactor

@dataclass
class MotionControl:
    speed: int
    features: MotionControlFeatures
    accelerometer: Accelerometer
# The motion control config has the following structure:
# {
#     "speed": 10,
#     "features": {
#         "spin": true,
#         "rotate": true,
#         "move": true
#     },
#     "accelerometer": {
#         "id": "accelerometer-1",
#         "offsets": {
#             "x": 0.0,
#             "y": 0.0,
#             "z": 0.0
#         },
#         "scaling_factor": {
#             "x": 1.0,
#             "y": 1.0,
#             "z": 1.0
#         }
# }