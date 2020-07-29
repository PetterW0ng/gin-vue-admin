<template>
    <div>
        <van-nav-bar title="评估结果"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem;padding-top:0.5rem"/>
        <div class="mainDiv" v-show="!dialogTableVisible">
            <!-- 雷达图 -->
            <div id="radarChartDiv" style="width: 100%;height: 380px"></div>
            <div class="recomand">
                <!-- 结果分析 -->
                <div class="mainTitle">结果分析</div>
                <van-cell :value="' 您在' + gtStandardStr + ' 方面的能力已满足职业要求，非常棒。'" v-if="gtStandard.length > 0"/>
                <van-cell :value="' 您在'+ ltStandardStr +' 方面的能力有所欠缺，还不能满足当前工作需要，建议学习以下内容进行提升：'"
                          v-if="ltStandard.length > 0"/>
                <!-- 培训课程 -->
                <div v-if="recommendCourses.length > 0">
                    <div class="title">建议培训课程</div>
                    <div v-for="item in recommendCourses" class="item">
                        <van-cell :title="item.recommendObject" icon="star-o" is-link
                                  style="padding-bottom: 0px;padding-top: 0.5rem" @click="goToStore(item.id)"/>
                    </div>
                </div>
                <!-- 阅读书籍 -->
                <div v-if="recommendBooks.length > 0">
                    <div class="title">建议阅读书籍</div>
                    <van-cell-group v-for="(item, index) in recommendBooks" class="item" :key="index">
                        <van-cell :title="item.recommendObject" :border="false" icon="notes-o"
                                  style="padding-bottom: 0px;padding-top: 0.7rem"/>
                        <van-cell :value="item.remark"
                                  style="font-size: 0.55rem;padding-left: 2.0rem;padding-top: 0.2rem"/>
                    </van-cell-group>
                </div>
                <!-- 干预工具 -->
                <div v-if="recommendTools.length > 0" style="padding-bottom: 1.5rem">
                    <div class="title">建议干预工具</div>
                    <div v-for="item in recommendTools" class="item">
                        <van-cell :title="item.recommendObject" icon="setting-o"
                                  style="padding-bottom: 0px;padding-top: 0.5rem" is-link @click="goToStore(item.id)"/>
                    </div>
                </div>
            </div>
        </div>
        <div ref="imageWrapper" class="shareBgDiv" v-show="showImageWrapper">
            <!-- 背景图 -->
            <div style="height: 100%;width: 100%">
                <img src="../../../images/mine/shareQr.jpg" style="height: 100%;width: 100%"/>
            </div>
            <!-- 雷达图 -->
            <div class="radarChartDiv2">
                <div id="radarChartDiv2" style="position: relative;height: 100%"></div>
            </div>
        </div>
        <div v-show="dialogTableVisible" align="center">
            <div style="align-content: center;padding-top: 3.5rem; padding-bottom: 1rem; font-size: 0.8rem">
                长按图片保存您的评估结果
            </div>
            <div style="align-content: center;"><img src="../../../images/mine/arrow.png"
                                                     style="width: 20px; height: 20px"/></div>
            <img :src="imgUrl" width="100%"/>
        </div>
        <van-button v-show="!dialogTableVisible" class="button" type="info" size="large" @click="toImage()">
            点击下载分享结果
        </van-button>
    </div>
</template>

