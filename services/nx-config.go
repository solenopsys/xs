package services

import (
	"encoding/json"
	"os"
	"xs/utils"
)

type Packages struct {
}

func InjectConfToJson(c *ConfLoader, fileName string) {

	existingJSON, err := utils.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var confData map[string]any
	err = json.Unmarshal([]byte(existingJSON), &confData)
	if err != nil {
		panic(err)
	}

	modulesConf := make(map[string][]string)

	groups := *c.data
	for _, group := range groups {
		for _, module := range group.Modules {

			path := c.targetDir + "/" + group.Dir + "/" + module.Directory
			tsFile := path + "/src/index.ts"

			println("Inject to config:", module.Name, tsFile)
			modulesConf[module.Name] = []string{tsFile}
		}
	}

	confData["paths"] = modulesConf

	newJSON, err := json.Marshal(confData)
	if err != nil {
		panic(err)
	}

	os.WriteFile(fileName, newJSON, 0644)
}
