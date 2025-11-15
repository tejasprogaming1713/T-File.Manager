#!/bin/bash

# -----------------
# T-File Manager (Bash)
# -----------------

cmd=$1

if [ -z "$cmd" ]; then
    echo "Usage: tfm <command> [args]"
    exit 1
fi

# -----------------
# List files/folders
# -----------------
if [ "$cmd" == "list" ]; then
    ls
fi

# -----------------
# Download file from URL
# -----------------
if [ "$cmd" == "download" ]; then
    url=$2
    if [ -z "$url" ]; then
        echo "Usage: tfm download <url>"
        exit 1
    fi
    name=$(basename "$url")
    curl -L "$url" -o "$name"
    echo "Downloaded: $name"
fi

# -----------------
# Zip a folder
# -----------------
if [ "$cmd" == "zip" ]; then
    folder=$2
    if [ -z "$folder" ]; then
        echo "Usage: tfm zip <folder>"
        exit 1
    fi
    zip -r "$folder.zip" "$folder"
    echo "Zipped: $folder.zip"
fi

# -----------------
# Unzip a file
# -----------------
if [ "$cmd" == "unzip" ]; then
    file=$2
    if [ -z "$file" ]; then
        echo "Usage: tfm unzip <file.zip>"
        exit 1
    fi
    unzip "$file"
    echo "Unzipped: $file"
fi

# -----------------
# Delete a file/folder
# -----------------
if [ "$cmd" == "delete" ]; then
    target=$2
    if [ -z "$target" ]; then
        echo "Usage: tfm delete <file/folder>"
        exit 1
    fi
    rm -rf "$target"
    echo "Deleted: $target"
fi
