# qrc721
--
    import "."


## Usage

#### type QRC721

```go
type QRC721 struct {
}
```

QRC721 smart contract standard

#### func  NewQRC721

```go
func NewQRC721(name string, symbol string, baseUri string) *QRC721
```
Creates the QRC721 smart contract

#### func (*QRC721) Approve

```go
func (token *QRC721) Approve(to string, tokenId uint64)
```
Give permission to "to" to transfer token id to another account

#### func (*QRC721) BalanceOf

```go
func (token *QRC721) BalanceOf(owner string) uint64
```
Retrieve the balance of owner

#### func (*QRC721) GetApproved

```go
func (token *QRC721) GetApproved(tokenId uint64) string
```
Returns the account approved for token id

#### func (*QRC721) IsApprovedForAll

```go
func (token *QRC721) IsApprovedForAll(owner string, operator string) bool
```
Returns if the operator is allowed to manage all of the assets of owner

#### func (*QRC721) IsApprovedOrOwner

```go
func (token *QRC721) IsApprovedOrOwner(spender string, tokenId uint64) bool
```
Returns if the operator is allowed to manage all of the assets of owner or is
the owner

#### func (*QRC721) Mint

```go
func (token *QRC721) Mint(to string, tokenId uint64)
```
Mints the token id and transfers it to "to"

#### func (*QRC721) Name

```go
func (token *QRC721) Name() string
```
Retrieve the name

#### func (*QRC721) OwnerOf

```go
func (token *QRC721) OwnerOf(tokenId uint64) string
```
Retrieve the owner of token id

#### func (*QRC721) SetApprovalForAll

```go
func (token *QRC721) SetApprovalForAll(operator string, approved bool)
```
Approve or remove an operator for the caller

#### func (*QRC721) Symbol

```go
func (token *QRC721) Symbol() string
```
Retrieve the symbol/ticker

#### func (*QRC721) TokenURI

```go
func (token *QRC721) TokenURI(tokenId uint64) string
```
Retrieve the token uri for token id

#### func (*QRC721) TransferFrom

```go
func (token *QRC721) TransferFrom(from string, to string, tokenId uint64)
```
Transfers token id from "from" to "to"

#### type QRC721Token

```go
type QRC721Token interface {
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
