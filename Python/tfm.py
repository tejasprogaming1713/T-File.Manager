#!/usr/bin/env python3
import os
import sys
import shutil
import requests
import zipfile

# -----------------
# Functions
# -----------------

def list_files():
    """List all files and folders in current directory"""
    for f in os.listdir('.'):
        print(f)

def download(url, name=None):
    """Download a file from URL"""
    if not name:
        name = url.split("/")[-1]
    r = requests.get(url)
    open(name, "wb").write(r.content)
    print("Downloaded:", name)

def zip_folder(folder):
    """Zip a folder"""
    shutil.make_archive(folder, 'zip', folder)
    print("Zipped:", folder + ".zip")

def unzip(file):
    """Unzip a zip file"""
    with zipfile.ZipFile(file, 'r') as z:
        z.extractall(".")
    print("Unzipped:", file)

def delete(path):
    """Delete a file or folder"""
    if os.path.isdir(path):
        shutil.rmtree(path)
    else:
        os.remove(path)
    print("Deleted:", path)

# -----------------
# Main
# -----------------

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: tfm <command> [args]")
        sys.exit(1)

    cmd = sys.argv[1]

    if cmd == "list":
        list_files()
    elif cmd == "download":
        if len(sys.argv) < 3:
            print("Usage: tfm download <url>")
            sys.exit(1)
        download(sys.argv[2])
    elif cmd == "zip":
        if len(sys.argv) < 3:
            print("Usage: tfm zip <folder>")
            sys.exit(1)
        zip_folder(sys.argv[2])
    elif cmd == "unzip":
        if len(sys.argv) < 3:
            print("Usage: tfm unzip <file.zip>")
            sys.exit(1)
        unzip(sys.argv[2])
    elif cmd == "delete":
        if len(sys.argv) < 3:
            print("Usage: tfm delete <file/folder>")
            sys.exit(1)
        delete(sys.argv[2])
    else:
        print("Unknown command:", cmd)
        print("Available commands: list, download, zip, unzip, delete")
