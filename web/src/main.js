import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import './style/element_visiable.scss'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
// 引入前端初始化相关内容
import './core/admin'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/admin.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
const app = createApp(App)
app.config.productionTip = false

app
  .use(run)
  .use(store)
  .use(auth)
  .use(router)
  .use(ElementPlus, { locale: zhCn })
  .mount('#app')

export default app
