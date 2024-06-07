<template>
    <div class="text-center text-[28px] mb-4 font-bold">登录</div>
    <div class="px-6 pb-1.5 text-[15px]">邮箱</div>
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

    <div class="px-6 pb-2 mt-6">
        <button 
            :disabled="(!id || !password)"
            :class="(!id || !password) ? 'bg-gray-200' : 'bg-[#F02C56]'"
            @click="$event => login()" 
            class="w-full text-[17px] font-semibold text-white py-3 rounded-sm"
        >
            登录
        </button>
    </div>
</template>

<script setup>
    const { $userStore, $generalStore } = useNuxtApp()
    // let email = ref(null)
    let id = ref(null)
    let password = ref(null)
    let errors = ref(null)

    const login = async () => {
        errors.value = null

        try {
            // await $userStore.getTokens()
            await $userStore.login(id.value, password.value)
            // await $userStore.getUser()
            // await $generalStore.getRandomUsers('suggested')
            // await $generalStore.getRandomUsers('following')
            $generalStore.isLoginOpen = false
        } catch (error) {
            // TODO:在后端加上错误码
            // errors.value = error.response.data.errors
        }
    }
</script>