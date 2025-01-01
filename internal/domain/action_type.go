package domain

type ActionType string

const (
	ActionCreate ActionType = "create"
	ActionUpdate ActionType = "update"
	ActionDelete ActionType = "delete"
)

func GetActionTypes() []ActionType {
	return []ActionType{ActionCreate, ActionUpdate, ActionDelete}
}
