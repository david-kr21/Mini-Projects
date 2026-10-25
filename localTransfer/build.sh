#!/bin/bash

echo "[*] Deleting old build files ..."
rm -rf ./build/*

echo "[*] Building new build files ..."
go build -o ./build/ .

echo "[*] Build finished!"