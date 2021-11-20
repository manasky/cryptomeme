// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cryptolist/common"
	"cryptolist/internal/ent/marketchart"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// MarketChart is the model entity for the MarketChart schema.
type MarketChart struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// Chart holds the value of the "chart" field.
	Chart common.MarketChart `json:"chart,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MarketChart) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case marketchart.FieldChart:
			values[i] = new([]byte)
		case marketchart.FieldName, marketchart.FieldCurrency:
			values[i] = new(sql.NullString)
		case marketchart.FieldCreateTime, marketchart.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case marketchart.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MarketChart", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MarketChart fields.
func (mc *MarketChart) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case marketchart.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				mc.ID = *value
			}
		case marketchart.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				mc.CreateTime = value.Time
			}
		case marketchart.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				mc.UpdateTime = value.Time
			}
		case marketchart.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mc.Name = value.String
			}
		case marketchart.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				mc.Currency = value.String
			}
		case marketchart.FieldChart:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field chart", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mc.Chart); err != nil {
					return fmt.Errorf("unmarshal field chart: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MarketChart.
// Note that you need to call MarketChart.Unwrap() before calling this method if this MarketChart
// was returned from a transaction, and the transaction was committed or rolled back.
func (mc *MarketChart) Update() *MarketChartUpdateOne {
	return (&MarketChartClient{config: mc.config}).UpdateOne(mc)
}

// Unwrap unwraps the MarketChart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mc *MarketChart) Unwrap() *MarketChart {
	tx, ok := mc.config.driver.(*txDriver)
	if !ok {
		panic("ent: MarketChart is not a transactional entity")
	}
	mc.config.driver = tx.drv
	return mc
}

// String implements the fmt.Stringer.
func (mc *MarketChart) String() string {
	var builder strings.Builder
	builder.WriteString("MarketChart(")
	builder.WriteString(fmt.Sprintf("id=%v", mc.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(mc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(mc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(mc.Name)
	builder.WriteString(", currency=")
	builder.WriteString(mc.Currency)
	builder.WriteString(", chart=")
	builder.WriteString(fmt.Sprintf("%v", mc.Chart))
	builder.WriteByte(')')
	return builder.String()
}

// MarketCharts is a parsable slice of MarketChart.
type MarketCharts []*MarketChart

func (mc MarketCharts) config(cfg config) {
	for _i := range mc {
		mc[_i].config = cfg
	}
}
