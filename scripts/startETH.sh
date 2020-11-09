#!/bin/bash

# Set the mnemoic to make sure that the address generated is the same
# Unlock the account to mint some coins
npm install -g ganache-cli
ganache-cli -q \
  -f http://watson.lowland.fun:19545/ \
  -m "clutch captain shoe salt awake harvest setup primary inmate ugly among become" \
  -i 1 \
  --unlock 0x39aa39c021dfbae8fac545936693ac917d5e7563 \
  --unlock 0xa0df350d2637096571F7A701CBc1C5fdE30dF76A \
  -u 0x9759A6Ac90977b93B58547b4A71c78317f391A28 &
