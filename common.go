package sl651common


type OpCode int

const (
	Op2F	OpCode	=	0x2F
	Op30		=	0x30
	Op32		=	0x32
	Op33		=	0x33
	Op34		=	0x34
	Op36		=	0x36
	Op37		=	0x37
	Op38		=	0x38
	Op3A		=	0x3A
	Op40		=	0x40
	Op41		=	0x41
	Op42		=	0x42
	Op43		=	0x43
	Op46		=	0x46
	Op4A		=	0x4A
	Op51		=	0x51
        OpE1		=       0xe1
	OpE5		=	0xe5
        OpE8		=       0xe8
        OpE9            =       0xe9
        OpEB		=       0xeb
	OpUnknown	=	0x00
)

type OpDirection int
const (
	OpDirectionUp	OpDirection	=	1
	OpDirectionDown	OpDirection	=	2
	OpDirectionUnknown	OpDirection	=	0
)

type H32Entry struct {
	MsgId		string	//流水，多个32Entry可能对应一个原始消息
	Region		string	//区域信息，该信息由接收程序启动时配置，无需从消息中解析
	GatewayId	string	//网关id（可以用监测站的id）
	GatewayType 	string	//遥测站类型
	ReportTimestamp	uint64	// 上报时间
	EventTimestamp	uint64	//观测时间
	ProcessTimestamp	uint64	// 处理时间，由框架生成
	Content		string	//其他属性，例如tag，json格式，也可以存放原始消息的json
}


type H33Entry struct {
	MsgId		string	//流水，多个33Entry可能对应一个原始消息
	Region		string	//区域信息，该信息由接收程序启动时配置，无需从消息中解析
	GatewayId	string	//网关id（可以用监测站的id）
	GatewayType 	string	//遥测站类型
	ReportTimestamp	uint64	// 上报时间
	EventTimestamp	uint64	//观测时间
	ProcessTimestamp	uint64	// 处理时间，由框架生成
	Content		string	//其他属性，例如tag，json格式，也可以存放原始消息的json
}

type H34Entry struct {
	MsgId		string	//流水，多个34Entry可能对应一个原始消息
	Region		string	//区域信息，该信息由接收程序启动时配置，无需从消息中解析
	GatewayId	string	//网关id（可以用监测站的id）
	GatewayType 	string	//遥测站类型
	ReportTimestamp	uint64	// 上报时间
	EventTimestamp	uint64	//观测时间
	ProcessTimestamp	uint64	// 处理时间，由框架生成
	Content		string	//其他属性，例如tag，json格式，也可以存放原始消息的json
}


type H36Entry struct {
	MsgId		string	//流水，多个36Entry可能对应一个原始消息
	Region		string	//区域信息，该信息由接收程序启动时配置，无需从消息中解析
	GatewayId	string	//网关id（可以用监测站的id）
	GatewayType 	string	//遥测站类型
	ReportTimestamp	uint64	// 上报时间
	EventTimestamp	uint64	//观测时间
	ProcessTimestamp	uint64	// 处理时间，由框架生成
	Content		string	//其他属性，例如tag，json格式，也可以存放原始消息的json
}


type HE8Entry struct {
	MsgId		string	//流水，多个E8Entry可能对应一个原始消息
	Region		string	//区域信息，该信息由接收程序启动时配置，无需从消息中解析
	GatewayId	string	//网关id（可以用监测站的id）
	GatewayType 	string	//遥测站类型
	ReportTimestamp	uint64	// 上报时间
	EventTimestamp	uint64	//观测时间
	ProcessTimestamp	uint64	// 处理时间，由框架生成
	Content		string	//其他属性，例如tag，json格式，也可以存放原始消息的json
}

type MsgBuffer interface {

	// VirtualLength 虚拟长度，虚读后剩余可读数据长度
	VirtualLength() int
	
	// 虚读，不移动 read 指针，需要配合 VirtualFlush 和 VirtualRevert 使用
	VirtualRead(p []byte) (n int, err error)

	// 还原虚读指针
	VirtualFlush()

	// VirtualRevert 还原虚读指针
        VirtualRevert() 
}


