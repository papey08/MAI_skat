package valid

import "brm-leads/internal/model"

func UpdateLead(upd model.UpdateLead) bool {
	return len([]rune(upd.Title)) <= 200 && len([]rune(upd.Description)) <= 1000
}
