<template>
    <div class="w-[100%] relative flex justify-center">
        <div
            class="
                absolute
                top-6
                z-50
                mx-auto
                text-white
                bg-[#F02C56]
                bg-opacity-70
                px-14
                py-3
                rounded-sm
            "    

            :class="errorType ? 'visible' : 'invisible'"
        >
            {{ error }}
        </div>
    </div>
</template>

<script setup>
const props = defineProps(['errorType'])
const { errorType } = toRefs(props)

let error = ref('')

watch (() => errorType.value,() => {
    if (errorType.value == 'caption'){
        error.value = '最大长度为150个字符'
    } else if (errorType.value == 'bio'){
        error.value = '最大长度为80个字符'
    } else if (errorType.value == 'file'){
        error.value = '不支持的文件类型。请使用.MP4文件!'
    } else {
        error.value = ''
    }
})
</script>
