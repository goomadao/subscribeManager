import {
  fetchGroups,
  addGroup,
  editGroup,
  updateGroup,
  deleteGroup,
  updateAllGroups,
  addNode,
  editNode,
  deleteNode,
} from '../service';

const Model = {
  namespace: 'groups',

  state: [],

  effects: {
    *fetchGroups(_: any, { call, put }: any) {
      const response = yield call(fetchGroups);
      yield put({
        type: 'saveGroups',
        payload: response.data,
      });
    },
    *addGroup({ payload, callback }: any, { call, put }: any) {
      const response = yield call(addGroup, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveGroups',
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
    *editGroup({ payload, callback }: any, { call, put }: any) {
      const response = yield call(editGroup, payload);
      console.log(response);
      if (response.status === 'success') {
        yield put({
          type: 'saveGroups',
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
    *updateGroup({ payload, callback }: any, { call, put }: any) {
      const response = yield call(updateGroup, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveGroup',
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
    *deleteGroup({ payload, callback }: any, { call, put }: any) {
      const response = yield call(deleteGroup, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveGroups',
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
    *updateAllGroups({ _, callback }: any, { call, put }: any) {
      const response = yield call(updateAllGroups);
      // if (response.status === 'success') {
      yield put({
        type: 'saveGroups',
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
    *addNode({ payload, callback }: any, { call, put }: any) {
      const response = yield call(addNode, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveGroup',
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
    *editNode({ payload, callback }: any, { call, put }: any) {
      const response = yield call(editNode, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveGroup',
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
    *deleteNode({ payload, callback }: any, { call, put }: any) {
      const response = yield call(deleteNode, payload);
      if (response.status === 'success') {
        yield put({
          type: 'saveGroup',
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
    saveGroups(state: any, action: any) {
      return action.payload;
    },
    saveGroup(state: any, action: any) {
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
