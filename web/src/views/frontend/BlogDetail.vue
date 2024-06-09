<template>
  <div>
    <h2 class="title">
      {{ title }}
    </h2>
    <div class="describe">
      <div>发表于: {{ publishAt }}</div>
      <a-divider direction="vertical" />
      <div>作者: {{ author }}</div>
    </div>
    <div style="width: 700px" v-html="trustedHtml"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const title = '远程过程调用 RPC——高性能的 gRPC（实践案例）'
const publishAt = '2024年6月7日(五月初二) 星期五 19:51:36'
const author = '左扬'

// 使用字符串来存储 HTML
const text = ref(`  
  <h1 style="text-align: center;">远程过程调用 RPC——高性能的 gRPC（实践案例）</h1>  
  <p style="margin-left: 30px;">  
    <span style="font-family: 宋体, 'Songti SC'; font-size: 16px;">  
      gRPC，这一由Google推出的高性能、开源、通用RPC框架，凭借其众多引人注目的特性，已成为业界瞩目的焦点。  
      它基于HTTP/2协议标准设计开发，并采用Protocol Buffers作为默认的数据序列化协议，广泛支持多种编程语言。  
      gRPC不仅简化了服务的精确定义，而且还能为客户端和服务端自动生成可靠的代码库，从而极大地提升了开发效率。  
    </span>  
  </p>  
`)

const trustedHtml = text
</script>

<style lang="less" scoped>
.title {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  color: #555;
}

.describe {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  color: #999;
  font-size: 12px;
}
</style>
