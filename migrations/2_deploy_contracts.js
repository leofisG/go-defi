var Proxy = artifacts.require("./Proxy.sol");
var Registry = artifacts.require("./Registry.sol");
var HCEther = artifacts.require("./handlers/compound/HCEther.sol");

module.exports = async function(deployer) {

    await deployer.deploy(Registry);
    registry = await Registry.deployed();
    await deployer.deploy(Proxy, registry.address);
    await deployer.deploy(HCEther);
    hCEther = HCEther.deployed();
    if (hCEther.address != undefined) {
        await registry.register(hCEther.address, "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
    }

};
