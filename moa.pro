MAKEFILE = Makefile.build

first.target = first
first.depends = build
first.CONFIG = phony

build.target = build
build.depends = all Moa.go
build.CONFIG = phony

app.target = Moa.go
app.commands = GODEBUG="cgocheck=0" go build -o $(TARGET) Moa.go
app.depends = Makefile components
app.CONFIG = phony

QMAKE_EXTRA_TARGETS += first build app