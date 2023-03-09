# qan721

--
import "."

## Example

```go

package main

import (

    "os"

    qan721 "qanx.space/qanx-contracts/go/QAN721"
    callintrpr "qanx.space/qanx-contracts/go/utils/CallInterpreter"

)

func main() {

    token := qan721.NewQAN721("Example", "XMPL", "https://example.com/")

    callintrpr.Interpret(token, os.Args[1:])

}

```

## Usage

#### type QAN721

```go
type QAN721 struct {
}
```

QAN721 smart contract standard

#### func NewQAN721

```go
func NewQAN721(name string, symbol string, baseUri string) *QAN721
```

Creates the QAN721 smart contract

#### func (\*QAN721) Approve

```go
func (token *QAN721) Approve(to string, tokenId uint64)
```

Give permission to "to" to transfer token id to another account

#### func (\*QAN721) BalanceOf

```go
func (token *QAN721) BalanceOf(owner string) uint64
```

Retrieve the balance of owner

#### func (\*QAN721) GetApproved

```go
func (token *QAN721) GetApproved(tokenId uint64) string
```

Returns the account approved for token id

#### func (\*QAN721) IsApprovedForAll

```go
func (token *QAN721) IsApprovedForAll(owner string, operator string) bool
```

Returns if the operator is allowed to manage all of the assets of owner

#### func (\*QAN721) IsApprovedOrOwner

```go
func (token *QAN721) IsApprovedOrOwner(spender string, tokenId uint64) bool
```

Returns if the operator is allowed to manage all of the assets of owner or is
the owner

#### func (\*QAN721) Mint

```go
func (token *QAN721) Mint(to string, tokenId uint64)
```

Mints the token id and transfers it to "to"

#### func (\*QAN721) Name

```go
func (token *QAN721) Name() string
```

Retrieve the name

#### func (\*QAN721) OwnerOf

```go
func (token *QAN721) OwnerOf(tokenId uint64) string
```

Retrieve the owner of token id

#### func (\*QAN721) SetApprovalForAll

```go
func (token *QAN721) SetApprovalForAll(operator string, approved bool)
```

Approve or remove an operator for the caller

#### func (\*QAN721) Symbol

```go
func (token *QAN721) Symbol() string
```

Retrieve the symbol/ticker

#### func (\*QAN721) TokenURI

```go
func (token *QAN721) TokenURI(tokenId uint64) string
```

Retrieve the token uri for token id

#### func (\*QAN721) TransferFrom

```go
func (token *QAN721) TransferFrom(from string, to string, tokenId uint64)
```

Transfers token id from "from" to "to"

#### type QAN721Token

```go
type QAN721Token interface {
	Name() string
	Symbol() string
	BalanceOf(owner string) uint64
	OwnerOf(tokenId uint64) string
	TokenURI(tokenId uint64) string
	Approve(to string, tokendId uint64)
	GetApproved(tokenId uint64) string
	SetApprovalForAll(operator string, approved bool)
	IsApprovedForAll(owner string, operator string) bool
	TransferFrom(from string, to string, tokenId uint64)
	Mint(to string, tokenId uint64)
}
```
