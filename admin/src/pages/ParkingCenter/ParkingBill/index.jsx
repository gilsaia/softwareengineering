/**
 * 账单管理
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

import { queryParkingBillList } from '@/services/parkingCenter';

const ParkingBill = () => {

    const [visible, setVisible] = useState(false);
    const [form] = Form.useForm();

    const actionRef = useRef();


    const columns = [
        {
            title: '账单ID',
            dataIndex: 'id',
        },
        {
            title: '用户ID',
            dataIndex: 'user_id',
        },
        {
            title: '停车场ID',
            dataIndex: 'park_id',
        },
        {
            title: '停车时长',
            dataIndex: 'duration',
        },
        {
            title: '停车场费用(元)',
            dataIndex: 'charge',
        },
        {
            title: '账单状态',
            dataIndex: 'status',
            valueEnum: {
                'Bill_Unknown': '未知',
                'Bill_WaitToPay': '待付款',
                'Bill_Paid': '已付款',
            },
        },
    ];

    return (
        <PageHeaderWrapper>
            <ProTable
                headerTitle="账单管理"
                rowKey="id"
                actionRef={actionRef}
                pagination={{
                    showSizeChanger: true,
                    defaultPageSize: 10,
                    pageSizeOptions: ['5', '10', '20', '50'],
                }}
                search={false}
                request={params => queryParkingBillList({ count: params.pageSize, num: params.current - 1 })}
                columns={columns}
            />
        </PageHeaderWrapper>
    );
}

export default ParkingBill;