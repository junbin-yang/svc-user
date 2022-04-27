<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <div class="login_panle_form_title">
          <img
            class="login_panle_form_title_logo"
            :src="$GIN_VUE_ADMIN.appLogo"
            alt
          >
          <p class="login_panle_form_title_p">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <!-- 登录表单 -->
        <el-form
          v-if="pageType==='login'"
          ref="loginForm"
          :model="loginFormData"
          :rules="rules"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="account">
            <el-input
              v-model="loginFormData.account"
              placeholder="请输入账号"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <user />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginFormData.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <component
                      :is="lock"
                      @click="changeLock"
                    />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="validateCode">
            <div class="vPicBox">
              <el-input
                v-model="loginFormData.validateCode"
                placeholder="请输入验证码"
                style="width: 60%"
              />
              <div class="vPic">
                <img
                  v-if="picPath"
                  :src="picPath"
                  alt="请输入验证码"
                  @click="loginVerify()"
                >
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              size="large"
              style="width: 100%"
              @click="submitForm"
            >登 录</el-button>
          </el-form-item>
          <el-form-item>
            <el-button size="large" style="width: 46%" plain @click="changePageType('registerUser')">注册用户</el-button>
            <el-button size="large" style="width: 46%;margin-left: 8%" plain @click="changePageType('retrievePass')">忘记密码</el-button>
          </el-form-item>
        </el-form>

        <!-- 注册表单 -->
        <el-form
          v-if="pageType==='registerUser'"
          ref="registerForm"
          :model="registerFormData"
          :rules="registerRules"
          label-width="80px"
        >
          <el-form-item label="邀请码" prop="invitationCode">
            <el-input
              v-model="registerFormData.invitationCode"
              placeholder="请输入邀请码"
              clearable
            />
          </el-form-item>
          <el-form-item label="账号" prop="account">
            <el-input
              v-model="registerFormData.account"
              placeholder="请输入账号"
              clearable
            />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input
              v-model="registerFormData.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              clearable
            />
          </el-form-item>
          <el-form-item label="确认密码" prop="password_ck">
            <el-input
              v-model="registerFormData.password_ck"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请再次输入密码"
              clearable
            />
          </el-form-item>
          <el-form-item label="用户昵称" prop="nickName">
            <el-input
              v-model="registerFormData.nickName"
              placeholder="请输入昵称"
              clearable
            />
          </el-form-item>
          <el-form-item label="手机号" prop="mobile">
            <el-input
              v-model="registerFormData.mobile"
              placeholder="请输入手机号"
              clearable
            />
          </el-form-item>
          <el-form-item label="邮箱地址" prop="email">
            <el-input
              v-model="registerFormData.email"
              placeholder="请输入邮箱地址"
              clearable
            />
          </el-form-item>
          <el-form-item label="验证码" prop="validateCode">
            <el-input
              v-model="registerFormData.validateCode"
              placeholder="请输入邮箱验证码"
              style="width: 46%"
              clearable
            />
            <el-button size="large" style="width: 46%;margin-left: 8%" type="success" :class="{gray:wait_timer>0}" :disabled="validateCodeButtonStatus" @click="getCode()">
              <span
                v-show="showNum"
                :class="{num:wait_timer>0}"
              >{{ wait_timer + "s" }}</span>
              <span
                :class="{gray_span:wait_timer>0}"
              >{{ getCodeText() }}</span>
            </el-button>
          </el-form-item>
          <el-form-item>
            <el-button size="large" style="width: 46%" type="primary" plain @click="submitRegisterForm">注册</el-button>
            <el-button size="large" style="width: 46%;margin-left: 8%" type="primary" plain @click="changePageType('login')">返回</el-button>
          </el-form-item>
        </el-form>

        <!-- 找回密码表单 -->
        <el-form
          v-if="pageType==='retrievePass'"
          ref="retrievePasswordForm"
          :model="retrievePasswordFormData"
          :rules="retrievePasswordRules"
          label-width="80px"
        >
          <el-form-item label="邮箱地址" prop="email">
            <el-input
              v-model="retrievePasswordFormData.email"
              placeholder="请输入邮箱地址"
              clearable
            />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input
              v-model="retrievePasswordFormData.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入新密码"
              clearable
            />
          </el-form-item>
          <el-form-item label="确认密码" prop="password_ck">
            <el-input
              v-model="retrievePasswordFormData.password_ck"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请再次输新入密码"
              clearable
            />
          </el-form-item>

          <el-form-item label="验证码" prop="validateCode">
            <el-input
              v-model="retrievePasswordFormData.validateCode"
              placeholder="请输入邮箱验证码"
              style="width: 46%"
              clearable
            />
            <el-button size="large" style="width: 46%;margin-left: 8%" type="success" :class="{gray:wait_timer>0}" :disabled="validateCodeButtonStatus" @click="getCode()">
              <span
                v-show="showNum"
                :class="{num:wait_timer>0}"
              >{{ wait_timer + "s" }}</span>
              <span
                :class="{gray_span:wait_timer>0}"
              >{{ getCodeText() }}</span>
            </el-button>
          </el-form-item>
          <el-form-item>
            <el-button size="large" style="width: 46%" type="primary" plain @click="submitRetrievePasswordForm">确定</el-button>
            <el-button size="large" style="width: 46%;margin-left: 8%" type="primary" plain @click="changePageType('login')">返回</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="login_panle_right" />
      <div class="login_panle_foot">
        <div class="copyright">
          <BottomInfo />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
}
</script>

