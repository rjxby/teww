#!/bin/bash

echo "build client started..."
cd ./teww-client
npm install >/dev/null
echo "client dependencies successfully installed..."
npm run build:prod
echo "client successfully builded..."
cd ../
echo "configuration copy..."
tee ./teww-auth/config.json ./teww-backend/config.json ./teww-client/config.json < ./config.json >/dev/null
echo "build compose file..."
docker-compose build
echo "build was successfully..."
