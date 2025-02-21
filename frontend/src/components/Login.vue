<template>
  <v-app style="background: linear-gradient(135deg, #f0f4ff 0%, #ffffff 100%); height: 100vh">
    <v-container 
      class="d-flex align-center justify-center" 
      style="height: 100%"
    >
      <v-card
        width="450"
        elevation="12" 
        class="pa-8 rounded-lg" 
        style="background: #ffffff;color: #000;"
      >
        <v-card-title
          class="text-h5"
          style="color: #002fa7;"
        >
          Pprof Cloud Platform
        </v-card-title>
        <!-- 登陆/注册切换 -->
        <v-tabs 
          v-model="tab" 
          center-active 
          background-color="transparent"
        >
          <v-tab
            class="text-h6"
            style="color: #2563eb"
          >
            登录
          </v-tab>
          <v-tab 
            class="text-h6" 
            style="color: #2563eb"
          >
            注册
          </v-tab>
        </v-tabs>
  
        <v-divider
          class="my-6"
          style="border-color: #90a4ae"
        />
  
        <!-- 登陆表单 -->
        <v-form
          v-if="tab === 0"
          @submit.prevent="handleLogin"
        >
          <v-text-field
            v-model="login.email"
            label="邮箱"
            outlined
            dense
            color="#2563eb"
            :rules="emailRules"
          />
  
          <v-text-field
            v-model="login.password"
            label="密码"
            outlined
            dense
            color="#2563eb"
            :rules="passwordRules"
            type="password"
          />
  
          <v-btn
            type="submit"
            block
            color="#2563eb"
            dark
            depressed
          >
            立即登录
          </v-btn>
        </v-form>
  
        <!-- 注册表单 -->
        <v-form
          v-if="tab === 1"
          @submit.prevent="handleRegister"
        >
          <v-text-field
            v-model="register.email"
            label="邮箱"
            outlined
            dense
            color="#2563eb"
            :rules="emailRules"
          />
  
          <v-text-field
            v-model="register.password"
            label="密码"
            outlined
            dense
            color="#2563eb"
            :rules="passwordRules"
            type="password"
          />
  
          <v-text-field
            v-model="register.confirmPassword"
            label="确认密码"
            outlined
            dense
            color="#2563eb"
            :rules="confirmPasswordRules"
            type="password"
          />
  
          <v-row>
            <v-col cols="8">
              <v-text-field
                v-model="register.verificationCode"
                label="验证码"
                outlined
                dense
                color="#2563eb"
                :rules="verificationCodeRules"
              />
            </v-col>
            <v-col
              cols="4"
              class="d-flex align-center"
            >
              <v-btn
                color="#2563eb"
                dark
                depressed
                :disabled="!isEmailValid||codeSent"
                block
                class="codebtn"
                @click="sendVerificationCode"
              >
                {{ codeSent ? `${countdown}s` : '获取验证码' }}
              </v-btn>
            </v-col>
          </v-row>
  
          <v-btn
            type="submit"
            block
            color="#2563eb"
            dark
            depressed
          >
            立即注册
          </v-btn>
        </v-form>
      </v-card>
    </v-container>
  </v-app>
</template>
  
  <script setup lang="ts">
 
    import { Login, Register, GetEmailCode } from '@/api/login/index'
    import type { LoginReq, RegisterReq, SendEmailReq } from '@/api/login/types'
    import { VSnackbarQueue } from 'vuetify/labs/VSnackbarQueue'
    import { saveUser, type User } from '@/stores/user'
    import router from '@/router'
    const tab = ref(0)
    const codeSent = ref(false)
    const countdown = ref(60)
  
    const login = reactive({
        email: '',
        password: ''
    })
  
    const register = reactive({
        email: '',
        password: '',
        confirmPassword: '',
        verificationCode: ''
    })
  
    const emailRules = [
    (v: string) => !!v || '邮箱不能为空',
    (v: string) => /.+@.+\..+/.test(v) || '邮箱格式不正确'
    ]

    const passwordRules = [
    (v: string) => !!v || '密码不能为空',
    (v: string) => (v && v.length >= 6) || '密码长度至少6位'
    ]

    const confirmPasswordRules = [
    (v: string) => !!v || '请确认密码',
    (v: string) => v === register.password || '两次输入密码不一致'
    ]

    const verificationCodeRules = [
    (v: string) => !!v || '验证码不能为空',
    (v: string) => /^\d{6}$/.test(v) || '验证码必须为6位数字'
    ]

    // 邮箱有效性计算属性
    const isEmailValid = computed(() => {
        return register.email.length > 0 
    })
    const sendVerificationCode = () => {
        codeSent.value = true
        const timer = setInterval(() => {
            if (countdown.value > 0) {
            countdown.value--
            } else {
            clearInterval(timer)
            codeSent.value = false
            countdown.value = 60
            }
        }, 1000)
        handleGetEmailCode()
    }

    const handleLogin = () => {
        // 登录逻辑
        const loginRequest: LoginReq = {
            email: login.email,
            password: login.password
        }
        Login(loginRequest).then((res) => {
            const data = res.data.data
            if (data.token) {
                const user: User = {
                    email: login.email,
                    token: data.token
                }
                saveUser(user)
                router.push('/user/home')
            }
        })
    }

    const handleRegister = () => {
        // 注册逻辑
        const registerRequest: RegisterReq = {
            code: register.verificationCode,
            email: register.email,
            password: register.password
        }
        Register(registerRequest).then((res) => {
            const data = res.data.data
            if (data.token) {
                const user: User = {
                    email: login.email,
                    token: data.token
                }
                saveUser(user)
                router.push('/user/home')
            }
        })
    }
    const handleGetEmailCode = () => {
        // 获取验证码逻辑
        const emailRequest: SendEmailReq = {
            email: register.email
        }
        GetEmailCode(emailRequest.email).then((res) => {
            const data = res.data.data
            const alert = VSnackbarQueue.alert
            if (data.msg === "success") {
                alert("验证码发送成功")
            } else {
                alert("验证码发送失败")
            }
        })
    }
</script>
  
<style>
.v-text-field--outlined fieldset {
    border-color: #90a4ae !important;
}

.v-input__slot:hover fieldset {
    border-color: #2563eb !important;
}

.v-btn--depressed {
    letter-spacing: 1px;
    transition: all 0.3s ease;
}

.v-btn--depressed:hover {
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(26, 35, 126, 0.2);
}

</style>
