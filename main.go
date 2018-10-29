package main

import "fmt"

var animalList = []string{"fly", "spider", "bird", "cat", "dog", "cow", "horse"}

func main() {
	generateSong(animalList)
}
func generateSong(animalList []string) {
	for i := 0; i < len(animalList); i++ {
		openingLyric(animalList[i])
		animalFeature(animalList[i])
		if animalList[i] == "horse" {
			break
		} else {
			for j := i; j >= 1; j-- {
				mainLyric(animalList[j], animalList[j-1])
			}
		}
		endingLyric()
	}
}
func animalFeature(animal string) {
	switch animal {
	case "fly":
	case "spider":
		fmt.Println("That wriggled and wiggled and tickled inside her.")
	case "bird":
		fmt.Println("How absurd to swallow a bird.")
	case "cat":
		fmt.Println("Fancy that to swallow a cat!")
	case "dog":
		fmt.Println("What a hog, to swallow a dog!")
	case "cow":
		fmt.Println("I don't know how she swallowed a cow!")
	case "horse":
		fmt.Println("...She's dead, of course!")
	}
}
func endingLyric() {
	fmt.Println("I don't know why she swallowed a fly - perhaps she'll die!")
}
func openingLyric(animal string) {
	if animal == "fly" {
		fmt.Println("There was an old lady who swallowed a " + animal + ".")
	} else {
		fmt.Println("There was an old lady who swallowed a " + animal + ";")
	}
}
func mainLyric(animal1 string, animal2 string) {
	if animal2 == "fly" {
		fmt.Println("She swallowed the " + animal1 + " to catch the " + animal2 + ";")
	} else {
		fmt.Println("She swallowed the " + animal1 + " to catch the " + animal2 + ",")
	}

}
