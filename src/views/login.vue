<template>
  <div class="login">
    <el-row class="w-full h-full">
      <el-col :span="17">
        <div class="logo h-full">
          <img src="@/assets/images/login-security.svg" class="logo-image" />
        </div>
      </el-col>
      <el-col :span="7" class="side h-full">
        <el-form ref="loginRef" :model="loginForm" :rules="loginRules" class="login-form" label-position="top">
          <h3 class="title">{{ $t("login.title1") }}</h3>
          <el-form-item label="用户名" prop="username">
            <el-input v-model="loginForm.username" type="text" size="large" auto-complete="off"
              :placeholder="$t('login.username')">
              <template #prefix><svg-icon icon-class="user" class="el-input__icon input-icon" /></template>
            </el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input v-model="loginForm.password" type="password" size="large" auto-complete="off"
              :placeholder="$t('login.password')" @keyup.enter="handleLogin" show-password>
              <template #prefix><svg-icon icon-class="password" class="el-input__icon input-icon" /></template>
            </el-input>
          </el-form-item>
          <el-form-item prop="code" v-if="captchaEnabled">
            <el-input v-model="loginForm.code" size="large" auto-complete="off" :placeholder="$t('login.code')"
              style="width: 63%" @keyup.enter="handleLogin">
              <template #prefix><svg-icon icon-class="validCode" class="el-input__icon input-icon" /></template>
            </el-input>
            <div class="login-code">
              <img :src="codeUrl" @click="getCode" class="login-code-img" />
            </div>
          </el-form-item>
          <el-checkbox v-model="loginForm.rememberMe" style="margin:0px 0px 25px 0px;">{{ $t("login.remember")
          }}</el-checkbox>
          <el-form-item style="width:100%;">
            <el-button :loading="loading" size="large" type="primary" style="width:100%;" @click.prevent="handleLogin">
              <span v-if="!loading">{{ $t("login.title") }}</span>
              <span v-else>{{ $t("login.loading") }}</span>
            </el-button>
            <div style="float: right;" v-if="register">
              <router-link class="link-type" :to="'/register'">{{ $t("login.register") }}</router-link>
            </div>
          </el-form-item>
        </el-form>
      </el-col>

    </el-row>
    <!--  底部  -->
    <div class="el-login-footer">
      <span>{{ $t("login.copyright") }}</span>
    </div>
  </div>
</template>

<script setup>
import { getCodeImg } from "@/api/login";
import Cookies from "js-cookie";
import { encrypt, decrypt } from "@/utils/jsencrypt";
import useUserStore from '@/store/modules/user'

const userStore = useUserStore()
const route = useRoute();
const router = useRouter();
const { proxy } = getCurrentInstance();

const loginForm = ref({
  username: "admin",
  password: "admin123",
  rememberMe: false,
  code: "",
  uuid: ""
});

const loginRules = {
  username: [{ required: true, trigger: "blur", message: "请输入您的账号" }],
  password: [{ required: true, trigger: "blur", message: "请输入您的密码" }],
  code: [{ required: true, trigger: "change", message: "请输入验证码" }]
};

const codeUrl = ref("");
const loading = ref(false);
// 验证码开关
const captchaEnabled = ref(true);
// 注册开关
const register = ref(false);
const redirect = ref(undefined);

watch(route, (newRoute) => {
  redirect.value = newRoute.query && newRoute.query.redirect;
}, { immediate: true });

function handleLogin() {
  proxy.$refs.loginRef.validate(valid => {
    if (valid) {
      loading.value = true;
      // 勾选了需要记住密码设置在 cookie 中设置记住用户名和密码
      if (loginForm.value.rememberMe) {
        Cookies.set("username", loginForm.value.username, { expires: 30 });
        Cookies.set("password", encrypt(loginForm.value.password), { expires: 30 });
        Cookies.set("rememberMe", loginForm.value.rememberMe, { expires: 30 });
      } else {
        // 否则移除
        Cookies.remove("username");
        Cookies.remove("password");
        Cookies.remove("rememberMe");
      }
      // 调用action的登录方法
      userStore.login(loginForm.value).then(() => {
        const query = route.query;
        const otherQueryParams = Object.keys(query).reduce((acc, cur) => {
          if (cur !== "redirect") {
            acc[cur] = query[cur];
          }
          return acc;
        }, {});
        router.push({ path: redirect.value || "/", query: otherQueryParams });
      }).catch(() => {
        loading.value = false;
        // 重新获取验证码
        if (captchaEnabled.value) {
          getCode();
        }
      });
    }
  });
}

function getCode() {
  getCodeImg().then(res => {
    captchaEnabled.value = res.captchaEnabled === undefined ? true : res.captchaEnabled;
    if (captchaEnabled.value) {
      codeUrl.value = "data:image/gif;base64," + res.img;
      loginForm.value.uuid = res.uuid;
    }
  });
}

function getCookie() {
  const username = Cookies.get("username");
  const password = Cookies.get("password");
  const rememberMe = Cookies.get("rememberMe");
  loginForm.value = {
    username: username === undefined ? loginForm.value.username : username,
    password: password === undefined ? loginForm.value.password : decrypt(password),
    rememberMe: rememberMe === undefined ? false : Boolean(rememberMe)
  };
}

getCode();
getCookie();
</script>

<style lang='scss' scoped>
.login {
  display: flex;
  // justify-content: center;
  align-items: center;
  height: 100%;
  background-color: #fff;

  &::before {
    content: "";
    position: absolute;
    height: 100%;
    width: 100%;
    opacity: .3;
    background: radial-gradient(#d2f1df, #d3d7fa, #bad8f4) 0 0 / 400% 400%;
    animation: 15sease 0s infinite normal none running gradient;
  }
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;

  .logo-image {
    width: 500px;
    height: auto;
  }
}

.side {
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.title {
  margin: 0px auto 30px auto;
  text-align: center;
  color: #333;
}

.login-form {
  width: 400px;

  form * {
    font-family: 'Poppins', sans-serif;
    color: #333;
    letter-spacing: 0.5px;
    outline: none;
    border: none;
  }

  .el-input {
    height: 40px;

    input {
      height: 40px;
    }
  }

  .input-icon {
    height: 39px;
    width: 14px;
    margin-left: 0px;
  }
}

.login-tip {
  font-size: 13px;
  text-align: center;
  color: #333;
}

.login-code {
  width: 33%;
  height: 40px;
  float: right;

  img {
    cursor: pointer;
    vertical-align: middle;
  }
}

.el-login-footer {
  height: 40px;
  line-height: 40px;
  position: fixed;
  bottom: 0;
  width: 100%;
  text-align: center;
  color: #fff;
  font-family: Arial;
  font-size: 12px;
  letter-spacing: 1px;
}

.login-code-img {
  height: 40px;
  padding-left: 12px;
}

.background {
  width: 430px;
  height: 460px;
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
}

.shape {
  height: 200px;
  width: 200px;
  position: absolute;
  border-radius: 50%;
}

.shape:first-child {
  background: linear-gradient(to left, #1845ad,
      #23a2f6);
  left: -80px;
  top: -80px;
}

.shape:last-child {
  background: linear-gradient(to right,
      #ff512f,
      #f09819);
  right: -30px;
  bottom: -80px;
}
</style>
