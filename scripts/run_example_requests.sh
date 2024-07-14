#!/bin/bash

# Load the token and user ID from the file
source ./auth_info.sh

# Function to generate a random suffix for asset_id
generate_asset_id() {
  echo "asset_$1"
}

# Check if sufficient arguments are provided
if [ $# -lt 1 ]; then
  echo "Usage: $0 <GET|UPDATE|DELETE|CREATE> [asset_id]"
  exit 1
fi

COMMAND=$1
ASSET_ID=$2

case $COMMAND in
  CREATE)
    # POST - Create a new favorite
    ASSET_ID=$(generate_asset_id $RANDOM)
    curl -X POST http://localhost:8080/favorites \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json" \
      -d "{
            \"asset_id\": \"$ASSET_ID\",
            \"asset_type\": \"chart\",
            \"description\": \"A sample chart\"
          }"
    ;;
  GET)
    # GET - Retrieve favorites with limit and offset
    curl -X GET "http://localhost:8080/favorites?limit=10&offset=0" \
      -H "Authorization: Bearer $TOKEN"
    ;;
  UPDATE)
    # Check if asset_id is provided
    if [ -z "$ASSET_ID" ]; then
      echo "Usage: $0 UPDATE <asset_id>"
      exit 1
    fi

    # PUT - Update the favorite
    curl -X PUT "http://localhost:8080/favorites?asset_id=$ASSET_ID" \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json" \
      -d "{
            \"description\": \"Updated description for the sample chart\"
          }"
    ;;
  DELETE)
    # Check if asset_id is provided
    if [ -z "$ASSET_ID" ]; then
      echo "Usage: $0 DELETE <asset_id>"
      exit 1
    fi

    # DELETE - Delete the favorite
    curl -X DELETE "http://localhost:8080/favorites?asset_id=$ASSET_ID" \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json"
    ;;
  *)
    echo "Invalid command. Usage: $0 <GET|UPDATE|DELETE|CREATE> [asset_id]"
    exit 1
    ;;
esac
