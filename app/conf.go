package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	User   user
	Server server
	File file
	Templates templates
}

type user struct {
	Name        string `toml: "name"`
	EmployeeNum int    `toml: "employeeNum"`
	Apartment   string `toml: "apartment"`
}

type server struct {
	SourcePath string `toml: "sourcePath"`
	BackupPath string `toml: "backupPath"`
}

type file struct {
	Name string		`toml: name`
}

type template struct {
	Pattern string 	`toml: "pattern"`
	Value string 	`toml: valie`
}

type templates []template

func (c Config) FileName() string {
	var now = time.Now()
	var year = now.Format("2006")
	var month = now.Format("01")
	var empNum = c.User.EmployeeNum
	var empName = c.User.Name

	return fmt.Sprintf("勤怠管理%s年(%s月)(%d_%s).xlsx", year, month, empNum, empName)
}

func (c Config) SourcePath() string {
	path, err := parsePath(c.Server.SourcePath)
	if err != nil {
		log.Fatal(err)
	}

	if !strings.HasSuffix(path, "/") {
		suffixAddSlash(&path)
	}

	return path
}

func (c Config) BackupPath() string {
	path, err := parsePath(c.Server.BackupPath)
	if err != nil {
		log.Fatal(err)
	}

	if !strings.HasSuffix(path, "/") {
		suffixAddSlash(&path)
	}

	return path
}

func parsePath(path string) (string, error) {
	reg, err := regexp.Compile(`({[Y|W]{4}}|{[M|D|W]{1,2}})`)
	if err != nil {
		return "", err
	}

	var result = path
	for _, match := range reg.FindAllString(result, -1) {
		if reg.MatchString(result) {
			parseDate(&result, match)
		}
	}

	return result, nil
}

func parseDate(src *string, format string) {
	var now = time.Now()
	switch format {
	case "{YYYY}":
		*src = strings.ReplaceAll(*src, format, now.Format("2006"))
	case "{MM}":
		*src = strings.ReplaceAll(*src, format, now.Format("01"))
	case "{M}":
		*src = strings.ReplaceAll(*src, format, now.Format("1"))
	case "{DD}":
		*src = strings.ReplaceAll(*src, format, now.Format("02"))
	case "{D}":
		*src = strings.ReplaceAll(*src, format, now.Format("2"))
	case "{W}":
		*src = strings.ReplaceAll(*src, format, dayOfTheWeekENtoJP(now.Format("Mon")))
	case "{WW}":
		*src = strings.ReplaceAll(*src, format, now.Format("Mon"))
	case "{WWWW}":
		*src = strings.ReplaceAll(*src, format, now.Format("Monday"))
	}
}

func dayOfTheWeekENtoJP(en string) string {
	switch strings.ToLower(en) {
	case "sun", "sunday":
		return "日"
	case "mon", "monday":
		return "月"
	case "tue", "tuesday":
		return "火"
	case "wed", "wednesday":
		return "水"
	case "thu", "thursday":
		return "木"
	case "fri", "friday":
		return "金"
	case "sat", "saturday":
		return "土"
	}
	return en
}

func suffixAddSlash(s *string) {
	*s = (*s)[:len(*s)] + "/"
}

func (c Config) parseTemplates(path string) (string, error) {
	var result = path
	for _, template := range c.Templates {
		reg, err := regexp.Compile(fmt.Sprintf(`({%s})`, template.Pattern))
		if err != nil {
			return "", err
		}

		for _, match := range reg.FindAllString(result, -1) {
			if reg.MatchString(c.File.Name) {
				result = strings.ReplaceAll(result, match, template.Value)
			}
		}
	}
fmt.Println(result)
	return result, nil
}