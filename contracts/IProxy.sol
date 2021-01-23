pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;


interface IProxy {
    function execs(address[] calldata tos, bytes[] calldata datas) external;
}
