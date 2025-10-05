package test

import (
	"interface/internal"
	"testing"
)

func TestDogSpeak(t *testing.T) {
	// Test case 1: Basic functionality
	dog := internal.Doge{
		Name:  "Buddy",
		Breed: "Golden Retriever",
		Age:   3,
	}

	expected := "Woof! Woof! Name is Buddy, Breed is Golden Retriever, Age is 3"
	result := dog.Speak()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestDogSpeakWithDifferentValues(t *testing.T) {
	// Test case 2: Different dog values
	dog := internal.Doge{
		Name:  "Max",
		Breed: "Labrador",
		Age:   5,
	}

	expected := "Woof! Woof! Name is Max, Breed is Labrador, Age is 5"
	result := dog.Speak()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestDogSpeakWithEmptyName(t *testing.T) {
	// Test case 3: Empty name
	dog := internal.Doge{
		Name:  "",
		Breed: "Beagle",
		Age:   2,
	}

	expected := "Woof! Woof! Name is , Breed is Beagle, Age is 2"
	result := dog.Speak()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestDogSpeakWithZeroAge(t *testing.T) {
	// Test case 4: Zero age
	dog := internal.Doge{
		Name:  "Puppy",
		Breed: "Poodle",
		Age:   0,
	}

	expected := "Woof! Woof! Name is Puppy, Breed is Poodle, Age is 0"
	result := dog.Speak()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestDogImplementsSpeakerInterface(t *testing.T) {
	// Test case 5: Verify that Doge implements Speaker interface
	var speaker internal.Speaker
	dog := internal.Doge{
		Name:  "Interface Test",
		Breed: "Test Breed",
		Age:   1,
	}

	// This should compile without error if Doge implements Speaker
	speaker = dog
	result := speaker.Speak()

	expected := "Woof! Woof! Name is Interface Test, Breed is Test Breed, Age is 1"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
