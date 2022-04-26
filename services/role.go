package services

import (
	"context"
	"github.com/junbin-yang/golib/json"
	"github.com/junbin-yang/see"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	structpb "google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
	"strconv"
	"svc-user/model"
	"svc-user/proto"
)

func (this *Svr) CreateRole(ctx context.Context, in *proto.RoleInfo) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)

	if find := model.M.Where("name = ? AND app_key = ?", in.Name, appid).First(&model.Role{}).RowsAffected; find > 0 {
		return nil, grpc.Errorf(codes.DataLoss, "已存在同名角色")
	}

	authoritys := []model.RoleAuthority{}
	for code, value := range in.Permissions {
		authoritys = append(authoritys, model.RoleAuthority{Code: code, ActionValue: value})
	}
	role := model.Role{
		Name:           in.Name,
		Remark:         in.Remark,
		AppKey:         appid,
		Status:         &in.Status,
		RoleAuthoritys: authoritys,
	}
	e := role.Create()
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	return ProtoSucceed(), nil
}

func (this *Svr) UpdateRole(ctx context.Context, in *proto.UpdateRoleRequest) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)
	authoritys := []model.RoleAuthority{}
	for code, value := range in.RoleInfo.Permissions {
		authoritys = append(authoritys, model.RoleAuthority{Code: code, ActionValue: value})
	}
	role := model.Role{
		Name:           in.RoleInfo.Name,
		Remark:         in.RoleInfo.Remark,
		AppKey:         appid,
		Status:         &in.RoleInfo.Status,
		Sequence:       in.RoleInfo.Sequence,
		RoleAuthoritys: authoritys,
	}
	e := role.Update(uint(in.Id))
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	return ProtoSucceed(), nil
}

func (this *Svr) DeleteRole(ctx context.Context, in *proto.DeleteRoleRequest) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)

	var total int64
	model.M.Model(&model.RoleUserMap{}).Where("app_key = ? AND role_id = ?", appid, in.Id).Count(&total)
	if total > 0 {
		return nil, grpc.Errorf(codes.DataLoss, "请先移除角色下用户")
	}

	role := new(model.Role)
	e := role.Delete(uint(in.Id), appid)
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	return ProtoSucceed(), nil
}

func (this *Svr) GetRoleList(ctx context.Context, in *proto.PaginateRequest) (*proto.GetRoleListReply, error) {
	appid := getMetaAppId(ctx)
	appid_str, _ := structpb.NewList([]interface{}{"app_key"})
	if in.Search == nil {
		in.Search = make(map[string]*structpb.ListValue)
	}
	in.Search[appid] = appid_str

	roles := []model.Role{}
	data, e := model.Paginate(&roles, int(in.Page), int(in.Limit), model.ConversionSearchType(in.Search), "RoleAuthoritys")
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	for i, item := range roles {
		// 角色关联的用户id
		var userIds []uint
		model.M.Model(&model.RoleUserMap{}).Where("role_id = ? AND app_key = ?", item.ID, appid).Pluck("user_id", &userIds)
		roles[i].AssUserIds = userIds
	}

	str, _ := json.ObjectToJson(roles)
	docs := []*proto.GetRoleListReply_RoleListData{}
	json.JsonToObject(str, &docs)

	return &proto.GetRoleListReply{
		Paginate: &proto.PaginateReply{
			Page:  uint32(data.Page),
			Pages: uint32(data.Pages),
			Total: uint32(data.Total),
			Limit: uint32(data.Limit),
		},
		Docs: docs,
	}, nil
}

type CreateRoleParam struct {
	Name        string            `json:"name" validate:"required"`
	Remark      string            `json:"remark" validate:"-"`
	Status      *bool             `json:"status" validate:"required"`
	Sequence    uint32            `json:"sequence" validate:"-"`
	Permissions map[string]uint32 `json:"permissions" validate:"-"`
	AppId       string            `json:"appid" validate:"required"`
}

