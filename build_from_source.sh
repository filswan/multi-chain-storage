#!/bin/bash

CONF_FILE_DIR_SRC=./config/config_file
CONF_FILE_DIR_DEST=${HOME}/.swan/mcs
mkdir -p ${CONF_FILE_DIR}

CONF_FILE_PATH=${CONF_FILE_DIR_DEST}/config_polygon.mumbai.toml
if [ -f "${CONF_FILE_PATH}" ]; then
    echo "${CONF_FILE_PATH} exists"
else
    cp ${CONF_FILE_DIR_SRC}/config_polygon.mumbai.toml.example $CONF_FILE_PATH
    echo "${CONF_FILE_PATH} created"
fi

CONF_FILE_PATH=${CONF_FILE_DIR_DEST}/config_polygon.mainnet.toml
if [ -f "${CONF_FILE_PATH}" ]; then
    echo "${CONF_FILE_PATH} exists"
else
    cp ${CONF_FILE_DIR_SRC}/config_polygon.mainnet.toml.example $CONF_FILE_PATH
    echo "${CONF_FILE_PATH} created"
fi

CONF_FILE_PATH=${CONF_FILE_DIR_DEST}/config_bsc.testnet.toml
if [ -f "${CONF_FILE_PATH}" ]; then
    echo "${CONF_FILE_PATH} exists"
else
    cp ${CONF_FILE_DIR_SRC}/config_bsc.testnet.toml.example $CONF_FILE_PATH
    echo "${CONF_FILE_PATH} created"
fi

git submodule update --init --recursive
make ffi
make