<script>
    import {userScore, toStore} from "../../../serve/api";
    import echarts from 'echarts'
    import html2canvas from "html2canvas"
    import {Toast} from 'vant'

    export default {
        name: "UserScore",
        data() {
            return {
                imgUrl: '',
                showImageWrapper: true,
                dialogTableVisible: false,
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
                        "show": false,
                        "trigger": "item"
                    },
                    "legend": {
                        "show": true,
                        "icon": "circle",
                        itemHeight: 8,
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
                        "radius": "53%",
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
                            "fontSize": 10,
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
                            borderWidth: 5,
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
                                "borderWidth": 5
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
                },
                option1: {
                    backgroundColor: 'white',
                    "normal": {
                        "top": 210,
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
                        "show": false,
                        "trigger": "item"
                    },
                    "legend": {
                        "show": true,
                        "icon": "circle",
                        itemHeight: 4,
                        "left": "80%",
                        "top": "0",
                        "orient": "horizontal",
                        "textStyle": {
                            "fontSize": 8,
                            "color": "black"
                        },
                        "data": ["基准分", "所得分"]
                    },
                    "radar": {
                        "center": ["50%", "50%"],
                        "radius": "60%",
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
                            "fontSize": 8,
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
                        "symbolSize": 5,
                        "areaStyle": {
                            "normal": {
                                "color": "rgba(245, 166, 35, 0.4)"
                            }
                        },
                        itemStyle: {
                            color: 'rgba(245, 166, 35, 1)',
                            borderColor: 'rgba(245, 166, 35, 0.3)',
                            borderWidth: 5,
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
                        "symbolSize": 5,
                        "itemStyle": {
                            "normal": {
                                color: 'rgba(19, 173, 255, 1)',
                                "borderColor": "rgba(19, 173, 255, 0.4)",
                                "borderWidth": 5
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
                        let hasV = false
                        arr.forEach(im => {
                            if (im.recommendObject == item.recommendObject) {
                                hasV = true
                            }
                        })
                        if (!hasV) {
                            arr.push(item)
                        }
                    }
                })
                return arr;
            },
            gtStandardStr() {
                return this.gtStandard.join("，");
            },
            ltStandardStr() {
                return this.ltStandard.join("，")
            },
            recommendBooks() {
                let arr = []
                this.topicRelatedList.forEach(item => {
                    if (item.recommendType == 2) {//课程推荐
                        let hasV = false
                        arr.forEach(im => {
                            if (im.recommendObject == item.recommendObject) {
                                hasV = true
                            }
                        })
                        if (!hasV) {
                            arr.push(item)
                        }
                    }
                })
                return arr;
            },
            recommendTools() {
                let arr = []
                this.topicRelatedList.forEach(item => {
                    if (item.recommendType == 3) {//课程推荐
                        let hasV = false
                        arr.forEach(im => {
                            if (im.recommendObject == item.recommendObject) {
                                hasV = true
                            }
                        })
                        if (!hasV) {
                            arr.push(item)
                        }
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
                if (r.data.url != "") {
                    location.href = r.data.url
                } else {
                    Toast.fail("暂无相关课程!")
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
                if (result.data.topicRelatedList) {
                    this.topicRelatedList = result.data.topicRelatedList
                }
                indicator.forEach((e) => {
                    this.option.radar.indicator.push({name: e.substr(0, 6) + "\n" + e.substr(6, e.length), max: 10})
                    this.option1.radar.indicator.push({name: e.substr(0, 6) + "\n" + e.substr(6, e.length), max: 10})
                })
                this.option.series[0]["data"].push(result.data.standardScores)
                this.option.series[1]["data"].push(result.data.userScores)

                this.option1.series[0]["data"].push(result.data.standardScores)
                this.option1.series[1]["data"].push(result.data.userScores)
                let radarChart = echarts.init(document.getElementById("radarChartDiv"))
                radarChart.setOption(this.option)
                let radarChart2 = echarts.init(document.getElementById("radarChartDiv2"))
                radarChart2.setOption(this.option1)
                let that = this;
                setTimeout(function () {
                    that.showImageWrapper = false;
                }, 500);
            },
            drawPng() {
                html2canvas(this.$refs.imageWrapper).then(canvas => {
                    let dataURL = canvas.toDataURL("image/png");
                    this.imgUrl = dataURL;
                    if (this.imgUrl !== "") {
                        this.dialogTableVisible = true;
                        this.showImageWrapper = false;
                    }
                });
            },
            toImage() {
                window.pageYOffset = 0;
                document.documentElement.scrollTop = 0;
                document.body.scrollTop = 0;
                this.showImageWrapper = true;
                let that = this;
                setTimeout(function () {
                    that.drawPng();
                }, 1000);
                Toast.loading({
                    duration: 1000, // 持续展示 toast
                    forbidClick: true,
                    message: 'loading...',
                });
            }
        }
    }
</script>

<style lang="less" scoped>
    .van-nav-bar__title {
        font-size: 1rem;
    }

    .shareBgDiv {
        position: relative;
        width: 100%;
        height: 100%;
    }

    .radarChartDiv2 {
        position: absolute;
        width: 84%;
        height: 35%;
        z-index: 1;
        top: 50%;
        left: 50%;
        transform: translateY(-50%) translateX(-50%);
    }

    .mainDiv {
        padding-top: 3rem;

        .recomand {
            .mainTitle {
                padding: 0.4rem 0.75rem;
                color: #108EE9;
                font-weight: bold;
            }

            .title {
                padding: 0.8rem 1rem;
                font-weight: bold;
            }

            .item {
                padding: 0rem 0.85rem;
            }
        }

    }
</style>
