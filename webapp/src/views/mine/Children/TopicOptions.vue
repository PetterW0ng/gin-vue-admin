<template>
    <div>
        <van-nav-bar :title="title"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem;padding-top:0.5rem"/>
        <div class="mainDiv">
            <!-- 问卷题 -->
            <div class="topicDiv" v-for="(item,index) in topicAndOptions" :key="item.ID">
                <van-cell-group v-if="index == page" :border=false>
                    <div class="topicTitle">
                        {{index+1}} . {{item.isRequired == 2 ? "【选填】" + (item.topicType == 2 ? "【多选】"+item.title :
                        item.title) : (item.topicType == 2 ? "【多选】"+item.title : item.title)}}
                    </div>

                    <van-radio-group v-model="selects.radio" v-if="item.topicType == 1" class="options">
                        <van-radio v-for="(it, ix) in item.options" :name="it.ID" :key="ix" style="padding: 0.3em"
                                   @click="handleShowOthers($event,it.title)">{{it.title}}
                        </van-radio>
                    </van-radio-group>

                    <van-checkbox-group v-model="selects.result" v-if="item.topicType == 2" class="options">
                        <van-checkbox v-for="(it, ix) in item.options" :name="it.ID" :key="ix" shape="square"
                                      style="padding: 0.3em"
                                      @click="handleShowOthers($event,it.title)">{{it.title}}
                        </van-checkbox>
                    </van-checkbox-group>

                    <van-field v-model="selects.others" v-if="item.topicType == 3 || othersSelected" class="options"
                               rows="2" autosize type="textarea"
                               maxlength="50"
                               placeholder="请填写"
                               show-word-limit/>

                    <div class="buttonDiv">
                        <van-cell :border="false" style="padding-bottom: 0.1rem"
                                  :value="'当前进度: ' + (index+1) + ' / '+ (topicAndOptions.length)"/>
                        <van-button class="button" type="info" size="large" @click="toBack(index)" v-if="index != 0">
                            上一步
                        </van-button>
                        <van-button class="button" type="info" size="large" @click="toNext(index,item.ID)"
                                    v-if="index != (topicAndOptions.length -1)">下一步
                        </van-button>
                        <van-button class="button" type="info" size="large" @click="saveUserAnswer(index)"
                                    v-if="index == (topicAndOptions.length -1)">提 交
                        </van-button>
                    </div>
                </van-cell-group>
            </div>

        </div>
    </div>
</template>

