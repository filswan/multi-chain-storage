#!/bin/bash

./script/install_pre-requisite.sh


CONF_FILE_DIR_DEST=${HOME}/.swan/mcp
mkdir -p ${CONF_FILE_DIR_DEST}

CONF_FILE_NAMES=(
    config.toml
    config_polygon.toml
    config_nbai.toml
    config_goerli.toml
    config_bsc.toml
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

echo $CONF_FILE_PATH

git submodule update --init --recursive
#make ffi
#make