// @Summary 创建角色
// @Description 创建应用角色
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body CreateRoleParam true "注册"
// @Success 200 {object} Response
// @Router /admin/role/createRole [post]
func (this *Svr) AdminCreateRole(c *see.Context) Response {
	var parameters CreateRoleParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if find := model.M.Where("name = ? AND app_key = ?", parameters.Name, parameters.AppId).First(&model.Role{}).RowsAffected; find > 0 {
		return Failure("已存在同名角色")
	}

	authoritys := []model.RoleAuthority{}
	for code, value := range parameters.Permissions {
		authoritys = append(authoritys, model.RoleAuthority{Code: code, ActionValue: value})
	}
	role := model.Role{
		Name:           parameters.Name,
		Remark:         parameters.Remark,
		AppKey:         parameters.AppId,
		Status:         parameters.Status,
		Sequence:       parameters.Sequence,
		RoleAuthoritys: authoritys,
	}
	err := role.Create()
	if err != nil {
		return Failure(err.Error())
	}
	return Success("创建成功")
}

type UpdateRoleParam struct {
	Id          uint              `json:"id" validate:"required"`
	Name        string            `json:"name" validate:"required"`
	Remark      string            `json:"remark" validate:"-"`
	Status      *bool             `json:"status" validate:"required"`
	Sequence    uint32            `json:"sequence" validate:"-"`
	Permissions map[string]uint32 `json:"permissions" validate:"-"`
	AppId       string            `json:"appid" validate:"required"`
}

// @Summary 更新角色信息
// @Description 更新角色信息和权限
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body UpdateRoleParam true "注册"
// @Success 200 {object} Response
// @Router /admin/role/updateRole [post]
func (this *Svr) AdminUpdateRole(c *see.Context) Response {
	var parameters UpdateRoleParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	authoritys := []model.RoleAuthority{}
	for code, value := range parameters.Permissions {
		authoritys = append(authoritys, model.RoleAuthority{Code: code, ActionValue: value})
	}
	role := model.Role{
		Name:           parameters.Name,
		Remark:         parameters.Remark,
		AppKey:         parameters.AppId,
		Status:         parameters.Status,
		Sequence:       parameters.Sequence,
		RoleAuthoritys: authoritys,
	}
	err := role.Update(parameters.Id)
	if err != nil {
		return Failure(err.Error())
	}

	return Success("修改成功")

}

// @Summary 获取角色列表
// @Description 获取角色列表
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body PaginateParam true "注册参数"
// @Success 200 {object} Response
// @Router /admin/role/list [post]
func (this *Svr) AdminGetRoleList(c *see.Context) Response {
	var parameters PaginateParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	roles := []model.Role{}
	data, e := model.Paginate(&roles, int(parameters.Page), int(parameters.Limit), model.ConversionSearchType2(parameters.Search), "RoleAuthoritys")
	if e != nil {
		return Failure(e.Error())
	}

	var appKey string
	if len(roles) > 0 {
		appKey = roles[0].AppKey
	}
	roleAuthTree, roleAuthTreeLeaf := getRoleAuthTree(0, appKey, 0)
	for i, item := range roles {
		// 权限树已选中节点
		checkedNode := []string{}
		for _, auth := range item.RoleAuthoritys {
			for _, leaf := range roleAuthTreeLeaf {
				if auth.ActionValue&(leaf["Value"].(uint32)) > 0 && auth.Code == leaf["PCode"].(string) {
					checkedNode = append(checkedNode, (leaf["ID"].(string)))
				}
			}
		}

		// 角色关联的用户id
		var userIds []uint
		model.M.Model(&model.RoleUserMap{}).Where("role_id = ? AND app_key = ?", item.ID, appKey).Pluck("user_id", &userIds)
		roles[i].UserCount = int64(len(userIds))

		roles[i].Other = see.H{"checkedNode": checkedNode, "assUserIds": userIds}
	}
	data.Docs = roles
	return RespData(see.H{"roles": data, "roleAuthTree": roleAuthTree})
}

