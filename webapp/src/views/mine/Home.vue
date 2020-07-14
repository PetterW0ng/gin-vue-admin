<template>
    <div class="mainDiv">
        <div class="topDiv">
            <div align="left" style="padding-left: 34px;padding-top: 56px;font-size: 16px;color: #FFFFFF;">孤独症康复从业人员
                <img src="../../images/home/tuichu3x.png"
                     style="vertical-align:middle;width: 20px;height: 20px;margin-left: 9.8rem" @click="logoutHandle"/>
            </div>
            <div align="left" style="padding-left: 34px;padding-top: 10px;font-size: 26px;color: #FFFFFF;">
                A-PKU职业能力评估系统
            </div>
            <div style="padding-left: 34px;padding-top: 20px;text-align: left">
                <van-button round type="info" color="#81D4FF" style="padding: 2px 30px" @click="goToBaseInfo"
                            text="完善基本信息"/>
            </div>
        </div>
        <div class="container">
            <div class="item1Div" style="background-color: #CEEDFD;" @click="goToPage('jobInfo')">
                <div style="padding-left: 37px;font-size: 18px;font-weight: bold">职业选择和体验 <img
                        src="../../images/home/link-arrow-blue3x.png"
                        style="vertical-align:middle;width: 20px;height: 20px;padding-bottom: 5px"/></div>
                <div style="padding-left: 37px;color: #ABB3BB;padding-top: 6px;font-size: 16px">Career choice and
                    experience
                </div>
            </div>
            <img src="../../images/home/duigou3x.png" class="duigou" v-if="this.userInfo.jobInfoNum != 0"/>
        </div>
        <div class="container">
            <div class="item2Div" style="background-color: #FFF4E5" @click="goToPage('perfession')">
                <div style="padding-left: 37px;font-size: 18px;font-weight: bold">专业理解 <img
                        src="../../images/home/link-arrow-orange3x.png"
                        style="vertical-align:middle;width: 20px;height: 20px;padding-bottom: 5px"/></div>
                <div style="padding-left: 37px;color: #ABB3BB;padding-top: 6px;font-size: 16px">Professional
                    understanding
                </div>
            </div>
            <img src="../../images/home/duigou3x.png" class="duigou" v-if="this.userInfo.perfessionNum != 0"/>
        </div>
        <div class="container">
            <div class="item3Div" style="background-color: #D0DAF2" @click="goToPage('industry')">
                <div style="padding-left: 37px;font-size: 18px;font-weight: bold">行业观点 <img
                        src="../../images/home/link-arrow-purple3x.png"
                        style="vertical-align:middle;width: 20px;height: 20px;padding-bottom: 5px"/></div>
                <div style="padding-left: 37px;color: #ABB3BB;padding-top: 6px;font-size: 16px">Industry point of view
                </div>
            </div>
            <img src="../../images/home/duigou3x.png" class="duigou" v-if="this.userInfo.industryNum != 0"/>
        </div>
        <van-button round type="info" size="large" @click="goToMyScore"
                    style="margin-top: 25px;width: 70%;margin-bottom: 20px" text="查看评估结果"/>
    </div>
</template>

<script type="text/javascript">
  // 引入vuex
  import {mapActions, mapState} from 'vuex'
  import {loadUserInfo} from "../../serve/api";

  export default {
    data() {
      return {
        // 版本信息
      }
    },
    computed: {
      ...mapState(['userInfo']),
    },
    created() {
      this.reloadUser()
    },
    methods: {
        ...mapActions(['syncuserInfo', 'logout']),
        goToPage(topicType) {
            if (this.userInfo[topicType + "Num"] > 0) {
                this.$router.push({path: `/userTopicAnswer/${topicType}`})
            } else {
                this.$router.push({path: `/topicOptions/${topicType}`})
            }
        },
        goToMyScore() {
            if (this.userInfo.perfessionNum == 0 || this.userInfo.industryNum == 0 || this.userInfo.jobInfoNum == 0 || this.userInfo.baseinfoId == 0) {
                this.$router.push({name: 'emptyScore'})
            } else {
                this.$router.push({name: 'userScore'})
            }
        },
        logoutHandle() {
            this.logout();
            this.$router.push({name: 'login'})
        },
        goToBaseInfo() {
            if (this.userInfo["baseinfoId"] > 0) {
                this.$router.push({
                    'name': 'userBaseinfo'
                });
            } else {
                this.$router.push({
                    'name': 'baseinfo'
                });
            }
      },
      async reloadUser() {
        let result = await loadUserInfo()
        this.syncuserInfo(result.data);
      }
    }
  }
</script>

<style lang="less" scoped>
    .mainDiv {
        text-align: center;
        align-items: center;
        align-content: center;

        .container {
            position: relative;

            .duigou {
                position: absolute;
                top: 0px;
                right: 21px;
                width: 42px;
                height: 27px;
            }
        }

        .topDiv {
            width: 100%;
            height: 281px;
            background-image: url("../../images/home/head-bg-img.png");
            background-repeat: no-repeat;
        }

        .item1Div {
            width: 90%;
            border-radius: 1em;
            padding-top: 33px;
            padding-bottom: 37px;
            margin: -3rem 1rem 0rem 1rem;
            text-align: left;
            -webkit-box-shadow: 0px 3px 0px #CEEDFD;
            -moz-box-shadow: 0px 3px 0px #CEEDFD;
            box-shadow: 0px 3px 0px #CEEDFD;
        }

        .item2Div {
            width: 90%;
            border-radius: 1em;
            padding-top: 33px;
            padding-bottom: 37px;
            margin: 1rem 1rem 0rem 1rem;
            text-align: left;
            -webkit-box-shadow: 0px 3px 0px #FFF4E5;
            -moz-box-shadow: 0px 3px 0px #FFF4E5;
            box-shadow: 0px 3px 0px #FFF4E5;
        }

        .item3Div {
            width: 90%;
            border-radius: 1em;
            padding-top: 33px;
            padding-bottom: 37px;
            margin: 1rem 1rem 0rem 1rem;
            text-align: left;
            -webkit-box-shadow: 0px 3px 0px #D0DAF2;
            -moz-box-shadow: 0px 3px 0px #D0DAF2;
            box-shadow: 0px 3px 0px #D0DAF2;
        }
    }

</style>
