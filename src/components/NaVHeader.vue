<template>
    <div class="bg-blue-300 fixed flex justify-between items-center top-0 left-0 right-0"
        style="height: 50px; width: 100%">
        <el-button type="primary" size="large" @click="toIndexPage" class="mx-5 rounded-lg">首页</el-button>
        <el-button type="primary" size="large" @click="toPicturePage" class="mx-5 rounded-lg">图片</el-button>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter()
function toIndexPage() {
    router.push("/")
}
</script>