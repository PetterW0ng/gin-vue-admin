<template>
    <div id="login">
        <div class="loginBox">
            <div class="auth-form">
                <van-tabs v-model="active" animated :border="false">
                    <!-- 登录 -->
                    <van-tab title="登录">
                        <van-cell-group class="cellGroup" :border="false">
                            <!-- 手机号快捷登录 -->
                            <van-field v-model="login_phone"
                                       required
                                       clearable
                                       maxlength="11"
                                       placeholder="请输入手机号"/>
                            <van-field center
                                       clearable
                                       required
                                       maxlength="6"
                                       v-model="login_captcha"
                                       placeholder="请输入验证码">
                                <van-button slot="button"
                                            size="small"
                                            type="primary"
                                            v-if="!loginCountDown"
                                            :disabled="loginCaptchaDisable"
                                            @click="sendVerifyCode(true)">发送验证码
                                </van-button>
                                <van-button slot="button"
                                            size="small"
                                            type="primary"
                                            disabled=""
                                            v-model="login_captcha"
                                            v-else>已发送 {{loginCountDown}} s
                                </van-button>
                            </van-field>

                            <van-button type="info"
                                        size="large"
                                        style="margin-top:4rem"
                                        @click='login'>立即登录
                            </van-button>
                        </van-cell-group>
                    </van-tab>
                    <!-- 注册 -->
                    <van-tab title="注册">
                        <van-cell-group class="cellGroup" :border="false">
                            <van-field v-model="register.username"
                                       clearable
                                       maxlength="6"
                                       placeholder="姓名"
                                       required/>
                            <van-field :value="gender"
                                       readonly
                                       is-link
                                       placeholder="性别"
                                       @click="showChooseSex = true"
                                       required/>
                            <van-field :value="birthday"
                                       placeholder="出生日期"
                                       required
                                       is-link
                                       @click="showDateTimePopView = true"
                                       readonly/>
                            <van-field v-model="register.telphone"
                                       clearable
                                       maxlength="11"
                                       placeholder="手机号"
                                       required/>
                            <van-field center
                                       clearable
                                       required
                                       maxlength="6"
                                       v-model="register.verificationCode"
                                       placeholder="请输入验证码">
                                <van-button slot="button"
                                            size="small"
                                            type="primary"
                                            v-if="!registerCountDown"
                                            :disabled="registerCaptchaDisable"
                                            @click="sendVerifyCode(false)">发送验证码
                                </van-button>
                                <van-button slot="button"
                                            size="small"
                                            type="primary"
                                            disabled=""
                                            v-model="register.verificationCode"
                                            v-else>已发送 {{registerCountDown}} s
                                </van-button>
                            </van-field>
                        </van-cell-group>

                        <!-- 底部声明 -->
                        <p class="agreement">
                            <van-checkbox v-model="isChecked" checked-color='#45c763' class="checkBoxS">我已阅读并同意<a
                                    @click.stop=agreement() class="agreement-box">《A-PKU能力评估同意书》</a></van-checkbox>
                        </p>
                        <van-button type="info"
                                    size="large"
                                    style="margin-top:4rem"
                                    @click='handleRegister'>注册
                        </van-button>
                    </van-tab>
                </van-tabs>
            </div>
        </div>

        <!-- 时间选择器 -->
        <van-popup v-model="showDateTimePopView"
                   round
                   position="bottom">
            <van-datetime-picker v-model="currentDate"
                                 type="date"
                                 @confirm="confirm"
                                 @cancel="showDateTimePopView = false"
                                 :formatter="formatter"
                                 :max-date="maxDate"
                                 :min-date="minDate"/>
        </van-popup>
        <!-- 性别选择器 -->
        <van-popup v-model="showChooseSex"
                   position="bottom"
                   :style="{ height: '25%' }">
            <van-radio-group v-model="register.gender">
                <van-cell-group style="margin-top:2rem"
                                @click="clickCell(radio)">
                    <van-cell title="男"
                              clickable
                              @click="radio = 1">
                        <van-radio slot="right-icon"
                                   name="1"
                                   checked-color="#07c160"/>
                    </van-cell>
                    <van-cell title="女"
                              clickable
                              @click="radio = 2">
                        <van-radio slot="right-icon"
                                   name="2"
                                   checked-color="#07c160"/>
                    </van-cell>
                </van-cell-group>
            </van-radio-group>
        </van-popup>
    </div>
