// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[INTERNAL_ERROR-1000]
	_ = x[PARAM_ERROR-1001]
	_ = x[REDIS_ERROR-1002]
	_ = x[MYSQL_ERROR-1003]
	_ = x[JSON_PARSE_FAIL-1004]
	_ = x[JWT_VERIFY_ERROR-1005]
	_ = x[API_RATELIMIT-1006]
	_ = x[JWT_RIGHT_NOT_ENOUGH-1007]
	_ = x[ERROR_JWT_GENERATE-1008]
	_ = x[CURRENCY_NOT_SUPPORT-2000]
	_ = x[WITHDRAW_STOP-2001]
	_ = x[BALANCE_NOT_ENOUGH-2002]
	_ = x[WITHDRAW_FORBID-2003]
	_ = x[ADDRESS_ILLEGAL-2005]
	_ = x[GET_RATE_ERROR-2009]
	_ = x[QUOTA_DAILY_LIMIT-2010]
	_ = x[QUOTA_TOTAL_LIMIT-2011]
	_ = x[AMOUNT_TOO_MUCH-2012]
	_ = x[SPOT_FORBIDDEN-2013]
	_ = x[SECURITY_SET_LIMIT-2014]
	_ = x[VCODE_ERROR-3000]
	_ = x[ACCOUNT_ALREADY_EXIST-3001]
	_ = x[CREATE_ENGINE_ACCOUNT_FAIL-3002]
	_ = x[CREATE_JWT_FAIL-3003]
	_ = x[ACCOUNT_NOT_EXIST-3004]
	_ = x[LOGIN_FAIL-3005]
	_ = x[ENGINE_ACCOUNT_NOT_EXSIST-3006]
	_ = x[ENGINE_ACCOUNT_ALREADY_EXIST-3007]
	_ = x[CONFIG_NOT_EXIST-3008]
	_ = x[GACODE_ALREADY_UESD-3009]
	_ = x[GACODE_VERIFY_FAILED-3010]
	_ = x[TRADE_PWD_VERIFY_FAILED-3011]
	_ = x[SMSCODE_VERIFY_FAILED-3012]
	_ = x[EMAILCODE_VERIFY_FAILED-3013]
	_ = x[SECURITY_VERIFY_FAILED-3014]
	_ = x[ALREADY_BINDED-3015]
	_ = x[MOBILE_OR_EMAIL_MUST_HAS_ONE_OPEN-3016]
	_ = x[ALREADY_KYC-3017]
	_ = x[ONE_LEVEL_KYC_MUST_AUTH-3018]
	_ = x[KYC_INFO_NOT_FOUND-3019]
	_ = x[EXCEPTIONAL_USER-3020]
	_ = x[CAPTCHA_ERROR-3021]
	_ = x[ONE_IDCARD_ERROR-3022]
	_ = x[FACEID_SIGN_FAILED-3023]
	_ = x[SPONSOR_IS_VERIFYING_OR_USING-3024]
	_ = x[SPONSOR_NOT_APPLY-3025]
	_ = x[SPONSOR_NOT_APPLY_OR_VERIFY-3026]
	_ = x[SPONSOR_RIGHT_NOT_EXIST-3027]
	_ = x[TRADE_PWD_MUST_SET-3028]
	_ = x[INVITE_CODE_NOT_EXIST-3029]
	_ = x[HAS_NO_USING_SPONSOR-3030]
	_ = x[COBBER_IS_APPLYING-3031]
	_ = x[SECURITY_TYPE_ERROR-3032]
	_ = x[GET_ENGINE_ACCOUNT_FAIL-4000]
	_ = x[ACCOUNT_BALANCE_NOT_ENOUGH-4001]
	_ = x[EMAIL_SERVICE_NOT_FOUND-5000]
	_ = x[EMAIL_TEMPLATE_NOT_FOUND-5001]
	_ = x[DINGDING_ROBOT_NOT_FOUND-5002]
	_ = x[DINGDING_ROBOT_SEND_FAIL-5003]
	_ = x[SMS_TYPE_NOT_EXISTS-5004]
	_ = x[VCODE_VERIFY_FAILED-5005]
	_ = x[EMAIL_TYPE_NOT_EXISTS-5006]
	_ = x[BANKCARD_AMOUNT_LIMIT-6000]
	_ = x[BANKCARD_EXISTS-6001]
	_ = x[BANKCARD_NOT_EXISTS-6002]
	_ = x[GLOBAL_CONFIG_ERROR-6003]
	_ = x[ORDER_NOT_EXISTS-6004]
	_ = x[ORDER_CANNOT_CANCEL-6005]
	_ = x[PENDING_ORDER_EXIST-6006]
	_ = x[NO_SYSTEM_ACCOUNT-6007]
	_ = x[C2C_FORBIDDEN-6010]
	_ = x[NEED_KYC-6011]
	_ = x[ENGINE_DEDUCT_WITHDRAW_FREEZE_FAIL-7000]
	_ = x[ENGINE_WITHDRAW_FREEZE_FAIL-7001]
	_ = x[ENGINE_DEPOSIT_FAIL-7002]
}

