# android ndk r24

build:
	go mod tidy
	fyne package -os android -appID darts.mobile -icon logo.png -name darts

debug:
	adb -d install -r darts.apk

install: build debug