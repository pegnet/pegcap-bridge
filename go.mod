module github.com/whosoup/pegcap-bridge

go 1.13

require (
	github.com/AdamSLevy/jsonrpc2 v2.0.0+incompatible // indirect
	github.com/Factom-Asset-Tokens/factom v0.0.0-20200212203234-df52cf9bebfb
	github.com/labstack/echo/v4 v4.1.10
	github.com/pegnet/pegnetd v0.5.1
)

replace github.com/Factom-Asset-Tokens/factom => github.com/Emyrk/factom v0.0.0-20200113153851-17d98c31e1bd

replace github.com/pegnet/pegnetd => github.com/WhoSoup/pegnetd v0.2.1-0.20200211091744-f09a03c279eb
