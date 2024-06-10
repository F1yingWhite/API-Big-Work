import { defineStore } from 'pinia'
import { useUserStore } from './user'
import axios from '../plugins/axios'

const $axios = axios().provide.axios

export const useGeneralStore = defineStore('general', {
  state: () => ({
    isLoginOpen: false,
    isEditProfileOpen: false,
    selectedPost: null,
    ids: null,
    isBackUrl: "/",
    posts: [],
    suggested: null,
    following: null,
  }),
  actions: {
    bodySwitch(val) {
      if (val) {
        document.body.style.overflow = 'hidden'
        return
      }
      document.body.style.overflow = 'visible'
    },

    allLowerCaseNoCaps(str) {
      return str.split(' ').join('').toLowerCase()
    },

    setBackUrl(url) {
        this.isBackUrl = url
    },

    async hasSessionExpired() {
      await $axios.interceptors.response.use((response) => {
        return response;
      }, (error) => {
          // TODO:确认错误码
          switch (error.response.status) {
              case 401:
              case 419:
              case 503:
                  useUserStore().resetUser()
                  window.location.href = '/';
                  break;
              case 500:
                  alert('Oops, something went wrong!  The team has been notified.');
                  break;
              default:
                  return Promise.reject(error);
          }
      });
    },

    async getPostById(id) {
      // TODO:这里的接口要改
      let res = await $axios.get(`/api/posts/${id}`)

      this.$state.selectedPost = res.data.post[0]
      this.$state.ids = res.data.ids
    },

    async getRandomUsers(type) {
      // TODO:这里的接口要改
      let res = await $axios.get(`/api/get-random-users`)

      if (type === 'suggested') {
        this.suggested = res.data.suggested
      }

      if (type === 'following') {
        this.following = res.data.following
      }
    },

    updateSideMenuImage(array, user) {
      for (let i = 0; i < array.length; i++) {
        const res = array[i];
        if (res.id == user.id) {
            res.image = user.image
        }
      }
    },

    // 获得所有的视频列表
    async getAllUsersAndPosts() {
      // TODO:这里的接口要改
      let res = await $axios.get('/api/movie/list')
      this.$state.posts = res.data.data

      console.log("Posts1:" + posts)
    }
  },
  persist: true,
})