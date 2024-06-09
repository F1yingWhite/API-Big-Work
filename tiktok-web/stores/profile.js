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
        const page = 1; // 页码
        const pageSize = 10; // 每页条数

        const response = await $axios.get(`/api/movie/author`, {
          params: {
            page: page,
            pageSize: pageSize
          }
        });
        
        if (response && response.data && response.data.data) {
          const videos = response.data.data;

          this.$state.posts = videos.map(video => ({
            path: video.path,
            title: video.title,
            like: video.like
          }));

          this.$state.id = id;
          this.$state.name = videos[0].author;

          console.log(this.$state.posts);
          console.log(this.$state.allLikes);
          // console.log(`http://127.0.0.1:8888/api/movie/${this.$state.posts[0].path}`);

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
