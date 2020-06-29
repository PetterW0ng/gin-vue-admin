<template>
  <div id="mine">
    <van-nav-bar title="A-PKU职业能力评估系统"
                 :fixed=true
                 :border=false
                 style="height:2.5rem"/>
    <van-cell-group style="margin-top:2.4rem">
      <van-button type="primary" size="large" @click="goToBaseInfo">完善基本信息</van-button>
    </van-cell-group>
    <!-- 订单相关-->
    <van-cell-group>
      <van-button type="primary" size="large" @click="goToPage('jobInfo')">职业选择和体验</van-button>
    </van-cell-group>

    <van-cell-group>
      <van-button type="primary" size="large" @click="goToPage('perfession')">专业理解</van-button>
    </van-cell-group>

    <van-cell-group>
      <van-button type="primary" size="large" @click="goToPage('industry')">行业观点</van-button>
    </van-cell-group>

    <!--路由的出口-->
    <transition name="router-slider"
                mode="out-in">
    </transition>
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
      ...mapActions(['syncuserInfo']),
      goToPage(topicType) {
        if (this.userInfo[topicType + "Num"] > 0) {
          this.$router.push({path: `/userTopicAnswer/${topicType}`})
        } else {
          this.$router.push({path: `/topicOptions/${topicType}`})
        }
        // this.$router.push({ path: `/topicOptions/${topicType}`})
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
  #mine {
    width: 100%;
    background-color: #f5f5f5;
    margin-bottom: 3rem;

    .version {
      margin: 0 auto;
      text-align: center;
      font-size: 0.6rem;
      background-color: #ffffff;
      height: 2rem;
      color: grey;
      line-height: 2rem;
    }

    .van-nav-bar {
      background-color: #3bba63;
      font-size: 0.6rem;
    }

    .van-nav-bar__title {
      color: white;
    }

    .personMsg {
      display: flex;
      align-items: center;

      .sex {
        position: absolute;
        top: 3.5rem;
        left: 3.8rem;
        width: 0.1rem;
        height: 0.1rem;

        img {
          width: 1rem;
          height: 1rem;
        }
      }

      img {
        width: 4rem;
        height: 4rem;
        border-radius: 50%;
      }

      .personInfo {
        display: flex;
        flex-direction: column;
        margin-left: 0.8rem;
      }
    }

    .van-cell__left-icon {
      color: #45c763;
      font-size: 1.2rem;
    }

    /*转场动画*/

    .router-slider-enter-active,
    .router-slider-leave-active {
      transition: all 0.3s;
    }

    .router-slider-enter,
    .router-slider-leave-active {
      transform: translate3d(2rem, 0, 0);
      opacity: 0;
    }
  }
</style>
