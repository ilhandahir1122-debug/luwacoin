#!/bin/bash

./build/bin/geth init genesis.json --datadir luwa-node

./build/bin/geth \
 --networkid 9882 \
 --datadir luwa-node \
 --http --http.addr 0.0.0.0 \
 --http.port 8545 \
 --http.api "eth,net,web3,personal" \
 --allow-insecure-unlock \
 --nodiscover \
 --miner.gaslimit 30000000