func getRoleAuthTree(pid uint, appid string, userId uint) ([]map[string]interface{}, []map[string]interface{}) {
	outInterFace := []map[string]interface{}{}
	leafNode := []map[string]interface{}{}
	menus := []model.Menu{}
	model.M.Where("p_id = ? AND app_key = ?", pid, appid).Find(&menus)
	for _, node := range menus {
		children := []map[string]interface{}{}
		var count int64
		model.M.Model(&model.Menu{}).Where("p_id = ? AND app_key = ?", node.ID, appid).Count(&count)
		if count > 0 {
			nextLeafNode := []map[string]interface{}{}
			children, nextLeafNode = getRoleAuthTree(node.ID, appid, userId)
			leafNode = append(leafNode, nextLeafNode...)
		}

		actions := []string{}
		json.JsonToObject(node.Actions, &actions)
		permissions := map[string]uint32{}
		var withPermissions bool
		if userId > 0 {
			permissions = GetUserPermissions(appid, userId)
			withPermissions = true
		}

		if len(children) == 0 { //将Actions转成Children用于角色权限树渲染
			for key, value := range CalculateBit(actions, withPermissions, node.Code, permissions) {
				leaf := map[string]interface{}{
					"ID":     node.Code + "-" + strconv.Itoa(int(value)),
					"Code":   key,
					"Name":   key,
					"Value":  value,
					"IsLeaf": true,
					"PCode":  node.Code,
				}
				leafNode = append(leafNode, leaf)
				children = append(children, leaf)
			}
		}

		item := map[string]interface{}{
			"ID":       node.Code,
			"Code":     node.Code,
			"Name":     node.Name,
			"Children": children,
		}
		outInterFace = append(outInterFace, item)
	}
	return outInterFace, leafNode
}

func (this *Svr) UnBindRoleUser(ctx context.Context, in *proto.BindRoleUserRequest) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)
	e := model.M.Unscoped().Where("role_id = ? AND app_key = ? AND user_id = ?", in.RoleId, appid, in.UserId).Delete(model.RoleUserMap{}).Error
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	return ProtoSucceed(), nil
}

func (this *Svr) BindRoleUser(ctx context.Context, in *proto.BindRoleUserRequest) (*proto.Reply, error) {
	roleUser := model.RoleUserMap{
		AppKey: getMetaAppId(ctx),
		UserID: uint(in.UserId),
		RoleID: uint(in.RoleId),
	}
	e := model.M.Where(roleUser).FirstOrCreate(&roleUser).Error
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	return ProtoSucceed(), nil
}

func (this *Svr) BindRoleUsers(ctx context.Context, in *proto.BindRoleUsersRequest) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)
	transErr := model.M.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("role_id = ? AND app_key = ?", in.RoleId, appid).Delete(model.RoleUserMap{}).Error; err != nil {
			return err
		}

		if len(in.UserId) > 0 {
			data := []model.RoleUserMap{}
			for _, userid := range in.UserId {
				data = append(data, model.RoleUserMap{
					AppKey: appid,
					UserID: uint(userid),
					RoleID: uint(in.RoleId),
				})
			}
			if err := tx.Create(&data).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if transErr != nil {
		return nil, grpc.Errorf(codes.DataLoss, transErr.Error())
	}

	return ProtoSucceed(), nil
}

type BindUserParam struct {
	AppId  string  `json:"appid" validate:"required"`
	RoleId uint    `json:"roleid" validate:"required"`
	UserId *[]uint `json:"userid" validate:"required"`
}

// @Summary 绑定用户
// @Description 绑定用户
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body BindUserParam true "注册"
// @Success 200 {object} Response
// @Router /admin/role/bindUser [post]
func (this *Svr) AdminBindUser(c *see.Context) Response {
	var parameters BindUserParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	transErr := model.M.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("role_id = ? AND app_key = ?", parameters.RoleId, parameters.AppId).Delete(model.RoleUserMap{}).Error; err != nil {
			return err
		}

		if len(*parameters.UserId) > 0 {
			data := []model.RoleUserMap{}
			for _, userid := range *parameters.UserId {
				data = append(data, model.RoleUserMap{
					AppKey: parameters.AppId,
					UserID: uint(userid),
					RoleID: parameters.RoleId,
				})
			}
			if err := tx.Create(&data).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if transErr != nil {
		return Failure(transErr.Error())
	}

	return Success("修改成功")
}

func (this *Svr) AuthorityVerify(ctx context.Context, in *proto.AuthorityVerifyRequest) (*proto.Reply, error) {
	appid := getMetaAppId(ctx)
	userPermissions := GetUserPermissions(appid, uint(in.UserId))
	if userPermissions[in.Code]&in.ActionValue > 0 {
		return ProtoSucceed(), nil
	}
	return ProtoFailure(), nil
}
