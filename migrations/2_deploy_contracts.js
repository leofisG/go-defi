var Proxy = artifacts.require("./Proxy.sol");
var Registry = artifacts.require("./Registry.sol");
var HCEther = artifacts.require("./handlers/compound/HCEther.sol");
var HSushiswap = artifacts.require("./handlers/sushiswap/HSushiswap.sol");
var HAave = artifacts.require("./handlers/aave/HAave.sol");
var HFunds = artifacts.require("./handlers/funds/HFunds.sol");
const AAVE_LENDING_POOL_ADDR = "0x398ec7346dcd622edc5ae82352f02be94c62d119"
const DUMMY_ADDR = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"


module.exports = async function(deployer) {

    await deployer.deploy(Registry);
    registry = await Registry.deployed();
    await deployer.deploy(Proxy, registry.address);
    
    await deployer.deploy(HCEther);
    hCEther = await HCEther.deployed();
    await registry.register(hCEther.address, DUMMY_ADDR)

    await deployer.deploy(HSushiswap);
    hSushiswap = await HSushiswap.deployed();
    await registry.register(hSushiswap.address, DUMMY_ADDR)

    await deployer.deploy(HAave);
    hAave = await HAave.deployed();
    await registry.register(hAave.address, DUMMY_ADDR)


    await deployer.deploy(HFunds);
    hFunds = await HFunds.deployed();
    await registry.register(hFunds.address, DUMMY_ADDR)

    // Aave lending pool
    await registry.register(AAVE_LENDING_POOL_ADDR, hAave.address)
};
