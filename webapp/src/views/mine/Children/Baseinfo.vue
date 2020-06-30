<template>
    <div>
        <van-nav-bar title="完善基本信息"
                     :fixed=true
                     :border=false
                     style="height:2.5rem"/>
        <div class="mainDiv">
            <!-- 教育背景 -->
            <van-cell-group v-show="currentStep == 1">
                <van-cell title="" value="" label="1. 基本信息" class="title"/>
                <van-field v-model="school"
                           label="毕业学校"
                           readonly is-link
                           @click="showSchoolPopView = true"
                           required/>
                <van-field :value="majorsStudied"
                           label="所修专业" is-link
                           @click="showMajorsStudiedView = true"
                           required/>
                <van-field :value="highestEducation"
                           label="最高学历"
                           required is-link
                           @click="showEducationPopView = true"
                           readonly/>
                <van-field :value="schoolSystem"
                           label="所上学制"
                           required is-link
                           @click="showSchoolSystemPopView = true"
                           readonly/>
            </van-cell-group>
            <!-- 工作单位 -->
            <van-cell-group v-show="currentStep == 2">
                <van-cell title="" value="" label="2. 工作单位"/>
                <van-field name="radio" required label="工作状态">
                    <template #input>
                        <van-radio-group v-model="baseinfo.workingState" direction="horizontal">
                            <van-radio v-for="(item, index) in workingStateOption" :name="item.propertyValue"
                                       :value="item.propertyValue" :key="item.propertyValue">{{item.text}}
                            </van-radio>
                        </van-radio-group>
                    </template>
                </van-field>

                <van-field v-model="baseinfo.company"
                           label="工作单位"
                           required/>
                <van-field v-model="areaSelected"
                           label="单位地域"
                           @click="showAreaPopView = true"
                           required readonly is-link/>
                <van-field v-model="baseinfo.jobTitle"
                           label="职务职位" required/>
                <van-field v-model="serviceType"
                           label="单位服务模式"
                           @click="showServiceTypePopView = true"
                           required readonly is-link/>
                <van-field v-model="income"
                           label="月度收入水平"
                           @click="showIncomePopView = true"
                           required readonly is-link/>
                <van-field name="benefits" label="福利待遇" required>
                    <template #input>
                        <van-checkbox-group v-model="baseinfo.benefits" direction="horizontal">
                            <van-checkbox v-for="(item, index) in benefitsOption" :name="item.propertyValue"
                                          :value="item.propertyValue" :key="item.propertyValue" shape="square">
                                {{item.text}}
                            </van-checkbox>
                        </van-checkbox-group>
                    </template>
                </van-field>
                <van-field name="childType" label="服务儿童类型" required>
                    <template #input>
                        <van-checkbox-group v-model="baseinfo.childType" direction="horizontal">
                            <van-checkbox v-for="(item, index) in childTypeOption" :name="item.propertyValue"
                                          :value="item.propertyValue" :key="item.propertyValue" shape="square">
                                {{item.text}}
                            </van-checkbox>
                        </van-checkbox-group>
                    </template>
                </van-field>
                <van-field v-model="childAge"
                           label="主要服务儿童年龄段"
                           @click="showChildAgePopView = true"
                           required readonly is-link/>
            </van-cell-group>
            <!-- 培训经历 -->
            <van-cell-group v-show="currentStep == 3">
                <van-cell title="" value="" label="3.培训经历"/>
                <van-field required label="我到目前为止参加培训数量（一周内的算作短期培训）">
                    <template #input>
                        <van-checkbox-group v-model="baseinfo.trainingNumber" direction="horizontal">
                            <van-checkbox v-for="(item, index) in trainingNumberOption" :name="item.propertyValue"
                                          :value="item.propertyValue" :key="item.propertyValue" shape="square">
                                {{item.text}}
                            </van-checkbox>
                        </van-checkbox-group>
                    </template>
                </van-field>
                <van-field v-model="trainingFee"
                           label="我去年在培训上的花费"
                           @click="showTrainingFeePopView = true"
                           required readonly is-link/>
                <van-cell title="" value="" label="我以往培训详情，请描述"/>
                <van-cell-group v-for="(item, index) in trainingInfos" :key="index">
                    <van-cell :title="'培训'+(index + 1)" icon="shop-o">
                        <template #right-icon>
                            <van-icon name="clear" style="line-height: inherit;" @click="removeTrainingInfo(index)"/>
                        </template>
                    </van-cell>
                    <van-field v-model="item.trainingCourse"
                               label="培训课程"/>
                    <van-field v-model="item.beginTimeStr" is-link readonly @click="handleSelectTime(1, index)"
                               label="开始时间"/>
                    <van-field v-model="item.endTimeStr" is-link readonly @click="handleSelectTime(2, index)"
                               label="结束时间"/>
                    <van-field v-model="item.paymentWayStr"
                               label="付费方式"
                               @click="handlePaymentWay(index)"
                               readonly is-link/>
                </van-cell-group>
                <van-button plain hairline type="info" size="large" @click="addTrainingInfo">添加培训经历</van-button>
            </van-cell-group>
        </div>
        <div class="buttonDiv">
            <van-button class="button" type="info" size="large" @click="previous" v-if="currentStep > 1">上一步
            </van-button>
            <van-button class="button" type="info" size="large" @click="next" v-if="currentStep < 3">下一步</van-button>
            <van-button class="button" type="info" size="large" @click="saveBaseinfo" v-if="currentStep == 3">提 交
            </van-button>
        </div>
        <!-- 弹框选项开始 -->
        <van-popup v-model="showSchoolPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="schoolColumns"
                        @cancel="showSchoolPopView = false"
                        @confirm="onConfirmSchool"/>
        </van-popup>

        <!-- 时间选择器 -->
        <van-popup v-model="showDateTimePopView" round position="bottom">
            <van-datetime-picker v-model="currentDate" type="date" @confirm="onConfirmDateTime"
                                 @cancel="showDateTimePopView = false"
                                 :formatter="formatter"
                                 :max-date="maxDate"
                                 :min-date="minDate"/>
        </van-popup>

        <van-popup v-model="showMajorsStudiedView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="majorsStudiedColumns"
                        @cancel="showMajorsStudiedView = false"
                        @confirm="onConfirmMajorsStudied"/>
        </van-popup>

        <van-popup v-model="showTrainingFeePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="trainingFeeColumns"
                        @cancel="showTrainingFeePopView = false"
                        @confirm="onConfirmTrainingFee"/>
        </van-popup>

        <van-popup v-model="showChildAgePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="childAgeColumns"
                        @cancel="showChildAgePopView = false"
                        @confirm="onConfirmChildAge"/>
        </van-popup>

        <van-popup v-model="showIncomePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="incomeColumns"
                        @cancel="showIncomePopView = false"
                        @confirm="onConfirmIncome"/>
        </van-popup>

        <van-popup v-model="showServiceTypePopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="serviceTypeColumns"
                        @cancel="showServiceTypePopView = false"
                        @confirm="onConfirmServiceType"/>
        </van-popup>

        <van-popup v-model="showAreaPopView" round position="bottom">
            <!--<van-picker show-toolbar  title=""
                    :columns="areaColumns"
                    @cancel="showAreaPopView = false"
                    @confirm="onConfirmArea" />-->
            <van-area :area-list="areas" ref="areaSelectPicker" title="" :value="baseinfo.area"
                      @cancel="showAreaPopView = false"
                      @confirm="onConfirmArea"/>
        </van-popup>

        <van-popup v-model="showSchoolSystemPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="schoolSystemColumns"
                        @cancel="showSchoolSystemPopView = false"
                        @confirm="onConfirmSchoolSystem"/>
        </van-popup>

        <van-popup v-model="showEducationPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="educationColumns"
                        @cancel="showEducationPopView = false"
                        @confirm="onConfirmEducation"/>
        </van-popup>

        <van-popup v-model="showPaymentWayPopView" round position="bottom">
            <van-picker show-toolbar title=""
                        :columns="paymentWayColumns"
                        @cancel="showPaymentWayPopView = false"
                        @confirm="onConfirmPaymentWay"/>
        </van-popup>
        <!-- 弹框选项结束 -->

    </div>
