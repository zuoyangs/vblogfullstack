import { createRouter, createWebHistory } from 'vue-router'
import { beforeEachHandler, afterEachHandler } from './permession'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/frontend/blogs'
    },
    {
      path: '/frontend',
      name: 'frontend',
      component: () => import('@/layout/FrontendLayout.vue'),
      children: [
        {
          path: 'blogs',
          name: 'BlogView',
          component: () => import('@/views/frontend/BlogView.vue')
        },
        {
          path: 'blogs/:id',
          name: 'BlogDetail',
          component: () => import('@/views/frontend/BlogDetail.vue')
        }
      ]
    },
    {
      path: '/login',
      name: 'LoginPage',
      component: () => import('@/views/login/LoginPage.vue')
    },
    {
      path: '/backend',
      name: 'backend',
      component: () => import('@/layout/BackendLayout.vue'),
      children: [
        {
          path: 'blogs',
          name: 'BlogList',
          component: () => import('@/views/backend/BlogList.vue')
        },
        {
          path: 'tags',
          name: 'TagList',
          component: () => import('@/views/backend/TagList.vue')
        }
      ]
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'notFound',
      component: () => import('@/views/errors/NotFound.vue')
    },
    {
      path: '/errors/403',
      name: 'PermissionDeny',
      component: () => import('@/views/errors/PermissionDeny.vue')
    }
  ]
})

// 补充导航守卫
router.beforeEach(beforeEachHandler)
router.afterEach(afterEachHandler)

export default router
