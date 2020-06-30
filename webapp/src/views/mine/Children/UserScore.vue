<template>
    <div>
        <van-nav-bar title="评估结果"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem"/>
        <div class="mainDiv">
            <!-- 雷达图 -->
            <div id="radarChartDiv" style="width: 100%;height: 400px"></div>
            <div id="recomand" style="padding:0.2em">
                <!-- 结果分析 -->
                <div class="title">结果分析</div>
                <van-cell :value="' 您在 ' + gtStandard + ' 方面的能力已满足职业要求，非常棒。'" v-if="gtStandard.length > 0"/>
                <van-cell :value="' 您在 '+ ltStandard+' 方面的能力有所欠缺，还不能满足当前工作需要，建议学习以下内容进行提升：'"
                          v-if="ltStandard.length > 0"/>
                <!-- 培训课程 -->
                <div v-if="recommendCourses.length > 0">
                    <div class="title">培训课程</div>
                    <div v-for="item in recommendCourses">
                        <van-cell :title="item.recommendObject" icon="star-o" @click="goToStore(item.id)"/>
                    </div>
                </div>
                <!-- 阅读书籍 -->
                <div v-if="recommendBooks.length > 0">
                    <div class="title">阅读书籍</div>
                    <div v-for="item in recommendBooks">
                        <van-cell :title="item.recommendObject" icon="notes-o"/>
                        <van-cell :value="item.remark"/>
                    </div>
                </div>
                <!-- 干预工具 -->
                <div v-if="recommendTools.length > 0">
                    <div class="title">干预工具</div>
                    <div v-for="item in recommendTools">
                        <van-cell :title="item.recommendObject" icon="setting-o"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {userScore, toStore} from "../../../serve/api";
    import echarts from 'echarts'
    import {Toast} from 'vant'

    export default {
        name: "UserScore",
        data() {
            return {
                gtStandard: [],
                ltStandard: [],
                topicRelatedList: [],
                option: {
                    backgroundColor: 'white',
                    "normal": {
                        "top": 200,
                        "left": 300,
                        "width": 500,
                        "height": 400,
                        "zIndex": 6,
                        "backgroundColor": ""
                    },
                    "color": ["rgba(245, 166, 35, 1)", "rgba(19, 173, 255, 1)"],
                    "title": {
                        "show": false,
                        "text": "测评结果",
                        "left": "80%",
                        "top": "2%",
                        "textStyle": {
                            "fontSize": 14,
                            "color": "black",
                            "fontStyle": "normal",
                            "fontWeight": "normal"
                        }
                    },
                    "tooltip": {
                        "show": true,
                        "trigger": "item"
                    },
                    "legend": {
                        "show": true,
                        "icon": "circle",
                        "left": "80%",
                        "top": "4%",
                        "orient": "horizontal",
                        "textStyle": {
                            "fontSize": 12,
                            "color": "black"
                        },
                        "data": ["基准分", "所得分"]
                    },
                    "radar": {
                        "center": ["50%", "50%"],
                        "radius": "70%",
                        "startAngle": 90,
                        "splitNumber": 4,
                        "shape": 'angle', // 拐点的样式，还可以取值'rect','angle'等
                        "splitArea": {
                            "areaStyle": {
                                "color": ["transparent"]
                            }
                        },
                        "axisLabel": {
                            "show": false,
                            "fontSize": 18,
                            "color": "#fff",
                            "fontStyle": "normal",
                            "fontWeight": "normal"
                        },
                        "axisLine": {
                            "show": true,
                            "lineStyle": {
                                "color": "grey" //
                            }
                        },
                        "splitLine": {
                            "show": true,
                            "lineStyle": {
                                "color": "grey" //
                            }
                        },
                        "indicator": [/*{
                            "name": "洪水",
                            "max": 88
                        }*/]
                    },
                    "series": [{
                        "name": "基准分",
                        "type": "radar",
                        "symbol": "circle",
                        "symbolSize": 10,
                        "areaStyle": {
                            "normal": {
                                "color": "rgba(245, 166, 35, 0.4)"
                            }
                        },
                        itemStyle: {
                            color: 'rgba(245, 166, 35, 1)',
                            borderColor: 'rgba(245, 166, 35, 0.3)',
                            borderWidth: 10,
                        },
                        "lineStyle": {
                            "normal": {
                                "type": "dashed",
                                "color": "rgba(245, 166, 35, 1)",
                                "width": 2
                            }
                        },
                        "data": [
                            /*[80, 50, 55, 80, 50, 80, 48, 43, 60, 78, 60, 40, 42, 44, 65]*/
                        ]
                    }, {
                        "name": "所得分",
                        "type": "radar",
                        "symbol": "circle",
                        "symbolSize": 10,
                        "itemStyle": {
                            "normal": {
                                color: 'rgba(19, 173, 255, 1)',
                                "borderColor": "rgba(19, 173, 255, 0.4)",
                                "borderWidth": 10
                            }
                        },
                        "areaStyle": {
                            "normal": {
                                "color": "rgba(19, 173, 255, 0.5)"
                            }
                        },
                        "lineStyle": {
                            "normal": {
                                "color": "rgba(19, 173, 255, 1)",
                                "width": 2,
                                "type": "dashed"
                            }
                        },
                        "data": [
                            /*[60, 60, 65, 60, 70, 40, 80, 63, 68, 60, 77, 60, 80, 62, 80]*/
                        ]
                    }]
                }
            }
        },
        created() {
        },
        computed: {
            recommendCourses() {
                let arr = []
                this.topicRelatedList.forEach(item => {
                    if (item.recommendType == 1) {//课程推荐
                        arr.push(item)
                    }
                })
                return arr;
            },
            recommendBooks() {
                let arr = []
                this.topicRelatedList.forEach(item => {
                    if (item.recommendType == 2) {//课程推荐
                        arr.push(item)
                    }
                })
                return arr;
            },
            recommendTools() {
                let arr = []
                this.topicRelatedList.forEach(item => {
                    if (item.recommendType == 3) {//课程推荐
                        arr.push(item)
                    }
                })
                return arr;
            },
        },
        mounted() {
            this.getUserScore();
        },
        methods: {
            async goToStore(relatedId) {
                let r = await toStore(relatedId)
                console.info(r + "--------------------")
                if (r.data.url != "") {
                    location.href = r.data.url
                } else {
                    Toast.fail("暂无相关课程")
                }
            },
            async getUserScore() {
                let result = await userScore()
                let indicator = result.data.surveyDimensions
                if (result.data.gtStandard) {
                    this.gtStandard = result.data.gtStandard
                }
                if (result.data.ltStandard) {
                    this.ltStandard = result.data.ltStandard
                }
                this.topicRelatedList = result.data.topicRelatedList
                indicator.forEach((e) => {
                    this.option.radar.indicator.push({name: e, max: 10})
                })
                this.option.series[0]["data"].push(result.data.standardScores)
                this.option.series[1]["data"].push(result.data.userScores)
                let radarChart = echarts.init(document.getElementById("radarChartDiv"))
                radarChart.setOption(this.option)
            },
        }
    }
</script>

<style lang="less" scoped>
    .mainDiv {
        padding-top: 3rem;
    }
</style>
