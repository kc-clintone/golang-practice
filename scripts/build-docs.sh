#!/bin/bash

# Remove the docs/ folder if it exists
if [ -d "../docs" ]; then
    rm -rf ../docs
fi

# Create a new docs/ folder
mkdir ../docs

# Copy all files from the root directory to docs/, except the content/ folder
shopt -s extglob
cd ..
cp -r !("contents"|"docs"|"scripts") docs/
cd scripts

echo "Documentation build complete"