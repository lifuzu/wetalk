// Copyright 2013 beebbs authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package utils implemented some useful functions.
package utils

import (
	"fmt"
	"os"

	"github.com/Unknwon/com"
	"github.com/Unknwon/goconfig"
)

var (
	Cfg     *goconfig.ConfigFile
	Message *goconfig.ConfigFile
)

// LoadConfig loads configuration file.
func LoadConfig(cfgPath string) (*goconfig.ConfigFile, error) {
	if !com.IsExist(cfgPath) {
		os.Create(cfgPath)
	}

	return goconfig.LoadConfigFile(cfgPath)
}

// I18n translates content to the target language.
func I18n(lang, format string, args ...interface{}) string {
	if lang != "en-US" {
		format = Message.MustValue(lang, format)
	}
	return fmt.Sprintf(format, args...)
}
