#!/bin/bash

git submodule update --init --recursive
make ffi
make

