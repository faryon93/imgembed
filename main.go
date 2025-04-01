package main

// imgembed
// Copyright (C) 2025 Maximilian Pachl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// ---------------------------------------------------------------------------------------
//  imports
// ---------------------------------------------------------------------------------------

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

// ---------------------------------------------------------------------------------------
//  constants
// ---------------------------------------------------------------------------------------

var (
	reImgTag = regexp.MustCompile(`<img([^>]*)src="([^"]*)"`)
)

// ---------------------------------------------------------------------------------------
//  application entry
// ---------------------------------------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: imgembed <html-file>")
		os.Exit(1)
	}

	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	output := reImgTag.ReplaceAllStringFunc(string(buf), func(match string) string {
		submatches := reImgTag.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}

		// Erster sub-match ist attribute vor src, zweiter sub-match ist Inhalt von src=''
		attributes := submatches[1]
		src := submatches[2]

		file, err := os.ReadFile(src)
		if err != nil {
			panic(err)
		}

		// Versuchen den Content-Type anhand der Magic Bytes zu erraten
		contentType := http.DetectContentType(file)
		if contentType == "application/octet-stream" {
			panic("unknown content type for file " + src)
		}

		return fmt.Sprintf(`<img%ssrc="data:%s;base64,%s"`,
			attributes,
			contentType,
			base64.StdEncoding.EncodeToString(file))
	})

	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}
