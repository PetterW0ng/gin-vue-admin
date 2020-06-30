<template>
    <div>
        <van-nav-bar title="评估结果"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem"/>
        <div class="mainDiv">
            <div align="center">题目全部答题完毕后，</div>
            <div align="center"> 可以查看结果！</div>

            <van-cell-group border=false class="mainCellGroup">
                <van-cell-group class="childCellGroup" v-if="this.userInfo.baseinfoId == 0">
                    <van-button type="info" size="large" plain @click="goToBaseInfo">完善基本信息</van-button>
                </van-cell-group>
                <!-- 订单相关-->
                <van-cell-group class="childCellGroup" v-if="this.userInfo.jobInfoNum == 0">
                    <van-button type="info" size="large" plain @click="goToPage('jobInfo')">职业选择和体验</van-button>
                </van-cell-group>

                <van-cell-group class="childCellGroup" v-if="this.userInfo.perfessionNum == 0">
                    <van-button type="info" size="large" plain @click="goToPage('perfession')">专业理解</van-button>
                </van-cell-group>

                <van-cell-group class="childCellGroup" v-if="this.userInfo.industryNum == 0">
                    <van-button type="info" size="large" plain @click="goToPage('industry')">行业观点</van-button>
                </van-cell-group>
            </van-cell-group>
        </div>
    </div>
</template>

<script>
    import {mapState} from "vuex";

    export default {
        name: "emptyScore",
        computed: {
            ...mapState(['userInfo']),
        },
        methods: {
            goToPage(topicType) {
                if (this.userInfo[topicType + "Num"] > 0) {
                    this.$router.push({path: `/userTopicAnswer/${topicType}`})
                } else {
                    this.$router.push({path: `/topicOptions/${topicType}`})
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
        }
    }
</script>

<style lang="less" scoped>
    .mainDiv {
        padding-top: 10rem;
        align-content: center;
        font-size: 1.8rem;
        height: 100%;

        .mainCellGroup {
            padding: 5rem 2rem 5rem 2rem;

            .childCellGroup {
                padding: 0.35rem 0rem;
            }
        }
    }
</style>
