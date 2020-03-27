# Echoing all commands gets boring
.SILENT:

# Makefile created by someone seriously out of shape
# so this makefile is seriously simplistic and inefficient (but works for now)
BIN = bin
SRC = src

all: \
	$(BIN)\compress.exe
	echo Done!


$(BIN)\compress.exe: \
	$(SRC)\main.go \
	$(SRC)\nodemap\nodemap.go \
	$(SRC)\palette\palette.go \
	$(SRC)\vector\vector.go
	echo Compress..
	go build -o $(BIN)\compress.exe $(SRC)\main.go

