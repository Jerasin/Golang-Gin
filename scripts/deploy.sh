#!/bin/bash

ls -la

git pull

docker-compose down

docker-compose up --build

