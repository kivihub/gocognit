package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

var ignoreFileContentRegexp *regexp.Regexp

func Init() {
	var err error
	ignoreFileContentRegexp, err = prepareRegexp(ignoreFileContentExpr)
	fmt.Printf("Ignore file content regex is: %s\n", ignoreFileContentExpr)
	if err != nil {
		log.Fatal(err)
	}
}

func IgnoreFileByContent(filePath string) bool {
	if ignoreFileContentRegexp == nil {
		return false
	}

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件时出错:%s, %v\n", filePath, err)
		return false
	}

	return ignoreFileContentRegexp.MatchString(string(contentBytes))
}
