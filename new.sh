#!/bin/bash

# 1. Validation: Ensure ID and Name are provided
if [ -z "$1" ] || [ -z "$2" ]; then
  echo "❌ Error: Missing arguments."
  echo "Usage: ./new.sh <id> \"<problem_name>\" [extension]"
  echo "Example: ./new.sh 1 \"Two Sum\" js"
  exit 1
fi

# 2. Variables Setup
ID=$(printf "%04d" $1) # Pads number with zeros (1 -> 0001)
RAW_NAME="$2"
# Converts "Two Sum" -> "two-sum" (lowercase + hyphens)
SLUG_NAME=$(echo "$RAW_NAME" | tr '[:upper:]' '[:lower:]' | tr -s ' ' '-') 
FOLDER_NAME="${ID}-${SLUG_NAME}"
EXT="${3:-js}"

if [ -d "$FOLDER_NAME" ]; then
  echo "⚠️  Folder '$FOLDER_NAME' already exists."
else
  mkdir -p "$FOLDER_NAME"
  echo "bh Create folder: $FOLDER_NAME"
fi

# 4. Create README.md Template
README_PATH="$FOLDER_NAME/README.md"
if [ ! -f "$README_PATH" ]; then
cat <<EOF > "$README_PATH"
# $1. $RAW_NAME

**Difficulty:** Medium/Easy/Hard

## Description
## Constraints
* ## Approach
EOF
fi

SOL_PATH="$FOLDER_NAME/Solution.$EXT"
if [ ! -f "$SOL_PATH" ]; then
  if [ "$EXT" == "js" ]; then
cat <<EOF > "$SOL_PATH"

// Time Complexity: 
// Space Complexity: 

EOF
  else
    touch "$SOL_PATH"
  fi
  echo "✅ Created $SOL_PATH"
else
  echo "⚠️  File $SOL_PATH already exists. Skipping."
fi