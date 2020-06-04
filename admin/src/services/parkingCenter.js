import request from '@/utils/request';
import { message } from 'antd';

// 获取车位列表
export async function queryParkingCarList(data) {
    return request('/apiv2/car/MGetCarPort', {
        method: 'POST',
        data,
    }).then(response => {
        if (response.err_no === 0) {
            return {
                data: response.ports.sort((a, b) => b.id - a.id),
                success: true,
                page: data.num,
                total: response.count,
            };
        } else {
            message.error(response.err_tips || 'Token令牌失效或者服务器接口错误');
        }
    });
}

// 新增车位
export async function createParkingCar(data) {
    return request('/apiv2/car/AddCarPort', {
        method: 'POST',
        data,
    });
}


/***********以下是停车场接口哦*************/


// 获取停车场列表
export async function queryParkingList(data) {
    return request('/apiv2/park/MGetPark', {
        method: 'POST',
        data,
    }).then(response => {
        if (response.err_no === 0) {
            return {
                data: response.parks.sort((a, b) => b.id - a.id),
                success: true,
                page: data.num,
                total: response.count,
            };
        } else {
            message.error(response.err_tips || 'Token令牌失效或者服务器接口错误');
        }
    });
}

// 新增停车场
export async function createParking(data) {
    return request('/apiv2/park/AddPark', {
        method: 'POST',
        data,
    });
}


/***********以下是账单管理接口哦*************/

// 获取账单管理列表
export async function queryParkingBillList(data) {
    return request('/apiv2/bill/MGetBill', {
        method: 'POST',
        data,
    }).then(response => {
        if (response.err_no === 0) {
            return {
                data: response.bills.sort((a, b) => b.id - a.id),
                success: true,
                page: data.num,
                total: response.count,
            };
        } else {
            message.error(response.err_tips || 'Token令牌失效或者服务器接口错误');
        }
    });
}