</template>

<script>
    import {mapState} from "vuex";
    import {selectOptions, areaList, saveBaseinfo, initUserBaseinfo} from "../../../serve/api";
    import areas from '../../../config/area.js'
    import {Toast} from "vant";
    import Moment from "moment";

    export default {
        data() {
            return {
                showSchoolPopView: false,
                schoolColumns: [],

                showMajorsStudiedView: false,
                majorsStudiedColumns: [],

                showTrainingFeePopView: false,
                trainingFeeColumns: [],

                showChildAgePopView: false,
                childAgeColumns: [],

                showIncomePopView: false,
                incomeColumns: [],

                showServiceTypePopView: false,
                serviceTypeColumns: [],

                showAreaPopView: false,
                areaColumns: [],

                showSchoolSystemPopView: false,
                schoolSystemColumns: [],

                showEducationPopView: false,
                educationColumns: [],

                showPaymentWayPopView: false,
                paymentWayNum: 0,
                paymentWayColumns: [],

                showDateTimePopView: false,
                currentDate: new Date('2000/01/01'),
                minDate: new Date('1949/01/01'),
                maxDate: new Date('2019/12/31'),
                //1:begin  2:end
                timePickerType: 0,
                timePickerNum: 0,

                trainingNumberOption: [],
                benefitsOption: [],
                childTypeOption: [],
                workingStateOption: [],
                currentStep: 1,
                areas: areas,
                areaSelected: '',
                modify: -1,
                // 基本信息
                baseinfo: {
                    // 教育背景
                    school: 0,
                    majorsStudied: 0,
                    highestEducation: 0,
                    schoolSystem: 0,
                    // 工作单位
                    workingState: 0,
                    company: '',
                    area: '',
                    jobTitle: '',
                    serviceType: 0,
                    income: 0,
                    benefits: [],
                    childType: [],
                    childAge: 0,
                    // 培训经历
                    trainingNumber: [],
                    trainingFee: 0,
                    trainingInfos: []
                },
            }
        },
        computed: {
            ...mapState(['userInfo']),
            school() {
                return this.getPropertyName(this.schoolColumns, this.baseinfo.school)
            },
            majorsStudied() {
                return this.getPropertyName(this.majorsStudiedColumns, this.baseinfo.majorsStudied)
            },
            trainingFee() {
                return this.getPropertyName(this.trainingFeeColumns, this.baseinfo.trainingFee)
            },
            childAge() {
                return this.getPropertyName(this.childAgeColumns, this.baseinfo.childAge)
            },
            income() {
                return this.getPropertyName(this.incomeColumns, this.baseinfo.income)
            },
            serviceType() {
                return this.getPropertyName(this.serviceTypeColumns, this.baseinfo.serviceType)
            },
            schoolSystem() {
                return this.getPropertyName(this.schoolSystemColumns, this.baseinfo.schoolSystem)
            },
            highestEducation() {
                return this.getPropertyName(this.educationColumns, this.baseinfo.highestEducation)
            },
            trainingInfos() {
                let newArr = [...this.baseinfo.trainingInfos]
                newArr.map(el => {
                    return el.beginTimeStr = Moment(el.beginTime).format('YYYY-MM-DD'), el.endTimeStr = Moment(el.endTime).format('YYYY-MM-DD'), el.paymentWayStr = this.getPropertyName(this.paymentWayColumns, el.paymentWay)
                })
                return newArr
            }
        },
        created() {
            this.modify = this.$route.query.modify
            this.initSelectOptions()
        },
        methods: {
            previous() {
                this.currentStep -= 1
            },
            next() {
                this.currentStep += 1
            },
            removeTrainingInfo(index) {
                this.baseinfo.trainingInfos.splice(index, 1)
            },
            addTrainingInfo() {
                this.baseinfo.trainingInfos.push({})
            },
            handlePaymentWay(index) {
                this.showPaymentWayPopView = true
                this.paymentWayNum = index
            },
            onConfirmPaymentWay(value) {
                this.showPaymentWayPopView = false
                let trainingInfo = this.baseinfo.trainingInfos[this.paymentWayNum]
                trainingInfo.paymentWay = value.propertyValue;
                this.$set(this.baseinfo.trainingInfos, this.paymentWayNum, trainingInfo);
            },
            // 格式化DateTime pickView
            formatter(type, value) {
                if (type === 'year') {
                    return `${value}`;
                } else if (type === 'month') {
                    return `${value}`
                } else if (type === 'day') {
                    return `${value}`
                }
                return value;
            },
            handleSelectTime(bOe, index) {
                this.showDateTimePopView = true
                this.timePickerNum = index
                this.timePickerType = bOe
            },
            onConfirmDateTime(value) {
                let trainingInfo = this.baseinfo.trainingInfos[this.timePickerNum]
                if (this.timePickerType == 1) {
                    trainingInfo.beginTime = Moment(value).format();
                } else {
                    trainingInfo.endTime = Moment(value).format();
                }
                this.$set(this.baseinfo.trainingInfos, this.timePickerNum, trainingInfo);
                this.showDateTimePopView = false;
            },
            onConfirmSchool(value, index) {
                this.baseinfo.school = value.propertyValue
                this.showSchoolPopView = false;
            },
            onConfirmMajorsStudied(value, index) {
                this.baseinfo.majorsStudied = value.propertyValue
                this.showMajorsStudiedView = false;
            },
            onConfirmTrainingFee(value, index) {
                this.baseinfo.trainingFee = value.propertyValue
                this.showTrainingFeePopView = false;
            },
            onConfirmChildAge(value, index) {
                this.baseinfo.childAge = value.propertyValue
                this.showChildAgePopView = false;
            },
            onConfirmIncome(value, index) {
                this.baseinfo.income = value.propertyValue
                this.showIncomePopView = false;
            },
            onConfirmServiceType(value, index) {
                this.baseinfo.serviceType = value.propertyValue
                this.showServiceTypePopView = false;
            },
            onConfirmArea(value, indexs) {
                this.areaSelected = value[0].name + "/" + value[1].name + "/" + value[2].name
                this.baseinfo.area = value[2].code
                this.showAreaPopView = false;
            },
            onConfirmSchoolSystem(value, index) {
                this.baseinfo.schoolSystem = value.propertyValue
                this.showSchoolSystemPopView = false;
            },
            onConfirmEducation(value, index) {
                this.baseinfo.highestEducation = value.propertyValue
                this.showEducationPopView = false;
            },

            goToPage(n) {
                this.$router.push({
                    'name': n
                });
            },
            getPropertyName(array, id) {
                let name = ""
                array.forEach(element => {
                    if (id == element.propertyValue) {
                        name = element.text;
                    }
                });
                return name
            },
            async initSelectOptions() {
                let result = await selectOptions();
                this.schoolColumns = result.data.dictGroup.school
                this.majorsStudiedColumns = result.data.dictGroup.study_major
                this.trainingFeeColumns = result.data.dictGroup.training_fee
                this.childAgeColumns = result.data.dictGroup.child_age
                this.incomeColumns = result.data.dictGroup.income
                this.serviceTypeColumns = result.data.dictGroup.service_type
                this.schoolSystemColumns = result.data.dictGroup.school_system
                this.paymentWayColumns = result.data.dictGroup.payment_way
                this.educationColumns = result.data.dictGroup.highest_education
                this.workingStateOption = result.data.dictGroup.working_state
                this.childTypeOption = result.data.dictGroup.child_type
                this.benefitsOption = result.data.dictGroup.benefits
                this.trainingNumberOption = result.data.dictGroup.training_number
                if (this.modify && this.modify == 1) {
                    // 获取基本信息
                    this.initUserBaseinfo();
                }
            },
            async initUserBaseinfo() {
                let result = await initUserBaseinfo()
                // 组装数据
                let obj = result.data.userBaseinfo
                obj.benefits = obj.benefits.split(',').map(Number)
                obj.childType = obj.childType.split(',').map(Number)
                obj.trainingNumber = obj.trainingNumber.split(',').map(Number)
                this.baseinfo = obj
                /*let v = this.$refs.areaSelectPicker;
                let value = v.getValues();
                this.areaSelected = value[0].name + "/" + value[1].name + "/" + value[2].name;
                console.log(this.baseinfo)*/
            },
            async initAreas() {
                let result = await areaList();
                this.areaColumns = result.data.provinces
            },
            async saveBaseinfo() {
                let result = await saveBaseinfo(this.baseinfo)
                Toast.success("保存成功！")
                this.$router.push({path: "/home"});
            }
        }
    }
</script>

<style lang="less" scoped>
    .mainDiv {
        padding: 2rem 1rem 2rem 1rem;

        .title {
            font-size: 1.8rem;
        }
    }

    .buttonDiv {
        padding: 2rem 1rem 2rem 1rem;

        .button {
            margin: 0.3rem 0rem;
        }
    }
</style>
