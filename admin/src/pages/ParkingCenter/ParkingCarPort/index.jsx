/**
 * 车位管理
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

import { queryParkingCarList, createParkingCar } from '@/services/parkingCenter';
import CreateParkingCar from '../components/CreateParkingCarPort';

const ParkingManage = () => {

    const [visible, setVisible] = useState(false);
    const [form] = Form.useForm();

    const actionRef = useRef();

    useEffect(() => {

    });

    const columns = [
        {
            title: '车位ID',
            dataIndex: 'id',
        },
        {
            title: '车位状态',
            dataIndex: 'state',
            valueEnum: {
                'Port_Empty': '车位空',
                'Port_Used': '已占用',
                'Port_Unknown': '未定义',
                'Port_Abandon': '已废弃',
            },
        },
        {
            title: '车位纬度',
            dataIndex: 'latitude',
        },
        {
            title: '车位经度',
            dataIndex: 'longitude',
        },
        {
            title: '车位所属停车场',
            dataIndex: 'park',
        },
    ];

    // 新建车位
    const handleCreateSubmit = () => {
        form.validateFields().then(values => {
            createParkingCar({
                port: { ...values },
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
                headerTitle="车位管理"
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
                        title='新建车位'
                        destroyOnClose
                        // width={700}
                        maskClosable={false}
                        onCancel={() => setVisible(false)}
                        onOk={handleCreateSubmit}
                    >
                        <CreateParkingCar form={form} />
                    </Modal>
                ]}
                request={params => queryParkingCarList({ count: params.pageSize, num: params.current - 1 })}
                columns={columns}
            />
        </PageHeaderWrapper>
    );
}

export default ParkingManage;