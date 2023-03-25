// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/zema1/watchvuln/ent/vulninformation"
)

// VulnInformation is the model entity for the VulnInformation schema.
type VulnInformation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Severity holds the value of the "severity" field.
	Severity string `json:"severity,omitempty"`
	// Cve holds the value of the "cve" field.
	Cve string `json:"cve,omitempty"`
	// Disclosure holds the value of the "disclosure" field.
	Disclosure string `json:"disclosure,omitempty"`
	// Solutions holds the value of the "solutions" field.
	Solutions string `json:"solutions,omitempty"`
	// References holds the value of the "references" field.
	References []string `json:"references,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VulnInformation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vulninformation.FieldReferences, vulninformation.FieldTags:
			values[i] = new([]byte)
		case vulninformation.FieldID:
			values[i] = new(sql.NullInt64)
		case vulninformation.FieldKey, vulninformation.FieldTitle, vulninformation.FieldDescription, vulninformation.FieldSeverity, vulninformation.FieldCve, vulninformation.FieldDisclosure, vulninformation.FieldSolutions, vulninformation.FieldFrom:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VulnInformation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VulnInformation fields.
func (vi *VulnInformation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vulninformation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vi.ID = int(value.Int64)
		case vulninformation.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				vi.Key = value.String
			}
		case vulninformation.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				vi.Title = value.String
			}
		case vulninformation.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				vi.Description = value.String
			}
		case vulninformation.FieldSeverity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field severity", values[i])
			} else if value.Valid {
				vi.Severity = value.String
			}
		case vulninformation.FieldCve:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cve", values[i])
			} else if value.Valid {
				vi.Cve = value.String
			}
		case vulninformation.FieldDisclosure:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field disclosure", values[i])
			} else if value.Valid {
				vi.Disclosure = value.String
			}
		case vulninformation.FieldSolutions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field solutions", values[i])
			} else if value.Valid {
				vi.Solutions = value.String
			}
		case vulninformation.FieldReferences:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field references", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &vi.References); err != nil {
					return fmt.Errorf("unmarshal field references: %w", err)
				}
			}
		case vulninformation.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &vi.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case vulninformation.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				vi.From = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this VulnInformation.
// Note that you need to call VulnInformation.Unwrap() before calling this method if this VulnInformation
// was returned from a transaction, and the transaction was committed or rolled back.
func (vi *VulnInformation) Update() *VulnInformationUpdateOne {
	return NewVulnInformationClient(vi.config).UpdateOne(vi)
}

// Unwrap unwraps the VulnInformation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vi *VulnInformation) Unwrap() *VulnInformation {
	_tx, ok := vi.config.driver.(*txDriver)
	if !ok {
		panic("ent: VulnInformation is not a transactional entity")
	}
	vi.config.driver = _tx.drv
	return vi
}

// String implements the fmt.Stringer.
func (vi *VulnInformation) String() string {
	var builder strings.Builder
	builder.WriteString("VulnInformation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vi.ID))
	builder.WriteString("key=")
	builder.WriteString(vi.Key)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(vi.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(vi.Description)
	builder.WriteString(", ")
	builder.WriteString("severity=")
	builder.WriteString(vi.Severity)
	builder.WriteString(", ")
	builder.WriteString("cve=")
	builder.WriteString(vi.Cve)
	builder.WriteString(", ")
	builder.WriteString("disclosure=")
	builder.WriteString(vi.Disclosure)
	builder.WriteString(", ")
	builder.WriteString("solutions=")
	builder.WriteString(vi.Solutions)
	builder.WriteString(", ")
	builder.WriteString("references=")
	builder.WriteString(fmt.Sprintf("%v", vi.References))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", vi.Tags))
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(vi.From)
	builder.WriteByte(')')
	return builder.String()
}

// VulnInformations is a parsable slice of VulnInformation.
type VulnInformations []*VulnInformation
