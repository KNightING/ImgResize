# ImgToMobile

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
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┣ 📂drawable-mdpi
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┣ 📂drawable-xhdpi
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┣ 📂drawable-xxhdpi
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┗ 📂drawable-xxxhdpi
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┣ 📂flutter
 ┃ ┣ 📂1.5x
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┣ 📂2.0x
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┣ 📂3.0x
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┣ 📂4.0x
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┗ 📜cat3.jpg
 ┃ ┣ 📜cat1.jpg
 ┃ ┣ 📜cat2.jpg
 ┃ ┗ 📜cat3.jpg
 ┣ 📂ios
 ┃ ┣ 📂cat1.imageset
 ┃ ┃ ┣ 📜cat1.jpg
 ┃ ┃ ┣ 📜cat1@2x.png
 ┃ ┃ ┣ 📜cat1@3x.png
 ┃ ┃ ┗ 📜Contents.json
 ┃ ┣ 📂cat2.imageset
 ┃ ┃ ┣ 📜cat2.jpg
 ┃ ┃ ┣ 📜cat2@2x.png
 ┃ ┃ ┣ 📜cat2@3x.png
 ┃ ┃ ┗ 📜Contents.json
 ┃ ┗ 📂cat3.imageset
 ┃ ┃ ┣ 📜cat3.jpg
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
