#!/usr/bin/env bash

if [[ "X$1" == "X" ]]; then
    echo "commit msg should not be empty!"
    exit 1
fi

changes=$(git status)
date_info=$(date)

git status

git add .
git commit -m "update at: ${date_info}, [$1]"
git push origin master

