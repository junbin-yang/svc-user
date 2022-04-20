<?php
namespace SvcUser;
require dirname(__FILE__).'/vendor/autoload.php';

class Client extends \Grpc\BaseStub {

    public $options = [];
    public $metadata = [];

    /**
     * @param string $hostname hostname 
     * @param array $metadata auth metadata 
     * @param array $opts channel options 
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $metadata, $opts, $channel = null) 
    {
        $this->metadata = $metadata;
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * 服务检测
     * @param \Proto\NullRequest $argument input argument
     * @return \Grpc\UnaryCall
     */
    public function Ping(\Proto\NullRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/Ping',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 生成验证码
     */
    public function CreateValidateCode(\Proto\NullRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/CreateValidateCode',
            $argument,
            ['\Proto\ValidateCodeReply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 校验验证码
     */
    public function VerifyValidateCode(\Proto\VerifyValidateCodeRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/VerifyValidateCode',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 创建新用户
     */
    public function CreateUser(\Proto\CreateUserRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/CreateUser',
            $argument,
            ['\Proto\CreateUserReply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 更新用户信息
     */
    public function UpdateUserInfo(\Proto\UpdateUserInfoRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/UpdateUserInfo',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 启用/禁用用户
     */
    public function ChangingUserStatus(\Proto\ChangingUserStatusRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/ChangingUserStatus',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 获取用户列表
     */
    public function GetUserList(\Proto\PaginateRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/GetUserList',
            $argument,
            ['\Proto\GetUserListReply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 校验用户密码
     */
    public function UserPassVerify(\Proto\UserPassVerifyRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/UserPassVerify',
            $argument,
            ['\Proto\Token', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * Token生成
     */
    public function TokenCreate(\Proto\TokenCreateRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/TokenCreate',
            $argument,
            ['\Proto\Token', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * Token验证
     */
    public function TokenVerify(\Proto\Token $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/TokenVerify',
            $argument,
            ['\Proto\TokenVerifyReply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * Token更新
     */
    public function TokenRefresh(\Proto\Token $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/TokenRefresh',
            $argument,
            ['\Proto\Token', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 菜单权限模板注入 - 多次调用覆盖
     */
    public function PermissionsTpl(\Proto\PermissionsTplRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/PermissionsTpl',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 启用/禁用菜单
     */
    public function DisableMenu(\Proto\DisableMenuRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/DisableMenu',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 获取应用菜单
     */
    public function GetAppMenus(\Proto\NullRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/GetAppMenus',
            $argument,
            ['\Proto\MenusReply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 创建应用角色
     */
    public function CreateRole(\Proto\RoleInfo $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/CreateRole',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 编辑应用角色
     */
    public function UpdateRole(\Proto\UpdateRoleRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/UpdateRole',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 删除应用角色
     */
    public function DeleteRole(\Proto\DeleteRoleRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/DeleteRole',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 获取应用角色列表
     */
    public function GetRoleList(\Proto\PaginateRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/GetRoleList',
            $argument,
            ['\Proto\GetRoleListReply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 用户绑定角色（多用户）
     */
    public function BindRoleUsers(\Proto\BindRoleUsersRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/BindRoleUsers',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 用户绑定角色
     */
    public function BindRoleUser(\Proto\BindRoleUserRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/BindRoleUser',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 用户解绑角色
     */
    public function UnBindRoleUser(\Proto\BindRoleUserRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/UnBindRoleUser',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }

    /**
     * 验证用户是否有某个菜单动作权限
     */
    public function AuthorityVerify(\Proto\AuthorityVerifyRequest $argument) 
    {
        return $this->_simpleRequest(
            '/proto.User/AuthorityVerify',
            $argument,
            ['\Proto\Reply', 'decode'],
            $this->metadata, $this->options
        );
    }
}