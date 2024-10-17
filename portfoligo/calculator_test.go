package portfoligo

import (
	"portfoli-go/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Calculator_should_complete_ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyA := domain.Holding{
		Name:     "companyA",
	}
	companyB := domain.Holding{
		Name:     "companyB",
	}
	holdingC := domain.Holding{
		Name:     "holdingC",
		Holdings: []domain.SubHolding{
			{Holding: &companyA, Weight: 0.6},
			{Holding: &companyB, Weight: 0.4},},
	}
	holdingB := domain.Holding{
		Name:     "holdingB",
		Holdings: []domain.SubHolding{
			{Holding: &holdingC, Weight: 0.4},
			{Holding: &companyA, Weight: 0.3},
			{Holding: &companyB, Weight: 0.3},},
	}
	holdingA := domain.Holding{
		Name: "holdingA",
		Holdings: []domain.SubHolding{
			{Holding: &holdingB, Weight: 0.5},
			{Holding: &holdingC, Weight: 0.3},
			{Holding: &companyA, Weight: 0.2},
		},
	}

	holdings := map[string]*domain.Holding{
		holdingA.Name: &holdingA,
		holdingB.Name: &holdingB,
		holdingC.Name: &holdingC,
	}

	// target := mocks.NewMockTargetWriter(ctrl)
	// target.EXPECT().WriteInfoLine(gomock.Any()).Return(nil)
	// target.EXPECT().WritePortfolioLine("afund_name", float64(0.1), false).Return(nil)

	config := CalculatorConfig{
		RootHoldingName: "holdingA",
	}
	result, err := config.Calculate(holdings)
	assert.Nil(t, err)

	assert.Equal(t, "companyA", result["companyA"].Name)
	assert.Equal(t, 0.65, result["companyA"].Weight)

	assert.Equal(t, "companyB", result["companyB"].Name)
	assert.Equal(t, 0.35, result["companyB"].Weight)
}
