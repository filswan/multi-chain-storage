#!/bin/bash

URL_PREFIX=https://github.com/filswan/multi-chain-storage/releases/download
BINARY_NAME=multi-chain-storage-2.0.0-linux-amd64
TAG_NAME=v2.0.0

wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/${BINARY_NAME}
wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/config.toml.example
wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/env.example
wget --no-check-certificate ${URL_PREFIX}/${TAG_NAME}/SwanPayment.json

CONF_FILE_DIR=${HOME}/.swan/mcs
mkdir -p ${CONF_FILE_DIR}

CONF_FILE_PATH=${CONF_FILE_DIR}/config.toml
if [ -f "${CONF_FILE_PATH}" ]; then
    echo "${CONF_FILE_PATH} exists"
else
    cp ./config.toml.example $CONF_FILE_PATH
    echo "${CONF_FILE_PATH} created"
fi

ENV_FILE_PATH=${CONF_FILE_DIR}/.env
if [ -f "${ENV_FILE_PATH}" ]; then
    echo "${ENV_FILE_PATH} exists"
else
    cp ./env.example $ENV_FILE_PATH
    echo "${ENV_FILE_PATH} created"
fi

cp ./SwanPayment.json $CONF_FILE_DIR_DEST/SwanPayment.json

chmod +x ./${BINARY_NAME}
