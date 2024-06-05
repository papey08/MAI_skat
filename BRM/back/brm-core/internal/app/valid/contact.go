package valid

import "brm-core/internal/model"

func UpdateContact(upd model.UpdateContact) bool {
	return len([]rune(upd.Notes)) <= 500
}
