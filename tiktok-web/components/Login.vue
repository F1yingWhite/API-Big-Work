<template>
    <div class="text-center text-[28px] mb-4 font-bold">登录</div>
    <div class="px-6 pb-1.5 text-[15px]">ID</div>
    <div class="px-6 pb-2 text-[15px]">
      <TextInput
        placeholder="输入ID"
        v-model:input="id"
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
        :disabled="(!id || !password)"
        :class="(!id || !password) ? 'bg-gray-200' : 'bg-[#F02C56]'"
        @click="login"
        class="w-full text-[17px] font-semibold text-white py-3 rounded-sm"
      >
        登录
      </button>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useNuxtApp } from '#app'
  
  const { $userStore, $generalStore } = useNuxtApp()
  let id = ref(null)
  let password = ref(null)
  let errors = ref(null)
  
  const login = async () => {
    errors.value = null
  
    try {
      await $userStore.login(id.value, password.value)
      $generalStore.isLoginOpen = false
    } catch (error) {
      errors.value = error.message
    }
  }
  </script>
  