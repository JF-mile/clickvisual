package view

import (
	db2 "github.com/clickvisual/clickvisual/api/internal/pkg/model/db"
)

type ReqCreateFolder struct {
	Iid       int `json:"iid" form:"iid" binding:"required"`
	Primary   int `json:"primary" form:"primary" binding:"required"`
	Secondary int `json:"secondary" form:"secondary"`

	ReqUpdateFolder
}

type ReqUpdateFolder struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Desc       string `json:"desc" form:"desc"`
	ParentId   int    `json:"parentId" form:"parentId"`
	WorkflowId int    `json:"workflowId" form:"workflowId"`
}

type RespListFolder struct {
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	Desc      string             `json:"desc"`
	Primary   int                `json:"primary"`
	Secondary int                `json:"secondary"`
	ParentId  int                `json:"parentId"`
	Children  []RespListFolder   `json:"children"`
	Nodes     []*db2.BigdataNode `json:"nodes"`
}

type RespInfoFolder struct {
	db2.BigdataFolder
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
}

type (
	ReqCreateSource struct {
		Iid int `json:"iid" form:"iid" binding:"required"`
		ReqUpdateSource
	}
	ReqUpdateSource struct {
		Name     string `json:"name" form:"name" binding:"required"`
		Desc     string `json:"desc" form:"desc"`
		URL      string `json:"url" form:"url"`
		UserName string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		Typ      int    `json:"typ" form:"typ"`
	}
	ReqListSource struct {
		Iid  int    `json:"iid" form:"iid" binding:"required"`
		Typ  int    `json:"typ" form:"typ"`
		Name string `json:"name" form:"name"`
	}
	ReqListSourceTable struct {
		Database string `json:"database" form:"database" binding:"required"`
	}
	ReqListSourceColumn struct {
		Database string `json:"database" form:"database" binding:"required"`
		Table    string `json:"table" form:"table" binding:"required"`
	}
	Column struct {
		Field   string `json:"field" form:"field"`
		Type    string `json:"type" form:"type"`
		Comment string `json:"comment" form:"comment"`
	}
)

type ReqCreateWorkflow struct {
	Iid int `json:"iid" form:"iid" binding:"required"`
	ReqUpdateSource
}

type ReqUpdateWorkflow struct {
	Name string `json:"name" form:"name" binding:"required"`
	Desc string `json:"desc" form:"desc"`
}

type ReqListWorkflow struct {
	Iid int `json:"iid" form:"iid" binding:"required"`
}

