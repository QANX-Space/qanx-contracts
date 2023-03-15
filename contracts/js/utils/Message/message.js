module.exports.sender = function sender() {
  let sender = process.env.MSG_SENDER.toLowerCase();

  if (!sender) {
    process.stderr.write("Message: Sender is not known\n");
    process.exit(1);
  }

  return sender;
};

module.exports.value = function value() {
  let value = process.env.MSG_VALUE || "0";
  return BigInt(value);
};
