module.exports.sender = function sender() {
  let sender = process.env.SENDER.toLowerCase();

  if (!sender) {
    process.stderr.write("Context: Sender is not known\n");
    process.exit(1);
  }

  return sender;
};
