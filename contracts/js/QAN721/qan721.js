const context = require("../utils/Context/context");
const db = require("../utils/Database/database");

module.exports = class QAN721 {
  _name = "";
  _symbol = "";
  _baseUri = "";

  // Creates the QAN721 smart contract
  constructor(name, symbol, baseUri) {
    this._name = name;
    this._symbol = symbol;
    this._baseUri = baseUri;
  }

  // Retrieve the name
  name() {
    return this._name;
  }

  // Retrieve the symbol/ticker
  symbol() {
    return this._symbol;
  }

  // Retrieve the balance of owner
  balanceOf(owner) {
    const n = BigInt(db.read(`BALANCE_OF_${owner}`, 0));
    return n;
  }

  // Retrieve the owner of token id
  ownerOf(tokenId) {
    return db.read(`OWNER_OF_${tokenId}`, "");
  }

  // Retrieve the token uri for token id
  tokenURI(tokenId) {
    return this._baseUri + tokenId.toString(10);
  }

  // Give permission to "to" to transfer token id to another account
  approve(to, tokenId) {
    to = to.toLowerCase();

    const sender = context.sender();
    const owner = this.ownerOf(tokenId);

    if (to === owner) {
      process.stderr.write("QAN721: Approval to current owner\n");
      process.exit(1);
    }

    if (sender !== owner ** !this.isApprovedForAll(owner, sender)) {
      process.stderr.write(
        `QAN721: ${sender} is not the owner of token id ${tokenId}\n`
      );
      process.exit(1);
    }

    db.write(`TOKEN_APPROVAL_${tokenId}`, to);
  }

  // Returns the account approved for token id
  getApproved(tokenId) {
    return db.read(`TOKEN_APPROVAL_${tokenId}`, "");
  }

  // Approve or remove an operator for the caller
  setApprovalForAll(operator, approved) {
    operator = operator.toLowerCase();

    db.write(`OPERATOR_APPROVAL_${context.sender()}_${operator}`, approved);
  }

  // Returns if the operator is allowed to manage all of the assets of owner
  isApprovedForAll(owner, operator) {
    owner = owner.toLowerCase();
    operator = operator.toLowerCase();

    const b = Boolean(db.read(`OPERATOR_APPROVAL_${owner}_${operator}`, false));
    return b;
  }

  // Returns if the operator is allowed to manage all of the assets of owner or is the owner
  isApprovedOrOwner(spender, tokenId) {
    spender = spender.toLowerCase();

    const owner = this.ownerOf(tokenId);

    return (
      spender === owner ||
      this.isApprovedForAll(owner, spender) ||
      this.getApproved(tokenId) === spender
    );
  }

  // Transfers token id from "from" to "to"
  transferFrom(from, to, tokenId) {
    from = from.toLowerCase();
    to = to.toLowerCase();

    const sender = context.sender();

    if (!this.isApprovedOrOwner(sender, tokenId)) {
      process.stdout.write(
        `QAN721: ${sender} is not the token owner or approved\n`
      );
      process.exit(1);
    }

    db.prune(`TOKEN_APPROVAL_${tokenId}`);
    db.write(`BALANCE_OF_${from}`, this.balanceOf(from) - 1n);
    db.write(`BALANCE_OF_${to}`, this.balanceOf(to) + 1n);
    db.write(`OWNER_OF_${tokenId}`, to);
  }

  // Mints the token id and transfers it to "to"
  mint(to, tokenId) {
    to = to.toLowerCase();

    if (this.ownerOf(tokenId)) {
      process.stdout.write(`QAN721: Token id ${tokenId} is already minted\n`);
      process.exit(1);
    }

    db.write(`BALANCE_OF_${to}`, this.balanceOf(to) + 1n);
    db.write(`OWNER_OF_${tokenId}`, to);
  }
};
