#!/bin/bash

CONF_FILE_DIR_DEST=${HOME}/.swan/mcs
mkdir -p ${CONF_FILE_DIR_DEST}

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


cp ./on-chain/contracts/abi/SwanPayment.json $CONF_FILE_DIR_DEST/SwanPayment.json

git submodule update --init --recursive
make ffi
make
