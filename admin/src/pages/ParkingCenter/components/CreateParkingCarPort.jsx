/**
 * 新建车位
 * @date 2020/06/02
 */
import React from 'react';
import {
    Form,
    Row,
    Col,
    Input,
    Select,
    InputNumber,
} from 'antd';

const { Option } = Select;
const { Item } = Form;

const CreateParkingCar = ({ form }) => {

    return (
        <Form layout='vertical' form={form}>
            <Row gutter={24}>
                <Col span={12}>
                    <Item
                        name='longitude'
                        label='车位经度'
                        rules={[{
                            message: '请输入车位经度',
                            required: true,
                            // min: 1,
                            // type: 'number',
                        }]}
                    >
                        <InputNumber style={{ width: '100%' }} placeholder='请输入车位经度' />
                    </Item>
                </Col>
                <Col span={12}>
                    <Item
                        name='latitude'
                        label='车位纬度'
                        rules={[{
                            message: '请输入车位纬度',
                            required: true,
                            // min: 1,
                            // type: 'number',
                        }]}
                    >
                        <InputNumber style={{ width: '100%' }} placeholder='请输入车位纬度' />
                    </Item>
                </Col>
                <Col span={24}>
                    <Item
                        name='park'
                        label='车位所属停车场'
                        rules={[{
                            message: '请输入车位所属停车场',
                            required: true,
                        }]}
                    >
                        <InputNumber style={{ width: '100%' }} placeholder='请输入车位所属停车场' />
                    </Item>
                </Col>
                <Col span={24}>
                    <Item
                        name='state'
                        label='车位状态'
                        rules={[{
                            message: '请选择车位状态',
                            required: true,
                        }]}
                    >
                        <Select
                            style={{ width: '100%' }}
                            placeholder='请选择车位状态'
                        >
                            <Option key='Port_Unknown' value='Port_Unknown'>未定义</Option>
                            <Option key='Port_Empty' value='Port_Empty'>车位空</Option>
                            <Option key='Port_Used' value='Port_Used'>已占用</Option>
                            <Option key='Port_Abandon' value='Port_Abandon'>已废弃</Option>
                        </Select>
                    </Item>
                </Col>
            </Row>
        </Form>
    );
}

export default CreateParkingCar;