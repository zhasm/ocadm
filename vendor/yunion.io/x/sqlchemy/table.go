// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlchemy

import (
	"fmt"
	"reflect"
	"strings"

	"yunion.io/x/pkg/utils"
)

type STableSpec struct {
	structType reflect.Type
	name       string
	columns    []IColumnSpec
	indexes    []STableIndex
	contraints []STableConstraint
}

type STable struct {
	spec  *STableSpec
	alias string
}

type STableField struct {
	table *STable
	spec  IColumnSpec
	alias string
}

func NewTableSpecFromStruct(s interface{}, name string) *STableSpec {
	val := reflect.Indirect(reflect.ValueOf(s))
	st := val.Type()
	if st.Kind() != reflect.Struct {
		panic("expect Struct kind")
	}
	table := &STableSpec{
		columns:    []IColumnSpec{},
		name:       name,
		structType: st,
	}
	struct2TableSpec(val, table)
	return table
}

func (ts *STableSpec) Name() string {
	return ts.name
}

func (ts *STableSpec) Columns() []IColumnSpec {
	return ts.columns
}

func (ts *STableSpec) PrimaryColumns() []IColumnSpec {
	ret := make([]IColumnSpec, 0)
	for i := range ts.columns {
		if ts.columns[i].IsPrimary() {
			ret = append(ret, ts.columns[i])
		}
	}
	return ret
}

func (ts *STableSpec) DataType() reflect.Type {
	return ts.structType
}

func (ts *STableSpec) CreateSQL() string {
	cols := make([]string, 0)
	primaries := make([]string, 0)
	indexes := make([]string, 0)
	for _, c := range ts.columns {
		cols = append(cols, c.DefinitionString())
		if c.IsPrimary() {
			primaries = append(primaries, fmt.Sprintf("`%s`", c.Name()))
		}
		if c.IsIndex() {
			indexes = append(indexes, fmt.Sprintf("KEY `ix_%s_%s` (`%s`)", ts.name, c.Name(), c.Name()))
		}
	}
	if len(primaries) > 0 {
		cols = append(cols, fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(primaries, ", ")))
	}
	if len(indexes) > 0 {
		cols = append(cols, indexes...)
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (\n%s\n) ENGINE=InnoDB DEFAULT CHARSET=utf8", ts.name, strings.Join(cols, ",\n"))
}

func (ts *STableSpec) Instance() *STable {
	table := STable{spec: ts, alias: getTableAliasName()}
	return &table
}

func (ts *STableSpec) ColumnSpec(name string) IColumnSpec {
	for _, c := range ts.Columns() {
		if c.Name() == name {
			return c
		}
	}
	return nil
}

func (tbl *STable) Field(name string, alias ...string) IQueryField {
	// name = reflectutils.StructFieldName(name)
	name = utils.CamelSplit(name, "_")
	spec := tbl.spec.ColumnSpec(name)
	if spec == nil {
		panic("column not found: " + name)
	}
	col := STableField{table: tbl, spec: spec}
	if len(alias) > 0 {
		col.Label(alias[0])
	}
	return &col
}

func (tbl *STable) Fields() []IQueryField {
	ret := make([]IQueryField, 0)
	for _, c := range tbl.spec.Columns() {
		ret = append(ret, tbl.Field(c.Name()))
	}
	return ret
}

func (c *STableField) Expression() string {
	if len(c.alias) > 0 {
		return fmt.Sprintf("%s.%s as `%s`", c.table.Alias(), c.spec.Name(), c.alias)
	} else {
		return fmt.Sprintf("%s.%s", c.table.Alias(), c.spec.Name())
	}
}

func (c *STableField) Name() string {
	if len(c.alias) > 0 {
		return c.alias
	} else {
		return c.spec.Name()
	}
}

func (c *STableField) Reference() string {
	return fmt.Sprintf("`%s`.`%s`", c.table.Alias(), c.Name())
}

func (c *STableField) Label(label string) IQueryField {
	if len(label) > 0 && label != c.spec.Name() {
		c.alias = label
	}
	return c
}
