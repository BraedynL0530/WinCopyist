package main

import "github.com/BraedynL0530/WinCotypist/internal"

func main() {
	buffer := &internal.RollingBuffer{} // add non zero values in prod
	internal.StartCapture(buffer)       // temp doesnt even have buffer
}
