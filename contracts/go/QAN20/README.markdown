# qan20

--
import "."

## Example

```go

package main

import (
	"math/big"
	"os"

	qan20 "qanx.space/qanx-contracts/go/QAN20"
	callintrpr "qanx.space/qanx-contracts/go/utils/CallInterpreter"
)

func main() {

	token := qan20.NewQAN20("Example", "XMPL", big.NewInt(18))

	callintrpr.Interpret(token, os.Args[1:])

}

```

## Usage

#### type QAN20

```go
type QAN20 struct {
}
```

QAN20 smart contract standard

#### func NewQAN20

```go
func NewQAN20(name string, symbol string, decimals uint64) *QAN20
```

Creates the QAN20 smart contract

#### func (\*QAN20) Allowance

```go
func (token *QAN20) Allowance(owner string, spender string) uint64
```

Retrieve the remaining amount of tokens allowed to spend by spender

#### func (\*QAN20) Approve

```go
func (token *QAN20) Approve(spender string, amount uint64) bool
```

Sets amount as the allowance of spender over the caller's tokens

#### func (\*QAN20) BalanceOf

```go
func (token *QAN20) BalanceOf(owner string) uint64
```

Retrieve the balance of owner

#### func (\*QAN20) Decimals

```go
func (token *QAN20) Decimals() uint64
```

Retrieve the decimal amount

#### func (\*QAN20) Mint

```go
func (token *QAN20) Mint(to string, amount uint64)
```

Mints the amount of tokens and transfers it to "to"

#### func (\*QAN20) Name

```go
func (token *QAN20) Name() string
```

Retrieve the name

#### func (\*QAN20) Symbol

```go
func (token *QAN20) Symbol() string
```

Retrieve the symbol/ticker

#### func (\*QAN20) TotalSupply

```go
func (token *QAN20) TotalSupply() uint64
```

Retrieve the total supply

#### func (\*QAN20) Transfer

```go
func (token *QAN20) Transfer(to string, amount uint64) bool
```

Transfers tokens from sender to "to"

#### func (\*QAN20) TransferFrom

```go
func (token *QAN20) TransferFrom(from string, to string, amount uint64) bool
```

Transfers tokens from "from" to "to"

#### type QAN20Token

```go
type QAN20Token interface {
	Name() string
	Symbol() string
	Decimals() uint64
	TotalSupply() uint64
	BalanceOf(owner string) uint64
	Transfer(to string, amount uint64) bool
	Allowance(owner string, spender string) uint64
	Approve(to string, amount uint64) bool
	TransferFrom(from string, to string, amount uint64) bool
	Mint(to string, tokenId uint64)
}
```
