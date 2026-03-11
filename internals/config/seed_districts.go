package config

import _ "embed"

//go:embed data/seed_districts.sql
var seedDistrictsSQL []byte
