import { UserInfoType } from "@/services/systemUser";
import { DeleteOutlined, EditOutlined, RedoOutlined } from "@ant-design/icons";
import { Button, message, Popconfirm, Space, Table, Tooltip } from "antd";
import { useState } from "react";
import { useIntl, useModel } from "umi";
import { userListType } from "..";
import EditUser from "./components/EditUser";

export interface EditUserInfoType extends UserInfoType {
  uid: number;
}

const UserTable = (props: {
  dataObj: { list: userListType[]; total: number };
  loadList: (currentPageInfo: { current?: number; pageSize?: number }) => void;
  setCurrentPagination: (data: { current: number; pageSize: number }) => void;
  currentPagination: { current: number; pageSize: number };
  copyInformation: (res: any, title: string) => void;
}) => {
  const [isEdit, setIsEdit] = useState<boolean>(false);
  const [editUserInfo, setEditUserInfo] = useState<EditUserInfoType>();
  const i18n = useIntl();
  const {
    dataObj,
    loadList,
    currentPagination,
    setCurrentPagination,
    copyInformation,
  } = props;

  const { sysUser } = useModel("system");
  const { doResetUserPassword, doDeleteUser } = sysUser;

  const confirm = (id: number) => {
    doResetUserPassword.run(id).then((res: any) => {
      if (res.code != 0) return;
      copyInformation(res, i18n.formatMessage({ id: "sys.user.resetSuccess" }));
    });
  };

  const deleteUser = (id: number) => {
    doDeleteUser.run(id).then((res: any) => {
      if (res.code != 0) return;
      loadList({});
      message.success(
        i18n.formatMessage({ id: "sys.user.deleteName.success" })
      );
    });
  };

  const column: any = [
    { key: "uid", title: "Uid", dataIndex: "uid" },
    { key: "username", title: "UserName", dataIndex: "username" },
    { key: "nickname", title: "NickName", dataIndex: "nickname" },
    {
      key: "phone",
      title: i18n.formatMessage({ id: "sys.user.table.phone" }),
      dataIndex: "phone",
    },
    { key: "email", title: "Email", dataIndex: "email" },
    {
      title: "Options",
      key: "options",
      render: (_: any, record: userListType) => (
        <Space>
          <Tooltip
            title={i18n.formatMessage({ id: "sys.user.resetPassword" })}
            placement="bottom"
          >
            <Popconfirm
              title={i18n.formatMessage(
                { id: "sys.user.resetTip" },
                { user: record.nickname }
              )}
              onConfirm={() => confirm(record.uid)}
              okText="Yes"
              cancelText="No"
            >
              <Button size={"small"} type={"link"} icon={<RedoOutlined />} />
            </Popconfirm>
          </Tooltip>
          <Tooltip
            title={i18n.formatMessage({ id: "sys.user.deleteName" })}
            placement="bottom"
          >
            <Popconfirm
              title={i18n.formatMessage(
                { id: "sys.user.deleteNameTips" },
                { user: record.nickname }
              )}
              onConfirm={() => deleteUser(record.uid)}
              okText="Yes"
              cancelText="No"
            >
              <Button size={"small"} type={"link"} icon={<DeleteOutlined />} />
            </Popconfirm>
          </Tooltip>
          <Button
            size={"small"}
            type={"link"}
            onClick={() => {
              setIsEdit(true);
              setEditUserInfo({
                uid: record.uid,
                nickname: record.nickname,
                email: record.email,
                phone: record.phone,
              });
            }}
            icon={<EditOutlined />}
          />
        </Space>
      ),
    },
  ];

  return (
    <>
      <EditUser
        open={isEdit}
        editUserInfo={editUserInfo}
        onChangeOpen={(flag: boolean) => {
          setIsEdit(flag);
        }}
        loadList={loadList}
      />

      <Table
        dataSource={dataObj?.list || []}
        columns={column}
        size="small"
        rowKey={(item: any) => item.uid}
        pagination={{
          responsive: true,
          showSizeChanger: true,
          size: "small",
          ...currentPagination,
          total: dataObj?.total || 0,
          onChange: (page, pageSize) => {
            setCurrentPagination({
              ...currentPagination,
              current: page,
              pageSize,
            });
            loadList({
              current: page,
              pageSize,
            });
          },
        }}
      />
    </>
  );
};
export default UserTable;
