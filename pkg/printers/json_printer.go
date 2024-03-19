/*
Copyright 2022 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package printers

import (
	"fmt"
	"io"

	"github.com/kiranmenon/jsonprinter/pkg/jsonwriter"
)

func JsonPrinter(header []string) PrinterFunc {
	return func(w io.Writer, args ...interface{}) error {
		var rows [][]string
		for _, arg := range args {
			switch arg := arg.(type) {
			case []interface{}:
				for _, v := range arg {
					s, ok := v.([][]string)
					if !ok {
						return fmt.Errorf("unsupported type %T", v)
					}
					rows = append(rows, s...)
				}
			default:
				return fmt.Errorf("unsupported type %T", arg)
			}
		}

		jsonDoc := jsonwriter.NewWriter(w)
		jsonDoc.SetHeader(header)
		jsonDoc.AppendBulk(rows)
		jsonDoc.Render()

		return nil
	}
}
