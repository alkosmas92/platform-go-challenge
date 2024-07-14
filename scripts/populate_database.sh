#!/bin/bash

# Database file
DB_FILE="./favorites.db"

# Function to generate a random suffix
generate_random_suffix() {
  echo $RANDOM
}

# Function to insert a favorite into the database
insert_favorite() {
  USER_ID=$1
  ASSET_ID=$2
  ASSET_TYPE=$3
  DESCRIPTION="Description for $ASSET_TYPE $4"

  sqlite3 $DB_FILE "INSERT INTO favorites (user_id, asset_id, asset_type, description) VALUES ('$USER_ID', '$ASSET_ID', '$ASSET_TYPE', '$DESCRIPTION');"
}

# Register a new user
register_user() {
  USERNAME="user$(generate_random_suffix)"
  RESPONSE=$(curl -s -X POST http://localhost:8080/register \
    -H "Content-Type: application/json" \
    -d "{
          \"username\": \"$USERNAME\",
          \"password\": \"password1\",
          \"firstname\": \"First\",
          \"lastname\": \"Last\"
        }")
  echo $USERNAME
}

# Login to get JWT token and user ID
login_user() {
  USERNAME=$1
  RESPONSE=$(curl -s -X POST http://localhost:8080/login \
    -H "Content-Type: application/json" \
    -d "{
          \"username\": \"$USERNAME\",
          \"password\": \"password1\"
        }")
  TOKEN=$(echo $RESPONSE | jq -r '.token')
  USER_ID=$(echo $RESPONSE | jq -r '.user_id')
  echo $TOKEN $USER_ID
}

# Create a user and login to get the token and user ID
USERNAME=$(register_user)
read TOKEN USER_ID <<<$(login_user $USERNAME)

# Verify token and user ID
if [ -z "$TOKEN" ] || [ -z "$USER_ID" ]; then
  echo "Failed to obtain token or user ID."
  exit 1
fi

echo "JWT Token: $TOKEN"
echo "User ID: $USER_ID"

# Check if the database file exists
if [ ! -f $DB_FILE ]; then
  echo "Database file not found!"
  exit 1
fi

# Insert 10 favorite entries for the user
for i in $(seq 1 10); do
  ASSET_ID="asset_$i"
  ASSET_TYPE="chart"
  if [ $((i % 3)) -eq 1 ]; then
    ASSET_TYPE="insight"
  elif [ $((i % 3)) -eq 2 ]; then
    ASSET_TYPE="audience"
  fi

  insert_favorite $USER_ID $ASSET_ID $ASSET_TYPE $i
done

echo "Created 10 favorites for user $USER_ID."

# Save token and user ID to a file for use in the example requests script
echo "TOKEN=\"$TOKEN\"" >> auth_info.sh
echo "USER_ID=\"$USER_ID\"" >> auth_info.sh
echo "USERNAME=\"$USERNAME\"" >> auth_info.sh
