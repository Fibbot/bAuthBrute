# bAuthBrute
basic auth endpoint brute forcer

# usage
go run main.go -users userListLocation -pass passwordListLocation -u https://site.com/thing

# TODO
* add some ability to slow this down, currently it's gonna get you ratelimited quickly
* add ability to add cookies in case these are needed to hit endpoint
