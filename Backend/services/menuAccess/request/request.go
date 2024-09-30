package request

type Request struct {
	PositionId string             `json:"positionId"`
	Access     []AccessPermission `json:"access"`
}

type AccessPermission struct {
	MainMenu string   `json:"mainMenu"`
	SubMenu  []string `json:"subMenu"`
	Action   []string `json:"action"`
}

type RequestUpdate struct {
	Access []AccessPermission `json:"access"`
}
