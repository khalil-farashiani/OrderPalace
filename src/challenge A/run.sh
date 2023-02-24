#!/bin/bash

# Function to run the app
function run_app() {

  # Read the .env file line by line and set environment variables
  while read -r line; do
      # shellcheck disable=SC2163
      export "$line"
  done < .env

  echo "$REDIS_DSN"
  # Check if environment variables are set for username and password, and use default values if not set
  if [[ -z "${DB_USERNAME}" ]]; then
    DB_USERNAME="default_username"
  fi

  if [[ -z "${DB_PASSWORD}" ]]; then
    DB_PASSWORD="default_password"
  fi

  # Set the environment variables for username and password in the docker-compose.yml file
  export DB_USERNAME="${DB_USERNAME}"
  export DB_PASSWORD="${DB_PASSWORD}"

  # Start the containers using docker-compose
  docker-compose up -d

  # shellcheck disable=SC2164
  cd ./bin/
  ./Sender 8000 &
  ./Receiver 8080
}

run_app
