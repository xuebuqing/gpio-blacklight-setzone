package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("this is blacklight brighten: ", ReadBackLight(0))
	SetBackLight(0, "100")
	fmt.Println("this is blacklight brighten: ", ReadBackLight(0))
}

func SetBackLight(gpioNum int, value string) {
	gpio, err := os.OpenFile(fmt.Sprintf("/sys/class/backlight/backlight%d/brightness", gpioNum), os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer gpio.Close()
	gpio.Write([]byte(value))
}

func ReadBackLight(gpioNum int) string {

	filePath := fmt.Sprintf("/sys/class/backlight/backlight%d/brightness", gpioNum)
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		result := string(contents)
		fmt.Println(result)
		return result
	} else {
		return "backlight err"
	}

}
