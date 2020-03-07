import {
  fetchGroups,
  fetchSelectors,
  addSelector,
  editSelector,
  updateSelector,
  deleteSelector,
  updateAllSelectors,
} from '../service';
import { ClashProxyGroupSelector } from '../typing';

const Model = {
  namespace: 'selectors',

  state: {
    selectors: [],
    groups: [],
  },

  effects: {
    *fetchGroups(_: any, { call, put }: any) {
      const response = yield call(fetchGroups);
      yield put({
        type: 'saveGroups',
        payload: response.data,
      });
    },
    *fetchSelectors(_: any, { call, put }: any) {
      const response = yield call(fetchSelectors);
      yield put({
        type: 'saveSelectors',
        payload: response.data,
      });
    },
    *addSelector({ payload, callback }: any, { call, put }: any) {
      const response = yield call(addSelector, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveSelectors',
          payload: response.data,
        });
      }
      if (callback) {
        if (response.status === 'success') {
          callback(true);
        } else {
          response.msg ? callback(false, response.msg) : callback(false);
        }
      }
    },
    *editSelector({ payload, callback }: any, { call, put }: any) {
      const response = yield call(editSelector, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveSelectors',
          payload: response.data,
        });
      }
      if (callback) {
        if (response.status === 'success') {
          callback(true);
        } else {
          response.msg ? callback(false, response.msg) : callback(false);
        }
      }
    },
    *updateSelector({ payload, callback }: any, { call, put }: any) {
      const response = yield call(updateSelector, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveSelector',
          payload: response.data,
        });
      }
      if (callback) {
        if (response.status === 'success') {
          callback(true);
        } else {
          response.msg ? callback(false, response.msg) : callback(false);
        }
      }
    },
    *deleteSelector({ payload, callback }: any, { call, put }: any) {
      const response = yield call(deleteSelector, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveSelectors',
          payload: response.data,
        });
      }
      if (callback) {
        if (response.status === 'success') {
          callback(true);
        } else {
          response.msg ? callback(false, response.msg) : callback(false);
        }
      }
    },
    *updateAllSelectors({ _, callback }: any, { call, put }: any) {
      const response = yield call(updateAllSelectors);
      // if (response.status === 'success') {
      yield put({
        type: 'saveSelectors',
        payload: response.data,
      });
      // }
      if (callback) {
        // if (response.status === 'success') {
        //   callback(true);
        // } else {
        //   response.msg ? callback(false, response.msg) : callback(false);
        // }
        response.msg ? callback(false, response.msg) : callback(true);
      }
    },
  },

  reducers: {
    saveGroups(state: any, action: any) {
      return { ...state, groups: action.payload };
    },
    saveSelectors(state: any, action: any) {
      return { ...state, selectors: action.payload };
    },
    saveSelector(state: any, action: any) {
      return {
        ...state,
        selectors: state.selectors.map((selector: ClashProxyGroupSelector) =>
          selector.name === action.payload.name ? action.payload : selector,
        ),
      };
    },
  },
};

export default Model;
