# qan20

## Example

```js
const callintrpr = require("./utils/CallInterpreter/callInterpreter");

const QAN20 = require("./QAN20/qan20");

const qan20 = new QAN20("Example", "XMPL", 18n);

callintrpr.interpret(qan20, process.argv.slice(2));
```

## Usage

#### Constructor

```js
const qan20 = new QAN20("Example", "XMPL", 18n);
```

Creates the QAN721 smart contract

#### Allowance

```go
qan20.allowance(owner, spender);
```

Retrieve the remaining amount of tokens allowed to spend by spender

#### Approve

```go
qan20.approve(spender, amount);
```

Sets amount as the allowance of spender over the caller's tokens

#### Balance Of

```go
qan20.balanceOf(owner);
```

Retrieve the balance of owner

#### Decimals

```go
qan20.decimals();
```

Retrieve the decimal amount

#### Mint

```go
qan20.mint(to, amount);
```

Mints the amount of tokens and transfers it to "to"

#### Name

```go
qan20.name();
```

Retrieve the name

#### Symbol

```go
qan20.symbol();
```

Retrieve the symbol/ticker

#### Total Supply

```go
qan20.totalSupply();
```

Retrieve the total supply

#### Transfer

```go
qan20.transfer(to, amount);
```

Transfers tokens from sender to "to"

#### Transfer From

```go
qan20.transferFrom(from, to, amount);
```

Transfers tokens from "from" to "to"
