# qrc20
--
    import "."

## Example

```go

package main

import (

    "os"

    qrc20 "qanx.space/qanx-contracts/go/QRC20"
    callintrpr "qanx.space/qanx-contracts/go/utils/CallInterpreter"

)

func main() {

    token := qrc20.NewQRC20("Example", "XMPL", 8)

    callintrpr.Interpret(token, os.Args[1:])

}

```

## Usage

#### type QRC20

```go
type QRC20 struct {
}
```

QRC20 smart contract standard

#### func  NewQRC20

```go
func NewQRC20(name string, symbol string, decimals uint64) *QRC20
```
Creates the QRC20 smart contract

#### func (*QRC20) Allowance

```go
func (token *QRC20) Allowance(owner string, spender string) uint64
```
Retrieve the remaining amount of tokens allowed to spend by spender

#### func (*QRC20) Approve

```go
func (token *QRC20) Approve(spender string, amount uint64) bool
```
Sets amount as the allowance of spender over the caller's tokens

#### func (*QRC20) BalanceOf

```go
func (token *QRC20) BalanceOf(owner string) uint64
```
Retrieve the balance of owner

#### func (*QRC20) Decimals

```go
func (token *QRC20) Decimals() uint64
```
Retrieve the decimal amount

#### func (*QRC20) Mint

```go
func (token *QRC20) Mint(to string, amount uint64)
```
Mints the amount of tokens and transfers it to "to"

#### func (*QRC20) Name

```go
func (token *QRC20) Name() string
```
Retrieve the name

#### func (*QRC20) Symbol

```go
func (token *QRC20) Symbol() string
```
Retrieve the symbol/ticker

#### func (*QRC20) TotalSupply

```go
func (token *QRC20) TotalSupply() uint64
```
Retrieve the total supply

#### func (*QRC20) Transfer

```go
func (token *QRC20) Transfer(to string, amount uint64) bool
```
Transfers tokens from sender to "to"

#### func (*QRC20) TransferFrom

```go
func (token *QRC20) TransferFrom(from string, to string, amount uint64) bool
```
Transfers tokens from "from" to "to"

#### type QRC20Token

```go
type QRC20Token interface {
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
