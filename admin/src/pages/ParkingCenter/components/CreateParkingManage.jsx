/**
 * 新建停车场
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

const CreateParking = ({ form }) => {

    return (
        <Form layout='vertical' form={form}>
            <Row gutter={24}>
                <Col span={24}>
                    <Item
                        name='name'
                        label='停车场名称'
                        rules={[{
                            message: '请输入停车场名称',
                            required: true,
                        }]}
                    >
                        <Input style={{ width: '100%' }} placeholder='请输入停车场名称' />
                    </Item>
                </Col>
                <Col span={12}>
                    <Item
                        name='longitude'
                        label='车位经度'
                        rules={[{
                            message: '请输入车位经度',
                            required: true,
                        }]}
                    >
                        <Input style={{ width: '100%' }} placeholder='请输入车位经度' />
                    </Item>
                </Col>
                <Col span={12}>
                    <Item
                        name='latitude'
                        label='车位纬度'
                        rules={[{
                            message: '请输入车位纬度',
                            required: true,
                        }]}
                    >
                        <Input style={{ width: '100%' }} placeholder='请输入车位纬度' />
                    </Item>
                </Col>
                <Col span={24}>
                    <Item
                        name='charge'
                        label='停车场费用(元)'
                        rules={[{
                            message: '请输入停车场费用',
                            required: true,
                        }]}
                    >
                        <Input style={{ width: '100%' }} placeholder='请输入停车场名称' suffix='元' />
                    </Item>
                </Col>

                <Col span={12}>
                    <Item
                        name='empty_ports'
                        label='空余车位'
                        rules={[{
                            message: '请输入空余车位',
                            required: true,
                        }]}
                    >
                        <InputNumber style={{ width: '100%' }} placeholder='请输入空余车位' />
                    </Item>
                </Col>

                <Col span={12}>
                    <Item
                        name='total_ports'
                        label='总车位'
                        rules={[{
                            message: '请输入总车位',
                            required: true,
                        }]}
                    >
                        <InputNumber style={{ width: '100%' }} placeholder='请输入总车位' />
                    </Item>
                </Col>
                
            </Row>
        </Form>
    );
}

export default CreateParking;