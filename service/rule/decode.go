package main

//func insensitiviseMap(m map[string]interface{}) {
//	for key, val := range m {
//		switch val.(type) {
//		case map[interface{}]interface{}:
//			// nested map: cast and recursively insensitivise
//			val = cast.ToStringMap(val)
//			insensitiviseMap(val.(map[string]interface{}))
//		case map[string]interface{}:
//			// nested map: recursively insensitivise
//			insensitiviseMap(val.(map[string]interface{}))
//		}
//
//		lower := strings.ToLower(key)
//		if key != lower {
//			// remove old key (not lower-cased)
//			delete(m, key)
//		}
//		// update map
//		m[lower] = val
//	}
//}
//
//func main() {
//	jsonEntry := `
//	{
//		"value": {
//			"Diags": ["R79.000x006", "XC02CAC0.12"],
//			"MedicalType": "住院",
//			"Items": [
//				{
//					"Code": "XA12CBL149"
//				},
//				{
//					"Code": "R79.000x006"
//				}
//			],
//			"Age": 121
//		},
//		"type": "MedicalRecord"
//	}`
//	entry := make(map[string]interface{})
//	if err := json.Unmarshal([]byte(jsonEntry), &entry); err != nil {
//		panic(err)
//	}
//	insensitiviseMap(entry)
//	fmt.Println(entry["Items"])
//}
