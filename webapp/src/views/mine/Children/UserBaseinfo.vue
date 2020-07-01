<template>
    <div>
        <van-nav-bar title="基本信息"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem;padding-top:0.5rem"/>
        <div class="mainDiv">
            <!-- 教育背景 -->
            <van-cell-group>
                <van-cell title="" value="" class="title" label="1. 教育背景"/>
                <van-cell title="毕业学校" :value="userBaseinfo.school"/>
                <van-cell title="所修专业" :value="userBaseinfo.majorsStudied"/>
                <van-cell title="最高学历" :value="userBaseinfo.highestEducation"/>
                <van-cell title="所上学制" :value="userBaseinfo.schoolSystem"/>
            </van-cell-group>
            <!-- 工作单位 -->
            <van-cell-group>
                <van-cell title="" value="" class="title" label="2. 工作单位"/>
                <van-cell title="工作状态" :value="userBaseinfo.workingState"/>
                <van-cell title="工作单位" :value="userBaseinfo.company"/>
                <van-cell title="单位地域" :value="userBaseinfo.area"/>
                <van-cell title="职务职位" :value="userBaseinfo.jobTitle"/>
                <van-cell title="所在单位服务模式" :value="userBaseinfo.serviceType"/>
                <van-cell title="我的月度收入水平" :value="userBaseinfo.income"/>
                <van-cell title="我的福利待遇情况" :value="userBaseinfo.benefits"/>
                <van-cell title="我目前服务的儿童类型" :value="userBaseinfo.childType"/>
                <van-cell title="我目前主要服务儿童年龄段" :value="userBaseinfo.childAge"/>
            </van-cell-group>
            <!-- 培训经历 -->
            <van-cell-group>
                <van-cell title="" value="" class="title" label="3. 培训经历"/>
                <van-cell title="我到目前为止参加培训数量（一周内的算作短期培训）" :value="userBaseinfo.trainingNumber"/>
                <van-cell title="我去年在培训上的花费" :value="userBaseinfo.trainingFee"/>
                <van-cell value="" :border="true" title="我以往培训详情"/>
                <van-cell-group v-for="(item, index) in userBaseinfo.trainingInfos" :key="index"
                                :title="'培训'+(index + 1)">
                    <van-cell title="参训课程" :value="item.trainingCourse"/>
                    <van-cell title="开始时间" :value="item.beginTime | dateFormat('YYYY-MM-DD')"/>
                    <van-cell title="结束时间" :value="item.endTime | dateFormat('YYYY-MM-DD')"/>
                    <van-cell title="付费方式" :value="item.paymentWay"/>
                </van-cell-group>
            </van-cell-group>
            <div class="buttonDiv">
                <van-button class="button" type="info" size="large" @click="goModifyBaseinfo">重新填写</van-button>
            </div>
        </div>
    </div>
</template>

<script>
    import {userBaseinfo} from "../../../serve/api";

    export default {
        name: "UserBaseinfo",
        data() {
            return {
                userBaseinfo: {}
            }
        },
        computed: {},
        created() {
            this.getUserBaseinfo()
        },
        methods: {
            goModifyBaseinfo() {
                this.$router.push({name: 'baseinfo', query: {modify: 1}});
            },
            async getUserBaseinfo() {
                let result = await userBaseinfo()
                this.userBaseinfo = result.data.userBaseinfo
            }
        }
    }
</script>

<style lang="less" scoped>
    .van-nav-bar__title {
        font-size: 1rem;
    }

    .mainDiv {
        padding-top: 3.6rem;

        .buttonDiv {
            padding: 2rem 1rem 2rem 1rem;

            .button {
                margin: 0.3rem 0rem;
            }
        }

        .title .van-cell__label {
            color: #108EE9;
            font-weight: bold;
            font-size: 0.8rem;
        }
    }
</style>
