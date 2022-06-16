export enum BusinessEngineEnum {
  Kafka = "Kafka",
  MergeTree = "MergeTree",
  Distributed = "Distributed",
}

export enum BigDataNavEnum {
  /**
   * 实时业务
   */
  RealTimeTrafficFlow = "realtime",
  /**
   * 临时查询
   */
  TemporaryQuery = "short",
  /**
   * 离线查询
   */
  OfflineManage = "offline",
  /**
   * 数据源管理
   */
  DataSourceManage = "datasourceManage",
}

export enum FolderEnums {
  /**
   * 节点 可在右侧打开
   */
  node = 1,

  /**
   * 文件夹 不可在右侧打开
   */
  folder = 2,
}

export enum PrimaryEnums {
  /**
   * 数据开发
   */
  mining = 1,

  /**
   * 临时查询
   */
  short = 3,
}

export enum SecondaryEnums {
  /**
   * 数据库
   */
  database = 1,
  /**
   * 数据集成
   */
  dataIntegration = 2,
}

export enum TertiaryEnums {
  /**
   * clickhouse
   */
  clickhouse = 10,
  /**
   * mysql
   */
  mysql = 11,
  /**
   * 离线分析
   */
  offline = 20,
  /**
   * 实时分析
   */
  realtime = 21,
}

export enum OfflineRightMenuClickSourceEnums {
  // 业务流程
  workflow = "workflow",
}