#!/bin/bash

ROOT_DIR=$(pwd)

# Iterate over all JSON files in the hermes/keys/ directory
for file in "$ROOT_DIR/"*.key; do
  # Extract the filename without the path and extension
  filename=$(basename "$file")
  name="${filename%.key}"  # This removes the '.json' extension
  echo "Processing file: $file with key name: $name"

  # Run the hermes command for each file
  hermes --config ../config.toml keys add --chain orbit-1 --key-name "$name" --mnemonic-file "$file"
  hermes --config ../config.toml keys add --chain spn-1 --key-name "$name" --mnemonic-file "$file"
done

echo "All keys have been added."
