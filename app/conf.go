package main

import (
	"fmt"
	"time"
)

type Config struct {
	User   user
	Server server
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

// func (c Config) User() user {
// 	return c.User
// }

// func (c Config) Server() server {
// 	return c.Server
// }

func (c Config) FileName() string {
	var now = time.Now()
	var year = now.Format("2006")
	var month = now.Format("01")
	var empNum = c.User.EmployeeNum
	var empName = c.User.Name

	return fmt.Sprintf("勤怠管理%s年(%s月)(%d_%s).xlsx", year, month, empNum, empName)
}

func (c Config) SourcePath() string {
	return c.Server.SourcePath
}

func (c Config) BackupPath() string {
	return c.Server.BackupPath
}
