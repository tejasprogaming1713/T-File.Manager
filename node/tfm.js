#!/usr/bin/env node

const fs = require("fs");
const https = require("https");
const { execSync } = require("child_process");

// -----------------
// Helper Functions
// -----------------

function listFiles() {
  fs.readdirSync(".").forEach(f => console.log(f));
}

function download(url, name) {
  if (!name) name = url.split("/").pop();
  const file = fs.createWriteStream(name);
  https.get(url, res => {
    res.pipe(file);
    file.on("finish", () => console.log("Downloaded:", name));
  });
}

function zipFolder(folder) {
  try {
    execSync(`zip -r "${folder}.zip" "${folder}"`);
    console.log("Zipped:", folder + ".zip");
  } catch (err) {
    console.error("Error zipping folder:", err.message);
  }
}

function unzipFile(file) {
  try {
    execSync(`unzip "${file}"`);
    console.log("Unzipped:", file);
  } catch (err) {
    console.error("Error unzipping file:", err.message);
  }
}

function deleteFileOrFolder(target) {
  try {
    fs.rmSync(target, { recursive: true, force: true });
    console.log("Deleted:", target);
  } catch (err) {
    console.error("Error deleting:", err.message);
  }
}

// -----------------
// Main
// -----------------

const args = process.argv.slice(2);
if (args.length < 1) {
  console.log("Usage: tfm <command> [args]");
  process.exit(1);
}

const cmd = args[0];

switch (cmd) {
  case "list":
    listFiles();
    break;
  case "download":
    if (!args[1]) {
      console.log("Usage: tfm download <url>");
      process.exit(1);
    }
    download(args[1]);
    break;
  case "zip":
    if (!args[1]) {
      console.log("Usage: tfm zip <folder>");
      process.exit(1);
    }
    zipFolder(args[1]);
    break;
  case "unzip":
    if (!args[1]) {
      console.log("Usage: tfm unzip <file.zip>");
      process.exit(1);
    }
    unzipFile(args[1]);
    break;
  case "delete":
    if (!args[1]) {
      console.log("Usage: tfm delete <file/folder>");
      process.exit(1);
    }
    deleteFileOrFolder(args[1]);
    break;
  default:
    console.log("Unknown command:", cmd);
    console.log("Available commands: list, download, zip, unzip, delete");
}
