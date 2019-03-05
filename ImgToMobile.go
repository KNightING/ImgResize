package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/peterhellberg/lossypng"

	"github.com/nfnt/resize"
)

//ImageType for check
type ImageType int

const (
	//Non do not know
	Non ImageType = iota

	//Png type of image
	Png
	//Jpeg type of image
	Jpeg
)

type IosContents struct {
	Images []IosImages `json:"images"`
	Info   IosInfo     `json:"info"`
}

type IosImages struct {
	Idiom    string `json:"idiom"`
	Filename string `json:"filename"`
	Scale    string `json:"scale"`
}

type IosInfo struct {
	Version int    `json:"version"`
	Author  string `json:"author"`
}

func main() {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Mission Start")

	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	root = root + "/"
	fmt.Println("MAP:" + root)

	androidRoot := root + "/android/"
	// ldpiPath := androidRoot + "drawable-ldpi/"
	mdpiPath := androidRoot + "drawable-mdpi/"
	hdpiPath := androidRoot + "drawable-hdpi/"
	xhdpiPath := androidRoot + "drawable-xhdpi/"
	xxhdpiPath := androidRoot + "drawable-xxhdpi/"
	xxxhdpiPath := androidRoot + "drawable-xxxhdpi/"

	iosRoot := root + "/ios/"

	flutterRoot := root + "/flutter/"
	flutter1_5Path := flutterRoot + "/1.5x/"
	flutter2_0Path := flutterRoot + "/2.0x/"
	flutter3_0Path := flutterRoot + "/3.0x/"
	flutter4_0Path := flutterRoot + "/4.0x/"

	// mkDir(ldpiPath)
	mkDir(mdpiPath)
	mkDir(hdpiPath)
	mkDir(xhdpiPath)
	mkDir(xxhdpiPath)
	mkDir(xxxhdpiPath)
	mkDir(iosRoot)
	mkDir(flutterRoot)
	mkDir(flutter1_5Path)
	mkDir(flutter2_0Path)
	mkDir(flutter3_0Path)
	mkDir(flutter4_0Path)

	files, err := ioutil.ReadDir(root)
	for _, f := range files {
		if !f.IsDir() {
			filePath := root + f.Name()
			file, err := os.Open(filePath)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
			_, err = file.Read(buff)

			filetype := http.DetectContentType(buff)

			file.Close()

			var t = Non

			switch filetype {

			case "image/jpeg", "image/jpg":
				t = Jpeg

			case "image/png":
				t = Png

			default:

			}

			if t != Non {
				fmt.Println("Find target, name:" + f.Name())

				getFilenameWithoutExt := getFilenameWithoutExt(f.Name())
				pngName := getFilenameWithoutExt + ".png"

				iOSInfo := IosInfo{
					Version: 1,
					Author:  "KNightING",
				}
				iosContents := IosContents{
					Info: iOSInfo,
				}
				iOSImgGroup := iosRoot + getFilenameWithoutExt + ".imageset/"
				mkDir(iOSImgGroup)

				//@4x,drawable-xxxhdpi
				targetPath := flutter4_0Path + pngName
				androidPath := xxxhdpiPath + pngName

				if t != Png {
					if imgResize(t, filePath, 1.0, targetPath) {
						copyFile(targetPath, androidPath)
					} else {
						fmt.Println("Mission fail at 4x")
					}
				} else {
					copyFile(filePath, targetPath)
					copyFile(filePath, androidPath)
				}

				//@3x,drawable-xxhdpi
				targetPath = flutter3_0Path + pngName
				iOSFileName := getFilenameWithoutExt + "@3x.png"
				iOSPath := iOSImgGroup + iOSFileName
				androidPath = xxhdpiPath + pngName
				if imgResize(t, filePath, 3.0/4, targetPath) {

					iosContents.Images = append(iosContents.Images, IosImages{
						Idiom:    "universal",
						Filename: iOSFileName,
						Scale:    "3x",
					})

					copyFile(targetPath, iOSPath)
					copyFile(targetPath, androidPath)
				} else {
					fmt.Println("Mission fail at 3x")
				}

				//@2x,drawable-xhdpi
				targetPath = flutter2_0Path + pngName
				iOSFileName = getFilenameWithoutExt + "@2x.png"
				iOSPath = iOSImgGroup + iOSFileName
				androidPath = xhdpiPath + pngName
				if imgResize(t, filePath, 2.0/4, targetPath) {

					iosContents.Images = append(iosContents.Images, IosImages{
						Idiom:    "universal",
						Filename: iOSFileName,
						Scale:    "2x",
					})

					copyFile(targetPath, iOSPath)
					copyFile(targetPath, androidPath)
				} else {
					fmt.Println("Mission fail at 2x")
				}

				//@1.5x,drawable-hdpi
				targetPath = flutter1_5Path + pngName
				androidPath = hdpiPath + pngName
				if imgResize(t, filePath, 3.0/8, targetPath) {
					copyFile(targetPath, androidPath)
				} else {
					fmt.Println("Mission fail at 1.5x")
				}

				//@1x,drawable-mhdpi
				targetPath = flutterRoot + pngName
				iOSPath = iOSImgGroup + pngName
				androidPath = mdpiPath + pngName
				if imgResize(t, filePath, 1.0/4, targetPath) {

					iosContents.Images = append(iosContents.Images, IosImages{
						Idiom:    "universal",
						Filename: pngName,
						Scale:    "1x",
					})

					copyFile(targetPath, iOSPath)
					copyFile(targetPath, androidPath)
				} else {
					fmt.Println("Mission fail at 1x")
				}

				// //@4x,drawable-xxxhdpi
				// targetPath = flutter4_0Path + f.Name()
				// androidPath = xxxhdpiPath + f.Name()
				// copyFile(filePath, targetPath)
				// copyFile(filePath, androidPath)

				writeIosContent(iOSImgGroup, iosContents)

				fmt.Println("Knock down !!!")
			}

		}
	}

	fmt.Println("Mission Complete!!!")
	fmt.Scanln()
	os.Exit(0)
}

func mkDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Create dir error")
		fmt.Println(err)
		return err
	}
	return nil
}

func imgResizeWork(imgType ImageType, path string) {
	// imgResize(imgType, path, 3.0/4)  //@3x,drawable-xxhdpi
	// imgResize(imgType, path, 2.0/4)  //@2x,drawable-xhdpi
	// imgResize(imgType, path, 3.0/8)  //@1.5x,drawable-hdpi
	// imgResize(imgType, path, 1.0/4)  //@1x,drawable-mdpi
	// imgResize(imgType, path, 3.0/16) //@0.75x,drawable-ldpi
}

func imgResize(imgType ImageType, path string, ratio float32, targetPath string) bool {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}

	var img image.Image

	if imgType == Png {
		// decode png into image.Image
		img, err = png.Decode(file)
	}

	if imgType == Jpeg {
		// decode jpeg into image.Image
		img, err = jpeg.Decode(file)
	}

	if err != nil {
		fmt.Println(err)
		return false
	}

	file.Close()

	b := img.Bounds()
	width := uint(float32(b.Max.X) * ratio)
	height := uint(float32(b.Max.Y) * ratio)
	img = resize.Resize(width, height, img, resize.Bilinear)

	// out, err := os.Create(fmt.Sprintf("test_resized_%f.png", ratio))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer out.Close()

	// jpeg.Encode(out, m, nil)

	// write new image to file
	img = lossypng.Optimize(img, lossypng.RGBAConversion, 10)

	buf := new(bytes.Buffer)
	// png.Encode(buf, img)

	pngEncoder := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	pngEncoder.Encode(buf, img)

	// fileName := fmt.Sprintf("%s_%d.png", getFilenameWithoutExt(file.Name()), width)
	err = ioutil.WriteFile(targetPath, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func getFilenameWithoutExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func copyFile(src string, dst string) {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		fmt.Println("Error creating", dst)
		fmt.Println(err)
		return
	}
}

func writeIosContent(groupPath string, contents IosContents) {
	json, _ := json.Marshal(contents)
	// fmt.Println(string(json))
	ioutil.WriteFile(groupPath+"Contents.json", json, 0644)
}

/* build any platform

On Linux and mac

===MAC===
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

===LINUX===
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

===WINDOWS===
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

On Windows

===MAC===
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

===LINUX===
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go

===WINDOWS===
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build main.go

GOOS：目標平台的操作系統（darwin、freebsd、linux、windows）
GOARCH：目標平台的體系架構（386、amd64、arm）
交叉編譯不支持CGO所以要禁用它

*/
