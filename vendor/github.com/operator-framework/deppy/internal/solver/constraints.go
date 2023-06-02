package solver

import (
	"github.com/go-air/gini/z"

	"github.com/operator-framework/deppy/pkg/deppy"
)

// zeroConstraint is returned by ConstraintOf in error cases.
type zeroConstraint struct{}

var _ deppy.Constraint = zeroConstraint{}

func (zeroConstraint) String(subject deppy.Identifier) string {
	return ""
}

func (zeroConstraint) Apply(lm deppy.LitMapping, subject deppy.Identifier) z.Lit {
	return z.LitNull
}

func (zeroConstraint) Order() []deppy.Identifier {
	return nil
}

func (zeroConstraint) Anchor() bool {
	return false
}
