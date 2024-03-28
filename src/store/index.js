import { createStore } from 'vuex';
import { getIndexInfo } from "~/api/user"
const store = createStore({
  state() {
    return {
      //首页数据
      indexData: [],
    }
  },
  mutations: {
    setIndexInfo(state, indexData) {//添加index数据
      state.indexData = indexData
    },
  },
  actions: {
    getIndexInfo({ commit }) {
      return new Promise((resolve, reject) => {
        getIndexInfo().then(res => {
          commit("setIndexInfo", res.dataAboutIndexMenu)
          resolve(res)
        }).catch(err => {
          reject(err)
        })
      })
    },
  }
});

export default store
