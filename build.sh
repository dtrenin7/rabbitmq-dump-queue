rm -rf ./rmqdump
clear
go build -o rmqdump
./rmqdump "$@"
