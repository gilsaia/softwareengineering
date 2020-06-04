/**
 * 更新用户
 * @date 2020/06/02
 */
import React from 'react';
import {
    Form,
    Row,
    Col,
    Select,
    InputNumber,
    Input,
} from 'antd';

const { Option } = Select;
const { Item } = Form;

const CreateParkingCar = ({ form }) => {

    return (
        <Form layout='vertical' form={form}>
            <Row gutter={24}>
                <Col span={24}>
                    <Item
                        name='id'
                        label='用户ID'
                    >
                        <Input placeholder='请输入用户ID' disabled />
                    </Item>
                </Col>
                <Col span={24}>
                    <Item
                        name='nick_name'
                        label='用户昵称'
                    >
                        <Input placeholder='请输入用户昵称' />
                    </Item>
                </Col>
                <Col span={24}>
                    <Item
                        name='password'
                        label='用户密码'
                    >
                        <Input placeholder='请输入用户密码(非必填)' />
                    </Item>
                </Col>
                <Col span={24}>
                    <Item
                        name='status'
                        label='账号状态'
                    >
                        <Select
                            style={{ width: '100%' }}
                            placeholder='请选择车位状态'
                        >
                            <Option key='User_Unknown' value='User_Unknown'>未定义</Option>
                            <Option key='User_Unverified' value='User_Unverified'>等待审核</Option>
                            <Option key='User_Verified' value='User_Verified'>审核通过</Option>
                        </Select>
                    </Item>
                </Col>
            </Row>
        </Form>
    );
}

export default CreateParkingCar;