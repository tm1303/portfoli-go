package portfoligo

import (
	"portfoli-go/domain"
	"portfoli-go/portfoligo/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Formater_should_complete_ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	portflio := map[string]domain.PortfolioLine{
		"afund": {Name: "afund_name", Weight: 0.1},
	}

	target := mocks.NewMockTargetWriter(ctrl)
	target.EXPECT().WriteInfoLine(gomock.Any()).Return(nil)
	target.EXPECT().WritePortfolioLine("afund_name", float64(0.1), false).Return(nil)

	config := NewOrderedFormatter(target, []string{})
	err := config.Output(portflio)
	assert.Nil(t, err)
}

func Test_Formater_should_order_results(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	portflio := map[string]domain.PortfolioLine{
		"mid_fund":   {Name: "mid_fund_name", Weight: 0.1},
		"small_fund": {Name: "small_fund_name", Weight: 0.01},
		"big_fund":   {Name: "big_fund_name", Weight: 0.5},
	}

	target := mocks.NewMockTargetWriter(ctrl)
	target.EXPECT().WriteInfoLine(gomock.Any()).Return(nil)

	// one way to achive this, but it's a bit smelly :/
	// in retro spec would be nicer to assert on the overall result rather than the call order
	call1 := target.EXPECT().WritePortfolioLine("big_fund_name", float64(0.5), false).Return(nil)
	call2 := target.EXPECT().WritePortfolioLine("mid_fund_name", float64(0.1), false).Return(nil).After(call1)
	target.EXPECT().WritePortfolioLine("small_fund_name", float64(0.01), false).Return(nil).After(call2)

	config := NewOrderedFormatter(target, []string{})
	err := config.Output(portflio)
	assert.Nil(t, err)
}


func Test_Formater_should_highlight_two_funds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	portflio := map[string]domain.PortfolioLine{
		"mid_fund":   {Name: "mid_fund_name", Weight: 0.1},
		"small_fund": {Name: "small_fund_name", Weight: 0.01},
		"big_fund":   {Name: "big_fund_name", Weight: 0.5},
	}

	target := mocks.NewMockTargetWriter(ctrl)
	target.EXPECT().WriteInfoLine(gomock.Any()).Return(nil)

	// still smelly >:)
	call1 := target.EXPECT().WritePortfolioLine("big_fund_name", float64(0.5), true).Return(nil)
	call2 := target.EXPECT().WritePortfolioLine("mid_fund_name", float64(0.1), false).Return(nil).After(call1)
	target.EXPECT().WritePortfolioLine("small_fund_name", float64(0.01), true).Return(nil).After(call2)

	config := NewOrderedFormatter(target, []string{"big_fund_name", "small_fund_name"})
	err := config.Output(portflio)
	assert.Nil(t, err)
}
