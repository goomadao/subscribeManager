import { fetchChangers, addChanger, editChanger, deleteChanger } from '../service';

const Model = {
  namespace: 'changers',

  state: [],

  effects: {
    *fetchChangers(_: any, { call, put }: any) {
      const response = yield call(fetchChangers);
      yield put({
        type: 'saveChangers',
        payload: response.data,
      });
    },
    *addChanger({ payload, callback }: any, { call, put }: any) {
      const response = yield call(addChanger, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveChangers',
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
    *editChanger({ payload, callback }: any, { call, put }: any) {
      const response = yield call(editChanger, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveChangers',
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
    *deleteChanger({ payload, callback }: any, { call, put }: any) {
      const response = yield call(deleteChanger, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveChangers',
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
  },

  reducers: {
    saveChangers(state: any, action: any) {
      return action.payload;
    },
  },
};

export default Model;
