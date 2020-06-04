// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';

const { REACT_APP_ENV } = process.env;
export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  locale: {
    // default zh-CN
    default: 'zh-CN',
    // default true, when it is true, will use `navigator.language` overwrite default
    antd: true,
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@/components/PageLoading/index',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes: [
    {
      path: '/user',
      component: '../layouts/UserLayout',
      routes: [
        {
          name: 'login',
          path: '/user/login',
          component: './user/login',
        },
      ],
    },
    {
      path: '/',
      component: '../layouts/SecurityLayout',
      routes: [
        {
          path: '/',
          component: '../layouts/BasicLayout',
          authority: ['admin', 'user'],
          routes: [
            {
              path: '/',
              redirect: '/ParkingCenter/ParkingManage',
            },
            {
              path: '/ParkingCenter',
              name: '车位中心',
              icon: 'appstore',
              routes: [
                {
                  path: '/ParkingCenter/ParkingManage',
                  name: '停车场管理',
                  component: './ParkingCenter/ParkingManage/index',
                },
                {
                  path: '/ParkingCenter/ParkingCarPort',
                  name: '车位管理',
                  component: './ParkingCenter/ParkingCarPort/index',
                },
                {
                  path: '/ParkingCenter/ParkingBill',
                  name: '账单管理',
                  component: './ParkingCenter/ParkingBill/index',
                },
              ],
            },
            {
              path: '/ParkingUser',
              name: '用户管理',
              icon: 'user',
              component: './ParkingUser/index',
            },
            {
              component: './404',
            },
          ],
        },
        {
          component: './404',
        },
      ],
    },
    {
      component: './404',
    },
  ],
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    // ...darkTheme,
    'primary-color': defaultSettings.primaryColor,
  },
  // @ts-ignore
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
});
