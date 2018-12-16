package msg

const(
	User              = 201
	Pass              = 202
	Cd                = 203
	Pwd               = 204
)

const(
	AuthFail          = 301
	AuthOk            = 302   //通过认证
	NoAccess          = 303  //无权限
	NoUser            = 304
    NotFound          = 305
    AlreadyExist      = 306
)