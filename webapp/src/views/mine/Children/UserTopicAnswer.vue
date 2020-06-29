<template>
    <div>
        <div v-for="(item,index) in getUserTopicAnswers" :key="index">
            <van-cell-group>
                <div class="content-title" style="padding-top: 1em">
                    {{index+1}} . {{item.topicTitle}}
                </div>

                <van-checkbox-group style="padding: 0.5em" disabled v-if="item.topicOptionSelected.length > 0">
                    <van-checkbox v-for="(it, ix) in item.topicOptionSelected" :key="ix" checked="true"
                                  style="padding: 0.3em">
                        {{it}}
                    </van-checkbox>
                </van-checkbox-group>

                <van-field style="padding: 0.5em" v-if="item.answer != ''" disabled v-model="item.answer"
                           rows="2" autosize type="textarea"
                           maxlength="50"
                           placeholder="请输入留言"
                           show-word-limit/>
            </van-cell-group>
        </div>

        <van-button type="primary" size="normal" @click="reAnswer">重新填写</van-button>
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
                    return el.topicOptionSelected = el.topicOptionSelected.split('[option]')
                })
                return newArr
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

<style scoped>

</style>
