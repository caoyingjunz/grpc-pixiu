<template>
    <div class="login_body_CFs_DNF">
        <div class="login_gbody_CFs_DNF">
            <div class="login_form_body_CFs_DNF">
                <div class="login_cont_body_CFs_DNF">
                    <div class="login_container_CFs_DNF">
                        <div class="login_container_CFs_DNF">
                            <div class="login_container_form_CFs_DNF">
                                <div class="login_login_form_CFs_DNF">
                                    <div class="login_form_CFs_DNF">
                                        <div direction="vertical" class="u-linear-layout_FaugwXHe" gap="large">
                                            <div class="u-tabs_3_DeFFee">
                                                <ul class="u-tabs_head_3_DeFFee">
                                                    <li title="用户登录" selected="selected" class="u-tabs_item_3_DeFFee">
                                                        <span>用户登录</span></li>
                                                </ul>
                                            </div>
                                            <el-input v-model="loginInfo.name" :prefix-icon="UserFilled"
                                                placeholder="请输入账号" clearable maxlength="128"
                                                @keyup.enter.native="login">
                                                <template #prefix>
                                                    <el-icon class="el-input__icon">
                                                        <user-filled />
                                                    </el-icon>
                                                </template>
                                            </el-input>

                                            <el-input v-model="loginInfo.password" :prefix-icon="Lock"
                                                placeholder="请输入密码" show-password clearable maxlength="128"
                                                @keyup.enter.native="login">
                                                <template #prefix>
                                                    <el-icon class="el-input__icon">
                                                        <lock />
                                                    </el-icon>
                                                </template>
                                            </el-input>

                                            <div direction="horizontal" class="u-linear-layout_FaugwXHe" type="flex"
                                                style="position: relative;">
                                                <el-button class="u-button_Forget" type="text" @click="forget"> 忘记密码
                                                </el-button>
                                            </div>

                                            <div direction="horizontal" class="u-linear-layout_FaugwXHe" type="flex"
                                                justify="center">
                                                <a target="_self" class="u-button_IeaEDOvW" color="primary" size="large"
                                                    :loading="load" @click="login"> 登陆 </a>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {
        UserFilled,
        Lock
    } from '@element-plus/icons-vue'

    export default {
        data() {
            return {
                loginInfo: {
                    name: "",
                    password: "",
                },
                load: false,
                // 图标高级
                Lock: '',
                UserFilled: ''
            }
        },
        methods: {
            forget() {
                this.$message.error("忘记密码，请联系管理员")
            },
            async login() {
                this.load = true

                // 发送登陆请求
                this.$http.post("/user/login", this.loginInfo)
                    .then((res) => {
                        const token = res.data.result
                        window.sessionStorage.setItem("token", token)
                        window.sessionStorage.setItem("account", this.loginInfo.name)

                        this.$message.success("登陆成功")
                        this.$router.push('/index')
                    })
                    .catch((err) => {
                        this.load = false
                        if (err.status == 500) {
                            return this.$message.error(err.statusText)
                        }

                        return this.$message.error(err.data.message)
                    })
            },
        },
        components: {
            Lock,
            UserFilled
        }
    }
</script>

<style>
    /* 有用 */
    .u-button_Forget {
        margin: 0;
        position: absolute;
        right: 0px;
        overflow: visible;
        font: inherit;
        color: #6383dc;

        padding: 0px 0px 12px;
    }

    /* 有用 */
    .u-button_IeaEDOvW {
        -webkit-appearance: none;
        margin: 0;
        border: none;
        overflow: visible;
        font: inherit;
        text-transform: none;
        text-decoration: none;
        cursor: pointer;
        /* background: none; */
        display: inline-block;
        text-align: center;
        padding: 0 12px;
        height: 36px;
        width: 3000px;
        line-height: 34px;
        /* background: rgb(38, 39, 110); */
        /* color: #0c1f7f; */
        border: 1px solid #e5e8ed;
    }

    /* 有用 */
    .el-input__inner {
        height: 40px;
        /* 边框 */
        border: 0px solid #dfe4ec;
    }

    /* 有用 */
    .u-button_IeaEDOvW[color=primary] {
        background: #508ae2;
        color: #fff;
        border: none;
        line-height: 34px;
    }

    /* 有用 */
    .u-tabs_item_3_DeFFee {
        cursor: pointer;
        display: inline-block;
        box-sizing: content-box;
        height: 30px;
        line-height: 30px;
        padding: 0 20px;
        border: 1px solid transparent;
        border-bottom: none;
        margin-bottom: -1px;
    }

    /* 有用 */
    .u-tabs_item_3_DeFFee[selected] {
        background: #fff;
        color: #508ae2;
        border-color: #e0e6ed;
        border-top: 2px solid #508ae2;
    }

    /* 有用 */
    .u-tabs_head_3_DeFFee {
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
        height: 30px;
        border-bottom: 1px solid #ece5e3;
    }

    /* 有用 */
    .u-linear-layout_FaugwXHe[direction=vertical][gap=large]>:not(:last-child) {
        margin-bottom: 30px;
    }

    /* 有用 */
    .u-linear-layout_FaugwXHe[type=flex] {
        display: flex;
        text-align: inherit;
    }

    /* 有用 */
    .login_form_CFs_DNF {
        width: 325px;
        margin: 0 auto;
        padding: 0;
        line-height: 1;
    }

    /* 有用 */
    .login_login_form_CFs_DNF {
        float: none;
        width: 450px;
        height: 380px;
        margin: auto;
        padding-top: 30px;
        /* 上下移动 */
        box-sizing: border-box;
        background: #fff;
        transition: box-shadow .2s;
        box-shadow: 0 0 10px 0 rgb(80 90 109 / 16%);
    }

    /* 有用 */
    .login_container_CFs_DNF {
        display: table;
        background: transparent;
        width: 1160px;
        /* min-height: 600px; */
        margin: 0 auto;
        padding-top: 26px;
    }

    * {
        margin: 0;
        padding: 0;
    }

    /* 有用 */
    .login_form_body_CFs_DNF {
        background: none;
        margin: 0;
        padding: 0;
        padding-top: 150px;
    }

    /* 有用 */
    .login_body_CFs_DNF {
        position: relative;
        box-sizing: border-box;
        padding-bottom: 133px;
        min-height: 100vh;
        background: url(../../../static/beijing.jpg) no-repeat;
        background-size: cover;
    }

    /* 有用 宽度*/
    .u-head_head_yingjun {
        width: 100%;
        min-width: 100px;
        height: 65px;
        margin: 0 1px;
        background-color: #eaedf3;
        box-shadow: 0 2px 4px -1px rgb(0 0 0 / 20%), 0 4px 6px 0 rgb(0 0 0 / 15%);
        position: fixed;
        z-index: 1000;
    }

    /* 头部logo */
    .u-head_logo_yingjun {
        margin-left: 10px;
        margin-right: 50px;
        display: inline-block;
        height: 65px;
    }
</style>