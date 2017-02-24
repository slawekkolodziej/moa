include Makefile.build

.PHONY: components

components:
	tools/components

run: components
	GODEBUG="cgocheck=0" go run Moa.go

prepare:
	qmake
