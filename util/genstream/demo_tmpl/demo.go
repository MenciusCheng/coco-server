package demo_tmpl

import _ "embed"

//go:embed debug.tmpl
var Debug string

//go:embed crud_model.tmpl
var CrudModel string

//go:embed crud_dao.tmpl
var CrudDao string

//go:embed crud_cache.tmpl
var CrudCache string
