// @ts-nocheck
import { ApplyPluginsType, dynamic } from 'D:/learn/softwareproject/softwareengineering/admin/node_modules/@umijs/runtime';
import { plugin } from './plugin';

const routes = [
  {
    "path": "/user",
    "component": dynamic({ loader: () => import(/* webpackChunkName: 'layouts__UserLayout' */'D:/learn/softwareproject/softwareengineering/admin/src/layouts/UserLayout'), loading: require('@/components/PageLoading/index').default}),
    "routes": [
      {
        "name": "login",
        "path": "/user/login",
        "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__user__login' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/user/login'), loading: require('@/components/PageLoading/index').default}),
        "exact": true
      }
    ]
  },
  {
    "path": "/",
    "component": dynamic({ loader: () => import(/* webpackChunkName: 'layouts__SecurityLayout' */'D:/learn/softwareproject/softwareengineering/admin/src/layouts/SecurityLayout'), loading: require('@/components/PageLoading/index').default}),
    "routes": [
      {
        "path": "/",
        "component": dynamic({ loader: () => import(/* webpackChunkName: 'layouts__BasicLayout' */'D:/learn/softwareproject/softwareengineering/admin/src/layouts/BasicLayout'), loading: require('@/components/PageLoading/index').default}),
        "authority": [
          "admin",
          "user"
        ],
        "routes": [
          {
            "path": "/",
            "redirect": "/ParkingCenter/ParkingManage",
            "exact": true
          },
          {
            "path": "/ParkingCenter",
            "name": "车位中心",
            "icon": "appstore",
            "routes": [
              {
                "path": "/ParkingCenter/ParkingManage",
                "name": "停车场管理",
                "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__ParkingCenter__ParkingManage__index' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/ParkingCenter/ParkingManage/index'), loading: require('@/components/PageLoading/index').default}),
                "exact": true
              },
              {
                "path": "/ParkingCenter/ParkingCarPort",
                "name": "车位管理",
                "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__ParkingCenter__ParkingCarPort__index' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/ParkingCenter/ParkingCarPort/index'), loading: require('@/components/PageLoading/index').default}),
                "exact": true
              },
              {
                "path": "/ParkingCenter/ParkingBill",
                "name": "账单管理",
                "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__ParkingCenter__ParkingBill__index' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/ParkingCenter/ParkingBill/index'), loading: require('@/components/PageLoading/index').default}),
                "exact": true
              }
            ]
          },
          {
            "path": "/ParkingUser",
            "name": "用户管理",
            "icon": "user",
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__ParkingUser__index' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/ParkingUser/index'), loading: require('@/components/PageLoading/index').default}),
            "exact": true
          },
          {
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__404' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/404'), loading: require('@/components/PageLoading/index').default}),
            "exact": true
          }
        ]
      },
      {
        "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__404' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/404'), loading: require('@/components/PageLoading/index').default}),
        "exact": true
      }
    ]
  },
  {
    "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__404' */'D:/learn/softwareproject/softwareengineering/admin/src/pages/404'), loading: require('@/components/PageLoading/index').default}),
    "exact": true
  }
];

// allow user to extend routes
plugin.applyPlugins({
  key: 'patchRoutes',
  type: ApplyPluginsType.event,
  args: { routes },
});

export { routes };
