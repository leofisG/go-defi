pragma solidity ^0.5.0;

import "../HandlerBase.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";



contract HFunds is HandlerBase {
    using SafeERC20 for IERC20;

    function inject(address[] calldata tokens, uint256[] calldata amounts)
        external
        payable
    {
        require(
            tokens.length == amounts.length,
            "token and amount does not match"
        );
        address sender = cache.getSender();
        for (uint256 i = 0; i < tokens.length; i++) {
            IERC20(tokens[i]).safeTransferFrom(
                sender,
                address(this),
                amounts[i]
            );

            // Update involved token
            _updateToken(tokens[i]);
        }
    }

    function sendToken(
        address token,
        uint256 amount,
        address receiver
    ) external payable {
        IERC20(token).safeTransfer(receiver, amount);
    }

    function send(uint256 amount, address payable receiver) external payable {
        receiver.transfer(amount);
    }
}
