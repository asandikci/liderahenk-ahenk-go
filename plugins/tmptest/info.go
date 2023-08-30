package main

func (p plug) Info() map[string]interface{} {
	return map[string]interface{}{
		"name":        "tmptest",
		"version":     "1.0.0",
		"support":     "any",
		"description": "Temporary testing",
		"developer":   "asandikci@aliberksandikci.com.tr",
	}
}
