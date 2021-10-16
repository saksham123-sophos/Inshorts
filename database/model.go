package database

func getObjectForIndia() map[string]interface{} {
	var doc = make(map[string]interface{})
	doc["state"] = "India"
	var total = make(map[string]float64)
	total["confirmed"] = 0
	total["tested"] = 0
	total["vaccinated1"] = 0
	total["vaccinated2"] = 0
	total["recovered"] = 0
	total["deceased"] = 0
	doc["total"] = total
	return doc
}

var stateCodeToNameMappings = make(map[string]string)

func initStateCodeMappings() {
	stateCodeToNameMappings["AN"] = "Andaman and Nicobar Islands"
	stateCodeToNameMappings["AP"] = "Andhra Pradesh"
	stateCodeToNameMappings["AR"] = "Arunachal Pradesh"
	stateCodeToNameMappings["AS"] = "Assam"
	stateCodeToNameMappings["BR"] = "Bihar"
	stateCodeToNameMappings["CH"] = "Chandigarh"
	stateCodeToNameMappings["CT"] = "Chhattisgarh"
	stateCodeToNameMappings["DN"] = "Dadra and Nagar Haveli"
	stateCodeToNameMappings["DD"] = "Daman and Diu"
	stateCodeToNameMappings["DL"] = "Delhi"
	stateCodeToNameMappings["GA"] = "Goa"
	stateCodeToNameMappings["GJ"] = "Gujarat"
	stateCodeToNameMappings["HR"] = "Haryana"
	stateCodeToNameMappings["HP"] = "Himachal Pradesh"
	stateCodeToNameMappings["JK"] = "Jammu and Kashmir"
	stateCodeToNameMappings["JH"] = "Jharkhand"
	stateCodeToNameMappings["KA"] = "Karnataka"
	stateCodeToNameMappings["KL"] = "Kerala"
	stateCodeToNameMappings["LD"] = "Lakshadweep"
	stateCodeToNameMappings["MP"] = "Madhya Pradesh"
	stateCodeToNameMappings["MH"] = "Maharashtra"
	stateCodeToNameMappings["MN"] = "Manipur"
	stateCodeToNameMappings["ML"] = "Meghalaya"
	stateCodeToNameMappings["MZ"] = "Mizoram"
	stateCodeToNameMappings["NL"] = "Nagaland"
	stateCodeToNameMappings["OR"] = "Odisha"
	stateCodeToNameMappings["PY"] = "Puducherry"
	stateCodeToNameMappings["PB"] = "Punjab"
	stateCodeToNameMappings["RJ"] = "Rajasthan"
	stateCodeToNameMappings["SK"] = "Sikkim"
	stateCodeToNameMappings["TN"] = "Tamil Nadu"
	stateCodeToNameMappings["TG"] = "Telangana"
	stateCodeToNameMappings["TR"] = "Tripura"
	stateCodeToNameMappings["UP"] = "Uttar Pradesh"
	stateCodeToNameMappings["UT"] = "Uttarakhand"
	stateCodeToNameMappings["WB"] = "West Bengal"
}

func getStateName(stateCode string) string {
	return stateCodeToNameMappings[stateCode]
}
