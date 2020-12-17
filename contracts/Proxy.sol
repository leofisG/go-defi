pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Config.sol";
import "./lib/libCache.sol";
import "./Cache.sol";
import "./interface/IRegistry.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

/**
 * @title The entrance of Furucombo
 * @author Ben Huang
 */
contract Proxy is Cache, Config {
    using Address for address;
    using SafeERC20 for IERC20;

    // keccak256 hash of "furucombo.handler.registry"
    bytes32 private constant HANDLER_REGISTRY = 0x6874162fd62902201ea0f4bf541086067b3b88bd802fac9e150fd2d1db584e19;

    constructor(address registry) public {
        bytes32 dummy_slot = HANDLER_REGISTRY;
        bytes32 slot = HANDLER_REGISTRY;
        assembly {
            sstore(slot, registry)
        }
    }

    /**
     * @notice Direct transfer from EOA should be reverted.
     * @dev Callback function will be handled here.
     */
    function() external payable {
        require(Address.isContract(msg.sender), "Not allowed from EOA");

        // If triggered by a function call, caller should be registered in registry.
        // The function call will then be forwarded to the location registered in
        // registry.
        if (msg.data.length != 0) {
            require(_isValid(msg.sender), "Invalid caller");

            address target = address(
                bytes20(IRegistry(_getRegistry()).getInfo(msg.sender))
            );
            _exec(target, msg.data);
        }
    }

    /**
     * @notice Combo execution function. Including three phases: pre-process,
     * exection and post-process.
     * @param tos The handlers of combo.
     * @param datas The combo datas.
     */
    function batchExec(address[] memory tos, bytes[] memory datas)
        public
        payable
    {
        _preProcess();
        _execs(tos, datas);
        _postProcess();
    }

    /**
     * @notice The execution interface for callback function to be executed.
     * @dev This function can only be called through the handler, which makes
     * the caller become proxy itself.
     */
    function execs(address[] memory tos, bytes[] memory datas) public payable {
        require(msg.sender == address(this), "Does not allow external calls");
        _execs(tos, datas);
    }

    /**
     * @notice The execution phase.
     * @param tos The handlers of combo.
     * @param datas The combo datas.
     */
    function _execs(address[] memory tos, bytes[] memory datas) internal {
        require(
            tos.length == datas.length,
            "Tos and datas length inconsistent"
        );
        for (uint256 i = 0; i < tos.length; i++) {
            _exec(tos[i], datas[i]);
            // Setup the process to be triggered in the post-process phase
            _setPostProcess(tos[i]);
        }
    }

    /**
     * @notice The execution of a single cube.
     * @param _to The handler of cube.
     * @param _data The cube execution data.
     */
    function _exec(address _to, bytes memory _data)
        internal
        returns (bytes memory result)
    {
        require(_isValid(_to), "Invalid handler");
        assembly {
            let succeeded := delegatecall(
                sub(gas, 5000),
                _to,
                add(_data, 0x20),
                mload(_data),
                0,
                0
            )
            let size := returndatasize

            result := mload(0x40)
            mstore(
                0x40,
                add(result, and(add(add(size, 0x20), 0x1f), not(0x1f)))
            )
            mstore(result, size)
            returndatacopy(add(result, 0x20), 0, size)

            switch iszero(succeeded)
                case 1 {
                    revert(add(result, 0x20), size)
                }
        }
    }

    /**
     * @notice Setup the post-process.
     * @param _to The handler of post-process.
     */
    function _setPostProcess(address _to) internal {
        // If the cache is empty, just skip
        // If the top is a custom post-process, replace it with the handler
        // address.
        require(cache.length > 0, "cache empty");
        if (cache.length == 1) return;
        else if (cache.peek() == bytes32(bytes12(uint96(HandlerType.Custom)))) {
            cache.pop();
            // Check if the handler is already set.
            if (bytes4(cache.peek()) != 0x00000000) cache.setAddress(_to);
            cache.setHandlerType(uint256(HandlerType.Custom));
        }
    }

    /// @notice The pre-process phase.
    function _preProcess() internal isCacheEmpty {
        // Set the sender on the top of cache.
        cache.setSender(msg.sender);
    }

    /// @notice The post-process phase.
    function _postProcess() internal {
        // If the top of cache is HandlerType.Custom (which makes it being zero
        // address when `cache.getAddress()`), get the handler address and execute
        // the handler with it and the post-process function selector.
        // If not, use it as token address and send the token back to user.
        while (cache.length > 1) {
            address addr = cache.getAddress();
            if (addr == address(0)) {
                addr = cache.getAddress();
                _exec(addr, abi.encodeWithSelector(POSTPROCESS_SIG));
            } else {
                uint256 amount = IERC20(addr).balanceOf(address(this));
                if (amount > 0) IERC20(addr).safeTransfer(msg.sender, amount);
            }
        }

        // Balance should also be returned to user
        uint256 amount = address(this).balance;
        if (amount > 0) msg.sender.transfer(amount);

        // Pop the msg.sender
        cache.pop();
    }

    /// @notice Get the registry contract address.
    function _getRegistry() internal view returns (address registry) {
        bytes32 slot = HANDLER_REGISTRY;
        assembly {
            registry := sload(slot)
        }
    }
    
    /// @notice Check if the handler is valid in registry.
    function _isValid(address handler) internal view returns (bool result) {
        return IRegistry(_getRegistry()).isValid(handler);
    }
}