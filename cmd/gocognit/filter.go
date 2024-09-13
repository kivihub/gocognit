package main

import (
	"fmt"
	"os"
	"regexp"
)

var ignoreFileContentRegexp *regexp.Regexp
var ignoreFilePathRegexp *regexp.Regexp

func InitIgnoreRegexBeforeAnalyze() {
	if verbose {
		fmt.Printf("--Ignore file path regex is: %s\n", defaultVal(ignoreFilePathExpr, "<not_set>"))
		fmt.Printf("--Ignore file content regex is: %s\n", defaultVal(ignoreFileContentExpr, "<not_set>"))
	}
	ignoreFilePathRegexp = prepareRegexp(ignoreFilePathExpr)
	ignoreFileContentRegexp = prepareRegexp(ignoreFileContentExpr)
}

func IgnoreFileByPath(filePath string) bool {
	ignore := ignoreFilePathRegexp != nil && ignoreFilePathRegexp.MatchString(filePath)
	if verbose && ignore {
		fmt.Printf("IgnoreFileByPath: %s -> %s\n", ignoreFilePathExpr, filePath)
	}
	return ignore
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

	ignore := ignoreFileContentRegexp.MatchString(string(contentBytes))
	if verbose && ignore {
		fmt.Printf("IgnoreFileByContent: %s -> %s\n", ignoreFileContentExpr, filePath)
	}
	return ignore
}

func defaultVal(val, def string) string {
	if val != "" {
		return val
	}
	return def
}
