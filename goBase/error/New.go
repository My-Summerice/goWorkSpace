package main

import (
    "fmt"
    "math"
    "errors"
)

func circleArea(radius float64) (float64, error) {
    if radius < 0 {
        return 0, errors.New("Area calculation failed, radius is less than zero")
    }
    return math.Pi * radius * radius, nil
}

func main() {
    radius := -3.25
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Area of circle", area)
}
