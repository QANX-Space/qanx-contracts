# callintrpr

## Example

```js
const callintrpr = require("./utils/CallInterpreter/callInterpreter");

const QAN721 = require("./QAN721/qan721");

const qan721 = new QAN721("Example", "XMPL", "https://example.com/");

callintrpr.interpret(qan721, process.argv.slice(2));
```

## Usage

#### Interpret

```js
callintrpr.interpret(contract, args);
```

Takes in a smart contract and calls its functions based on the arguments given
