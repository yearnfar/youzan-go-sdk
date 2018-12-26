package lib

const AppIDKey = "app_id"
const MethodKey = "method"
const TimestampKey = "timestamp"
const FormatKey = "format"
const VersionKey = "v"
const SignKey = "sign"
const SignMethodKey = "sign_method"
const TokenKey = "access_token"
const AllowedDeviateSeconds = 600

const ErrSystem = -1
const ErrInvalid_app_id = 40001
const ErrInvalid_app = 40002
const ErrInvalid_timestamp = 40003
const ErrEmpty_signature = 40004
const ErrInvalid_signature = 40005
const ErrInvalid_method_name = 40006
const ErrInvalid_method = 40007
const ErrInvalid_team = 40008
const ErrParameter = 41000
const ErrLogic = 50000
