#!/bin/bash

if [ -z "$1" ]; then
    echo "missing app name argument!"
    echo "try with ./boilerplate.sh <your app name>"
    exit 1
fi

# Create new temp repo
mkdir -p ./$1
cd ./_services-templates/go-api/go-app
cp -r . ../../../$1

# replace sauron-app to args name
cd ../../../$1
find . -type f -exec sed -i '' s/go-app/$1/g {} +

echo "Your app was generated on folder $1"
echo "To start run: cd $1 && docker-compose up"