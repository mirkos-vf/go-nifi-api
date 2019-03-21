// this documentation
package go_nifi_api

import (
	"encoding/json"
	"fmt"

	"./types"
)

// WARNING: This file is generated!

// ControllerBulletin this godoc
func (a *app) ControllerBulletin(token, method string) *types.BulletinEntity {
	variables := types.BulletinEntity{}

	url := fmt.Sprintf("%s/%s/%s", a.host, "controller/bulletin", "")
	bytes, _ := a.Do(url, token, method, nil)
	_ = json.Unmarshal(bytes, &variables)

	return &variables

}

// ControllerCluster this godoc
func (a *app) ControllerCluster(token, method string) *types.ClusterEntity {
	variables := types.ClusterEntity{}

	url := fmt.Sprintf("%s/%s/%s", a.host, "controller/cluster", "")
	bytes, _ := a.Do(url, token, method, nil)
	_ = json.Unmarshal(bytes, &variables)

	return &variables

}
