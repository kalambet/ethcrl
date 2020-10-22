const Migrations = artifacts.require("Migrations");
const CRLv0 = artifacts.require("CRLv0");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(CRLv0);
};
