# Delivery Logistics Library

Пакет для расчета логистических параметров: объемного веса и стоимости доставки.

## Установка
`go get github.com/твой_логин/delivery/logistics`

## Быстрый старт
```go
import "[github.com/твой_логин/delivery/logistics](https://github.com/твой_логин/delivery/logistics)"

func main() {
    vWeight, _ := logistics.VolumetricWeight(50, 40, 30, 5000)
    fmt.Println("Объемный вес:", vWeight)
}
