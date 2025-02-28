import miruconfigclient


def main():
    client = miruconfigclient.Client(
        ca_file="./path/to/ca.pem",
        cert_file="./path/to/cert.pem",
        key_file="./path/to/key.pem",
    )

    motion_control_config = client.get_config(
        config_name="Motion Control Config",
        git_commit_sha="a1b2c3d4e5f6g7h8i9j0",
        default_config=default_motion_control_config,
    )

    print(motion_control_config.speed)

if __name__ == "__main__":
    main()  
