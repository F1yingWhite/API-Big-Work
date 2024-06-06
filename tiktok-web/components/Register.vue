<template>
    <div class="text-center text-[28px] mb-4 font-bold">注册</div>

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
            placeholder="输入邮箱"
            v-model:input="email"
            inputType="email"
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

    <div class="px-6 pb-2">
        <TextInput 
            placeholder="确认密码"
            v-model:input="confirmPassword"
            inputType="password"
        />
    </div>

    <div class="px-6 text-[12px] text-gray-600">忘记密码?</div>

    <div class="px-6 pb-2 mt-6">
        <button 
            :disabled="(!name || !email || !password || !confirmPassword)"
            :class="(!name || !email || !password || !confirmPassword) ? 'bg-gray-200' : 'bg-[#F02C56]'"
            @click="register()" 
            class="w-full text-[17px] font-semibold text-white bg-[#F02C56] py-3 rounded-sm"
        >
            注册
        </button>
    </div>
</template>

<script setup>
let name = ref(null)
let email = ref(null)
let password = ref(null)
let errors = ref(null)
let confirmPassword = ref(null)


</script>