<template>
  <div>
      <van-nav-bar title="A-PKU职业能力评估系统"
                   :fixed=true
                   :border=false
                   style="height:2.5rem"/>
      <van-cell-group class="mainCellGroup">
          <van-cell-group class="childCellGroup">
              <van-button type="info" size="large" plain @click="goToBaseInfo">完善基本信息</van-button>
          </van-cell-group>
          <!-- 订单相关-->
          <van-cell-group class="childCellGroup">
              <van-button type="info" size="large" plain @click="goToPage('jobInfo')">职业选择和体验</van-button>
          </van-cell-group>

          <van-cell-group class="childCellGroup">
              <van-button type="info" size="large" plain @click="goToPage('perfession')">专业理解</van-button>
          </van-cell-group>

          <van-cell-group class="childCellGroup">
              <van-button type="info" size="large" plain @click="goToPage('industry')">行业观点</van-button>
          </van-cell-group>

          <van-cell-group class="childCellGroup">
              <van-button type="info" size="large" plain @click="goToMyScore">查看评估结果</van-button>
          </van-cell-group>
      </van-cell-group>
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
        },
        goToMyScore() {
            if (this.userInfo.perfessionNum == 0 || this.userInfo.industryNum == 0 || this.userInfo.jobInfoNum == 0 || this.userInfo.baseinfoId == 0) {
                this.$router.push({name: 'emptyScore'})
            } else {
                this.$router.push({name: 'userScore'})
            }
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
    .mainCellGroup {
        padding: 5rem 2rem 5rem 2rem;

        .childCellGroup {
            padding: 0.35rem 0rem;
        }
    }
</style>
