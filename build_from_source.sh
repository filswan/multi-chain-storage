#!/bin/bash

CONF_FILE_DIR=${HOME}/.swan/mcs
mkdir -p ${CONF_FILE_DIR}

CONF_FILE_PATH=${CONF_FILE_DIR}/config.toml
if [ -f "${CONF_FILE_PATH}" ]; then
    echo "${CONF_FILE_PATH} exists"
else
    cp ./config/config.toml.example $CONF_FILE_PATH
    echo "${CONF_FILE_PATH} created"
fi

ENV_FILE_PATH=${CONF_FILE_DIR}/.env
if [ -f "${ENV_FILE_PATH}" ]; then
    echo "${ENV_FILE_PATH} exists"
else
    cp ./config/env.example $ENV_FILE_PATH
    echo "${ENV_FILE_PATH} created"
fi

cp ./on-chain/contracts/abi/SwanPayment.json $CONF_FILE_DIR_DEST/SwanPayment.json
cp ./on-chain/contracts/abi/FilswanOracle.json $CONF_FILE_DIR_DEST/FilswanOracle.json

git submodule update --init --recursive
make ffi
make
