package agentscmdapi

var (
	// APIMap - map of api verions
	APIMap = map[string]string{
		"jeep-v0.2.2": "jeep-api-v0.1.0",
	}
	// Comms - available commands
	Comms = map[string]map[string][]string{
		"jeep-api-v0.1.0": map[string][]string{
			"s":          []string{"CMD010", "CMD020"},
			"w":          []string{"CMD010", "CMD020", "CMD012", "CMD022"},
			"a":          []string{"CMD010", "CMD020", "CMD012", "CMD122"},
			"d":          []string{"CMD010", "CMD020", "CMD112", "CMD022"},
			"x":          []string{"CMD010", "CMD020", "CMD112", "CMD122"},
			"l":          []string{"CMDLOK"},
			"disconnect": []string{"CMDDIS"},
		},
	}
)
