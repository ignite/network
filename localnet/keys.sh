#!/bin/bash

# This script imports all account keys from the specified YAML configuration file
# into Ignite accounts using the provided names and mnemonics.

# Path to the YAML configuration file
YAML_FILE="config.yml"

# Check if 'yq' is installed, which is required for parsing the YAML file
if ! command -v yq &> /dev/null
then
    echo "The 'yq' command-line tool is required to run this script."
    echo "You can install it with 'brew install yq' (on macOS) or see https://github.com/mikefarah/yq for other options."
    exit 1
fi

# Loop through each account entry in the YAML file
for index in $(yq eval '.accounts | keys' "$YAML_FILE" -o=json | jq '.[]'); do
  # Extract the account's name and mnemonic based on the current index
  NAME=$(yq eval ".accounts[$index].name" "$YAML_FILE")
  MNEMONIC=$(yq eval ".accounts[$index].mnemonic" "$YAML_FILE")

  # Run the import command for each account
  echo "Importing account: $NAME"
  ignite a import "$NAME" --secret "$MNEMONIC"
done
