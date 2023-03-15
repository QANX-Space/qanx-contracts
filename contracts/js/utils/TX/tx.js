module.exports.gasPrice = function gasPrice() {
  let gasPrice = process.env.TX_GASPRICE || "0";
  return BigInt(gasPrice);
};

module.exports.origin = function origin() {
  let origin = process.env.TX_ORIGIN.toLowerCase();

  if (!origin) {
    process.stderr.write("TX: Origin is not known\n");
    process.exit(1);
  }

  return origin;
};
