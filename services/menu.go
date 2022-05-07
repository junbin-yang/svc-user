package services

import (
	"context"
	"github.com/junbin-yang/golib/json"
	"github.com/junbin-yang/golib/logger"
	"github.com/junbin-yang/see"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"svc-user/model"
	"svc-user/proto"
)

func (this *Svr) GetAdminMenus(c *see.Context) Response {
	menus := []map[string]interface{}{
		map[string]interface{}{
			"ID":        1,
			"parentId":  0,
			"path":      "dashboard",
			"name":      "dashboard",
			"hidden":    false,
			"component": "view/dashboard/index.vue",
			"sort":      1,
			"meta": map[string]interface{}{
				"keepAlive":   false,
				"defaultMenu": false,
				"title":       "工作台",
				"icon":        "odometer",
				"closeTab":    false,
			},
			"authoritys": nil,
			"menuBtn":    nil,
			"menuId":     "1",
			"children":   nil,
			"parameters": []map[string]interface{}{},
			"btns":       nil,
		},
		map[string]interface{}{
			"ID":        2,
			"parentId":  0,
			"path":      "person",
			"name":      "person",
			"hidden":    true,
			"component": "view/person/person.vue",
			"sort":      4,
			"meta": map[string]interface{}{
				"keepAlive":   false,
				"defaultMenu": false,
				"title":       "个人信息",
				"icon":        "message",
				"closeTab":    false,
			},
			"authoritys": nil,
			"menuBtn":    nil,
			"menuId":     "2",
			"children":   nil,
			"parameters": []map[string]interface{}{},
			"btns":       nil,
		},
		map[string]interface{}{
			"ID":        3,
			"parentId":  0,
			"path":      "applist",
			"name":      "applist",
			"hidden":    false,
			"component": "view/appAdmin/app.vue",
			"sort":      3,
			"meta": map[string]interface{}{
				"keepAlive":   true,
				"defaultMenu": false,
				"title":       "应用管理",
				"icon":        "platform",
				"closeTab":    false,
			},
			"authoritys": nil,
			"menuBtn":    nil,
			"menuId":     "3",
			"children":   nil,
			"parameters": []map[string]interface{}{},
			"btns":       nil,
		},
		map[string]interface{}{
			"ID":        4,
			"parentId":  0,
			"path":      "userlist",
			"name":      "userlist",
			"hidden":    false,
			"component": "view/userAdmin/user.vue",
			"sort":      3,
			"meta": map[string]interface{}{
				"keepAlive":   true,
				"defaultMenu": false,
				"title":       "用户管理",
				"icon":        "coordinate",
				"closeTab":    false,
			},
			"authoritys": nil,
			"menuBtn":    nil,
			"menuId":     "4",
			"children":   nil,
			"parameters": []map[string]interface{}{},
			"btns":       nil,
		},
		map[string]interface{}{
			"ID":        5,
			"parentId":  0,
			"path":      "appDetail/:id",
			"name":      "appDetail",
			"hidden":    true,
			"component": "view/appAdmin/detail.vue",
			"sort":      4,
			"meta": map[string]interface{}{
				"keepAlive":   false,
				"defaultMenu": false,
				"title":       "角色管理",
				"icon":        "order",
				"closeTab":    false,
			},
			"authoritys": nil,
			"menuBtn":    nil,
			"menuId":     "5",
			"children":   nil,
			"parameters": []map[string]interface{}{},
			"btns":       nil,
		},
	}
	return RespData(see.H{"menus": menus})
}

func (this *Svr) GetAppMenus(ctx context.Context, in *proto.NullRequest) (*proto.MenusReply, error) {
	appid := getMetaAppId(ctx)
	return &proto.MenusReply{Node: getMenus(0, appid, 0)}, nil
}

func (this *Svr) Ping(ctx context.Context, in *proto.NullRequest) (*proto.Reply, error) {
	return ProtoSucceed(), nil
}

