package main

func (p plug) Info() map[string]interface{} {
	return map[string]interface{}{
		"name":    "resources",
		"version": "0.1.0",
		"support": map[string]interface{}{
			"linux":   "debian",
			"windows": "10",
		},
		"description": "Resource Usage Information and Controls",
		"developer":   "asandikci@aliberksandikci.com.tr",
	}
}
