# twister
tiny android application for basic automating the game of twister <br>
supports only Russian

### build:
    fyne package -os android -appID twister.mobile -icon logo.png -name twister

#### or:
    make build

### debug:
    adb -d install -r twister.apk

#### or:
    make debug

### install:
    make install

<br>

REQUIREMENTS: Golang, Android NDK, fyne, Android itself