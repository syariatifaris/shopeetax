package tax

import (
	taxmodel "github.com/syariatifaris/shopeetax/app/model/tax"
)

//Holder holds persistent data
type Holder struct {
	Taxes      []*taxmodel.Tax
	Components []*taxmodel.Component
	TaxRules   map[int64]*taxmodel.Tax
}

func (h *Holder) mapRules() {
	h.TaxRules = make(map[int64]*taxmodel.Tax)
	for _, rule := range h.Taxes {
		h.TaxRules[rule.ID] = rule
	}
}
