<template>
  <div class="background">
    <router-view v-if="show!=false">
    </router-view>
  </div>
</template>

<script setup>
import { onMounted,ref } from 'vue';
import { ElLoading } from 'element-plus'
const show = ref(false)
onMounted(() => {
    const loading = ElLoading.service({
        lock: true,
        text: '正在努力加载中 Loading...',
        background: 'rgba(0, 0, 0, 0.7)',
    })
    setTimeout(async () => {
        // while (satisfyData.value === null) {
        //     console.log(satisfyData.value)
        //     await new Promise(resolve => setTimeout(resolve, 100)); // 等待100毫秒
        // }
        loading.close()
        show.value = true;
    }, 1000)
});
</script>

<style>
.background {
  /* 设置背景图片 */
  /* background-image: url('./static/background.jpg'); */
  /* 背景图片覆盖整个容器 */
  background-size: cover;
  /* 背景图片不重复 */
  background-repeat: no-repeat;
  /* 背景图片居中显示 */
  background-position: center;
  /* 设置容器的最小高度为视窗的高度，这样背景可以覆盖整个视窗 */
  min-height: 100vh;
  /* 在背景之上添加一层半透明的遮罩，实现透明效果 */
  position: relative;
  z-index: 1;
}

.background::before {
  /* 在背景之上创建一个伪元素作为遮罩层 */
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  /* 这里的rgba(255, 255, 255, 0.5)代表半透明的白色遮罩，最后一个参数0.5是透明度 */
  /* 调整这个值来改变透明感的强度，数值越小透明感越强 */
  background: rgba(255, 255, 255, 0.2);
  z-index: -1;
}

/* Webkit浏览器样式 */
::-webkit-scrollbar {
  width: 2px;
  height: 2px;
  /* 设置滚动条高度，以模拟底部滚动条 */
}

::-webkit-scrollbar-thumb {
  background-color: #888888;
  /* 滚动条拖动块的颜色 */
}

/* IE浏览器样式 */
::-ms-scrollbar {
  width: 2px;
  height: 2px;
  /* 设置滚动条高度，以模拟底部滚动条 */

}

::-ms-scrollbar-thumb {
  background-color: #888888;
  /* 滚动条拖动块的颜色 */
}
</style>

