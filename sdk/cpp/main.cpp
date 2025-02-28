// sdk/cpp/src/production.cpp
#include <iostream>
#include <string>
#include "miru/miru.h"

int main() {
    Client client = Client::from_cert("./path/to/ca.crt", "./path/to/client.crt", "./path/to/client.key");

    MotionControl motion_control_data = {
        10,  // speed
        { // Features
            true, // spin
            true, // jump
            true, // backflip
        },  
        { // Accelerometer
            "accelerometer-1", // id
            { // Offsets
                0.0f, // x
                0.0f, // y
                0.0f // z
            },
            { // ScalingFactor
                1.0f, // x
                1.0f, // y
                1.0f  // z
            },
        }  
    };

    Config<MotionControl> motion_control_config = client.init_config(
        "Motion Control Config",
        "./path/to/development/file",
        motion_control_data,
        true,
        "./path/to/development/file"
    );

    // Retrieve the config data
    MotionControl motion_control_config_data = motion_control_config.data();
    std::cout << "speed: " << motion_control_config_data.speed << std::endl;
    std::cout << "spin enabled?: " << motion_control_config_data.features.spin << std::endl;
    std::cout << "accelerometer offsets: "
              << motion_control_config_data.accelerometer.offsets.x << ", "
              << motion_control_config_data.accelerometer.offsets.y << ", "
              << motion_control_config_data.accelerometer.offsets.z << std::endl;

    // Update the config data
    motion_control_config_data.features.spin = false;
    motion_control_config_data.speed = 20;
    motion_control_config.set(motion_control_config_data);

    return 0;
}

// Define the MotionControlFeatures struct
struct MotionControlFeatures {
    bool jump;
    bool spin;
    bool backflip;
};

// Define the Offsets struct
struct Offsets {
    float x;
    float y;
    float z;
};

// Define the ScalingFactor struct
struct ScalingFactor {
    float x;
    float y;
    float z;
};

// Define the Accelerometer struct
struct Accelerometer {
    std::string id;
    Offsets offsets;
    ScalingFactor scaling_factor;
};

// Define the MotionControl struct
struct MotionControl {
    int speed;
    MotionControlFeatures features;
    Accelerometer accelerometer;
};