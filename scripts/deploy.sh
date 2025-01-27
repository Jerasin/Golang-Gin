#!/bin/bash

cd ~/workspace/Golang-Gin

git pull

docker-compose down

docker-compose up --build

