import { defineStore } from 'pinia'
import axios from '../plugins/axios'
import { useGeneralStore } from './general'

const $axios = axios().provide.axios

export const useUserStore = defineStore('user', {
  state: () => ({
    id: '',
    name: '',
    token: ''
  }),
  getters: {
    isAuthenticated: (state) => !!state.token
  },
  actions: {
    // 登录接口
    async login(id, password) {
      try {
        const response = await $axios.post('/api/user/login', {
          id: id,
          password: password
        })
    
        this.token = response.data.data[0].token
        this.id = id
    
        localStorage.setItem('token', this.token)
        $axios.defaults.headers.common['Authorization'] = `${this.token}`
      } catch (error) {
        if (error.response && error.response.status === 400) {
          throw new Error(error.response.data.error)
        } else {
          console.error(error)
          throw new Error('登录失败，请检查账号和密码。')
        }
      }
    },
    
    async register(id, password, name) {
      try {
        await $axios.post('/api/user/register', {
          name: name,
          id: id,
          password: password
        })
        await this.login(id, password)
      } catch (error) {
        if (error.response && error.response.status === 400) {
          throw new Error(error.response.data.error)
        } else {
          console.error(error)
          throw new Error('注册失败，账号已存在。')
        }
      }
    },
    
    async getUser() {
      // TODO:接口
      let res = await $axios.get('/api/user')

      this.$state.id = res.data.id
      this.$state.name = res.data.username
    },

    async updateUserImage(data) {
      // TODO:接口
      return await $axios.post('/api/user/update-user-image', data)
    },

    async updateUser(name) {
      // TODO:接口
      return await $axios.put('/api/user/username', {
        name: name,
      })
    },

    async updateUserPassword(password) {
      // TODO:接口
      return await $axios.put('/api/user/password', {
        password: password,
      })
    },

    // 上传视频,加上参数(video:file,name:string)
    async createPost(data) {
      // TODO:接口
      return await $axios.post('/api/movie', data)
    },

    async deletePost(post) {
      // TODO:接口
      return await $axios.delete(`/api/movie/${post.id}`)
    },

    async addComment(post, comment) {
      // TODO:接口
      let res = await $axios.post('/api/comments', {
        post_id: post.id,
        comment: comment
      })

      if (res.status === 200) {
        await this.updateComments(post)
      }
    },

    async deleteComment(post, commentId) {
      // TODO:接口
      let res = await $axios.delete(`/api/comments/${commentId}`, {
        post_id: post.id
      })

      if (res.status === 200) {
        await this.updateComments(post)
      }
    },

    async updateComments(post) {
      // TODO:接口
      let res = await $axios.get(`/api/profiles/${post.user.id}`)

      for (let i = 0; i < res.data.posts.length; i++) {
          const updatePost = res.data.posts[i];

          if (post.id == updatePost.id) {
              useGeneralStore().selectedPost.comments = updatePost.comments
          }
      }
    },

    // 确认用户是否对这个视频点赞
    async likeOrNot(post) {
      try {
        let res = await $axios.get(`api/movie/like/${post.ID}`);
        // return res.data.data.like === "true";
        console.log("1:" + res.data.data.like)
        return res.data.data.like === "true";
      } catch (error) {
        console.error(error);
      }
    },

    // 用户对视频点赞以及取消点赞
    async like(post) {
      try {
        let res = await $axios.post(`api/movie/like/${post.ID}`);
      } catch (error) {
        console.error(error);
      }
    },

    async logout() {
      this.resetUser()
      localStorage.removeItem('token')
      delete $axios.defaults.headers.common['Authorization']
    },

    resetUser() {
      this.id = ''
      this.name = ''
      this.token = ''
    }
  },
  persist: true,
})
