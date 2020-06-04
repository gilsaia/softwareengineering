/**
 * 用户管理
 * @date 2020/06/02
 */
import React, { useEffect, useState, useRef } from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { PlusOutlined } from '@ant-design/icons';
import {
    Button,
    Modal,
    Form,
    message,
} from 'antd';

import { queryUserList, updateParkingUser } from '@/services/parkingUser';
import CreateParkingUser from './components/CreateParkingUser';

const ParkingUserManage = () => {

    const [visible, setVisible] = useState(false);
    const [form] = Form.useForm();

    const actionRef = useRef();

    const columns = [
        {
            title: '用户ID',
            dataIndex: 'id',
        },
        {
            title: '用户昵称',
            dataIndex: 'nick_name',
        },
        {
            title: '手机号码',
            dataIndex: 'cellphone',
        },
        {
            title: '用户状态',
            dataIndex: 'status',
            valueEnum: {
                'User_Unknown': '未定义',
                'User_Unverified': '等待审核',
                'User_Verified': '审核通过',
            },
        },
        {
            title: '操 作',
            render: (text, record) => {
                return (
                    <div>
                        <Button type='primary' onClick={() => handleEdit(record)}>更新</Button>
                    </div>
                );
            },
        }
    ];

    // 编辑提交
    const handleUpdateSubmit = () => {
        form.validateFields().then(values => {
            const queryParams = { ...values };
            if (queryParams.password) delete queryParams.password;
            updateParkingUser({
                password: values.password ? values.password : undefined,
                user: queryParams,
            }).then((response) => {
                if (response && response.err_no === 0) {
                    message.success('更新成功');
                    setVisible(false);
                    actionRef.current.reload();
                    form.resetFields();
                } else {
                    message.error(response.err_tips || 'Token令牌失效或者服务器接口错误');
                }
            });
        });
    }

    // 打开弹窗
    const handleEdit = (record) => {
        form.setFieldsValue({
            id: record.id,
            nick_name: record.nick_name,
            status: record.status,
        });
        setVisible(true);
    }

    return (
        <PageHeaderWrapper>
            <ProTable
                headerTitle="用户管理"
                rowKey="id"
                actionRef={actionRef}
                pagination={{
                    showSizeChanger: true,
                    defaultPageSize: 10,
                    pageSizeOptions: ['5', '10', '20', '50'],
                }}
                search={false}
                toolBarRender={(action, { selectedRows }) => [
                    <Modal
                        visible={visible}
                        title='更新用户'
                        destroyOnClose
                        maskClosable={false}
                        onCancel={() => setVisible(false)}
                        onOk={handleUpdateSubmit}
                    >
                        <CreateParkingUser form={form} />
                    </Modal>
                ]}
                request={params => queryUserList({ count: params.pageSize, num: params.current - 1 })}
                columns={columns}
            />
        </PageHeaderWrapper>
    );
}

export default ParkingUserManage;