module.exports.prevHash = function prevHash() {
  let prevhash = process.env.BLOCK_PREVHASH || "0";
  return BigInt(prevhash);
};

module.exports.number = function number() {
  let number = process.env.BLOCK_NUMBER || "0";
  return BigInt(number);
};

module.exports.chainId = function chainId() {
  let chainid = process.env.BLOCK_CHAINID || "0";
  return BigInt(chainid);
};

module.exports.coinbase = function coinbase() {
  let coinbase = process.env.BLOCK_COINBASE || "0";
  return BigInt(coinbase);
};

module.exports.difficulty = function difficulty() {
  let difficulty = process.env.BLOCK_DIFFICULTY || "0";
  return BigInt(difficulty);
};

module.exports.gasLimit = function gasLimit() {
  let gaslimit = process.env.BLOCK_GASLIMIT || "0";
  return BigInt(gaslimit);
};

module.exports.timestamp = function timestamp() {
  let timestamp = process.env.BLOCK_TIMESTAMP || "0";
  return BigInt(timestamp);
};
