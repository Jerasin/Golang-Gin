#!/bin/bash

set -e  # ถ้าคำสั่งไหนล้มเหลว ให้หยุดทันที
set -x

echo "Changing to the project directory..."
cd ~/workspace/Golang-Gin || exit 1  # ถ้า cd ไม่สำเร็จ ให้หยุด script

echo "Pulling latest changes from Git..."
git pull || exit 1  # ถ้า git pull ไม่สำเร็จ ให้หยุด

echo "Stopping and removing containers..."
docker-compose down || exit 1  # ถ้า docker-compose down ไม่สำเร็จ ให้หยุด

echo "Building and starting containers..."
docker-compose up --build -d || exit 1  # ถ้า docker-compose up ไม่สำเร็จ ให้หยุด
