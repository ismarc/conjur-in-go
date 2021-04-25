#!/usr/bin/env bash

docker run --rm -e "POSTGRES_DB=conjur" -e "POSTGRES_HOST_AUTH_METHOD=trust" -p 5432:5432 --name postgres postgres