#! /bin/sh
IMAGE_NAME=fiagram_account_service
IMAGE_VERSION=$(cat ./VERSION)

docker build --target deployment -t $IMAGE_NAME:$IMAGE_VERSION .