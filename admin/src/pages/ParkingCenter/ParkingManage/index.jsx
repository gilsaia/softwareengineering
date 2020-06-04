/**
 * 停车场管理
 * @date 2020/06/01
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

import { queryParkingList, createParking } from '@/services/parkingCenter';
import CreateParkingManage from '../components/CreateParkingManage';

const ParkingManage = () => {

    const [visible, setVisible] = useState(false);
    const [form] = Form.useForm();

    const actionRef = useRef();


    const columns = [
        {
            title: '停车场ID',
            dataIndex: 'id',
        },
        {
            title: '停车场名称',
            dataIndex: 'name',
        },
        {
            title: '停车场纬度',
            dataIndex: 'latitude',
        },
        {
            title: '停车场经度',
            dataIndex: 'longitude',
        },
        {
            title: '停车场费用(元)',
            dataIndex: 'charge',
        },
        {
            title: '空余车位',
            dataIndex: 'empty_ports',
        },
        {
            title: '总车位',
            dataIndex: 'total_ports',
        },
    ];

    // 新建停车场
    const handleCreateSubmit = () => {
        form.validateFields().then(values => {
            createParking({
                park: { ...values },
            }).then((response) => {
                if (response && response.err_no === 0) {
                    message.success('新建成功');
                    setVisible(false);
                    actionRef.current.reload();
                    form.resetFields();
                } else {
                    message.error(response.err_tips || 'Token令牌失效或者服务器接口错误');
                }
            });
        });

    }
    return (
        <PageHeaderWrapper>
            <ProTable
                headerTitle="停车场管理"
                rowKey="id"
                actionRef={actionRef}
                pagination={{
                    showSizeChanger: true,
                    defaultPageSize: 10,
                    pageSizeOptions: ['5', '10', '20', '50'],
                }}
                search={false}
                toolBarRender={(action, { selectedRows }) => [
                    <Button type="primary" onClick={() => setVisible(true)}>
                        <PlusOutlined /> 新建
                    </Button>,
                    <Modal
                        visible={visible}
                        title='新建停车场'
                        destroyOnClose
                        width={700}
                        maskClosable={false}
                        onCancel={() => setVisible(false)}
                        onOk={handleCreateSubmit}
                    >
                        <CreateParkingManage form={form} />
                    </Modal>
                ]}
                request={params => queryParkingList({ count: params.pageSize, num: params.current - 1 })}
                columns={columns}
            />
        </PageHeaderWrapper>
    );
}

export default ParkingManage;