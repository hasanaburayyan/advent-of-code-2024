#! /bin/bash

# Check that an argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <problem-name>"
    exit 1
fi


problem_name=$1
github_username="hasanaburayyan"
github_project="advent-of-code-2024"

# Create the problem directory
mkdir $problem_name

# Create the problem file
touch $problem_name/problem.md
echo "## Problem Details\n\nFill in the problem details here." >> $problem_name/problem.md
touch $problem_name/notes.md
echo "## Notes\n\nFill in the notes here." >> $problem_name/notes.md

# Create the Go Project
cd $problem_name
go mod init github.com/$github_username/$github_project/$problem_name

touch main.go
echo "package main\n" >> main.go

touch input.txt
