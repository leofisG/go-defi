var Proxy = artifacts.require("./Proxy.sol");
var Registry = artifacts.require("./Registry.sol");
var HCEther = artifacts.require("./handlers/compound/HCEther.sol");

module.exports = function(deployer) {
    deployer.deploy(Registry).then(function() {
        return deployer.deploy(Proxy, Registry.address);
    });
    deployer.deploy(HCEther);
};