type (
	// ReqCreateNode Node
	ReqCreateNode struct {
		Primary    int `json:"primary" form:"primary" binding:"required"`
		Secondary  int `json:"secondary" form:"secondary" binding:"required"`
		Tertiary   int `json:"tertiary" form:"tertiary"`
		Iid        int `json:"iid" form:"iid" binding:"required"`
		WorkflowId int `json:"workflowId" form:"workflowId"`
		SourceId   int `json:"sourceId" form:"sourceId"`
		ReqUpdateNode
	}

	ReqUpdateNode struct {
		FolderId int    `json:"folderId" form:"folderId"`
		Name     string `json:"name" form:"name"`
		Desc     string `json:"desc" form:"desc"`
		Content  string `json:"content" form:"content"`
		SourceId int    `json:"sourceId" form:"sourceId"`
		Tertiary int    `json:"tertiary" form:"tertiary"`
	}

	RespCreateNode struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Desc    string `json:"desc"`
		Content string `json:"content"`
		LockUid int    `json:"lockUid"`
		LockAt  int64  `json:"lockAt"`
	}

	ReqListNode struct {
		Iid        int `json:"iid" form:"iid"  binding:"required"`
		Primary    int `json:"primary" form:"primary" binding:"required"`
		Secondary  int `json:"secondary" form:"secondary"`
		FolderId   int `json:"folderId" form:"folderId"`
		WorkflowId int `json:"workflowId" form:"workflowId"`
	}

	RespListNode struct {
		FolderId int    `json:"folderId"`
		Name     string `json:"name"`
		Desc     string `json:"desc"`
		Uid      int    `json:"uid"`
		UserName string `json:"userName"`
	}

	RespInfoNode struct {
		Id              int    `json:"id"`
		Name            string `json:"name"`
		Desc            string `json:"desc"`
		Content         string `json:"content"`
		LockUid         int    `json:"lockUid"`
		LockAt          int64  `json:"lockAt"`
		Username        string `json:"username"`
		Nickname        string `json:"nickname"`
		Status          int    `json:"status"`
		PreviousContent string `json:"previousContent"`
		Result          string `json:"result"`
	}

	RunNodeResult struct {
		Logs           []map[string]interface{} `json:"logs"`
		InvolvedSQLs   map[string]string        `json:"involvedSQLs"`
		Message        string                   `json:"message"`
		DagFailedNodes map[int]string           `json:"dagFailedNodes"`
	}

	RespRunNode struct {
		Result string `json:"result"`
		Status int    `json:"status"`
	}

	SyncContent struct {
		Source  IntegrationFlat      `json:"source"`
		Target  IntegrationFlat      `json:"target"`
		Mapping []IntegrationMapping `json:"mapping"`
	}
	// IntegrationFlat integration offline sync step 1
	IntegrationFlat struct {
		Typ      string `json:"typ"` // clickhouse mysql
		SourceId int    `json:"sourceId"`
		Cluster  string `json:"cluster"`
		Database string `json:"database"`
		Table    string `json:"table"`

		SourceFilter string `json:"sourceFilter"`

		// Deprecated
		TargetBefore string `json:"targetBefore"`
		// Deprecated
		TargetAfter string `json:"targetAfter"`

		TargetBeforeList []string `json:"targetBeforeList"`
		TargetAfterList  []string `json:"targetAfterList"`
	}
	// IntegrationMapping integration offline sync step 2
	IntegrationMapping struct {
		Source     string `json:"source"`
		SourceType string `json:"sourceType"`
		Target     string `json:"target"`
		TargetType string `json:"targetType"`
	}

	InnerNodeRun struct {
		N  *db2.BigdataNode
		NC *db2.BigdataNodeContent
	}

	ReqNodeRunOpenAPI struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	ReqNodeHistoryList struct {
		db2.ReqPage

		IsExcludeCrontabResult int `json:"isExcludeCrontabResult" form:"isExcludeCrontabResult"`
	}

	NodeHistoryItem struct {
		UUID     string `json:"uuid"`
		Utime    int64  `json:"utime"`
		Uid      int    `json:"uid"`
		UserName string `json:"userName"`
		Nickname string `json:"nickname"`
	}

	RespNodeHistoryList struct {
		Total int64             `json:"total"`
		List  []NodeHistoryItem `json:"list"`
	}

	ReqNodeResultList struct {
		db2.ReqPage
	}

	RespNodeResult struct {
		ID           int    `json:"id"`
		Ctime        int64  `json:"ctime"`
		NodeId       int    `json:"nodeId"`
		Content      string `json:"content,omitempty"`
		Result       string `json:"result,omitempty"`
		Cost         int64  `json:"cost,omitempty"`
		ExcelProcess string `json:"excelProcess,omitempty"`
		Status       int    `json:"status"`
		RespUserSimpleInfo
	}

	RespNodeResultList struct {
		Total int64            `json:"total"`
		List  []RespNodeResult `json:"list"`
	}
)

