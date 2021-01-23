# Deployment Guide for Contracts

1. Deploy the `Registry.sol` contract and get its address.
2. Deploy the `Proxy.sol` contract and providing the addres of the `Registry`.
3. Deploy the `handlers/compound/HCEther.sol` and get its address `HCEtherAddr`.
4. Call the `register(HCEtherAddr, bytes32data)` of the `Registry` to record the handler.