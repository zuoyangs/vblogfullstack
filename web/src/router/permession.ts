// 定义导航守卫
export async function beforeEachHandler(to: any, from: any, next: any) {
  if (to.fullPath.indexOf('/backend') === 0) {
    // 如果未登陆 重定向到登录页面, 并且把目标页面作为重定向参数传递下去
    const username = localStorage.getItem('username')
    const password = localStorage.getItem('password')
    if (username === null || password === null || username === '' || password === '') {
      console.log('not login')
      next({
        path: '/login',
        query: {
          redirect: to.name,
          ...to.query
        }
      })
    } else {
      // 已经登录的用户直接放行
      next()
    }
  } else {
    // 不属于/backend的页面 直接放开
    next()
  }
}

export function afterEachHandler(to: any, from: any) {
  console.log(to, from)
}
