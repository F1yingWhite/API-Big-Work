<template>
    <div
        id="EditProfileOverlay" 
        class="fixed flex justify-center pt-14 md:pt-[105px] z-50 top-0 left-0 w-full h-full bg-black bg-opacity-50 overflow-auto"
    >
        <div 
            class="relative bg-white w-full max-w-[700px] sm:h-[580px] h-[655px] mx-3 p-4 rounded-lg mb-10"
            :class="!uploadedImage ? 'h-[655px]' : 'h-[580px]'"
        >
            
            <div class="absolute flex items-center justify-between w-full p-5 left-0 top-0 border-b border-b-gray-300">
                <div class="text-[22px] font-medium">
                    编辑简介
                </div>
                <button @click="$generalStore.isEditProfileOpen = false">
                    <Icon name="mdi:close" size="25"/>
                </button>
            </div>

        </div>
    </div>
</template>

<script setup>
const { $generalStore } = useNuxtApp()
let uploadedImage = ref(null)
</script>
