#!/usr/bin/env python3
"""
OpenAPI Model Regeneration Script

This script regenerates OpenAPI models and copies them to the appropriate
directories in the internal packages.
"""

import shutil
import subprocess
import sys
from pathlib import Path
from typing import Optional


def run_command(cmd: str, cwd: Optional[Path] = None) -> str:
    """Run a shell command and return the result."""
    try:
        result = subprocess.run(
            cmd, shell=True, cwd=cwd, check=True,
            capture_output=True, text=True
        )
        return result.stdout
    except subprocess.CalledProcessError as e:
        print(f"Error running command: {cmd}")
        print(f"Error: {e.stderr}")
        sys.exit(1)


def run_command_stream(cmd: str, cwd: Optional[Path] = None) -> None:
    """Run a shell command and stream output to stdout."""
    try:
        subprocess.run(
            cmd, shell=True, cwd=cwd, check=True,
            text=True
        )
    except subprocess.CalledProcessError as e:
        print(f"Error running command: {cmd}")
        print(f"Error: {e.stderr}")
        sys.exit(1)


def regen_openapi_specs(openapi_dir: Path) -> None:
    """Regenerate the api.yaml file."""
    configs_dir: Path = openapi_dir / "configs"
    backend_server_dir: Path = configs_dir / "backend-server"
    webhooks_dir: Path = configs_dir / "webhooks"
    agent_server_dir: Path = configs_dir / "agent-server"

    # regenerate the backend server api spec
    run_command_stream(f"cd {backend_server_dir} && make bundle-public")
    run_command_stream(f"cd {webhooks_dir} && make bundle-webhooks")
    run_command_stream(f"cd {agent_server_dir} && make bundle-all")


def refresh_server_api_yaml(api_dir: Path, openapi_dir: Path) -> None:
    """Refresh the api.yaml file."""
    backend_server_dir: Path = openapi_dir / "configs" / "backend-server"
    public_server_dir: Path = backend_server_dir / "public"
    public_openapi_spec: Path = public_server_dir / "openapi.gen.yaml"

    target_api_file: Path = api_dir / "server-api.yaml"

    # delete the target api file if it exists
    if target_api_file.exists():
        target_api_file.unlink()
    shutil.copy(public_openapi_spec, target_api_file)


def refresh_webhooks_yaml(api_dir: Path, openapi_dir: Path) -> None:
    """Refresh the webhooks.yaml file."""
    webhooks_dir: Path = openapi_dir / "configs" / "webhooks"
    webhooks_file: Path = webhooks_dir / "openapi.gen.yaml"

    target_webhooks_file: Path = api_dir / "webhooks.yaml"

    # delete the target webhooks file if it exists
    if target_webhooks_file.exists():
        target_webhooks_file.unlink()
    shutil.copy(webhooks_file, target_webhooks_file)


def refresh_agent_api_yaml(api_dir: Path, openapi_dir: Path) -> None:
    """Refresh the api.yaml file."""
    agent_server_dir: Path = openapi_dir / "configs" / "agent-server"
    public_openapi_spec: Path = agent_server_dir / "openapi.gen.yaml"

    target_api_file: Path = api_dir / "agent-api.yaml"

    # delete the target api file if it exists
    if target_api_file.exists():
        target_api_file.unlink()
    shutil.copy(public_openapi_spec, target_api_file)


def main() -> None:
    # Get repository root and set up paths
    repo_root: Path = Path(
        run_command("git rev-parse --show-toplevel").strip()
    )
    api_dir: Path = repo_root / "api"
    openapi_dir: Path = api_dir / "openapi"

    # ensure the openapi specs are up to date
    regen_openapi_specs(openapi_dir)

    # refresh the yaml files in the api directory
    refresh_server_api_yaml(api_dir, openapi_dir)
    refresh_agent_api_yaml(api_dir, openapi_dir)
    refresh_webhooks_yaml(api_dir, openapi_dir)


if __name__ == "__main__":
    main()