<script type="text/javascript">
    import {queryTopicList, addUserAnswer} from "../../../serve/api";
    import {Toast} from "vant";

    export default {
        data() {
            return {
                isShow: true,
                othersSelected: false,
                // 当前题的选项
                selects: {
                    radio: '',
                    result: [],
                    others: ''
                },
                page: 0,
                answers: [],  // 已答题目的列表
                topicAndOptions: [],
                businessType: 0
            }
        },
        computed: {
            title() {
                return this.businessType == 3 ? "行业观点" : (this.businessType == 2 ? "专业理解" : "职业选择和体验")
            }
        },
        created() {
            this.getTopicAndOptions()
        },
        methods: {
            // 返回
            onClickLeft() {
                this.$router.back();
            },
            handleShowOthers(e, title) {
                let tgClassName = e.currentTarget.className;
                if (tgClassName === "van-radio") {
                    if (title == "其他") {
                        this.othersSelected = true
                    } else {
                        this.othersSelected = false
                        this.selects.others = ""
                    }
                } else if (tgClassName === "van-checkbox") {
                    if (title == "其他" && e.currentTarget.ariaChecked === "false") {
                        this.othersSelected = true
                    } else if (title == "其他" && e.currentTarget.ariaChecked === "true") {
                        this.othersSelected = false
                        this.selects.others = ""
                    }
                }
            },
            // 获取问卷数据
            async getTopicAndOptions() {
                let topicType = this.$route.params.topicType
                if (topicType == 'industry') {
                    this.businessType = 3
                } else if (topicType == 'perfession') {
                    this.businessType = 2
                } else if (topicType == 'jobInfo') {
                    this.businessType = 1
                } else {
                    Toast.fail("没有对应的问题哦！")
                    return
                }

                const result = await queryTopicList(this.businessType)
                this.topicAndOptions = result.data.topicOptions
            },
            async saveUserAnswer(index) {
                if (!this.saveLocalAnswer(index)) {
                    return
                }
                let data = {businessType: this.businessType, answers: this.answers}
                let result = await addUserAnswer(data)
                Toast.success("保存成功！")
                this.$router.push({path: "/home"});
            },
            // 点击下一题
            toNext(index, id) {
                // 1. 判断改题是否必答，是如果没有回答则无法跳转到下一题
                if (!this.saveLocalAnswer(index)) {
                    return
                }
                // 3. index ++；
                index++;
                // 4. 判断已答题目列表是否包含下标为index的题目，有则回显
                let ans = this.answers[index]
                if (ans !== undefined && ans !== '') {
                    this.selects.others = ans.others;
                    if (ans.others !== '') {
                        this.othersSelected = true
                    } else {
                        this.othersSelected = false
                    }
                    if (this.topicAndOptions[index].topicType == 1) {
                        this.selects.radio = ans.optionIds[0];
                        this.selects.result = []
                    } else {
                        this.selects.radio = '';
                        this.selects.result = ans.optionIds
                    }
                } else {
                    this.selects.radio = '';
                    this.selects.result = []
                    this.selects.others = ''
                    this.othersSelected = false
                }
                // 5. page ++
                this.page++;
                // 6. 其他（如按钮隐藏）
                // 当前页等于最后一题 下一题按钮隐藏
                if (this.page == this.topicAndOptions.length - 1) {
                    return this.isShow = false
                } else {
                    return this.isShow = true
                }
            },
            saveLocalAnswer(index) {
                let tAo = this.topicAndOptions[index]
                let topicType = tAo.topicType
                if (tAo.isRequired == 1) { // 必填项
                    if (topicType == 1 && this.selects.radio.length == 0) {// 单选
                        Toast.fail("该题为必填项！")
                        return false;
                    } else if (topicType == 2 && this.selects.result.length == 0) {// 多选
                        Toast.fail("该题为必填项！")
                        return false;
                    } else if (topicType == 3 && this.selects.others.length == 0) {// 问答
                        Toast.fail("该题为必填项！")
                        return false;
                    }
                    if (topicType == 1) {
                        this.selects.result[0] = this.selects.radio
                    }
                }

                let answer = {topicId: tAo.ID, others: this.selects.others, optionIds: this.selects.result}
                this.answers[index] = answer;
                return true
            },
            // 点击上一题
            toBack(index) {
                // 1. index --；
                index--;
                // 2. 从已回答列表中获取答题数据，回显
                let ans = this.answers[index]
                this.selects.others = ans.others;
                if (ans.others !== '') {
                    this.othersSelected = true
                } else {
                    this.othersSelected = false
                }
                if (this.topicAndOptions[index].topicType == 1) {
                    this.selects.radio = ans.optionIds[0];
                    this.selects.result = []
                } else {
                    this.selects.radio = '';
                    this.selects.result = ans.optionIds
                }
                // 3. page --
                this.page--;
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

        .topicDiv {
            .topicTitle {
                font-size: 0.85rem;
                padding: 0.3rem 1.5rem;
                line-height: 1.4rem;
            }

            .options {
                font-size: 0.8rem;
                padding: 0.5rem 2rem;
                line-height: 1.3rem;
            }

            .buttonDiv {
                padding: 0.3rem 1rem 2rem 1rem;

                .button {
                    margin: 0.2rem 0rem;
                }

                .van-cell__value--alone {
                    text-align: center;
                }
            }
        }
    }

</style>
