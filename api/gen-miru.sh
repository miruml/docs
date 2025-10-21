#!/bin/bash
set -e

# use a virtual environment to run the python script
if [ ! -d ".venv" ]; then
    echo "Creating virtual environment..."
    python3 -m venv .venv
fi
. .venv/bin/activate

# ensure python3 and pyyaml are installed
if ! command -v python3 >/dev/null 2>&1; then
    echo "Python 3 is required but not installed"
    exit 1
fi
if ! python3 -c "import pyyaml" 2>/dev/null; then
    echo "Installing pyyaml..."
    if command -v pip3 >/dev/null 2>&1; then
        python3 -m pip install pyyaml types-pyyaml
    else
        echo "pip not found"
        exit 1
    fi
fi

# remove the webhooks section from the server-api.yaml file
python3 gen-miru.py