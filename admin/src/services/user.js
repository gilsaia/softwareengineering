import request from '@/utils/request';
import qs from 'querystring';

export async function query() {
  return request('/api/users');
}
export async function queryCurrent() {
  return request('/api/currentUser');
}
export async function queryNotices() {
  return request('/api/notices');
}

// 获取登录人信息
export async function queryUserInfo(params) {
  return request(`/apiv2/user/GetUser?${qs.stringify({ ...params, id: 0 })}`);
}