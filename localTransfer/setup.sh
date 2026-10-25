#!/bin/bash

echo "[*] Starting the setup process ..."

echo "[*] Copying executable to ~/.local/bin . . ."
mkdir -p ~/.local/bin
cp ./build/localTransfer ~/.local/bin/localTransfer

echo "[*] Adding executable to PATH . . ."
echo 'export PATH="$PATH:~/.local/bin"' >> ~/.bashrc

echo "[*] Reloading console . . ."
source ~/.bashrc

echo "[*] Finished! You can use localTransfer"