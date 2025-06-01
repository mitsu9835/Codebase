var abi = JSON.parse("'"${ABI}"'");
var bin = "0x'"${BIN}"'";
var Simple = eth.contract(abi);
var tx = Simple.new({from: eth.accounts[0], data: bin, gas: 3000000});
console.log("TxHash:", tx.transactionHash);
