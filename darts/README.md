# darts
tiny android application for basic automating the game of darts <br>
supports only Russian

### build:
    fyne package -os android -appID darts.mobile -icon logo.png -name darts

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