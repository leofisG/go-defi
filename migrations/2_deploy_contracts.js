var Proxy = artifacts.require("./Proxy.sol");
var Registry = artifacts.require("./Registry.sol");
var HCEther = artifacts.require("./handlers/compound/HCEther.sol");
var HSushiswap = artifacts.require("./handlers/sushiswap/HSushiswap.sol");

module.exports = async function(deployer) {

    await deployer.deploy(Registry);
    registry = await Registry.deployed();
    await deployer.deploy(Proxy, registry.address);
    
    await deployer.deploy(HCEther);
    hCEther = await HCEther.deployed();
    await registry.register(hCEther.address, "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

    await deployer.deploy(HSushiswap);
    hSushiswap = await HSushiswap.deployed();
    await registry.register(hSushiswap.address, "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

};
