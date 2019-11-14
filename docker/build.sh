#!/bin/bash
docker build -t shunf4/filestash:latest -f Dockerfile .
docker-compose up -d
