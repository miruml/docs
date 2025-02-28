import miruconfigclient


def main():
    client = miruconfigclient.Client(
        sdk_key="<your-sdk-key>",
        client_id="<your-client-id>",
    )

    motion_control_config = client.get_config(
        config_name="Motion Control Config",
        development_file="./path/to/development/file"
        development=true,
    )

    print(motion_control_config.speed)

if __name__ == "__main__":
    main()  
