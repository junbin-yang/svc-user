# SDK使用说明 

> 实现php调用grpc服务

## 服务状态检测


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$res = $connect->Ping();
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php ping.php
1
```

## 导入应用菜单


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$menu = '{"Node":[{"Code":"用户管理","Name":"用户管理","Sequence":1,"Uri":"/baidu","Remark":"","Actions":[],"Children":[{"Code":"公众用户","Name":"公众用户","Remark":"","Sequence":1,"Uri":"/123","Actions":["编辑","删除"],"Children":[]},{"Code":"村委用户","Name":"村委用户","Remark":"","Sequence":2,"Uri":"/123/123","Actions":["新建","查看"],"Children":[]}]},{"Code":"角色管理","Name":"角色管理","Sequence":2,"Uri":"/role","Remark":"","Actions":["新增角色","删除角色"],"Children":[]}]}';
	$res = $connect->PermissionsTpl($menu);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php importMenus.php
1
```

## 获取应用菜单


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$res = $connect->GetAppMenus();
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	print_r($res->GetData());
?>
```

```shell
php getMenus.php
Array
(
    [Node] => Array
        (
            [0] => Array
                (
                    [Code] => 用户管理
                    [Name] => 用户管理
                    [Sequence] => 1
                    [Uri] => /baidu
                    [Children] => Array
                        (
                            [0] => Array
                                (
                                    [Code] => 公众用户
                                    [Name] => 公众用户
                                    [Sequence] => 1
                                    [Uri] => /123
                                    [Actions] => Array
                                        (
                                            [编辑] => 1
                                            [删除] => 2
                                        )

                                    [Status] => 1
                                )

                            [1] => Array
                                (
                                    [Code] => 村委用户
                                    [Name] => 村委用户
                                    [Sequence] => 2
                                    [Uri] => /123/123
                                    [Actions] => Array
                                        (
                                            [查看] => 2
                                            [新建] => 1
                                        )

                                    [Status] => 1
                                )

                        )

                    [Status] => 1
                )

            [1] => Array
                (
                    [Code] => 角色管理
                    [Name] => 角色管理
                    [Sequence] => 2
                    [Uri] => /role
                    [Actions] => Array
                        (
                            [新增角色] => 1
                            [删除角色] => 2
                        )

                    [Status] => 1
                )

        )

)
```

## 修改菜单状态

允许修改某个菜单节点的状态。

注意如某个节点的状态是false，该节点下的权限值在获取角色权限时也会返回。


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$menu_code = "角色管理";
	$status = false;
	$res = $connect->DisableMenu($menu_code,$status);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php setMenuNodeStatus.php
1
```

## 创建应用角色


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$name = "SDK角色1";
	$desc = "这是一个测试角色";
	$status = true;
    $sequence = 0;
	$permissions = array("村委用户"=>3,"公众用户"=>1);//key为菜单的Code，value为Actions的相或值
	//$res = $connect->CreateRole($name);	//创建角色，没有权限
	$res = $connect->CreateRole(
        $name,$desc,
        $status,$sequence,
        $permissions
    );		//创建角色，并设置权限
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php newRole.php
1
```

## 编辑应用角色


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$role_id = 7;
	$name = "SDK角色..";
	$desc = "更新测试角色";
	$status = true;
    $sequence = 0;
	$permissions = array("村委用户"=>3,"公众用户"=>1);
	$res = $connect->UpdateRole(
        $role_id,
        $name,$desc,
        $status,$sequence,
        $permissions
    );
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php editRole.php
1
```

## 删除应用角色


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$role_id = 6;
	$res = $connect->DeleteRole($role_id);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php delRole.php
1
```

## 获取应用角色


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$page = 1;
	$limit = 10;
	$search = array("SDK角色"=>array("name","remark")); //搜索字段
	$res = $connect->GetRoleList($page, $limit, $search);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	print_r($res->GetData());
?>
```

```shell
php getRoleList.php
Array
(
    [Paginate] => Array
        (
            [Page] => 1
            [Pages] => 1
            [Total] => 3
            [Limit] => 10
        )

    [Docs] => Array
        (
            [0] => Array
                (
                    [Id] => 7
                    [CreatedAt] => 2022-04-17T14:01:03.869+08:00
                    [UpdatedAt] => 2022-04-17T14:18:33.675+08:00
                    [Name] => SDK角色..
                    [Remark] => 更新测试角色
                    [RoleAuthoritys] => Array
                        (
                            [0] => Array
                                (
                                    [Code] => 公众用户
                                    [ActionValue] => 1
                                )

                            [1] => Array
                                (
                                    [Code] => 村委用户
                                    [ActionValue] => 3
                                )

                        )

                )

            [1] => Array
                (
                    [Id] => 5
                    [CreatedAt] => 2022-04-17T13:50:37.437+08:00
                    [UpdatedAt] => 2022-04-17T13:50:37.437+08:00
                    [Name] => SDK角色2
                    [Remark] => 这是一个测试角色
                    [Status] => 1
                    [RoleAuthoritys] => Array
                        (
                            [0] => Array
                                (
                                    [Code] => 公众用户
                                    [ActionValue] => 1
                                )

                            [1] => Array
                                (
                                    [Code] => 村委用户
                                    [ActionValue] => 3
                                )

                        )

                )

            [2] => Array
                (
                    [Id] => 4
                    [CreatedAt] => 2022-04-17T13:49:19.916+08:00
                    [UpdatedAt] => 2022-04-17T13:49:19.916+08:00
                    [Name] => SDK角色1
                    [Remark] => 这是一个测试角色
                    [Status] => 1
                    [RoleAuthoritys] => Array
                        (
                            [0] => Array
                                (
                                    [Code] => 村委用户
                                    [ActionValue] => 3
                                )

                        )

                )

        )

)
```

