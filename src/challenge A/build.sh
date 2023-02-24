#!/bin/bash

# Function to build the app
function build_app() {
  docker-compose build

  # Set variables
  program_name="Sender"
  output_directory="./bin"

  # Create output directory if it doesn't exist
  # shellcheck disable=SC2164
  cd "./sender/cmd"

  # Build program
  go build -o "../../$output_directory/$program_name" .

  # Print success message
  echo "Build successful. Output file: $output_directory/$program_name"
}


build_app