</template>

<script type="text/javascript">
    // 引入Vant的组件
    import {Toast} from 'vant'
    // 引入API调用接口
    import {getPhoneCaptcha, register, login} from '../../serve/api/index.js'
    // 引入vuex
    import {mapActions} from 'vuex'
    import Moment from "moment";

    export default {
        data() {
            return {
                active: 0,
                loginCountDown: 0,                 // 倒计时
                login_phone: '',              // 手机号码
                login_captcha: '',            // 登录短信验证码

                register: {
                    username: '',           // 注册用户名
                    gender: 0,              // 性别
                    birthday: '',           // 生日
                    telphone: '',           // 手机号
                    verificationCode: ''
                },
                registerCaptcha: '',               // 注册短信验证码
                registerCountDown: 0,         // 倒计时

                userInfo: null,
                isChecked: false,
                // 性别、生日组件
                showDateTimePopView: false,
                showChooseSex: false,
                currentDate: new Date('1989/01/01'),
                // 最小时间
                minDate: new Date('1949/01/01'),
                maxDate: new Date('2019/12/31'),
            };
        },
        computed: {
            // 2.发送验证码按钮显示
            loginCaptchaDisable() {
                if (this.login_phone.length > 10 && this.phoneRegex(this.login_phone)) {
                    return false;
                } else {
                    return true;
                }
            },
            gender() {
                return this.register.gender == 0 ? "" : (this.register.gender == 1 ? "男" : "女");
            },
            birthday() {
                if (this.register.birthday == "") {
                    return "";
                }
                return Moment(this.register.birthday).format("YYYY-MM-DD");
            },
            // 2.发送验证码按钮显示
            registerCaptchaDisable() {
                if (this.register.telphone.length > 10 && this.phoneRegex(this.register.telphone)) {
                    return false;
                } else {
                    return true;
                }
            }
        },
        methods: {
            // 0.mapActions 同步用户信息
            ...mapActions(['syncuserInfo', 'synToken']),
            clickCell(radio) {
                this.register.gender = radio
                this.showChooseSex = false
            },
            // DateTime pcikView 确定
            confirm(value) {
                this.register.birthday = Moment(value).format();
                this.showDateTimePopView = false;
            },
            // 格式化DateTime pickView
            formatter(type, value) {
                if (type === 'year') {
                    return `${value}`;
                } else if (type === 'month') {
                    return `${value}`
                } else if (type === 'day') {
                    return `${value}`
                }
                return value;
            },
            // 4.获取手机验证码
            async sendVerifyCode(isLogin) {
                // 4.2 获取短信验证码
                let params = {}
                if (isLogin) {
                    this.loginCountDown = 60;
                    this.timeIntervalID1 = setInterval(() => {
                        this.loginCountDown--;
                        // 4.1 如果减到0 则清除定时器
                        if (this.loginCountDown == 0) {
                            clearInterval(this.timeIntervalID1);
                        }
                    }, 1000)

                    params.phone = this.login_phone;
                    params.isLogin = isLogin;
                } else {
                    this.registerCountDown = 60;
                    this.timeIntervalID2 = setInterval(() => {
                        this.registerCountDown--;
                        // 4.1 如果减到0 则清除定时器
                        if (this.registerCountDown == 0) {
                            clearInterval(this.timeIntervalID2);
                        }
                    }, 1000)

                    params.phone = this.register.telphone;
                    params.isLogin = isLogin;
                }
                let result = await getPhoneCaptcha(params);
                if (result.code == 0) {
                    Toast.success("验证码发送成功")
                }
            },
            // 5.登录
            async login() {
                // 5.1手机验证码登录
                // 5.1.1 验证手机号
                if (!this.phoneRegex(this.login_phone)) {
                    Toast({
                        message: "手机号不正确",
                        duration: 800
                    });
                    return;
                }
                // 5.1.3 请求后台登录接口
                let ref = await login(this.login_phone, this.login_captcha);
                // 设置userInfo 保存到vuex和本地
                this.syncuserInfo(ref.data.user);
                this.synToken(ref.data.token);
                this.$router.push({path: "/home"});
            },
            // 6.注册
            async handleRegister() {
                if (this.register.username.length < 1) {
                    Toast({
                        message: '姓名不能为空',
                        duration: 800
                    })
                } else if (!this.phoneRegex(this.register.telphone)) {
                    Toast({
                        message: '手机格式不正确',
                        duration: 800
                    })
                } else if (this.register.gender == 0) {
                    Toast({
                        message: '性别不能为空',
                        duration: 800
                    })
                } else if (this.register.birthday.length < 6) {
                    Toast({
                        message: '出生日期不能为空',
                        duration: 800
                    })
                } else if (this.register.verificationCode.length < 6) {
                    Toast({
                        message: '验证码填写不正确',
                        duration: 800
                    })
                } else if (!this.isChecked) {
                    Toast({
                        message: '请勾选用户协议',
                        duration: 800
                    })
                } else {
                    // 6.1 请求后台登录接口
                    let result = await register(this.register);
                    // 设置userInfo\token 保存到本地并跳转至home页
                    this.syncuserInfo(result.data.user);
                    this.synToken(result.data.token);
                    this.$router.push({path: "/home"});
                }
            },
            // 7.用户协议
            agreement() {
                Toast({
                    message: "A-PKU职业能力评估系统同意书\n" +
                        " \n" +
                        "很高兴邀请您使用全国“孤独症从业人员能力分析”系统，随着孤独症谱系障碍儿童患病率逐年增高（美国CDC2020年报道），我国对于孤独症康复从业人员的数量需求和质量需求也在逐年递增。而目前我国缺少对于从业人员职业等级的认证标准，及相关培训标准，从业人员（老师）不容易判断自身能力水平以及如何提升自己。\n" +
                        " \n" +
                        "我们旨在通过本系统，为孤独症康复从业人员完成职业等级的建议性评估，并根据评估结果，给予进一步学习和提升的建议。\n" +
                        " \n" +
                        "系统中将需要您提供您的个人信息、学历信息、工作经历、职业相关信息、职业能力信息、个人观点及态度等，并依据这些信息完成建议性评估。您填写的信息将严格保密，不会泄露，除非经过您本人的许可。该系统汇总的报告、或用于商用时，将不会披露您的个人任何信息。参与的全过程中，您可以自愿选择参加或中途退出。是否参加完全取决于您的意愿。如果您对本系统有任何疑问，您可以向工作人员提出来。您也可以拒绝参加，或在使用过程中的任何时间退出本系统，这都不会影响您的正常权益。\n" +
                        " \n" +
                        "为确保评估结果的准确性，需要您根据事实进行填写。\n" +
                        " \n" +
                        "在参与过程中，如果您有任何与本系统有关的疑问或不理解的事情，或如果您认为在参与本系统中您受到非公正对待，可向教务微信号（微信名称：tonglaoshi_njk）咨询。\n",
                    duration: 8000
                })
            },
            // 正则验证
            phoneRegex(number) {
                return (/[1][3,4,5,6,7,8][0-9]{9}$/.test(number));
            }
        }
    }
</script>
<style lang="less" scoped>
    #login {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        background-attachment: fixed;

        .loginBox {
            display: flex;
            /*align-items: center;*/
            justify-content: center;
            position: fixed;
            opacity: 0.95;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            z-index: 500;

            .auth-form {
                position: relative;
                padding: 5rem 2rem 5rem 2rem;
                width: 26.5rem;
                max-width: 100%;
                font-size: 1.167rem;
                background-color: #fff;
                border-radius: 8px;
                box-sizing: border-box;

                .cellGroup {
                    padding-top: 2rem;
                }

                .van-tab .van-tab--active {
                    font-size: 22px;
                }
            }

            .verificationImage {
                position: absolute;
                right: 0;
                width: 8rem;
                height: 3rem;
                margin-left: 3rem;
            }

            .switchLogin {
                margin-top: 1rem;
                font-size: 0.78rem;
            }
        }
    }

    .agreement {
        // margin-top: 1.667rem;
        line-height: 1rem;
        color: #767676;
        font-size: 0.867rem;
        padding-top: 2rem;

        .agreement-box {
            color: blue;
        }
    }

    .checkBoxS {
        padding: 0.1rem;
        font-size: 0.425rem;
    }

</style>
