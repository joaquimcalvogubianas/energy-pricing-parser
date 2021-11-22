package resources_manager

import (
	"os"
)

func GetTestResource(filePath string) []byte {
	filePathMustPointToSamePackage(filePath)
	fc, err := os.ReadFile(filePath)

	if err != nil {
		panic("Something went wrong trying to get test resource using location " + filePath)
	}

	return fc
}

func filePathMustPointToSamePackage(filePath string) {
	firstCharacter := filePath[0:0]
	if firstCharacter == "/" {
		panic("Only package resources can be obtained!")
	}
}
