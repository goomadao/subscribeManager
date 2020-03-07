import { fetchRules, addRule, editRule, updateRule, deleteRule, updateAllRules } from '../service';

const Model = {
  namespace: 'rules',

  state: [],

  effects: {
    *fetchRules(_: any, { call, put }: any) {
      const response = yield call(fetchRules);
      yield put({
        type: 'saveRules',
        payload: response.data,
      });
    },
    *addRule({ payload, callback }: any, { call, put }: any) {
      const response = yield call(addRule, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveRules',
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
    *editRule({ payload, callback }: any, { call, put }: any) {
      const response = yield call(editRule, payload);
      console.log('edit');
      console.log(response);
      if (response.status === 'success') {
        yield put({
          type: 'saveRules',
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
    *updateRule({ payload, callback }: any, { call, put }: any) {
      const response = yield call(updateRule, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveRule',
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
    *deleteRule({ payload, callback }: any, { call, put }: any) {
      const response = yield call(deleteRule, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveRules',
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
    *updateAllRules({ _, callback }: any, { call, put }: any) {
      const response = yield call(updateAllRules);
      // if (response.status === 'success') {
      yield put({
        type: 'saveRules',
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
    saveRules(state: any, action: any) {
      return action.payload;
    },
    saveRule(state: any, action: any) {
      let result = state;
      for (let i = 0; i < result.length; i++) {
        if (result[i].name === action.payload.name) {
          result[i] = action.payload;
        }
      }
      return result;
    },
  },
};

export default Model;
