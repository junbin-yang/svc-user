/**
 * 网站配置文件
 */

const config = {
  appName: 'Svc-User-Admin',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用SvcUser`
      )
    )
    console.log('\n')
  }
}

export default config
