import { updateAll } from './service';

const Model = {
  namespace: 'subs',

  state: {},

  effects: {
    *updateAll({ _, callback }: any, { call }: any) {
      const response = yield call(updateAll);
      if (callback) {
        if (response.msg) {
          callback(false, response.msg);
        } else {
          callback(true);
        }
      }
    },
  },
};

export default Model;
