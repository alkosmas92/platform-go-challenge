#!/bin/bash

# Define the database file
DB_FILE="favorites.db"

# Remove the existing database file if it exists
if [ -f $DB_FILE ]; then
  rm $DB_FILE
fi

# Create the SQLite database and the favorites table
sqlite3 $DB_FILE <<EOF
CREATE TABLE IF NOT EXISTS favorites (
    user_id TEXT,
    asset_id TEXT,
    asset_type TEXT,
    description TEXT,
    PRIMARY KEY (user_id, asset_id)
);
EOF

# Insert 1000 records into the favorites table
for i in {1..300}; do
  USER_ID="user$i"
  ASSET_ID="asset$i"
  if (( i % 3 == 1 )); then
    ASSET_TYPE="chart"
    DESCRIPTION="Favorite chart asset $i"
  elif (( i % 3 == 2 )); then
    ASSET_TYPE="insight"
    DESCRIPTION="Favorite insight asset $i"
  else
    ASSET_TYPE="audience"
    DESCRIPTION="Favorite audience asset $i"
  fi

  sqlite3 $DB_FILE <<EOF
INSERT INTO favorites (user_id, asset_id, asset_type, description)
VALUES ('$USER_ID', '$ASSET_ID', '$ASSET_TYPE', '$DESCRIPTION');
EOF
done

echo "Database initialization complete."
