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

    // 未登录，获得所有的视频列表
    async getNoauthPosts() {
      // TODO:这里的接口要改
      try {
        // TODO:后期调整下页数和每页数目
        const page = 1; // 页码
        const pageSize = 5; // 每页条数
        const response = await $axios.get(`/api/movie/list`, {
          params: {
            page: page,
            pageSize: pageSize
          }
        });
        
        if (response && response.data && response.data.data) {
          const videos = response.data.data;

          this.$state.posts = videos.map(video => ({
            ID: video.ID,
            path: video.path,
            title: video.title,
            like: video.like,
            author: video.author
          }));
        } else {
          console.error('Unexpected response structure:', response);
        }
      } catch (error) {
        console.error('Error fetching profile:', error);
      }
    },

    // 登录，给用户推荐视频
    async getAuthPosts() {
      try {
        const response = await $axios.get(`/api/movie/recommend`);
        
        if (response && response.data && response.data.data) {
          const videos = response.data.data;

          this.$state.posts = videos.map(video => ({
            ID: video.ID,
            path: video.path,
            title: video.title,
            like: video.like,
            author: video.author
          }));
        } else {
          console.error('Unexpected response structure:', response);
        }
      } catch (error) {
        console.error('Error fetching profile:', error);
      }
    }
  },
  persist: true,
})