#!/bin/bash
# Usage: ./push.sh "commit message"

if [ -z "$1" ]; then
    echo " Error: Missing commit message"
    echo "Usage: ./push.sh \"feat: add solution for X\""
    exit 1
fi

echo "Pulling latest changes from GitHub (Rebase)"
git pull --rebase origin main

echo "Adding files..."
git add .

echo "Committing..."
git commit -m "$1"

echo "Pushing..."
git push

echo "Done!"