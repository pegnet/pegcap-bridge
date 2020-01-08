module github.com/whosoup/pegcap-bridge

go 1.13

require (
	github.com/AdamSLevy/jsonrpc2 v2.0.0+incompatible // indirect
	github.com/Factom-Asset-Tokens/factom v0.0.0-20190911201853-7b283996f02a
	github.com/labstack/echo/v4 v4.1.10
	github.com/pegnet/pegnetd v0.4.0
)

replace github.com/Factom-Asset-Tokens/factom => github.com/Emyrk/factom v0.0.0-20191001194233-40c0cdc2f2a0

replace github.com/pegnet/pegnetd => github.com/WhoSoup/pegnetd v0.2.1-0.20200108035536-cba3a5ba91e4
