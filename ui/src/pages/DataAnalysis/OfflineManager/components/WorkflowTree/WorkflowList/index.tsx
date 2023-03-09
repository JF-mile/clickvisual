import WorkflowLine from "@/pages/DataAnalysis/OfflineManager/components/WorkflowTree/WorkflowList/WorkflowLine";
import offlineStyles from "@/pages/DataAnalysis/OfflineManager/index.less";
import { useModel } from "@umijs/max";
import { useMemo } from "react";

export interface WorkflowListType {}

const WorkflowList = (props: WorkflowListType) => {
  const { workflowList } = useModel("dataAnalysis", (model) => ({
    workflowList: model.workflow.workflowList,
  }));

  const List = useMemo(() => {
    if (workflowList.length <= 0) return null;
    return workflowList.map((workflow) => (
      <WorkflowLine key={workflow.id} workflow={workflow} />
    ));
  }, [workflowList]);

  return <div className={offlineStyles.workflowList}>{List}</div>;
};

export default WorkflowList;
