<template>
    <div>
        <van-nav-bar :title="title"
                     :fixed=true
                     :border=false
                     style="height: 2.5rem;padding-top:0.5rem"/>
        <div class="mainDiv">
            <div v-for="(item,index) in getUserTopicAnswers" :key="index">
                <van-cell-group :border=false>
                    <div class="topicTitle">
                        {{index+1}} . {{item.topicTitle}}
                    </div>

                    <van-checkbox-group disabled v-if="item.topicOptionSelected.length > 0">
                        <van-checkbox v-for="(it, ix) in item.topicOptionSelected" :key="ix" checked="true"
                                      class="options">
                            {{it}}
                        </van-checkbox>
                    </van-checkbox-group>

                    <van-field v-if="item.answer != ''" disabled :value="item.answer" class="options"
                               rows="2" autosize type="textarea"
                               maxlength="50"
                               placeholder="请输入留言"
                               show-word-limit/>
                </van-cell-group>
            </div>
            <div class="buttonDiv">
                <van-button class="button" type="info" size="large" @click="reAnswer">重新填写</van-button>
            </div>
        </div>
    </div>
</template>

<script>
    import {loadUserTopicAnswers} from "../../../serve/api";
    import {Toast} from "vant";
    import {mapState} from "vuex";

    export default {
        name: "UserTopicAnswer",
        data() {
            return {
                topicType: '',
                checked: true,
                userTopicAnswers: []
            }
        },
        created() {
            this.topicType = this.$route.params.topicType
            this.getUserTopicAnswer()
        },
        computed: {
            ...mapState(['userInfo']),
            getUserTopicAnswers() {
                let newArr = [...this.userTopicAnswers]
                newArr.map(el => {
                    return el.topicOptionSelected == '' ? el.topicOptionSelected = [] : el.topicOptionSelected = el.topicOptionSelected.split('[option]')
                })
                return newArr
            },
            title() {
                return this.businessType == 3 ? "行业观点" : (this.businessType == 2 ? "专业理解" : "职业选择和体验")
            }
        },
        methods: {
            reAnswer() {
                this.$router.push({path: `/topicOptions/${this.topicType}`})
            },
            async getUserTopicAnswer() {
                if (this.topicType == 'industry') {
                    this.businessType = 3
                } else if (this.topicType == 'perfession') {
                    this.businessType = 2
                } else if (this.topicType == 'jobInfo') {
                    this.businessType = 1
                } else {
                    Toast.fail("没有对应的问题哦！")
                    return
                }
                let result = await loadUserTopicAnswers(this.businessType, this.userInfo[this.topicType + "Num"])
                this.userTopicAnswers = result.data.topicAnswers
            }
        }
    }
</script>

<style lang="less" scoped>
    .van-nav-bar__title {
        font-size: 1rem;
    }

    .mainDiv {
        padding-top: 3rem;

        .topicTitle {
            font-size: 0.85rem;
            padding: 0.3rem 1.5rem;
            line-height: 1.4rem;
        }

        .options {
            font-size: 0.8rem;
            padding: 0.2rem 2rem;
            line-height: 1.3rem;
        }
        .buttonDiv {
            padding: 2rem 1rem 2rem 1rem;

            .button {
                margin: 0.3rem 0rem;
            }
        }
    }
</style>