type (
	WorkerStats struct {
		Iid  int
		Uid  int
		Data map[int64]WorkerStatsRow
	}
	WorkerStatsRow struct {
		Timestamp int64 `json:"timestamp"`
		Unknown   int   `json:"unknown"`
		Failed    int   `json:"failed"`
		Success   int   `json:"success"`
	}
	// ReqWorkerDashboard Request start and end time
	ReqWorkerDashboard struct {
		Start      int64 `json:"start" form:"start"`
		End        int64 `json:"end" form:"end"`
		Iid        int   `json:"iid" form:"iid"`
		IsInCharge int   `json:"isInCharge" form:"isInCharge"`
	}
	RespWorkerDashboard struct {
		NodeFailed    int              `json:"nodeFailed"`    // node status
		NodeSuccess   int              `json:"nodeSuccess"`   // node status
		NodeUnknown   int              `json:"nodeUnknown"`   // node status
		WorkerFailed  int              `json:"workerFailed"`  // Execution status of each periodic schedule
		WorkerSuccess int              `json:"workerSuccess"` // Execution status of each periodic schedule
		WorkerUnknown int              `json:"workerUnknown"` // Execution status of each periodic schedule
		Flows         []WorkerStatsRow `json:"flows"`         // Execution trend chart
	}

	ReqWorkerList struct {
		Start    int    `json:"start" form:"start"`
		End      int    `json:"end" form:"end"`
		NodeName string `json:"nodeName" form:"nodeName"`
		Tertiary int    `json:"tertiary" form:"tertiary"` // ClickHouse 10; MySQL 11; OfflineSync 20
		Iid      int    `json:"iid" form:"iid"`
		Status   int    `json:"status" form:"status"` // 0 未知；1 成功；2 失败
		Pagination
	}

	ReqStructuralTransfer struct {
		Source  string   `json:"source" form:"source"`
		Target  string   `json:"target" form:"target"`
		Columns []Column `json:"columns" form:"columns"`
	}

	RespWorkerRow struct {
		NodeName  string `json:"nodeName"`
		Status    int    `json:"status"` // unknown 0; success 1; failed 2
		Tertiary  int    `json:"tertiary"`
		Crontab   string `json:"crontab"`
		StartTime int64  `json:"startTime"`
		EndTime   int64  `json:"endTime"`
		Iid       int    `json:"iid"`
		ID        int    `json:"id"`
		NodeId    int    `json:"nodeId"`
		Cost      int64  `json:"cost"`

		ChargePerson RespUserSimpleInfo `json:"chargePerson"`
	}

	RespWorkerList struct {
		Total int64           `json:"total"`
		List  []RespWorkerRow `json:"list"`
	}
)

func (s *SyncContent) Cluster() string {
	if s.Target.Typ == "clickhouse" {
		return s.Target.Cluster
	}
	if s.Source.Typ == "clickhouse" {
		return s.Source.Cluster
	}
	return ""
}

// crontab struct
type (
	ReqCreateCrontab struct {
		ReqUpdateCrontab
	}
	ReqUpdateCrontab struct {
		Desc          string          `json:"desc" form:"desc"`
		DutyUid       int             `json:"dutyUid" form:"dutyUid" required:"true"`
		Cron          string          `json:"cron" form:"cron" required:"true"`
		Typ           int             `json:"typ" form:"typ"`
		Args          []ReqCrontabArg `json:"args" form:"args"`
		IsRetry       int             `json:"isRetry" form:"isRetry"` // isRetry: 0 no 1 yes
		RetryTimes    int             `json:"retryTimes" form:"retryTimes"`
		RetryInterval int             `json:"retryInterval" form:"retryInterval"` // retryInterval: the unit is in seconds, 100 means 100s
		ChannelIds    []int           `json:"channelIds" form:"channelIds"`
	}
	ReqCrontabArg struct {
		Key string `json:"key" form:"key"`
		Val string `json:"val" form:"val"`
	}
	ReqNodeRunResult struct {
		ExcelProcess string `json:"excelProcess" form:"excelProcess"`
	}
)

// DAG ...
type (
	ReqDAG struct {
		BoardNodeList []ReqDagNode `json:"boardNodeList"`
		BoardEdges    []ReqDagEdge `json:"boardEdges"`
	}
	ReqDagNode struct {
		Id int `json:"id"` // node id
	}
	ReqDagEdge struct {
		Source string `json:"source"`
		Target string `json:"target"`
	}
	DagExecFlow struct {
		NodeId   int           `json:"nodeId"`
		Children []DagExecFlow `json:"children"`
	}
)

type ReqTableDependencies struct {
	DatabaseName string `json:"databaseName" form:"databaseName" binding:"required"`
	TableName    string `json:"tableName" form:"tableName" binding:"required"`
}

type RespTableDependencies struct {
	Utime int64           `json:"utime"`
	Data  []RespTableDeps `json:"data"`
}
