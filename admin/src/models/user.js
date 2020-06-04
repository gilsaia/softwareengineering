import { queryCurrent, queryUserInfo, query as queryUsers } from '@/services/user';

const UserModel = {
  namespace: 'user',
  state: {
    currentUser: {},
  },
  effects: {
    *fetch(_, { call, put }) {
      const response = yield call(queryUsers);
      yield put({
        type: 'save',
        payload: response,
      });
    },

    *fetchCurrent({ payload }, { call, put }) {
      const response = yield call(queryUserInfo, payload);
      if(response && response.user){
        yield put({
          type: 'save',
          payload: {
            currentUser: {
              name: response.user.nick_name,
            },
          },
        });
      }
    },
  },
  reducers: {
    save(state, action) {
      return { ...state, ...action.payload };
    },

    saveCurrentUser(state, action) {
      return { ...state, currentUser: action.payload || {} };
    },

    changeNotifyCount(
      state = {
        currentUser: {},
      },
      action,
    ) {
      return {
        ...state,
        currentUser: {
          ...state.currentUser,
          notifyCount: action.payload.totalCount,
          unreadCount: action.payload.unreadCount,
        },
      };
    },
  },
};
export default UserModel;
