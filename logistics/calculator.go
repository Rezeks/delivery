// Package logistics предоставляет функции для расчета логистических параметров,
// таких как объемный вес, стоимость доставки и формирование отчетов.
package logistics

import (
	"fmt"
)

// VolumetricWeight рассчитывает объемный вес на основе размеров и делителя.
// Возвращает ошибку, если размеры или делитель меньше или равны нулю.
func VolumetricWeight(length, width, height, divisor float64) (float64, error) {
	if length <= 0 || width <= 0 || height <= 0 {
		return 0, fmt.Errorf("dimensions must be positive: got %.2f x %.2f x %.2f", length, width, height)
	}
	if divisor <= 0 {
		return 0, fmt.Errorf("divisor must be greater than zero: got %.2f", divisor)
	}
	return (length * width * height) / divisor, nil
}

// DeliveryCost вычисляет итоговую стоимость, выбирая наибольший вес между
// физическим и объемным.
func DeliveryCost(weight, volumeWeight, pricePerKg float64) (float64, error) {
	if weight < 0 || volumeWeight < 0 || pricePerKg < 0 {
		return 0, fmt.Errorf("weight and price cannot be negative")
	}

	actualWeight := weight
	if volumeWeight > weight {
		actualWeight = volumeWeight
	}

	return actualWeight * pricePerKg, nil
}

// ApplyRemoteAreaFee добавляет фиксированную комиссию к стоимости по указателю.
// Это изменяет исходную переменную в вызывающем коде.
func ApplyRemoteAreaFee(cost *float64, fee float64) error {
	if cost == nil {
		return fmt.Errorf("nil pointer: cost must be a valid float64 pointer")
	}
	if fee < 0 {
		return fmt.Errorf("fee cannot be negative: %.2f", fee)
	}
	*cost += fee
	return nil
}

// FormatDeliveryReport создает отформатированную строку с информацией о доставке.
// Использует fmt.Sprintf для выравнивания данных.
func FormatDeliveryReport(city string, weight, cost float64) (string, error) {
	if city == "" {
		return "", fmt.Errorf("city name is empty")
	}
	return fmt.Sprintf("ID: [City: %-12s] | Weight: %8.3f kg | Total: %10.2f USD", city, weight, cost), nil
}
