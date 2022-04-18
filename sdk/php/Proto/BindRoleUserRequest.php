<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/user.proto

namespace Proto;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>proto.BindRoleUserRequest</code>
 */
class BindRoleUserRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * 角色id
     *
     * Generated from protobuf field <code>uint32 RoleId = 1;</code>
     */
    protected $RoleId = 0;
    /**
     * 用户id
     *
     * Generated from protobuf field <code>uint32 UserId = 2;</code>
     */
    protected $UserId = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $RoleId
     *           角色id
     *     @type int $UserId
     *           用户id
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
     * 用户id
     *
     * Generated from protobuf field <code>uint32 UserId = 2;</code>
     * @return int
     */
    public function getUserId()
    {
        return $this->UserId;
    }

    /**
     * 用户id
     *
     * Generated from protobuf field <code>uint32 UserId = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setUserId($var)
    {
        GPBUtil::checkUint32($var);
        $this->UserId = $var;

        return $this;
    }

}

