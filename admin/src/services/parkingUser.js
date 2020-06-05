import request from '@/utils/request';
import { message } from 'antd';

// 获取用户列表
export async function queryUserList(data) {
    return request('/apiv2/user/MGetUser', {
        method: 'POST',
        data,
    }).then(response => {
        if (response.err_no === 0) {
            return {
                data: response.users.sort((a, b) => b.id - a.id),
                success: true,
                page: data.num,
                total: response.count,
            };
        } else {
            message.error(response.err_tips || 'Token令牌失效或者服务器接口错误');
        }
    });
}

// 更新用户
export async function updateParkingUser(data) {
    return request('/apiv2/user/UpdateUser', {
        method: 'POST',
        data,
    });
}