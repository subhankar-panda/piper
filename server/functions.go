package main

import (
    "strings"
    "html/template"
)

var (
    templateMap = template.FuncMap{
        "Upper" : func(s string) string {
            return strings.ToUpper(s)
        },
    }
)
