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


def main() -> None:
    # Get repository root and set up paths
    repo_root: Path = Path(
        run_command("git rev-parse --show-toplevel").strip()
    )
    api_dir: Path = repo_root / "api"
    openapi_dir: Path = api_dir / "openapi"
    backend_server_dir: Path = openapi_dir / "configs" / "backend-server"
    public_server_dir: Path = backend_server_dir / "public"
    public_openapi_spec: Path = public_server_dir / "openapi.gen.yaml"

    target_api_file: Path = api_dir / "api.yaml"

    # delete the target api file if it exists
    if target_api_file.exists():
        target_api_file.unlink()
    shutil.copy(public_openapi_spec, target_api_file)


if __name__ == "__main__":
    main()
