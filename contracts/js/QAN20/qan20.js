const message = require("../utils/Message/message");
const db = require("../utils/Database/database");

module.exports = class QAN20 {
  _name = "";
  _symbol = "";
  _decimals = 0n;

  // Creates the QAN20 smart contract
  constructor(name, symbol, decimals) {
    this._name = name;
    this._symbol = symbol;
    this._decimals = decimals;
  }

  // Retrieve the name
  name() {
    return this._name;
  }

  // Retrieve the symbol/ticker
  symbol() {
    return this._symbol;
  }

  // Retrieve the decimal amount
  decimals() {
    return this._decimals;
  }

  // Retrieve the total supply
  totalSupply() {
    const n = BigInt(db.read("TOTAL_SUPPLY", 0));
    return n;
  }

  // Retrieve the balance of owner
  balanceOf(owner) {
    owner = owner.toLowerCase();

    const n = BigInt(db.read(`BALANCE_OF_${owner}`, 0));
    return n;
  }

  // Transfers tokens from sender to "to"
  transfer(to, amount) {
    to = to.toLowerCase();

    const sender = message.sender();

    return transfer(sender, to, amount).bind(this);
  }

  // Retrieve the remaining amount of tokens allowed to spend by spender
  allowance(owner, spender) {
    owner = owner.toLowerCase();
    spender = spender.toLowerCase();

    const n = BigInt(db.read(`TOKEN_ALLOWANCE_${owner}_${spender}`, 0));
    return n;
  }

  // Sets amount as the allowance of spender over the caller's tokens
  approve(spender, amount) {
    spender = spender.toLowerCase();

    const sender = message.sender();
    db.write(`TOKEN_ALLOWANCE_${sender}_${spender}`, amount);
    return true;
  }

  // Transfers tokens from "from" to "to"
  transferFrom(from, to, amount) {
    from = from.toLowerCase();
    to = to.toLowerCase();

    const sender = message.sender();
    const allowance = this.allowance(from, sender);

    if (allowance < amount) {
      process.stderr.write("QAN20: Insufficient allowance\n");
      process.exit(1);
    }

    db.write(`TOKEN_ALLOWANCE_${from}_${sender}`, allowance - amount);

    return transfer(from, to, amount).bind(this);
  }

  // Mints the amount of tokens and transfers it to "to"
  mint(to, amount) {
    db.write(`BALANCE_OF_${to}`, this.balanceOf(to) + amount);
    db.write(`TOTAL_SUPPLY`, this.totalSupply() + amount);
  }
};

function transfer(from, to, amount) {
  const fromBalance = this.balanceOf(from);

  if (fromBalance < amount) {
    process.stderr.write("QAN20: Transfer amount exceeds balance\n");
    process.exit(1);
  }

  db.write(`BALANCE_OF_${from}`, fromBalance - amount);
  db.write(`BALANCE_OF_${to}`, this.balanceOf(to) + amount);
}