func getMenus(pid uint, appid string, userId uint) []*proto.MenusReply_MenusNode {
	outInterFace := []*proto.MenusReply_MenusNode{}
	menus := []model.Menu{}
	model.M.Where("p_id = ? AND app_key = ?", pid, appid).Find(&menus)
	for _, node := range menus {
		children := []*proto.MenusReply_MenusNode{}
		var count int64
		model.M.Model(&model.Menu{}).Where("p_id = ? AND app_key = ?", node.ID, appid).Count(&count)
		if count > 0 {
			children = getMenus(node.ID, appid, userId)
			// 如果下级是叶子节点且没有权限。当前节点设置为false
			if node.Status {
				childrenStatus := []bool{}
				for i, next := range children {
					if len(next.Actions) == 0 && len(next.Children) == 0 {
						children[i].Status = false
						childrenStatus = append(childrenStatus, false)
					}
				}
				if len(childrenStatus) == len(children) {
					node.Status = false
				}
			}
		}

		actions := []string{}
		json.JsonToObject(node.Actions, &actions)
		permissions := map[string]uint32{}
		var withPermissions bool
		if userId > 0 {
			permissions = GetUserPermissions(appid, userId)
			withPermissions = true
		}

		item := &proto.MenusReply_MenusNode{
			Code:     node.Code,
			Name:     node.Name,
			Sequence: node.Sequence,
			Uri:      node.Uri,
			Actions:  CalculateBit(actions, withPermissions, node.Code, permissions),
			Remark:   node.Remark,
			Status:   node.Status,
			Children: children,
		}
		outInterFace = append(outInterFace, item)
	}
	return outInterFace
}

func CalculateBit(actions []string, withPermissions bool, code string, permissions map[string]uint32) map[string]uint32 {
	data := map[string]uint32{}
	var i uint32 = 1
	for _, v := range actions {
		if withPermissions {
			if permissions[code]&i > 0 {
				data[v] |= i
			}
		} else {
			data[v] |= i
		}
		i *= 2
	}
	return data
}

func (this *Svr) DisableMenu(ctx context.Context, in *proto.DisableMenuRequest) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)
	result := model.M.Model(model.Menu{}).Where("code = ? AND app_key = ?", in.Code, appid).Updates(map[string]interface{}{"status": in.Status})
	if result.Error != nil {
		return nil, grpc.Errorf(codes.DataLoss, result.Error.Error())
	}
	return ProtoSucceed(), nil
}

func (this *Svr) PermissionsTpl(ctx context.Context, in *proto.PermissionsTplRequest) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)
	repeatedCode := duplicateInArray(findMenuCode(in.Node))
	if repeatedCode != "" {
		return nil, grpc.Errorf(codes.DataLoss, "重复的Code["+repeatedCode+"],同一份菜单数据要求Code字段唯一")
	}
	model.M.Unscoped().Where("app_key = ?", appid).Delete(model.Menu{})
	err := importMenus(in.Node, 0, appid)

	if err != nil {
		return nil, grpc.Errorf(codes.DataLoss, err.Error())
	}

	str, _ := json.ObjectToJson(in)
	logger.Info(str)
	return ProtoSucceed(), nil
}

func importMenus(in []*proto.PermissionsTplRequest_PermissionsNode, pid uint, appid string) error {
	for _, node := range in {
		actionStr, _ := json.ObjectToJson(node.Actions)
		menu := model.Menu{
			PID:      pid,
			Code:     node.Code,
			Name:     node.Name,
			Uri:      node.Uri,
			Remark:   node.Remark,
			Sequence: node.Sequence,
			Actions:  actionStr,
			AppKey:   appid,
		}
		e := menu.Create()
		if e != nil {
			return e
		}

		if node.Children != nil {
			e := importMenus(node.Children, menu.ID, appid)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

func findMenuCode(in []*proto.PermissionsTplRequest_PermissionsNode) []string {
	out := []string{}
	for _, node := range in {
		out = append(out, node.Code)
		if node.Children != nil {
			children := findMenuCode(node.Children)
			out = append(out, children...)
		}
	}
	return out
}

func duplicateInArray(arrInit []string) string {
	arrTop := []string{}

	if len(arrInit) == 0 {
		return ""
	}

	arrMap1 := make(map[string]int)
	for _, value := range arrInit {
		if arrMap1[value] != 0 {
			arrMap1[value]++
		} else {
			arrMap1[value] = 1
		}
	}

	arrMap2 := arrMap1
	for _ = range arrMap1 {
		var maxCountKey string
		var maxCountVal int = 1
		for key, val := range arrMap2 {
			if val > maxCountVal {
				maxCountVal = val
				maxCountKey = key
			}
		}
		arrTop = append(arrTop, maxCountKey)
		if len(arrTop) >= 1 {
			return maxCountKey
		}
		delete(arrMap2, maxCountKey)
	}
	return ""
}