搜索字段说明：

key为要搜索的值,value为数据库字段，多个value匹配其中一个,如模糊搜索手机号或账号。

如下面的例子中：

{"SDK角色":["name","remark"]}	=>	... where name like %SDK角色% or remark like %SDK角色%

如果是多条件查询：

{"SDK角色":["name"],"测试":["remark"]}	=>	... where name like %SDK角色% and remark like %测试%

可查询的字段：

| Field      | Type     | Description |
| ---------- | -------- | ----------- |
| id         | int      | 用户id      |
| created_at | datetime | 创建时间    |
| updated_at | datetime | 更新时间    |
| name       | string   | 角色名      |
| remark     | string   | 描述信息    |
| sequence   | int      | 排序        |
| status     | int      | 状态        |

## 创建新用户

除了各应用通过调用服务创建用户，在系统管理后台登录页面可以进行用户注册和密码找回。


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$account = "admin";
	$password = "jqweaD9w1zsi#02";
	$nick_name = "普通管理员";
    $mobile = "15012345671";
	$email = "123@qq.com";
	$info = array(
    	"realname" => "姓名",
        "age" => "18",
        "sex" => "1",
        "nation" => "汉族",
        "address" => "地址",
        "idcard" => "身份证号码",
    );	// 用户扩展信息，非必传参数
	$res = $connect->CreateUser(
        $account,$password,
        $nick_name,$mobile,
        $email,$info
    );		//创建角色，并设置权限
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// userId
?>
```

```shell
php newUser.php
4
```

## 更新用户信息


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$user_id = 4;
	$info = array(
    	"realname" => "姓名1",
        "age" => "19",
        "sex" => "1",
        "nation" => "汉族",
        "address" => "地址",
        "idcard" => "身份证号码",
    );
	$res = $connect->UpdateUserInfo($user_id,$info);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php editUser.php
1
```

## 启用/禁用用户


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$user_id = 4;
	$status = false;
	$res = $connect->ChangingUserStatus($user_id,$status);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php stopUser.php
1
```

## 获取用户列表


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$page = 1;
	$limit = 10;
	$search = array("150"=>array("mobile","email")); //搜索字段
	$res = $connect->GetUserList($page, $limit, $search);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	print_r($res->GetData());
?>
```

```shell
php getUserList.php
Array
(
    [Paginate] => Array
        (
            [Page] => 1
            [Pages] => 1
            [Total] => 2
            [Limit] => 10
        )

    [Docs] => Array
        (
            [0] => Array
                (
                    [Id] => 4
                    [CreatedAt] => 2022-04-17T16:05:09.855+08:00
                    [UpdatedAt] => 2022-04-17T16:05:09.855+08:00
                    [Account] => admin
                    [NickName] => 普通管理员
                    [Mobile] => 15012345671
                    [Email] => 123@qq.com
                    [Status] => 1
                    [LoginedAt] => 0001-01-01T00:00:00Z
                    [Info] => Array
                        (
                            [RealName] => 姓名1
                            [Age] => 19
                            [Sex] => 1
                            [Nation] => 汉族
                            [Address] => 地址
                            [IdCard] => 身份证号码
                        )

                )

            [1] => Array
                (
                    [Id] => 1
                    [CreatedAt] => 2022-04-10T21:13:01.05+08:00
                    [UpdatedAt] => 2022-04-17T19:56:06.22+08:00
                    [Account] => root
                    [NickName] => 超级管理员
                    [Mobile] => 15000000011
                    [Email] => zhang@yaoapps.com
                    [Status] => 1
                    [LoginedAt] => 2022-04-17T19:56:06.219+08:00
                    [Info] => Array
                        (
                            [RealName] => zhangsan
                            [Age] => 18
                            [Sex] => 1
                            [Nation] => 汉
                            [Address] => 中国-广东-深圳
                            [IdCard] => 55554444
                        )

                )

        )

)
```

搜索字段说明：

key为要搜索的值,value为数据库字段，多个value匹配其中一个,如模糊搜索手机号或账号。

如下面的例子中：

