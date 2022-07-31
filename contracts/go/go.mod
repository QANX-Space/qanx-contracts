module main

go 1.18

replace qanx.space/qanx-contracts/go/QRC721 => ./QRC721

replace qanx.space/qanx-contracts/go/utils/Context => ./utils/Context

replace qanx.space/qanx-contracts/go/utils/Database => ./utils/Database

replace qanx.space/qanx-contracts/go/utils/CallInterpreter => ./utils/CallInterpreter

require (
	qanx.space/qanx-contracts/go/QRC721 v0.0.0-00010101000000-000000000000
	qanx.space/qanx-contracts/go/utils/CallInterpreter v0.0.0-00010101000000-000000000000
)

require (
	qanx.space/qanx-contracts/go/utils/Context v0.0.0-00010101000000-000000000000 // indirect
	qanx.space/qanx-contracts/go/utils/Database v0.0.0-00010101000000-000000000000 // indirect
)