<script setup>
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import { captcha, registerUser, getEmailCaptcha, retrievePassword } from '@/api/user'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
// import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import md5 from 'js-md5'

const pageType = ref('login')
const changePageType = (v) => {
  pageType.value = v
}

// const router = useRouter()
// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 4) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

// 获取验证码
const loginVerify = () => {
  captcha({}).then((ele) => {
    picPath.value = ele.data.b64s
    loginFormData.validateCodeId = ele.data.key
  })
}
loginVerify()

// 登录相关操作
const lock = ref('lock')
const changeLock = () => {
  lock.value = lock.value === 'lock' ? 'unlock' : 'lock'
}

const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  account: 'root',
  password: '',
  validateCode: '',
  validateCodeId: '',
})
const rules = reactive({
  account: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  validateCode: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()
const login = async() => {
  loginFormData.password = md5(loginFormData.password)
  return await userStore.LoginIn({ loginType: 'account', loginValue: loginFormData.account, password: loginFormData.password, validateCodeId: loginFormData.validateCodeId, validateCode: loginFormData.validateCode })
}
const submitForm = () => {
  loginForm.value.validate(async(v) => {
    if (v) {
      const flag = await login()
      if (!flag) {
        loginVerify()
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}

const showNum = ref(false)
const wait_timer = ref(false)
const validateCodeButtonStatus = ref(false)
const registerForm = ref(null)
const registerFormData = reactive({
  invitationCode: '',
  account: '',
  password: '',
  password_ck: '',
  nickName: '',
  mobile: '',
  email: '',
  validateCode: '',
  validateCodeId: '0',
})
const checkPasswordNext = (rule, value, callback) => {
  var checkpass = ''
  if (pageType.value === 'registerUser') {
    checkpass = registerFormData.password
  } else if (pageType.value === 'retrievePass') {
    checkpass = retrievePasswordFormData.password
  }
  if (value !== checkpass) {
    return callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}
const checkRegisterPassword = (rule, value, callback) => {
  const passwordreg = /(?=.*\d)(?=.*[a-zA-Z])(?=.*[^a-zA-Z0-9]).{6,30}/
  const isValid = passwordreg.test(value)
  if (isValid !== true) {
    return callback(new Error('必须是大写字母,小写字母,数字,特殊字符组成,且长度为6到30位'))
  } else {
    callback()
  }
}
const checkEmail = (rule, value, callback) => {
  const mailReg = /^([a-zA-Z0-9_-.])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
  if (!value) {
    return callback(new Error('邮箱不能为空'))
  }
  setTimeout(() => {
    if (mailReg.test(value)) {
      callback()
    } else {
      callback(new Error('请输入正确的邮箱格式'))
    }
  }, 100)
}
const checkPhone = (rule, value, callback) => {
  const mobileReg = /^(1[3456789]\d{9})$/
  if (!value) {
    return callback(new Error('手机号不能为空'))
  }
  setTimeout(() => {
    if (mobileReg.test(value)) {
      callback()
    } else {
      callback(new Error('请输入正确的手机号格式'))
    }
  }, 100)
}
const registerRules = reactive({
  invitationCode: [{ required: true, message: '请输入邀请码', trigger: 'blur' }],
  account: [{ required: true, validator: checkUsername, trigger: 'blur' }],
  password: [{ required: true, validator: checkRegisterPassword, trigger: 'blur' }],
  password_ck: [{ required: true, validator: checkPasswordNext, trigger: 'blur' }],
  nickName: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
  mobile: [{ required: true, validator: checkPhone, trigger: 'blur' }],
  email: [{ required: true, validator: checkEmail, trigger: 'blur' }],
  validateCode: [{ required: true, message: '请输入邮箱验证码', trigger: 'blur' }],
})

const getCode = async() => {
  if (wait_timer.value > 0) {
    return false
  }
  const mailReg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
  var ckthis = false
  if (pageType.value === 'registerUser') {
    ckthis = mailReg.test(registerFormData.email)
  } else if (pageType.value === 'retrievePass') {
    ckthis = mailReg.test(retrievePasswordFormData.email)
  }
  if (!ckthis) {
    ElMessage({
      type: 'error',
      message: '邮箱格式错误',
      showClose: true
    })
    return false
  }
  validateCodeButtonStatus.value = true
  showNum.value = true
  wait_timer.value = 30

  var timer_interval = setInterval(function() {
    if (wait_timer.value > 0) {
      wait_timer.value--
    } else {
      clearInterval(timer_interval)
      validateCodeButtonStatus.value = false
    }
  }, 1000)

  // 获取验证码
  if (pageType.value === 'registerUser') {
    const res = await getEmailCaptcha({ email: registerFormData.email, invitationCode: registerFormData.invitationCode })
    if (res.code === 200) {
      registerFormData.validateCodeId = res.data.key
      ElMessage({
        type: 'success',
        message: '验证码发送成功'
      })
    }
  } else if (pageType.value === 'retrievePass') {
    const res = await getEmailCaptcha({ email: retrievePasswordFormData.email })
    if (res.code === 200) {
      retrievePasswordFormData.validateCodeId = res.data.key
      ElMessage({
        type: 'success',
        message: '验证码发送成功'
      })
    }
  }
}

const getCodeText = () => {
  if (wait_timer.value > 0) {
    return '已发送'
  }
  if (wait_timer.value === 0) {
    showNum.value = false
    return '重新发送验证码'
  }
  if (wait_timer.value === false) {
    return '发送验证码'
  }
}

const submitRegisterForm = () => {
  registerForm.value.validate(async(v) => {
    if (v) {
      const res = await registerUser(registerFormData)
      if (res.code === 200) {
        ElMessage({
          type: 'success',
          message: res.message
        })
        clearFormData()
        changePageType('login')
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      return false
    }
  })
}

const retrievePasswordForm = ref(null)
const retrievePasswordFormData = reactive({
  password: '',
  password_ck: '',
  email: '',
  validateCode: '',
  validateCodeId: '0',
})
const retrievePasswordRules = reactive({
  password: [{ required: true, validator: checkRegisterPassword, trigger: 'blur' }],
  password_ck: [{ required: true, validator: checkPasswordNext, trigger: 'blur' }],
  email: [{ required: true, validator: checkEmail, trigger: 'blur' }],
  validateCode: [{ required: true, message: '请输入邮箱验证码', trigger: 'blur' }],
})

const submitRetrievePasswordForm = () => {
  retrievePasswordForm.value.validate(async(v) => {
    if (v) {
      const res = await retrievePassword(retrievePasswordFormData)
      if (res.code === 200) {
        ElMessage({
          type: 'success',
          message: res.message
        })
        clearFormData()
        changePageType('login')
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写用户信息',
        showClose: true,
      })
      return false
    }
  })
}

const clearFormData = () => {
  if (pageType.value === 'registerUser') {
    registerFormData.invitationCode = ''
    registerFormData.account = ''
    registerFormData.password = ''
    registerFormData.password_ck = ''
    registerFormData.nickName = ''
    registerFormData.mobile = ''
    registerFormData.email = ''
    registerFormData.validateCode = ''
    registerFormData.validateCodeId = '0'
  } else if (pageType.value === 'retrievePass') {
    retrievePasswordFormData.password = ''
    retrievePasswordFormData.password_ck = ''
    retrievePasswordFormData.email = ''
    retrievePasswordFormData.validateCode = ''
    retrievePasswordFormData.validateCodeId = '0'
  }
}

</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";
</style>
