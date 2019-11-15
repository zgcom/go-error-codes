package error

type ErrCode int

//go:generate stringer -type ErrCode -linecomment
const (
	INTERNAL_ERROR       ErrCode = 1000 // 内部错误
	PARAM_ERROR          ErrCode = 1001 // 参数错误
	REDIS_ERROR          ErrCode = 1002 // 内部错误
	MYSQL_ERROR          ErrCode = 1003 // 内部错误
	JSON_PARSE_FAIL      ErrCode = 1004 // 内部错误
	JWT_VERIFY_ERROR     ErrCode = 1005 // jwt权限验证失败
	API_RATELIMIT        ErrCode = 1006 // api频率限制
	JWT_RIGHT_NOT_ENOUGH ErrCode = 1007 // jwt权限不足
	ERROR_JWT_GENERATE   ErrCode = 1008 // jwt生成错误

	CURRENCY_NOT_SUPPORT ErrCode = 2000 // 币种不支持
	WITHDRAW_STOP        ErrCode = 2001 // 提现已暂停
	BALANCE_NOT_ENOUGH   ErrCode = 2002 // 余额不足
	WITHDRAW_FORBID      ErrCode = 2003 // 禁止提现
	ADDRESS_ILLEGAL      ErrCode = 2005 // 地址非法
	GET_RATE_ERROR       ErrCode = 2009 // 获取汇率失败
	QUOTA_DAILY_LIMIT    ErrCode = 2010 // 到达每日限额
	QUOTA_TOTAL_LIMIT    ErrCode = 2011 // 到达总限额
	AMOUNT_TOO_MUCH      ErrCode = 2012 // 提现数量太大
	SPOT_FORBIDDEN       ErrCode = 2013 // 用户被禁止币币提现
	SECURITY_SET_LIMIT   ErrCode = 2014 // 安全设置24小时内不能提现

	VCODE_ERROR                       ErrCode = 3000 //验证码错误
	ACCOUNT_ALREADY_EXIST             ErrCode = 3001 //账户已存在
	CREATE_ENGINE_ACCOUNT_FAIL        ErrCode = 3002 //创建账户错误
	CREATE_JWT_FAIL                   ErrCode = 3003 //创建jwt失败
	ACCOUNT_NOT_EXIST                 ErrCode = 3004 //账户不存在
	LOGIN_FAIL                        ErrCode = 3005 //登陆失败
	ENGINE_ACCOUNT_NOT_EXSIST         ErrCode = 3006 //账户不存在
	ENGINE_ACCOUNT_ALREADY_EXIST      ErrCode = 3007 //账户已存在
	CONFIG_NOT_EXIST                  ErrCode = 3008 //配置不存在
	GACODE_ALREADY_UESD               ErrCode = 3009 //谷歌验证码已经使用了
	GACODE_VERIFY_FAILED              ErrCode = 3010 //谷歌验证码验证失败
	TRADE_PWD_VERIFY_FAILED           ErrCode = 3011 //交易密码验证失败
	SMSCODE_VERIFY_FAILED             ErrCode = 3012 //手机验证码验证失败
	EMAILCODE_VERIFY_FAILED           ErrCode = 3013 //邮箱验证码验证失败
	SECURITY_VERIFY_FAILED            ErrCode = 3014 //安全验证失败
	ALREADY_BINDED                    ErrCode = 3015 //用户已经绑定
	MOBILE_OR_EMAIL_MUST_HAS_ONE_OPEN ErrCode = 3016 //手机或者邮箱必须有一个开启
	ALREADY_KYC                       ErrCode = 3017 //用户已经实名验证
	ONE_LEVEL_KYC_MUST_AUTH           ErrCode = 3018 //必须通过一级实名验证
	KYC_INFO_NOT_FOUND                ErrCode = 3019 //用户认证信息没有找到
	EXCEPTIONAL_USER                  ErrCode = 3020 //
	CAPTCHA_ERROR                     ErrCode = 3021 //图形验证码错误或失效
	ONE_IDCARD_ERROR                  ErrCode = 3022 //一级实名认证失败
	FACEID_SIGN_FAILED                ErrCode = 3023 //face++签名错误
	SPONSOR_IS_VERIFYING_OR_USING     ErrCode = 3024 //申请的保荐机构正在审核或者已经在使用
	SPONSOR_NOT_APPLY                 ErrCode = 3025 //还没有申请保荐机构
	SPONSOR_NOT_APPLY_OR_VERIFY       ErrCode = 3026 //还没有申请或者审核保荐机构
	SPONSOR_RIGHT_NOT_EXIST           ErrCode = 3027 //还没有配置保荐机构的权益
	TRADE_PWD_MUST_SET                ErrCode = 3028 //交易密码必须设置
	INVITE_CODE_NOT_EXIST             ErrCode = 3029 //邀请码不存在
	HAS_NO_USING_SPONSOR              ErrCode = 3030 //还没有正在使用的保荐机构
	COBBER_IS_APPLYING                ErrCode = 3031 //合伙人正在申请升级中
	SECURITY_TYPE_ERROR               ErrCode = 3032 // 安全验证类型错误

	GET_ENGINE_ACCOUNT_FAIL    ErrCode = 4000 // 获取引擎账户信息失败
	ACCOUNT_BALANCE_NOT_ENOUGH ErrCode = 4001 // 账户余额不足

	EMAIL_SERVICE_NOT_FOUND  ErrCode = 5000 // 没找到邮箱服务供应商
	EMAIL_TEMPLATE_NOT_FOUND ErrCode = 5001 // 邮件模板未找到
	DINGDING_ROBOT_NOT_FOUND ErrCode = 5002 // 没找到dingding机器人
	DINGDING_ROBOT_SEND_FAIL ErrCode = 5003 // dingding机器人发送失败
	SMS_TYPE_NOT_EXISTS      ErrCode = 5004 // 短信类型不存在
	VCODE_VERIFY_FAILED      ErrCode = 5005 // 验证码验证失败
	EMAIL_TYPE_NOT_EXISTS    ErrCode = 5006 // 邮件类型不存在

	BANKCARD_AMOUNT_LIMIT ErrCode = 6000 // 银行卡总数限制
	BANKCARD_EXISTS       ErrCode = 6001 // 银行卡已经存在
	BANKCARD_NOT_EXISTS   ErrCode = 6002 // 银行卡不存在
	GLOBAL_CONFIG_ERROR   ErrCode = 6003 // 配置错误
	ORDER_NOT_EXISTS      ErrCode = 6004 // 订单不存在
	ORDER_CANNOT_CANCEL   ErrCode = 6005 // 订单不可取消
	PENDING_ORDER_EXIST   ErrCode = 6006 // 已存在pending订单
	NO_SYSTEM_ACCOUNT     ErrCode = 6007 // 没有系统银行账户
	C2C_FORBIDDEN         ErrCode = 6010 // 用户c2c被禁用
	NEED_KYC              ErrCode = 6011

	ENGINE_DEDUCT_WITHDRAW_FREEZE_FAIL ErrCode = 7000 // 扣除冻结失败
	ENGINE_WITHDRAW_FREEZE_FAIL        ErrCode = 7001 // 冻结失败
	ENGINE_DEPOSIT_FAIL                ErrCode = 7002 // 充值失败
)
