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

    async getPostById(ID) {
      try {
        const response = await $axios.get(`/api/movie`, {
          params: {
            id: ID
          }
        });;

        if (response && response.data && response.data.data) {
          this.$state.selectedPost = response.data.data;
          console.log(this.$state.selectedPost)
        } else {
          console.error('Unexpected response structure:', response);
        }
      } catch (error) {
        console.error('Error fetching post:', error);
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
    },

    // 获取下一个视频
    async fetchNextVideo (currentId) {
        try {
            const response = await $axios.post('/api/movie/down', {
                movieId: parseInt(currentId)
            });
            return response.data.data;
        } catch (error) {
            console.error('Error fetching next video:', error);
        }
    },
  
    // 获取上一个视频
    async fetchPreviousVideo (currentId) {
        try {
            const response = await $axios.post('/api/movie/up', {
                movieId: parseInt(currentId)
            });
            return response.data.data;
        } catch (error) {
            console.error('Error fetching previous video:', error);
        }
    },

    async deleteVideo (id) {
      await $axios.delete(`/api/movie/${id}`);
    }
  },

  persist: true,
})