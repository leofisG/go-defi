/**
 *Submitted for verification at Etherscan.io on 2020-03-26
*/

pragma solidity 0.5.16;

contract Context {
    // Empty internal constructor, to prevent people from mistakenly deploying
    // an instance of this contract, which should be used via inheritance.
    constructor () internal { }
    // solhint-disable-previous-line no-empty-blocks

    function _msgSender() internal view returns (address payable) {
        return msg.sender;
    }

    function _msgData() internal view returns (bytes memory) {
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        return msg.data;
    }
}

contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor () internal {
        address msgSender = _msgSender();
        _owner = msgSender;
        emit OwnershipTransferred(address(0), msgSender);
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(isOwner(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Returns true if the caller is the current owner.
     */
    function isOwner() public view returns (bool) {
        return _msgSender() == _owner;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public onlyOwner {
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     */
    function _transferOwnership(address newOwner) internal {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }
}

contract Registry is Ownable {
    mapping(address => bytes32) handlers;

    bytes32 constant DEPRECATED = bytes10(0x64657072656361746564);

    function register(address registration, bytes32 info) external onlyOwner {
        require(registration != address(0), "zero address");
        require(handlers[registration] == bytes32(0), "registered");
        handlers[registration] = info;
    }

    function unregister(address registration) external onlyOwner {
        require(registration != address(0), "zero address");
        require(handlers[registration] != bytes32(0), "no registration");
        require(handlers[registration] != DEPRECATED, "unregistered");
        handlers[registration] = DEPRECATED;
    }

    function isValid(address handler) external view returns (bool result) {
        if (handlers[handler] == 0 || handlers[handler] == DEPRECATED)
            return false;
        else
            return true;
    }

    function getInfo(address handler) external view returns (bytes32 info) {
        return handlers[handler];
    }
}