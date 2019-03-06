# ImgResize

Resize 4x image(png/jpeg) for ios/android/flutter

#### Move application (ImgToMobile.exe for windows, ImgToMobile_mac_os for mac os) to image folder and execute it.

#### Will resize image to 3x and 2x and 1.5x and 1x and generate png file to android/flutter/ios folder.

## Filetree

<pre>
📦example
 ┣ 📜cat1.jpg
 ┣ 📜cat2.jpg
 ┣ 📜cat3.jpg
 ┣ 📜ImgToMobile.exe
 ┗ 📜ImgToMobile_mac_os
</pre>

become to

<pre>
📦example
 ┣ 📂android
 ┃ ┣ 📂drawable-hdpi
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┣ 📂drawable-mdpi
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┣ 📂drawable-xhdpi
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┣ 📂drawable-xxhdpi
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┗ 📂drawable-xxxhdpi
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┣ 📂flutter
 ┃ ┣ 📂1.5x
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┣ 📂2.0x
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┣ 📂3.0x
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┣ 📂4.0x
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┗ 📜cat3.png
 ┃ ┣ 📜cat1.png
 ┃ ┣ 📜cat2.png
 ┃ ┗ 📜cat3.png
 ┣ 📂ios
 ┃ ┣ 📂cat1.imageset
 ┃ ┃ ┣ 📜cat1.png
 ┃ ┃ ┣ 📜cat1@2x.png
 ┃ ┃ ┣ 📜cat1@3x.png
 ┃ ┃ ┗ 📜Contents.json
 ┃ ┣ 📂cat2.imageset
 ┃ ┃ ┣ 📜cat2.png
 ┃ ┃ ┣ 📜cat2@2x.png
 ┃ ┃ ┣ 📜cat2@3x.png
 ┃ ┃ ┗ 📜Contents.json
 ┃ ┗ 📂cat3.imageset
 ┃ ┃ ┣ 📜cat3.png
 ┃ ┃ ┣ 📜cat3@2x.png
 ┃ ┃ ┣ 📜cat3@3x.png
 ┃ ┃ ┗ 📜Contents.json
 ┣ 📜cat1.jpg
 ┣ 📜cat2.jpg
 ┣ 📜cat3.jpg
 ┣ 📜ImgToMobile.exe
 ┗ 📜ImgToMobile_mac_os
</pre>

### Use this [resize package](https://github.com/nfnt/resize) to resize image.

### Use [lossypng](https://github.com/peterhellberg/lossypng) to compression image.