const (
	_ErrCode_name_0 = "内部错误参数错误内部错误内部错误内部错误jwt权限验证失败api频率限制jwt权限不足jwt生成错误"
	_ErrCode_name_1 = "币种不支持提现已暂停余额不足禁止提现"
	_ErrCode_name_2 = "地址非法"
	_ErrCode_name_3 = "获取汇率失败到达每日限额到达总限额提现数量太大用户被禁止币币提现安全设置24小时内不能提现"
	_ErrCode_name_4 = "验证码错误账户已存在创建账户错误创建jwt失败账户不存在登陆失败账户不存在账户已存在配置不存在谷歌验证码已经使用了谷歌验证码验证失败交易密码验证失败手机验证码验证失败邮箱验证码验证失败安全验证失败用户已经绑定手机或者邮箱必须有一个开启用户已经实名验证必须通过一级实名验证用户认证信息没有找到图形验证码错误或失效一级实名认证失败face++签名错误申请的保荐机构正在审核或者已经在使用还没有申请保荐机构还没有申请或者审核保荐机构还没有配置保荐机构的权益交易密码必须设置邀请码不存在还没有正在使用的保荐机构合伙人正在申请升级中安全验证类型错误"
	_ErrCode_name_5 = "获取引擎账户信息失败账户余额不足"
	_ErrCode_name_6 = "没找到邮箱服务供应商邮件模板未找到没找到dingding机器人dingding机器人发送失败短信类型不存在验证码验证失败邮件类型不存在"
	_ErrCode_name_7 = "银行卡总数限制银行卡已经存在银行卡不存在配置错误订单不存在订单不可取消已存在pending订单没有系统银行账户"
	_ErrCode_name_8 = "用户c2c被禁用NEED_KYC"
	_ErrCode_name_9 = "扣除冻结失败冻结失败充值失败"
)

var (
	_ErrCode_index_0 = [...]uint8{0, 12, 24, 36, 48, 60, 81, 96, 111, 126}
	_ErrCode_index_1 = [...]uint8{0, 15, 30, 42, 54}
	_ErrCode_index_3 = [...]uint8{0, 18, 36, 51, 69, 96, 131}
	_ErrCode_index_4 = [...]uint16{0, 15, 30, 48, 63, 78, 90, 105, 120, 135, 165, 192, 216, 243, 270, 288, 306, 345, 369, 399, 429, 429, 459, 483, 501, 555, 582, 621, 657, 681, 699, 735, 765, 789}
	_ErrCode_index_5 = [...]uint8{0, 30, 48}
	_ErrCode_index_6 = [...]uint8{0, 30, 51, 77, 106, 127, 148, 169}
	_ErrCode_index_7 = [...]uint8{0, 21, 42, 60, 72, 87, 105, 127, 151}
	_ErrCode_index_8 = [...]uint8{0, 18, 26}
	_ErrCode_index_9 = [...]uint8{0, 18, 30, 42}
)

func (i ErrCode) String() string {
	switch {
	case 1000 <= i && i <= 1008:
		i -= 1000
		return _ErrCode_name_0[_ErrCode_index_0[i]:_ErrCode_index_0[i+1]]
	case 2000 <= i && i <= 2003:
		i -= 2000
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case i == 2005:
		return _ErrCode_name_2
	case 2009 <= i && i <= 2014:
		i -= 2009
		return _ErrCode_name_3[_ErrCode_index_3[i]:_ErrCode_index_3[i+1]]
	case 3000 <= i && i <= 3032:
		i -= 3000
		return _ErrCode_name_4[_ErrCode_index_4[i]:_ErrCode_index_4[i+1]]
	case 4000 <= i && i <= 4001:
		i -= 4000
		return _ErrCode_name_5[_ErrCode_index_5[i]:_ErrCode_index_5[i+1]]
	case 5000 <= i && i <= 5006:
		i -= 5000
		return _ErrCode_name_6[_ErrCode_index_6[i]:_ErrCode_index_6[i+1]]
	case 6000 <= i && i <= 6007:
		i -= 6000
		return _ErrCode_name_7[_ErrCode_index_7[i]:_ErrCode_index_7[i+1]]
	case 6010 <= i && i <= 6011:
		i -= 6010
		return _ErrCode_name_8[_ErrCode_index_8[i]:_ErrCode_index_8[i+1]]
	case 7000 <= i && i <= 7002:
		i -= 7000
		return _ErrCode_name_9[_ErrCode_index_9[i]:_ErrCode_index_9[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
