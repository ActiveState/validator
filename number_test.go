package validator_test

import (
	"strings"
	"testing"

	"github.com/go-json-schema/schema"
	"github.com/go-json-schema/schema/draft04"
	"github.com/go-json-schema/validator"
	"github.com/go-json-schema/validator/builder"
	"github.com/stretchr/testify/assert"
)

func TestNumericConstraint(t *testing.T) {
	var _ validator.NumericConstraint = validator.Integer()
	var _ validator.NumericConstraint = validator.Number()
}

func TestNumberFromSchema(t *testing.T) {
	const src = `{
  "type": "number",
  "minimum": 5,
  "maximum": 15,
  "default": 10
}`

	s, err := schema.Parse(strings.NewReader(src), schema.WithSchemaID(draft04.SchemaID))
	if !assert.NoError(t, err, "schema.Parse should succeed") {
		return
	}

	b := builder.New()
	v, err := b.Build(s)
	if !assert.NoError(t, err, "Builder.Build should succeed") {
		return
	}

	c2 := validator.Number()
	c2.Default(float64(10)).Maximum(15).Minimum(5)
	if !assert.Equal(t, c2, v.Root(), "constraints are equal") {
		return
	}
}

func TestNumber(t *testing.T) {
	c := validator.Number()
	c.Default(float64(10)).Maximum(15)

	if !assert.True(t, c.HasDefault(), "HasDefault is true") {
		return
	}

	if !assert.Equal(t, c.DefaultValue(), float64(10), "DefaultValue returns expected value") {
		return
	}

	var s float64
	if !assert.NoError(t, c.Validate(s), "validate should succeed") {
		return
	}

	c.Minimum(5)
	if !assert.Error(t, c.Validate(s), "validate should fail") {
		return
	}

	s = 10
	if !assert.NoError(t, c.Validate(s), "validate should succeed") {
		return
	}
}
