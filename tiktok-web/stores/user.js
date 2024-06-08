import { defineStore } from 'pinia'
import axios from '../plugins/axios'
import { useGeneralStore } from './general'

const $axios = axios().provide.axios

export const useUserStore = defineStore('user', {
  state: () => ({
    id: '',
    name: '',
    bio: '',
    image: '',
    token: ''
  }),
  actions: {

    async getTokens() {
      // TODO:接口
      await $axios.get('/api/sanctum/csrf-cookie')
    },
    async login(id, password) {
      try {
        const response = await $axios.post('/api/user/login', {
          id: id,
          password: password
        })
    
        this.token = response.data.token
        this.id = id
    
        localStorage.setItem('token', this.token)
        $axios.defaults.headers.common['Authorization'] = `Bearer ${this.token}`
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
      
      // this.$state.token = res.data.Authorization
      this.$state.id = res.data.id
      this.$state.name = res.data.username
      // this.$state.bio = res.data.bio
      // this.$state.image = res.data.image
    },

    async updateUserImage(data) {
      // TODO:接口
      return await $axios.post('/api/user/update-user-image', data)
    },

    async updateUser(name, bio) {
      // TODO:接口
      return await $axios.put('/api/user', {
        name: name,
        // bio: bio
      })
    },

    async createPost(data) {
      // TODO:接口
      return await $axios.post('/api/posts', data)
    },

    async deletePost(post) {
      // TODO:接口
      return await $axios.delete(`/api/posts/${post.id}`)
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

    async likePost(post, isPostPage) {
      // TODO:接口
      let res = await $axios.post('/api/likes', {
        post_id: post.id,
      })

      console.log(res)

      let singlePost = null

      if (isPostPage) {
        singlePost = post
      } else {
        singlePost = useGeneralStore().posts.find(p => p.id === post.id)
      }
      console.log(singlePost)
      singlePost.likes.push(res.data.like)
    },

    async unlikePost(post, isPostPage) {
      let deleteLike = null
      let singlePost = null

      if (isPostPage) {
        singlePost = post
      } else {
        singlePost = useGeneralStore().posts.find(p => p.id === post.id)
      }

      singlePost.likes.forEach(like => {
        if (like.user_id === this.id) { deleteLike = like }
      });
      
      // TODO: 接口
      let res = await $axios.delete('/api/likes/' + deleteLike.id)

      for (let i = 0; i < singlePost.likes.length; i++) {
        const like = singlePost.likes[i];
        if (like.id === res.data.like.id) { singlePost.likes.splice(i, 1); }
      }
    },

    async logout() {
      await $axios.post('/api/logout')
      this.resetUser()
      localStorage.removeItem('token')
      delete $axios.defaults.headers.common['Authorization']
    },

    resetUser() {
      this.id = ''
      this.name = ''
      this.email = ''
      this.bio = ''
      this.image = ''
      this.token = ''
    }
  },
  persist: true,
})
