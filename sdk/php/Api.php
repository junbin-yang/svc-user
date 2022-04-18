<?php
namespace SvcUser;
require dirname(__FILE__).'/vendor/autoload.php';
require dirname(__FILE__).'/Client.php';

class Response {
	public $Code = null;
	public $Error = null;
	public $Data = null;

	public function SetCode($code)
	{
		$this->Code = $code;
	}

	public function GetCode()
	{
		return $this->Code;
	}

	public function SetError($error)
	{
		$this->Error = $error;
	}

	public function GetError()
	{
		return $this->Error;
	}

	public function SetData($data)
	{
		$this->Data = $data;
	}

	public function GetData()
	{
		return $this->Data;
	}
}

class Api {
    private $client = null;

    public function __construct($hostname, $metadata) 
    {
    	$options = ['credentials' => \Grpc\ChannelCredentials::createInsecure()];
        $this->client = new Client($hostname, $metadata, $options);
    }

    public function Test($json)
    {
    	$response = new Response();
    	$request = new \Proto\CreateUserRequest();
    	$request->mergeFromJsonString($json);
    	echo $request->serializeToJsonString();
    }

    public function Ping()
    {
    	$response = new Response();
    	$request = new \Proto\NullRequest();
    	list($res, $status) = $this->client->Ping($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    }

    /**
     * @param string $json
     * {
     *     "Node":[
     *         {
     *             "Code":"用户管理",
     *             "Name":"用户管理",
     *             "Sequence":1,
     *             "Uri":"/#",
     *             "Remark":"",
     *             "Actions":["新增","编辑"],
     *             "Children":[]
     *         }
     *     ]
     * } 
     */
    public function PermissionsTpl($json)
    {
    	$response = new Response();
    	$request = new \Proto\PermissionsTplRequest();
    	$request->mergeFromJsonString($json);
    	list($res, $status) = $this->client->PermissionsTpl($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    } 

    public function GetAppMenus()
    {
    	$response = new Response();
    	$request = new \Proto\NullRequest();
    	list($res, $status) = $this->client->GetAppMenus($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData(json_decode($res->serializeToJsonString(),true));
    	return $response;
    }

    public function DisableMenu($menu_code, $status = false)
    {
    	$response = new Response();
    	$request = new \Proto\DisableMenuRequest();
    	$request->setCode($menu_code);
    	$request->setStatus((bool)$status);
    	list($res, $status) = $this->client->DisableMenu($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    }

    public function CreateRole($name, $desc = '', $status = true, $sequence = 0,$permissions = [])
    {
    	$response = new Response();
    	$request = new \Proto\RoleInfo();
    	$request->setName($name);
    	$request->setRemark($desc);
    	$request->setStatus((bool)$status);
    	$request->setSequence((int)$sequence);
		$request->setPermissions($permissions);
    	list($res, $status) = $this->client->CreateRole($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    }

    public function UpdateRole($id, $name, $desc = '', $status = true, $sequence = 0,$permissions = [])
    {
    	$response = new Response();
    	$request = new \Proto\UpdateRoleRequest();
    	$request->setId($id);
    	$role_info = new \Proto\RoleInfo();
    	$role_info->setName($name);
    	$role_info->setRemark($desc);
    	$role_info->setStatus((bool)$status);
    	$role_info->setSequence((int)$sequence);
		$role_info->setPermissions($permissions);
    	$request->setRoleInfo($role_info);
    	list($res, $status) = $this->client->UpdateRole($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    }

    public function DeleteRole($role_id)
    {
    	$response = new Response();
    	$request = new \Proto\DeleteRoleRequest();
    	$request->setId($role_id);
    	list($res, $status) = $this->client->DeleteRole($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    }

    public function GetRoleList($page = 1, $limit = 10, $search = [])
    {
    	$response = new Response();
    	$request = new \Proto\PaginateRequest();
    	$request->setPage($page);
		$request->setLimit($limit);

		$encode_search = array();
		foreach ($search as $key => $value) {
			$listValue = new \Google\Protobuf\ListValue();
			$listValue->mergeFromJsonString(json_encode($value,true));
			$encode_search[$key] = $listValue;
		}
		$request->setSearch($encode_search);

    	list($res, $status) = $this->client->GetRoleList($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData(json_decode($res->serializeToJsonString(),true));
    	return $response;
    }

    public function CreateUser($account, $password, $nick_name, $mobile, $email, $info = [])
    {
    	
    	$response = new Response();
    	$request = new \Proto\CreateUserRequest();
    	$request->setAccount($account);
    	$request->setPassword($password);
    	$request->setNickName($nick_name);
    	$request->setMobile($mobile);
    	$request->setEmail($email);
    	if (!empty($info))
    	{
    		$info = array_change_key_case($info, CASE_LOWER);
    		$user_info = new \Proto\UserInfo();
    		$user_info->setRealName($info["realname"]);
    		$user_info->setAge($info["age"]);
    		$user_info->setSex($info["sex"]);
    		$user_info->setNation($info["nation"]);
    		$user_info->setAddress($info["address"]);
    		$user_info->setIdCard($info["idcard"]);
    		$request->setUserInfo($user_info);
    	}
    	list($res, $status) = $this->client->CreateUser($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getUserId());
    	return $response;
    }

    public function UpdateUserInfo($user_id, $info = [])
    {
    	$response = new Response();
    	$request = new \Proto\UpdateUserInfoRequest();
    	$request->setId($user_id);
    	if (!empty($info))
    	{
    		$info = array_change_key_case($info, CASE_LOWER);
    		$user_info = new \Proto\UserInfo();
    		$user_info->setRealName($info["realname"]);
    		$user_info->setAge($info["age"]);
    		$user_info->setSex($info["sex"]);
    		$user_info->setNation($info["nation"]);
    		$user_info->setAddress($info["address"]);
    		$user_info->setIdCard($info["idcard"]);
    		$request->setUserInfo($user_info);
    	}
    	list($res, $status) = $this->client->UpdateUserInfo($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    }

    public function ChangingUserStatus($user_id, $status)
    {
    	$response = new Response();
    	$request = new \Proto\ChangingUserStatusRequest();
    	$request->setId($user_id);
    	$request->setStatus((bool)$status);
    	list($res, $status) = $this->client->ChangingUserStatus($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData($res->getStatus());
    	return $response;
    }

    public function GetUserList($page = 1, $limit = 10, $search = [])
    {
    	$response = new Response();
    	$request = new \Proto\PaginateRequest();
    	$request->setPage($page);
		$request->setLimit($limit);

		$encode_search = array();
		foreach ($search as $key => $value) {
			$listValue = new \Google\Protobuf\ListValue();
			$listValue->mergeFromJsonString(json_encode($value,true));
			$encode_search[$key] = $listValue;
		}
		$request->setSearch($encode_search);

    	list($res, $status) = $this->client->GetUserList($request)->wait();
    	$response->SetCode($status->code);
    	if ($status->code !== \Grpc\STATUS_OK)
    	{
    		$response->SetError($status->details);
    		return $response;
    	}
    	$response->SetData(json_decode($res->serializeToJsonString(),true));
    	return $response;
    }

    public function BindRoleUsers($role_id, $user_id = [])
    {
        $response = new Response();
        $request = new \Proto\BindRoleUsersRequest();
        $request->setRoleId($role_id);
        $request->setUserId($user_id);
        list($res, $status) = $this->client->BindRoleUsers($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getStatus());
        return $response;
    }

    public function BindRoleUser($role_id, $user_id)
    {
        $response = new Response();
        $request = new \Proto\BindRoleUserRequest();
        $request->setRoleId($role_id);
        $request->setUserId($user_id);
        list($res, $status) = $this->client->BindRoleUser($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getStatus());
        return $response;
    }

    public function UnBindRoleUser($role_id, $user_id)
    {
        $response = new Response();
        $request = new \Proto\BindRoleUserRequest();
        $request->setRoleId($role_id);
        $request->setUserId($user_id);
        list($res, $status) = $this->client->UnBindRoleUser($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getStatus());
        return $response;
    }

    public function AuthorityVerify($user_id, $menu_code, $action_value)
    {
        $response = new Response();
        $request = new \Proto\AuthorityVerifyRequest();
        $request->setCode($menu_code);
        $request->setActionValue($action_value);
        $request->setUserId($user_id);
        list($res, $status) = $this->client->AuthorityVerify($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getStatus());
        return $response;
    }

    public function UserPassVerify($user_type, $type_value, $password)
    {
        $response = new Response();
        $request = new \Proto\UserPassVerifyRequest();
        $request->setType($user_type);
        $request->setValue($type_value);
        $request->setPassword($password);
        list($res, $status) = $this->client->UserPassVerify($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getStatus());
        return $response;
    }

    public function TokenCreate($user_id, $issuer)
    {
        $response = new Response();
        $request = new \Proto\TokenCreateRequest();
        $request->setUserId($user_id);
        $request->setIssuer($issuer);
        list($res, $status) = $this->client->TokenCreate($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getToken());
        return $response;
    }

    public function TokenVerify($token)
    {
        $response = new Response();
        $request = new \Proto\Token();
        $request->setToken($token);
        list($res, $status) = $this->client->TokenVerify($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData(json_decode($res->serializeToJsonString(),true));
        return $response;
    }

    public function TokenRefresh($token)
    {
        $response = new Response();
        $request = new \Proto\Token();
        $request->setToken($token);
        list($res, $status) = $this->client->TokenRefresh($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getToken());
        return $response;
    }

    public function CreateValidateCode()
    {
        $response = new Response();
        $request = new \Proto\NullRequest();
        list($res, $status) = $this->client->CreateValidateCode($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData(json_decode($res->serializeToJsonString(),true));
        return $response;
    }

    public function VerifyValidateCode($code_id, $code_value)
    {
        $response = new Response();
        $request = new \Proto\VerifyValidateCodeRequest();
        $request->setCodeId($code_id);
        $request->setCodeValue($code_value);
        list($res, $status) = $this->client->VerifyValidateCode($request)->wait();
        $response->SetCode($status->code);
        if ($status->code !== \Grpc\STATUS_OK)
        {
            $response->SetError($status->details);
            return $response;
        }
        $response->SetData($res->getStatus());
        return $response;
    }
}