import { defineStore } from 'pinia'
import axios from '../plugins/axios'

const $axios = axios().provide.axios

export const useProfileStore = defineStore('profile', {
  state: () => ({
    id: '',
    name: '',
    posts: [], // 用于存储视频列表的数组
    post: null,
    allLikes: 0,
  }),

  actions: {
    async getProfile(id) {
      this.resetUser()
      try {
        // TODO:后期调整下页数和每页数目
        const page = 1; // 页码
        const pageSize = 100; // 每页条数

        const response = await $axios.get(`/api/movie/author`, {
          params: {
            page: page,
            pageSize: pageSize
          }
        });
        
        if (response && response.data && response.data.data) {
          const videos = response.data.data;

          this.$state.posts = videos.map(video => ({
            ID: video.ID,
            id: id,
            path: video.path,
            title: video.title,
            author: video.author,
            like: video.like
          }));

          this.$state.id = id;
          this.$state.name = videos[0].author;

          this.allLikesCount();
        } else {
          console.error('Unexpected response structure:', response);
        }
      } catch (error) {
        console.error('Error fetching profile:', error);
      }
    },

    allLikesCount() {
      this.allLikes = this.$state.posts.reduce((sum, post) => sum + post.like, 0);
    },

    resetUser() {
      this.$state.id = ''
      this.$state.name = ''
      this.$state.posts = []
      this.$state.allLikes = 0
    }
  },
  persist: true,
})
