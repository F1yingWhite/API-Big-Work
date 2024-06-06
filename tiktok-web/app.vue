<template>
  <NuxtPage />

  <AuthOverlay v-if="isLoginOpen"/>
  <!-- <EditProfileOverlay v-if="isEditProfile"/> -->
  <EditProfileOverlay />
</template>

<script setup>
import { storeToRefs } from 'pinia';
const { $generalStore } = useNuxtApp()
const { isLoginOpen,isEditProfile } = storeToRefs($generalStore)
</script>
