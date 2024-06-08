<template>
    <div class="text-center text-[28px] mb-4 font-bold">注册</div>
  
    <div class="px-6 pb-2 text-[15px]">
      <TextInput
        placeholder="用户ID"
        v-model:input="id"
        inputType="text"
        :autoFocus="true"
        error=""
      />
    </div>
  
    <div class="px-6 pb-2 text-[15px]">
      <TextInput
        placeholder="用户名"
        v-model:input="name"
        inputType="text"
        :autoFocus="true"
        error=""
      />
    </div>
  
    <div class="px-6 pb-2">
      <TextInput 
        placeholder="密码"
        v-model:input="password"
        inputType="password"
      />
    </div>
  
    <div class="px-6 text-[12px] text-gray-600">忘记密码?</div>
  
    <div v-if="errors" class="px-6 text-red-500">{{ errors }}</div>
  
    <div class="px-6 pb-2 mt-6">
      <button 
        :disabled="(!id || !password || !name)"
        :class="(!id || !password || !name) ? 'bg-gray-200' : 'bg-[#F02C56]'"
        @click="register" 
        class="w-full text-[17px] font-semibold text-white py-3 rounded-sm"
      >
        注册
      </button>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useNuxtApp } from '#app'
  
  const { $userStore, $generalStore } = useNuxtApp()
  
  let name = ref(null)
  let password = ref(null)
  let id = ref(null)
  let errors = ref(null)
  
  const register = async () => {
    errors.value = null
  
    try {
      await $userStore.register(id.value, password.value, name.value)
      await $userStore.login(id.value, password.value)
      $generalStore.isLoginOpen = false
    } catch (error) {
      errors.value = error.message
    }
  }
  </script>
  