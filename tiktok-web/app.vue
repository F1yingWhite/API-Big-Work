<template>
  <NuxtPage />

  <AuthOverlay v-if="isLoginOpen" />
  <EditProfileOverlay v-if="isEditProfileOpen" />
</template>

<script setup>
import { storeToRefs } from 'pinia';
const { $userStore, $generalStore } = useNuxtApp()
const { isLoginOpen, isEditProfileOpen } = storeToRefs($generalStore)
import axios from '../plugins/axios'

const $axios = axios().provide.axios

onMounted(async () => {
  $generalStore.bodySwitch(false)
  isLoginOpen.value = false
  isEditProfileOpen.value = false

  // 从localStorage中恢复token并设置axios的Authorization头
  const token = localStorage.getItem('token')
  if (token) {
      $userStore.token = token
      $axios.defaults.headers.common['Authorization'] = `${token}`
      await $userStore.getUser()
  }

  try {
      await $generalStore.hasSessionExpired()
  } catch (error) {
      console.log(error)
  }
})

watch(() => isLoginOpen.value, (val) => $generalStore.bodySwitch(val) )
watch(() => isEditProfileOpen.value, (val) => $generalStore.bodySwitch(val) )
</script>