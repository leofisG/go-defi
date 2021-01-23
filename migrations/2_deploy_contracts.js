var Proxy = artifacts.require("./Proxy.sol");
var Registry = artifacts.require("./Registry.sol");
var HSushiswap = artifacts.require("./handlers/sushiswap/HSushiswap.sol");
var UniswapFlashSwapper = artifacts.require("./handlers/uniswap/UniswapFlashSwapper.sol");
const AAVE_LENDING_POOL_ADDR = "0x398ec7346dcd622edc5ae82352f02be94c62d119"
const DUMMY_ADDR = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const hCEtherAddr   = "0x9A1049f7f87Dbb0468C745d9B3952e23d5d6CE5e"
const hCTokenAddr   = "0x8973D623d883c5641Dd3906625Aac31cdC8790c5"
const hMakerDaoAddr = "0x294fbca49c8a855e04d7d82b28256b086d39afea"
const hUniswapAddr  = "0x58a21cfcee675d65d577b251668f7dc46ea9c3a0"
const hCurveAddr    = "0xa36dfb057010c419c5917f3d68b4520db3671cdb"
const hYearnAddr    = "0xC50C8F34c9955217a6b3e385a069184DCE17fD2A"
const hAaveAddr     = "0xf579b009748a62b1978639d6b54259f8dc915229"
const hOneInch      = "0x783f5c56e3c8b23d90e4a271d7acbe914bfcd319"
const hFunds        = "0xf9b03e9ea64b2311b0221b2854edd6df97669c09"
const hKyberAddr    = "0xe2a3431508cd8e72d53a0e4b57c24af2899322a0"


module.exports = async function(deployer) {

    await deployer.deploy(Registry);
    registry = await Registry.deployed();
    await deployer.deploy(Proxy, registry.address);

    await deployer.deploy(HSushiswap);
    hSushiswap = await HSushiswap.deployed();
    await registry.register(hSushiswap.address, DUMMY_ADDR)

    await deployer.deploy(UniswapFlashSwapper);
    uniswapFlashSwapper = await UniswapFlashSwapper.deployed();
    await registry.register(uniswapFlashSwapper.address, DUMMY_ADDR)

    // Aave lending pool
    await registry.register(AAVE_LENDING_POOL_ADDR, hAaveAddr)
    // register a dummy address for uniswap flash swapper.
    await registry.register("0x1111111111111111111111111111111111111111", uniswapFlashSwapper.address)
    await registry.register(hCEtherAddr, DUMMY_ADDR)
    await registry.register(hCTokenAddr, DUMMY_ADDR)
    await registry.register(hMakerDaoAddr, DUMMY_ADDR)
    await registry.register(hUniswapAddr, DUMMY_ADDR)
    await registry.register(hCurveAddr, DUMMY_ADDR)
    await registry.register(hYearnAddr, DUMMY_ADDR)
    await registry.register(hAaveAddr, DUMMY_ADDR)
    await registry.register(hOneInch, DUMMY_ADDR)
    await registry.register(hFunds, DUMMY_ADDR)
    await registry.register(hKyberAddr, DUMMY_ADDR)
};
