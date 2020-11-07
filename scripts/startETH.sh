#!/bin/bash

npm install -g ganache-cli
ganache-cli -q \
  -f http://watson.lowland.fun:19545/ \
  -m "clutch captain shoe salt awake harvest setup primary inmate ugly among become" \
  -i 1 \
  -u 0x9759A6Ac90977b93B58547b4A71c78317f391A28 &
