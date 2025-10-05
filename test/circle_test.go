package test

import (
	"interface/internal"
	"testing"
)

func TestArea(t *testing.T) {
	// Test case 1: Basic functionality
	circle := internal.Circle{
		Radius: 5,
	}

	expected := 78.53975 // π * r^2 = 3.14159 * 5 * 5
	result := circle.Area()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}

func TestPerimeter(t *testing.T) {
	// Test case 2: Basic functionality
	circle := internal.Circle{
		Radius: 5,
	}

	expected := 31.4159 // 2 * π * r = 2 * 3.14159 * 5
	result := circle.Perimeter()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}

func TestAreaWithZeroRadius(t *testing.T) {
	// Test case 3: Zero radius
	circle := internal.Circle{
		Radius: 0,
	}

	expected := 0.0
	result := circle.Area()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}

func TestPerimeterWithZeroRadius(t *testing.T) {
	// Test case 4: Zero radius
	circle := internal.Circle{
		Radius: 0,
	}

	expected := 0.0
	result := circle.Perimeter()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}

func TestAreaWithNegativeRadius(t *testing.T) {
	// Test case 5: Negative radius
	circle := internal.Circle{
		Radius: -5,
	}

	expected := 78.53975 // Area should be positive even for negative radius
	result := circle.Area()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}

func TestPerimeterWithNegativeRadius(t *testing.T) {
	// Test case 6: Negative radius
	circle := internal.Circle{
		Radius: -5,
	}

	expected := 31.4159 // Perimeter should be positive even for negative radius
	result := circle.Perimeter()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}

func TestAreaWithLargeRadius(t *testing.T) {
	// Test case 7: Large radius
	circle := internal.Circle{
		Radius: 1e6,
	}

	expected := 3.14159e12 // π * r^2 = 3.14159 * (1e6)^2
	result := circle.Area()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}

func TestPerimeterWithLargeRadius(t *testing.T) {
	// Test case 8: Large radius
	circle := internal.Circle{
		Radius: 1e6,
	}

	expected := 6.28318e6 // 2 * π * r = 2 * 3.14159 * 1e6
	result := circle.Perimeter()

	if result != expected {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}