{"150":["mobile","email"]}	=>	... where mobile like %150% or email like %150%

如果是多条件查询：

{"150":["mobile"],"yao":["email"]}	=>	... where mobile like %150% and email like %yao%

可查询的字段：

| Field      | Type     | Description  |
| ---------- | -------- | ------------ |
| id         | int      | 用户id       |
| created_at | datetime | 创建时间     |
| updated_at | datetime | 更新时间     |
| account    | string   | 账号         |
| nick_name  | string   | 昵称         |
| mobile     | string   | 手机号       |
| email      | string   | 邮箱         |
| status     | int      | 状态         |
| logined_at | datetime | 最后登录时间 |

##  角色用户绑定（多用户） 

此操作会将角色下对应的绑定用户全部清除。然后按提交的用户id重新绑定


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$role_id = 5;
	$user_id = array(2,3);
	$res = $connect->BindRoleUsers($role_id, $user_id);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php bindUsers.php
1
```

##  角色用户绑定

给角色新增一个绑定用户，不会清除现有绑定用户。


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$role_id = 5;
	$user_id = 3;
	$res = $connect->BindRoleUser($role_id, $user_id);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php bindUser.php
1
```

##  角色用户解绑


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$role_id = 5;
	$user_id = 3;
	$res = $connect->UnBindRoleUser($role_id, $user_id);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php unBindUser.php
1
```

##   验证用户菜单权限 

 验证用户是否有某个菜单动作权限 


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	/*
	Array
    (
    	[Code] => 公众用户
        [Name] => 公众用户
        [Sequence] => 1
        [Uri] => /123
        [Actions] => Array
        (
        	[编辑] => 1
            [删除] => 2
        )
        [Status] => 1
    )
	*/
	$user_id = 3;
	$menu_code = "公众用户";
	$action_value = 1;
	$res = $connect->AuthorityVerify($user_id, $menu_code, $action_value);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php authMenu.php
1
```

##  验证用户密码 


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$user_type = "account";   // account|email|mobile
	$type_value = "root";
	$password = "123456";
	$res = $connect->UserPassVerify($user_type, $type_value, $password);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// true or false
?>
```

```shell
php authUser.php
1
```

##  创建Token 


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$user_id = 2;
	$issuer = "desc";
	$res = $connect->TokenCreate($user_id, $issuer);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// token string
?>
```

```shell
php newToken.php
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjIsImFkbWluIjpmYWxzZSwiZXhwIjoxNjUwMjgwMzg4LCJpc3MiOiJkZXNjIn0.QX9cyS4rekF8yO6ol53s4FuQD7tDe8ZTJRiLwVvQmNs
```

## 验证Token 


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjIsImFkbWluIjpmYWxzZSwiZXhwIjoxNjUwMjgwMzg4LCJpc3MiOiJkZXNjIn0.QX9cyS4rekF8yO6ol53s4FuQD7tDe8ZTJRiLwVvQmNs";
	$res = $connect->TokenVerify($token);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	print_r($res->GetData());		// token object
?>
```

```shell
php verifyToken.php
Array
(
    [UserId] => 2
    [ExpiresAt] => 1650280388
    [Issuer] => desc
    [MenuPermissions] => []
)
```

## 更新Token 


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjIsImFkbWluIjpmYWxzZSwiZXhwIjoxNjUwMjgwMzg4LCJpc3MiOiJkZXNjIn0.QX9cyS4rekF8yO6ol53s4FuQD7tDe8ZTJRiLwVvQmNs";
	$res = $connect->TokenRefresh($token);
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData() . PHP_EOL;		// token string
?>
```

```shell
php refreshToken.php
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjIsImFkbWluIjpmYWxzZSwiZXhwIjoxNjUwMjgwODQ0LCJpc3MiOiJkZXNjIn0.spiQXxdh29TbX8cOke7JyOEtPMVCCnqrNdCbIFsWv6w
```

##  生成验证码 


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$res = $connect->CreateValidateCode();
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	print_r($res->GetData());		// validate code object
?>
```

```shell
php newValidateCode.php
Array
(
    [CodeId] => qbCj8b8ekScSEZDhA4xR
    [ImageData] => data:image/png;base64,...
)
```

##  校验验证码 


```php
<?php
	require dirname(__FILE__).'/Api.php';

	$meta = [
        "appid" => ["620f9a6fb47b05e7"], 
        "appkey" => ["65356533-3539-6464-6330-65653738"]
    ];
	$hostname = "127.0.0.1:8888";

	$connect = new SvcUser\Api($hostname, $meta);
	$code_id = "qbCj8b8ekScSEZDhA4xR";
    $code_value = "ua82";
	$res = $connect->VerifyValidateCode();
	if ($res->GetError() != null) {
        echo $res->GetError() . PHP_EOL;
        exit(1);
	}
	echo $res->GetData();		// true or false
?>
```

```shell
php verifyValidateCode.php
1
```

## 









































