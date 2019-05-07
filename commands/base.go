// pmm-admin
// Copyright (C) 2018 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package commands provides base commands and helpers.
package commands

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
)

// Ctx is a shared context for all requests.
var Ctx = context.Background()

// Result is a common interface for all command results.
//
// In addition to methods of this interface, result is expected to work with json.Marshal.
type Result interface {
	Result()
	fmt.Stringer
}

// Command is a common interface for all commands.
//
// Command should:
//  * use logrus.Trace/Debug functions for debug logging;
//  * return result on success;
//  * return error on failure.
//
// Command should not:
//  * return both result and error;
//  * exit with logrus.Fatal, os.Exit, etc;
//  * use logrus.Print, logrus.Info and higher levels.
type Command interface {
	Run() (Result, error)
}

type ErrorResponse interface {
	error
	Code() int
}

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func GetError(err ErrorResponse) Error {
	v := reflect.ValueOf(err)
	p := v.Elem().FieldByName("Payload")
	e := p.Elem().FieldByName("Error")
	return Error{
		Code:  err.Code(),
		Error: e.String(),
	}
}

func ParseTemplate(text string) *template.Template {
	t := template.New("").Option("missingkey=error")
	return template.Must(t.Parse(strings.TrimSpace(text)))
}

// RenderTemplate renders given template with given data and returns result as string.
func RenderTemplate(t *template.Template, data interface{}) string {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		logrus.Panicf("Failed to render response.\n%s.\nPlease report this bug.", err)
	}
	return buf.String()
}

func ParseCustomLabels(labels string) (map[string]string, error) {
	if labels == "" {
		return map[string]string{}, nil
	}
	regex := regexp.MustCompile(`(\w+)=(\w+)`)
	result := make(map[string]string)
	parts := strings.Split(labels, ",")
	for _, part := range parts {
		if !regex.MatchString(part) {
			return nil, fmt.Errorf("wrong custom label format")
		}
		submatches := regex.FindStringSubmatch(part)
		result[submatches[1]] = submatches[2]
	}
	return result, nil
}
