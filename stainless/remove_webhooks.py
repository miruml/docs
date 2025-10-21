import yaml
import sys
import os


def remove_webhooks_from_yaml(file_path):
    """Remove x-webhooks section from OpenAPI YAML file."""

    # Read the YAML file
    with open(file_path, 'r', encoding='utf-8') as file:
        data = yaml.safe_load(file)

    # Check if x-webhooks exists and remove it
    if 'webhooks' in data:
        print(f"Found webhooks section, removing it...")
        del data['webhooks']
        print("✅ Successfully removed webhooks section")
    else:
        print("ℹ️  No webhooks section found in the file")

    # Write the modified YAML back to file
    with open(file_path, 'w', encoding='utf-8') as file:
        yaml.dump(
            data, file, default_flow_style=False,
            sort_keys=False, allow_unicode=True,
        )

    print(f"✅ Updated file: {file_path}")


def main():
    if len(sys.argv) != 2:
        print("Usage: python remove_webhooks.py <yaml_file_path>")
        sys.exit(1)

    file_path = sys.argv[1]

    if not os.path.exists(file_path):
        print(f"❌ Error: File {file_path} does not exist")
        sys.exit(1)

    try:
        remove_webhooks_from_yaml(file_path)
    except yaml.YAMLError as e:
        print(f"❌ Error parsing YAML: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"❌ Error: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()
