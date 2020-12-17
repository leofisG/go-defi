pragma solidity ^0.5.0;

import "../Config.sol";
import "../Cache.sol";


contract HandlerBase is Cache, Config {
    function postProcess() external payable {
        revert("Invalid post process");
        /* Implementation template
        bytes4 sig = cache.getSig();
        if (sig == bytes4(keccak256(bytes("handlerFunction_1()")))) {
            // Do something
        } else if (sig == bytes4(keccak256(bytes("handlerFunction_2()")))) {
            bytes32 temp = cache.get();
            // Do something
        } else revert("Invalid post process");
        */
    }

    function _updateToken(address token) internal {
        cache.setAddress(token);
        // Ignore token type to fit old handlers
        // cache.setHandlerType(uint256(HandlerType.Token));
    }

    function _updatePostProcess(bytes32[] memory params) internal {
        for (uint256 i = params.length; i > 0; i--) {
            cache.set(params[i - 1]);
        }
        cache.set(msg.sig);
        cache.setHandlerType(uint256(HandlerType.Custom));
    }
}
