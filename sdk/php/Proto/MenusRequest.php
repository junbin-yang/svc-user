<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/user.proto

namespace Proto;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>proto.MenusRequest</code>
 */
class MenusRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .proto.MenusRequest.MenusNode Node = 1;</code>
     */
    private $Node;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Proto\MenusRequest\MenusNode[]|\Google\Protobuf\Internal\RepeatedField $Node
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\User::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .proto.MenusRequest.MenusNode Node = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getNode()
    {
        return $this->Node;
    }

    /**
     * Generated from protobuf field <code>repeated .proto.MenusRequest.MenusNode Node = 1;</code>
     * @param \Proto\MenusRequest\MenusNode[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setNode($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Proto\MenusRequest\MenusNode::class);
        $this->Node = $arr;

        return $this;
    }

}

