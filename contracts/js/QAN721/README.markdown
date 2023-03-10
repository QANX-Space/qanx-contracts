# qan721

## Example

```js
const callintrpr = require("./utils/CallInterpreter/callInterpreter");

const QAN721 = require("./QAN721/qan721");

const qan721 = new QAN721("Example", "XMPL", "https://example.com/");

callintrpr.interpret(qan721, process.argv.slice(2));
```

## Usage

#### Constructor

```js
const qan721 = new QAN721(name, symbol, baseUri);
```

Creates the QAN721 smart contract

#### Approve

```js
qan721.approve(to, tokenId);
```

Give permission to "to" to transfer token id to another account

#### Balance Of

```js
qan721.balanceOf(owner);
```

Retrieve the balance of owner

#### Get Approved

```js
qan721.getApproved(tokenId);
```

Returns the account approved for token id

#### Is Approved For All

```js
qan721.isApprovedForAll(owner, operator);
```

Returns if the operator is allowed to manage all of the assets of owner

#### Is Approved Or Owner

```js
qan721.isApprovedOrOwner(spender, tokenId);
```

Returns if the operator is allowed to manage all of the assets of owner or is
the owner

#### Mint

```js
qan721.mint(to, tokenId);
```

Mints the token id and transfers it to "to"

#### Name

```js
qan721.name();
```

Retrieve the name

#### Owner Of

```js
qan721.ownerOf(tokenId);
```

Retrieve the owner of token id

#### Set Approval For All

```js
qan721.setApprovalForAll(operator, approved);
```

Approve or remove an operator for the caller

#### Symbol

```js
qan721.symbol();
```

Retrieve the symbol/ticker

#### Token URI

```js
qan721.tokenURI(tokenId);
```

Retrieve the token uri for token id

#### Transfer From

```js
qan721.transferFrom(from, to, tokenId);
```

Transfers token id from "from" to "to"
