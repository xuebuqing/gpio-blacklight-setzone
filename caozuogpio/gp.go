package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	var x int
	var y int
	var z string
	for {
		fmt.Println("1:ExportGpio(y)        2:UnexportGpio(y)")
		fmt.Println("3:ReadGpioDirection(y) 4:SetGpioDirection(y, z)")
		fmt.Println("5:ReadGpioValue(y)     6:SetGpioValue(y, z)")
		fmt.Println("请输入function num(x),input num(y),value or direction(z)")
		fmt.Scanln(&x, &y, &z)
		fmt.Printf("function num:%d,input num:%d,value:%s\n", x, y, z)
		switch x {
		case 1:
			ExportGpio(y)
		case 2:
			UnexportGpio(y)
		case 3:
			fmt.Println(ReadGpioDirection(y))
		case 4:
			SetGpioDirection(y, z)
		case 5:
			ReadGpioValue(y)
		case 6:
			SetGpioValue(y, z)
		default:
			fmt.Println("function num err")
		}
	}

}
func UnexportGpio(gpioNum int) {
	export, err := os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("failed to open gpio unexport file for writing\n")
		os.Exit(1)
	}
	defer export.Close()
	export.Write([]byte(strconv.Itoa(gpioNum)))
}

func ExportGpio(gpioNum int) {
	export, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("failed to open gpio export file for writing\n")
		os.Exit(1)
	}
	defer export.Close()
	export.Write([]byte(strconv.Itoa(gpioNum)))
}

func SetGpioDirection(gpioNum int, direction string) {
	gpio, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", gpioNum), os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer gpio.Close()

	gpio.Write([]byte(direction))
}

func ReadGpioDirection(gpioNum int) string {
	filePath := fmt.Sprintf("/sys/class/gpio/gpio%d/direction", gpioNum)
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		result := string(contents)
		return result
	}
	return "ReadGpioDirection err"
}

func SetGpioValue(gpioNum int, value string) {

	if strings.HasPrefix(ReadGpioDirection(gpioNum), "out") {
		gpio, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/value", gpioNum), os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("open file err", err)
		}
		defer gpio.Close()

		gpio.Write([]byte(value))
	} else {
		fmt.Println("please set the direction :out firstly\n")
	}

}

func ReadGpioValue(gpioNum int) string {

	filePath := fmt.Sprintf("/sys/class/gpio/gpio%d/value", gpioNum)
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		result := string(contents)
		fmt.Println(result)
		return result
	} else {
		return "ReadGpioValue err"
	}

}
