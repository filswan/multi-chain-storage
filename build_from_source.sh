#!/bin/bash

CONF_FILE_DIR=${HOME}/.swan/mcp
mkdir -p ${CONF_FILE_DIR}

CONF_FILE_NAMES=(
    config.toml, 
    config_polygon.toml,
    config_nbai.toml,
    config_goerli.toml,
    config_bsc.toml
    )

for CONF_FILE_NAME in ${CONF_FILE_NAMES[@]}; do
  echo $CONF_FILE_NAME
done

echo $CONF_FILE_PATH

if [ -f "${CONF_FILE_PATH}" ]; then
    echo "${CONF_FILE_PATH} exists"
else
    cp ./config/config.toml.example $CONF_FILE_PATH
    echo "${CONF_FILE_PATH} created"
fi

git submodule update --init --recursive
make ffi
make

