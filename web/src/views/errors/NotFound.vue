<template>
  <div class="content">
    <a-result class="result" status="404" :subtitle="'页面未找到'">
      <template #extra>
        <a-button key="back" type="primary" @click="back">返回主页</a-button>
      </template>
    </a-result>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
const back = () => {
  router.push({ name: 'home' })
}
</script>

<style scoped lang="less">
.content {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh; // 使用视口高度单位来确保内容始终垂直居中
  text-align: center;
}

.result {
  margin-bottom: 20px; // 为按钮添加一些间距，使其看起来更整洁
}
</style>
