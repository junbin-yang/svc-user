<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/user.proto

namespace Proto;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>proto.BindRoleUsersRequest</code>
 */
class BindRoleUsersRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * 角色id
     *
     * Generated from protobuf field <code>uint32 RoleId = 1;</code>
     */
    protected $RoleId = 0;
    /**
     * 角色绑定的所有用户id，每次覆盖
     *
     * Generated from protobuf field <code>repeated uint32 UserId = 2;</code>
     */
    private $UserId;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $RoleId
     *           角色id
     *     @type int[]|\Google\Protobuf\Internal\RepeatedField $UserId
     *           角色绑定的所有用户id，每次覆盖
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\User::initOnce();
        parent::__construct($data);
    }

    /**
     * 角色id
     *
     * Generated from protobuf field <code>uint32 RoleId = 1;</code>
     * @return int
     */
    public function getRoleId()
    {
        return $this->RoleId;
    }

    /**
     * 角色id
     *
     * Generated from protobuf field <code>uint32 RoleId = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setRoleId($var)
    {
        GPBUtil::checkUint32($var);
        $this->RoleId = $var;

        return $this;
    }

    /**
     * 角色绑定的所有用户id，每次覆盖
     *
     * Generated from protobuf field <code>repeated uint32 UserId = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getUserId()
    {
        return $this->UserId;
    }

    /**
     * 角色绑定的所有用户id，每次覆盖
     *
     * Generated from protobuf field <code>repeated uint32 UserId = 2;</code>
     * @param int[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setUserId($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::UINT32);
        $this->UserId = $arr;

        return $this;
    }

}

