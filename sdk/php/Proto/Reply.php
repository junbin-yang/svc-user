<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/user.proto

namespace Proto;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>proto.Reply</code>
 */
class Reply extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bool Status = 1;</code>
     */
    protected $Status = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type bool $Status
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\User::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bool Status = 1;</code>
     * @return bool
     */
    public function getStatus()
    {
        return $this->Status;
    }

    /**
     * Generated from protobuf field <code>bool Status = 1;</code>
     * @param bool $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkBool($var);
        $this->Status = $var;

        return $this;
    }

}

