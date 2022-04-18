<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/user.proto

namespace Proto;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>proto.TokenCreateRequest</code>
 */
class TokenCreateRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * 用户id
     *
     * Generated from protobuf field <code>uint32 UserId = 1;</code>
     */
    protected $UserId = 0;
    /**
     * 发行标记
     *
     * Generated from protobuf field <code>string Issuer = 2;</code>
     */
    protected $Issuer = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $UserId
     *           用户id
     *     @type string $Issuer
     *           发行标记
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\User::initOnce();
        parent::__construct($data);
    }

    /**
     * 用户id
     *
     * Generated from protobuf field <code>uint32 UserId = 1;</code>
     * @return int
     */
    public function getUserId()
    {
        return $this->UserId;
    }

    /**
     * 用户id
     *
     * Generated from protobuf field <code>uint32 UserId = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setUserId($var)
    {
        GPBUtil::checkUint32($var);
        $this->UserId = $var;

        return $this;
    }

    /**
     * 发行标记
     *
     * Generated from protobuf field <code>string Issuer = 2;</code>
     * @return string
     */
    public function getIssuer()
    {
        return $this->Issuer;
    }

    /**
     * 发行标记
     *
     * Generated from protobuf field <code>string Issuer = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setIssuer($var)
    {
        GPBUtil::checkString($var, True);
        $this->Issuer = $var;

        return $this;
    }

}

