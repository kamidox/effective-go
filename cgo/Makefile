# CC_FOR_linux_arm=/pitools/arm-bcm2708/gcc-linaro-arm-linux-gnueabihf-raspbian-x64/bin/arm-linux-gnueabihf-gcc
export CC = /rpxc/bin/arm-linux-gnueabihf-gcc
export CGO_LDFLAGS=--sysroot=/rpxc/sysroot
export CGO_CFLAGS=--sysroot=/rpxc/sysroot
export CGO_CPPFLAGS=--sysroot=/rpxc/sysroot
export CGO_CXXFLAGS=--sysroot=/rpxc/sysroot
export PATH := $(PATH):/usr/local/go/bin

all:
	GOOS=linux GOARCH=arm CGO_ENABLED=1 GO111MODULE=on go build
