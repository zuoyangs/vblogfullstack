<template>
  <div class="Login-Form">
    <a-form
      :model="form"
      :style="{ width: '400px', height: '100%', justifyContent: 'center' }"
      @submit="handleSubmit"
    >
      <!--标题 -->
      <a-form-item>
        <div class="title">登录博客管理后台</div>
      </a-form-item>

      <!-- 用户名输入框 -->
      <a-form-item
        field="username"
        validate-trigger="input"
        hide-asterisk
        :rules="[{ required: true, message: '请输入用户名' }]"
      >
        <a-input v-model="form.username" placeholder="请输入邮箱/手机号/用户名" />
      </a-form-item>

      <!--密码输入框 -->
      <a-form-item
        field="password"
        validate-trigger="input"
        hide-asterisk
        :rules="[{ required: true, message: '请输入登录密码' }]"
      >
        <a-input v-model="form.password" type="password" placeholder="请输入登录密码" />
      </a-form-item>

      <!-- 是否已读 -->
      <a-form-item field="isRead">
        <a-checkbox v-model="form.isRead">已阅读并同意在线服务协议和隐私政策</a-checkbox>
      </a-form-item>

      <!-- 登录按钮 -->
      <a-form-item>
        <a-button type="primary" style="width: 100%" html-type="submit"> 登录 </a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useRouter } from 'vue-router'

// 定义表单数据的接口
interface LoginForm {
  username: string
  password: string
  isRead: boolean
}

interface ValidatedError {}

// 使用定义的接口来约束form数据
const form = reactive<LoginForm>({
  username: '',
  password: '',
  isRead: false
})

const router = useRouter()

// 定义提交表单的处理函数，并添加类型注解
const handleSubmit = (
  data: { values: Record<string, any>; errors?: Record<string, ValidatedError> },
  event: Event
) => {
  event.preventDefault()

  // 使用类型断言来确保 values 是一个 LoginForm 类型的对象
  const formValues = data.values as LoginForm

  if (!data.errors) {
    const { username, password } = formValues

    if (validateCredentials(username, password)) {
      localStorage.setItem('username', username)
      localStorage.setItem('password', password)

      // 登录成功后的页面跳转逻辑
      const { redirect, ...othersQuery } = router.currentRoute.value.query as Record<
        string,
        string | undefined
      >
      router.push({
        name: redirect || 'BlogList',
        query: {
          ...othersQuery
        }
      })

      Message.success('登录成功')
    } else {
      Message.error('用户名或密码错误')
    }
  }
}
function validateCredentials(username: string, password: string): boolean {
  return (
    (username === 'admin' && password === 'admin') || (username === 'user' && password === 'user')
  )
}
</script>

<style scoped>
/* 登录表单样式 */
.Login-Form {
  display: flex;
  align-items: center;
  justify-content: center;
  align-content: center;
  height: 100%;
}

.title {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  font-size: 30px;
  font-weight: bold;
  color: #000;
  margin-bottom: 20px;
  width: 100%;
}
</style>
