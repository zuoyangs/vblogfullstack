<template>
  <div>
    <!-- 顶部导航 -->
    <div class="header">
      <div class="logo">你们的胃叫胃，孤的胃叫胃PLUS的博客</div>
      <div class="right-header">
        <div>
          <!-- 登录后台进行博客管理 -->
          <a-button size="mini" type="text" @click="jumpToBackend">跳转到博客后台</a-button>
        </div>
      </div>
    </div>
    <!-- 显示博客列表 -->
    <div class="content">
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'

const router = useRouter()
const jumpToBackend = () => {
  const username = localStorage.getItem('username')
  const password = localStorage.getItem('password')
  if (
    (username === 'admin' && password === 'admin') ||
    (username === 'user' && password === 'user')
  ) {
    // 直接跳转后台管理页面
    router.push('/backend/blogs')
  } else {
    // 跳转去登录页面
    router.push('/login')
  }
}
</script>

<style scoped>
.header {
  display: flex;
  align-content: center;
  justify-content: flex-start;
  align-items: center;
  border-bottom: 1px solid rgb(229, 230, 235);
  height: 45px;
}

.logo {
  margin-left: 8px;
  font-size: 14px;
  font-weight: 500;
}

.right-header {
  margin-left: auto;
}

.content {
  margin: 20px;
  display: flex;
  align-content: center;
  justify-content: center;
  align-items: center;
}
</style>
