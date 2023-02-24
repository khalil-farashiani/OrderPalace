#!/bin/bash

# Function to build the app
function build_app() {
  docker-compose build

  # Set variables
  program_name1="Sender"
  program_name2="Receiver"
  output_directory="./bin"

  # Create output directory if it doesn't exist
  # shellcheck disable=SC2164
  cd "./sender/cmd"
  # Build sender program
  go build -o "../../$output_directory/$program_name1" .

  # shellcheck disable=SC2164
  cd "../../receiver/cmd"
  # Build receiver program
  go build -o "../../$output_directory/$program_name2" .

  # Print success message
  echo "Build successful. Output files: $output_directory/$program_name1 and $output_directory/$program_name2"
}


build_app

