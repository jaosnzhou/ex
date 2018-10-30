package main

import "fmt"

var animalList = []string{"fly", "spider", "bird", "cat", "dog", "cow", "horse"}

func main() {
	generateSong(animalList)
}
func generateSong(aList []string) {
	for i := 0; i < len(aList); i++ {
		openingLyric(aList[i])
		animalFeature(aList[i])
		for j := i; j >= 1 && j < len(aList)-1; j-- {
			mainLyric(aList[j], aList[j-1])
		}
		endingLyric(aList[i])
	}
}
func openingLyric(animal string) {
	if animal == "fly" {
		fmt.Println("There was an old lady who swallowed a " + animal + ".")
	} else {
		fmt.Println("There was an old lady who swallowed a " + animal + ";")
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
	}
}
func mainLyric(animal1 string, animal2 string) {
	if animal2 == "fly" {
		fmt.Println("She swallowed the " + animal1 + " to catch the " + animal2 + ";")
	} else {
		fmt.Println("She swallowed the " + animal1 + " to catch the " + animal2 + ",")
	}
}
func endingLyric(animal string) {
	if animal == "horse" {
		fmt.Println("...She's dead, of course!")
	} else {
		fmt.Println("I don't know why she swallowed a fly - perhaps she'll die!")
	}
}

//
// There was an old lady who swallowed a fly.                    -----openingLyric
// I don't know why she swallowed a fly - perhaps she'll die!    -----endingLyric
//
// There was an old lady who swallowed a spider;                 -----openingLyric
// That wriggled and wiggled and tickled inside her.             -----animalFeature
// She swallowed the spider to catch the fly;                    -----mainLyric
// I don't know why she swallowed a fly - perhaps she'll die!    -----endingLyric
//
// There was an old lady who swallowed a horse...                -----openingLyric
// ...She's dead, of course!                                     -----animalFeature(endingLyric?)

