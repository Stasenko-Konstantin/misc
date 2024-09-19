# android ndk r24

build:
	go mod tidy
	fyne package -os android -appID twister.mobile -icon logo.png -name twister

debug:
	adb -d install -r twister.apk

install: build debug