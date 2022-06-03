#!/bin/bash

URL_PREFIX=https://github.com/filswan/multi-chain-storage/releases/download
BINARY_NAME=multi-chain-storage-2.0.0-linux-amd64
TAG_NAME=v2.0.0

wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/${BINARY_NAME}
wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/config.toml.example
wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/.env.example
wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/SwanPayment.json

CONF_FILE_DIR=${HOME}/.swan/mcs
mkdir -p ${CONF_FILE_DIR}

CONF_FILE_NAMES=(
    .env
    config.toml
    )

for CONF_FILE_NAME in ${CONF_FILE_NAMES[@]}; do
    CONF_FILE_PATH_SRC=./config/${CONF_FILE_NAME}.example
    CONF_FILE_PATH_DEST=${CONF_FILE_DIR_DEST}/${CONF_FILE_NAME}

    if [ -f "$CONF_FILE_PATH_DEST" ]; then
        echo "$CONF_FILE_PATH_DEST exists"
    else
        cp $CONF_FILE_PATH_SRC $CONF_FILE_PATH_DEST
        echo "copied $CONF_FILE_PATH_SRC to $CONF_FILE_PATH_DEST"
    fi
done

cp ./SwanPayment.json $CONF_FILE_DIR_DEST/SwanPayment.json

chmod +x ./${BINARY_NAME}
