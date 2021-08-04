#!/bin/sh

docker buildx build -t mainnet/smartbch --platform=linux/arm64/v8 --load .
docker buildx build -t mainnet/smartbch --platform=linux/amd64 --load .

docker buildx build -t mainnet/smartbch --platform=linux/arm64/v8,linux/amd64 --push .
