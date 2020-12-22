pragma solidity ^0.5.0;

interface ILendingPoolCore {
	function getReserveATokenAddress(address _reserve) external view returns (address);
}
