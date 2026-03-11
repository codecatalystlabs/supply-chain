package config

import _ "embed"

//go:embed data/seed_facilities.sql
var seedFacilitiesSQL []byte

