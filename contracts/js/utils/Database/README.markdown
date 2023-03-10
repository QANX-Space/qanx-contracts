# db

## Example

```js
const db = require("./utils/Database/database");

const totalSupply = db.read("TOTAL_SUPPLY", 0); // 0 if TOTAL_SUPPLY not set.
```

## Usage

#### Prune

```js
db.prune(key);
```

Removes the key from the database

#### Read

```js
db.read(key, defaultsTo);
```

Read key from database

#### Write

```js
db.write(key, value);
```

Write key and value to the database
