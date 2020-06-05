import { stringify } from 'querystring';
import { history } from 'umi';
import { fakeAccountLogin, login } from '@/services/login';
import { setAuthority } from '@/utils/authority';
import { getPageQuery } from '@/utils/utils';
import { message } from 'antd';
import md5 from 'md5';

const Model = {
  namespace: 'login',
  state: {
    status: undefined,
    cellphone: null,
  },
  effects: {
    *login({ payload }, { call, put }) {

      // 密码进行md5加密
      const queryParams = { ...payload, password: md5(payload.password) };
      console.log(payload);
      // const response = yield call(login, queryParams); // 需要加密时，用该行代码
      const response = yield call(login, payload); // 需要加密时，将该行代码删除
      console.log(response);
      if (response.err_no === 0) {
        console.log(33);
        yield put({
          type: 'save',
          payload: {
            token: response.token,
            status: 'success', // 登录成功
            cellphone: payload.cellphone,
          },
        }); // Login successfully
        console.log(66);
        history.push('/');
      } else {
        message.error(response.err_tips);
      }

    },

    logout() {
      const { redirect } = getPageQuery(); // Note: There may be security issues, please note
      localStorage.removeItem('token'); // 移除token
      localStorage.removeItem('cellphone'); // 移除phone
      if (window.location.pathname !== '/user/login' && !redirect) {
        history.replace({
          pathname: '/user/login',
          search: stringify({
            redirect: window.location.href,
          }),
        });
      }
    },
  },
  reducers: {
    save(state, { payload }) {
      localStorage.setItem('token', payload.token);
      localStorage.setItem('cellphone', payload.cellphone);
      return { ...state, ...payload };
    },
  },
};
export default Model;
