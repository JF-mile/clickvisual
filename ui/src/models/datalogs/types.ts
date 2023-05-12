import {
  HighChartsResponse,
  IndexInfoType,
  LogsResponse,
  StatisticalTableResponse,
  TablesResponse,
} from "@/services/dataLogs";

export interface QueryParams {
  logLibrary?: TablesResponse;
  page?: number;
  pageSize?: number;
  st?: number;
  et?: number;
  kw?: string;
  filters?: string[];
}

export type PaneType = {
  pane: string;
  paneId: string;
  paneType: number;
  start?: number;
  end?: number;
  keyword?: string;
  activeTabKey?: string;
  activeIndex?: number;
  queryType?: string;
  page?: number;
  pageSize?: number;
  logs: LogsResponse | undefined;
  highCharts: HighChartsResponse | undefined;
  querySql?: string;
  logChart?: StatisticalTableResponse;
  desc: string;
  histogramChecked: boolean;
  foldingChecked: boolean;
  mode?: number;
  baseFieldsIndexList?: IndexInfoType[];
  logFieldsIndexList?: IndexInfoType[];
  logState: number;
  relTraceTableId: number;
  columsList: string[];
};

export enum hashType {
  noneSet = 0,
  siphash = 1,
  urlhash = 2,
}

export interface Extra {
  isPaging?: boolean; // 是否是切换页面
  isDisableFilter?: boolean; // 是否启用filter功能 分享页面禁用
  isOnlyLog?: boolean;
  reqParams?: QueryParams; // 请求参数
  analysisField?: string[]; // 拥有的字段
}
