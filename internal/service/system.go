package service

import (
	"LavaPanel/config"
	"LavaPanel/internal/model"
	"encoding/json"
	"fmt"
	"strconv"
)

func getDefaultSystemConfig(panelPort int, authCode int) *config.SystemConfig {
	return &config.SystemConfig{
		PanelPort: strconv.Itoa(panelPort),
		AuthCode:  strconv.Itoa(authCode),
	}
}

func GetSystemConfig() *config.SystemConfig {
	CSystemConfig := new(config.SystemConfig)
	MSystemConfig := new(model.SystemConfig)

	result := model.DB.Model(new(model.SystemConfig)).Where("key = ?", "system").Find(MSystemConfig).RowsAffected
	var panelPort, authCode int
	if result != 1 {
		fmt.Println("")
		fmt.Print("[LavaPanel·Install]: 请输入面板端口: ")
		fmt.Scanln(&panelPort)
		fmt.Print("[LavaPanel·Install]: 请输入面板安全码: ")
		fmt.Scanln(&authCode)
		fmt.Println("")
	}

	dscByte, _ := json.Marshal(getDefaultSystemConfig(panelPort, authCode))
	err := model.DB.Model(new(model.SystemConfig)).
		Where("key = ?", "system").
		Attrs(map[string]interface{}{"key": "system", "value": string(dscByte)}).
		FirstOrCreate(MSystemConfig).Error
	if err != nil {
		panic("[LavaPlane][DB]|[ERROR]:" + err.Error())
	}
	err = json.Unmarshal([]byte(MSystemConfig.Value), CSystemConfig)
	if err != nil {
		panic("[LavaPlane][UNMARSHAL]|[ERROR]:" + err.Error())
	}
	return CSystemConfig
}
