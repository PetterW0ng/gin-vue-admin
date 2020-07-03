<template>
    <div>
        <van-nav-bar title="评估结果"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem;padding-top:0.5rem"/>
        <div class="mainDiv">
            <div align="center"><img src="../../../images/mine/main.png"/></div>
            <div align="center" class="gantan"><img src="../../../images/mine/gantan.png"/></div>
            <div align="center" style="padding-top: 10px;font-size: 22px">还有未完成内容</div>
            <div align="center" style="padding-top: 75px;font-size: 18px;color: #A0A0A0">填写完成后可查看评估结果</div>
            <div align="center"><img src="../../../images/mine/xiayibu.png"/></div>
            <div align="center" class="topicDiv" v-if="this.userInfo.baseinfoId == 0" @click="goToBaseInfo">《完善基本信息》
            </div>
            <div align="center" class="topicDiv" v-if="this.userInfo.jobInfoNum == 0" @click="goToPage('jobInfo')">
                《职业选择和体验》
            </div>
            <div align="center" class="topicDiv" v-if="this.userInfo.perfessionNum == 0"
                 @click="goToPage('perfession')">《专业理解》
            </div>
            <div align="center" class="topicDiv" v-if="this.userInfo.industryNum == 0" @click="goToPage('industry')">
                《行业观点》
            </div>
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
        created() {

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
    .van-nav-bar__title {
        font-size: 1rem;
    }

    .mainDiv {
        padding-top: 5rem;
        align-content: center;
        font-size: 1.8rem;
        height: 100%;

        .gantan {
            padding-top: 62px;
        }

        .topicDiv {
            font-size: 18px;
            color: #0F8EE9;
        }

        .mainCellGroup {
            padding: 5rem 2rem 5rem 2rem;

            .childCellGroup {
                padding: 0.35rem 0rem;
            }
        }
    }
</style>
